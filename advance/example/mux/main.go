package main

import (
	"fmt"
	"math"
	"os"
	"syscall"
	"unsafe"

	"github.com/Lensual/go-libav/advance/goavcodec"
	"github.com/Lensual/go-libav/avcodec"
	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/swresample"
	"github.com/Lensual/go-libav/swscale"
)

const STREAM_DURATION = 10.0
const STREAM_FRAME_RATE = 25                     /* 25 images/s */
const STREAM_PIX_FMT = avutil.AV_PIX_FMT_YUV420P /* default pix_fmt */

const SCALE_FLAGS = swscale.SWS_BICUBIC

// a wrapper around a single output AVStream
type OutputStream struct {
	st  *avformat.CAVStream
	enc *avcodec.CAVCodecContext

	/* pts of the next frame that will be generated */
	nextPts      int64
	samplesCount int

	frame    *avutil.CAVFrame
	tmpFrame *avutil.CAVFrame

	tmpPkt *goavcodec.AVPacket

	t, tincr, tincr2 float32

	swsCtx *swscale.CSwsContext
	swrCtx *swresample.CSwrContext
}

func log_packet(fmt_ctx *avformat.CAVFormatContext, pkt *avcodec.CAVPacket) {
	var time_base *avutil.CAVRational = unsafe.Slice(fmt_ctx.GetStreams(), fmt_ctx.GetNbStreams())[pkt.GetStreamIndex()].GetTimeBasePtr()

	fmt.Printf("pts:%s pts_time:%s dts:%s dts_time:%s duration:%s duration_time:%s stream_index:%d\n",
		avutil.AvTs2Str(pkt.GetPts()), avutil.AvTs2Timestr(pkt.GetPts(), time_base),
		avutil.AvTs2Str(pkt.GetDts()), avutil.AvTs2Timestr(pkt.GetDts(), time_base),
		avutil.AvTs2Str(pkt.GetDuration()), avutil.AvTs2Timestr(pkt.GetDuration(), time_base),
		pkt.GetStreamIndex())
}

func write_frame(fmt_ctx *avformat.CAVFormatContext, c *avcodec.CAVCodecContext,
	st *avformat.CAVStream, frame *avutil.CAVFrame, pkt *avcodec.CAVPacket) int {

	var ret int

	// send the frame to the encoder
	ret = avcodec.AvcodecSendFrame(c, frame)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Error sending a frame to the encoder: %s\n",
			avutil.AvErr2str(ret))
		os.Exit(1)
	}

	for ret >= 0 {
		ret = avcodec.AvcodecReceivePacket(c, pkt)
		if ret == avutil.AVERROR(int(syscall.EAGAIN)) || ret == avutil.AVERROR_EOF {

			break
		} else if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error encoding a frame: %s\n", avutil.AvErr2str(ret))
			os.Exit(1)
		}

		/* rescale output packet timestamp values from codec to stream timebase */
		avcodec.AvPacketRescaleTs(pkt, c.GetTimeBase(), st.GetTimeBase())
		pkt.SetStreamIndex(st.GetIndex())

		/* Write the compressed frame to the media file. */
		log_packet(fmt_ctx, pkt)
		ret = avformat.AvInterleavedWriteFrame(fmt_ctx, pkt)
		/* pkt is now blank (av_interleaved_write_frame() takes ownership of
		 * its contents and resets pkt), so that no unreferencing is necessary.
		 * This would be different if one used av_write_frame(). */
		if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error while writing output packet: %s\n", avutil.AvErr2str(ret))
			os.Exit(1)
		}
	}

	if ret == avutil.AVERROR_EOF {
		return 1
	}
	return 0
}

/* Add an output stream. */
func add_stream(ost *OutputStream, oc *avformat.CAVFormatContext, codec **avcodec.CAVCodec, codecId avcodec.CAVCodecID) {
	var c *avcodec.CAVCodecContext = nil
	// var i int = 0

	/* find the encoder */
	*codec = avcodec.AvcodecFindEncoder(codecId)
	if *codec == nil {
		fmt.Fprintf(os.Stderr, "Could not find encoder for '%s'\n",
			avcodec.AvcodecGetName(codecId))
		os.Exit(1)
	}

	ost.tmpPkt = goavcodec.AllocAvPacket()
	if ost.tmpPkt == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate AVPacket\n")
		os.Exit(1)
	}

	ost.st = avformat.AvformatNewStream(oc, nil)
	if ost.st == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate stream\n")
		os.Exit(1)
	}
	ost.st.SetId(int(oc.GetNbStreams()) - 1)
	c = avcodec.AvcodecAllocContext3(*codec)
	if c == nil {
		fmt.Fprintf(os.Stderr, "Could not alloc an encoding context\n")
		os.Exit(1)
	}
	ost.enc = c

	switch (*codec).GetType() {
	case avutil.AVMEDIA_TYPE_AUDIO:
		var sampleFmt avutil.CAVSampleFormat
		if (*codec).GetSampleFmts() != nil {
			sampleFmt = *(*codec).GetSampleFmts() //(*codec)->sample_fmts[0]
		} else {
			sampleFmt = avutil.AV_SAMPLE_FMT_FLTP
		}
		c.SetSampleFmt(sampleFmt)
		c.SetBitRate(64000)
		c.SetSampleRate(44100)
		if (*codec).GetSupportedSamplerates() != nil {
			sampleRateCArr := (*codec).GetSupportedSamplerates()
			sampleRateCPtr := unsafe.Pointer(sampleRateCArr)
			sampleRateCSize := int(unsafe.Sizeof(*sampleRateCArr))
			c.SetSampleRate(int(*sampleRateCArr)) //c->sample_rate = (*codec)->supported_samplerates[0];
			for i := 0; ; i++ {                   //for (i = 0; (*codec)->supported_samplerates[i]; i++) {
				sampleRate := *(*int)(unsafe.Add(sampleRateCPtr, i*sampleRateCSize))
				if sampleRate == 0 {
					break
				}
				if sampleRate == 44100 {
					c.SetSampleRate(44100)
				}
			}
		}
		src := avutil.AV_CHANNEL_LAYOUT_STEREO()
		avutil.AvChannelLayoutCopy(c.GetChLayoutPtr(), &src)
		timebase := avutil.CAVRational{}
		timebase.SetNum(1)
		timebase.SetDen(c.GetSampleRate())
		ost.st.SetTimeBase(timebase)
		break
	case avutil.AVMEDIA_TYPE_VIDEO:
		c.SetCodecId(codecId)

		c.SetBitRate(400000)

		/* Resolution must be a multiple of two. */
		c.SetWidth(352)
		c.SetHeight(288)

		timebase := avutil.CAVRational{}
		timebase.SetNum(1)
		timebase.SetDen(STREAM_FRAME_RATE)

		/* timebase: This is the fundamental unit of time (in seconds) in terms
		 * of which frame timestamps are represented. For fixed-fps content,
		 * timebase should be 1/framerate and timestamp increments should be
		 * identical to 1. */
		ost.st.SetTimeBase(timebase)
		c.SetTimeBase(ost.st.GetTimeBase())

		c.SetGopSize(12) /* emit one intra frame every twelve frames at most */
		c.SetPixFmt(STREAM_PIX_FMT)
		if c.GetCodecId() == avcodec.AV_CODEC_ID_MPEG2VIDEO {
			/* just for testing, we also add B-frames */
			c.SetMaxBFrames(2)
		}
		if c.GetCodecId() == avcodec.AV_CODEC_ID_MPEG1VIDEO {
			/* Needed to avoid using macroblocks in which some coeffs overflow.
			 * This does not happen with normal video, it just happens here as
			 * the motion of the chroma plane does not match the luma plane. */
			c.SetMbDecision(2)
		}
		break
	default:
		break

	}

	/* Some formats want stream headers to be separate. */
	if (oc.GetOFormat().GetFlags() & avformat.AVFMT_GLOBALHEADER) != 0 {
		c.SetFlags(c.GetFlags() | avcodec.AV_CODEC_FLAG_GLOBAL_HEADER)
	}
}

/**************************************************************/
/* audio output */

func alloc_audio_frame(sampleFmt avutil.CAVSampleFormat, channelLayout *avutil.CAVChannelLayout, sampleRate int, nbSamples int) *avutil.CAVFrame {
	var frame *avutil.CAVFrame = avutil.AvFrameAlloc()
	if frame == nil {
		fmt.Fprintf(os.Stderr, "Error allocating an audio frame\n")
		os.Exit(1)
	}

	frame.SetFormat(int(sampleFmt))
	avutil.AvChannelLayoutCopy(frame.GetChLayoutPtr(), channelLayout)
	frame.SetSampleRate(sampleRate)
	frame.SetNbSamples(nbSamples)

	if nbSamples != 0 {
		if avutil.AvFrameGetBuffer(frame, 0) < 0 {
			fmt.Fprintf(os.Stderr, "Error allocating an audio buffer\n")
			os.Exit(1)
		}
	}

	return frame
}

func open_audio(oc *avformat.CAVFormatContext, codec *avcodec.CAVCodec, ost *OutputStream, optArg *avutil.CAVDictionary) {
	var c *avcodec.CAVCodecContext
	var nb_samples int
	var ret int
	var opt *avutil.CAVDictionary = nil

	c = ost.enc

	/* open it */
	avutil.AvDictCopy(&opt, optArg, 0)
	ret = avcodec.AvcodecOpen2(c, codec, &opt)
	avutil.AvDictFree(&opt)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open audio codec: %s\n", avutil.AvErr2str(ret))
		os.Exit(1)
	}

	/* init signal generator */
	ost.t = 0
	ost.tincr = 2 * math.Pi * 110.0 / float32(c.GetSampleRate())
	/* increment frequency by 110 Hz per second */
	ost.tincr2 = 2 * math.Pi * 110.0 / float32(c.GetSampleRate()) / float32(c.GetSampleRate())

	if (c.GetCodec().GetCapabilities() & avcodec.AV_CODEC_CAP_VARIABLE_FRAME_SIZE) != 0 {
		nb_samples = 10000
	} else {
		nb_samples = c.GetFrameSize()
	}

	cChLayout := c.GetChLayout()
	ost.frame = alloc_audio_frame(c.GetSampleFmt(), &cChLayout,
		c.GetSampleRate(), nb_samples)
	ost.tmpFrame = alloc_audio_frame(avutil.AV_SAMPLE_FMT_S16, &cChLayout,
		c.GetSampleRate(), nb_samples)

	/* copy the stream parameters to the muxer */
	ret = avcodec.AvcodecParametersFromContext(ost.st.GetCodecPar(), c)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not copy the stream parameters\n")
		os.Exit(1)
	}

	/* create resampler context */
	ost.swrCtx = swresample.SwrAlloc()
	if ost.swrCtx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate resampler context\n")
		os.Exit(1)
	}

	/* set options */
	avutil.AvOptSetChlayout(unsafe.Pointer(ost.swrCtx), "in_chlayout", &cChLayout, 0)
	avutil.AvOptSetInt(unsafe.Pointer(ost.swrCtx), "in_sample_rate", int64(c.GetSampleRate()), 0)
	avutil.AvOptSetSampleFmt(unsafe.Pointer(ost.swrCtx), "in_sample_fmt", avutil.AV_SAMPLE_FMT_S16, 0)
	avutil.AvOptSetChlayout(unsafe.Pointer(ost.swrCtx), "out_chlayout", &cChLayout, 0)
	avutil.AvOptSetInt(unsafe.Pointer(ost.swrCtx), "out_sample_rate", int64(c.GetSampleRate()), 0)
	avutil.AvOptSetSampleFmt(unsafe.Pointer(ost.swrCtx), "out_sample_fmt", c.GetSampleFmt(), 0)

	/* initialize the resampling context */
	ret = swresample.SwrInit(ost.swrCtx)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Failed to initialize the resampling context\n")
		os.Exit(1)
	}
}

/* Prepare a 16 bit dummy audio frame of 'frame_size' samples and
 * 'nb_channels' channels. */
func get_audio_frame(ost *OutputStream) *avutil.CAVFrame {
	var frame *avutil.CAVFrame = ost.tmpFrame
	var j, i, v int
	var q *int16 = (*int16)(frame.GetData()[0])

	/* check if we want to generate more frames */
	rational := avutil.CAVRational{}
	rational.SetNum(1)
	rational.SetDen(1)
	if avutil.AvCompareTs(ost.nextPts, ost.enc.GetTimeBase(),
		STREAM_DURATION, rational) > 0 {
		return nil
	}

	for j = 0; j < frame.GetNbSamples(); j++ {
		v = (int)(math.Sin(float64(ost.t)) * 10000)
		for i = 0; i < ost.enc.GetChLayout().GetNbChannels(); i++ {
			*q = (int16)(v)
			*q++
		}
		ost.t += ost.tincr
		ost.tincr += ost.tincr2
	}

	frame.SetPts(ost.nextPts)
	ost.nextPts += int64(frame.GetNbSamples())

	return frame
}

/*
 * encode one audio frame and send it to the muxer
 * return 1 when encoding is finished, 0 otherwise
 */
func write_audio_frame(oc *avformat.CAVFormatContext, ost *OutputStream) int {
	var c *avcodec.CAVCodecContext
	var frame *avutil.CAVFrame
	var ret int
	var dst_nb_samples int

	c = ost.enc

	frame = get_audio_frame(ost)

	if frame != nil {
		/* convert samples from native format to destination codec format, using the resampler */
		/* compute destination number of samples */
		dst_nb_samples = int(avutil.AvRescaleRnd(swresample.SwrGetDelay(ost.swrCtx, int64(c.GetSampleRate()))+int64(frame.GetNbSamples()),
			int64(c.GetSampleRate()), int64(c.GetSampleRate()), avutil.AV_ROUND_UP))
		avutil.AvAssert0(dst_nb_samples == frame.GetNbSamples())

		/* when we pass a frame to the encoder, it may keep a reference to it
		 * internally;
		 * make sure we do not overwrite it here
		 */
		ret = avutil.AvFrameMakeWritable(ost.frame)
		if ret < 0 {
			os.Exit(1)
		}

		/* convert to destination format */
		ostFrameData := ost.frame.GetData()
		frameData := frame.GetData()
		ret = swresample.SwrConvert(ost.swrCtx,
			(**uint8)(unsafe.Pointer((unsafe.SliceData(ostFrameData[:])))), dst_nb_samples,
			(**uint8)(unsafe.Pointer((unsafe.SliceData(frameData[:])))), frame.GetNbSamples())
		if ret < 0 {
			fmt.Fprintf(os.Stderr, "Error while converting\n")
			os.Exit(1)
		}
		frame = ost.frame

		rational := avutil.CAVRational{}
		rational.SetNum(1)
		rational.SetDen(c.GetSampleRate())

		frame.SetPts(avutil.AvRescaleQ(int64(ost.samplesCount), rational, c.GetTimeBase()))
		ost.samplesCount += dst_nb_samples
	}

	return write_frame(oc, c, ost.st, frame, ost.tmpPkt.CAVPacket)
}

/**************************************************************/
/* video output */

func alloc_picture(pixFmt avutil.CAVPixelFormat, width int, height int) *avutil.CAVFrame {
	var picture *avutil.CAVFrame
	var ret int = 0

	picture = avutil.AvFrameAlloc()
	if picture == nil {
		return nil
	}

	picture.SetFormat(int(pixFmt))
	picture.SetWidth(width)
	picture.SetHeight(height)

	/* allocate the buffers for the frame data */
	ret = avutil.AvFrameGetBuffer(picture, 0)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate frame data.\n")
		os.Exit(1)
	}

	return picture
}

func open_video(oc *avformat.CAVFormatContext, codec *avcodec.CAVCodec, ost *OutputStream, optArg *avutil.CAVDictionary) {
	var ret int = 0
	var c *avcodec.CAVCodecContext = ost.enc
	var opt *avutil.CAVDictionary = nil

	avutil.AvDictCopy(&opt, optArg, 0)

	/* open the codec */
	ret = avcodec.AvcodecOpen2(c, codec, &opt)
	avutil.AvDictFree(&opt)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open video codec: %s\n", avutil.AvErr2str(ret))
		os.Exit(1)
	}

	/* allocate and init a re-usable frame */
	ost.frame = alloc_picture(c.GetPixFmt(), c.GetWidth(), c.GetHeight())
	if ost.frame == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate video frame\n")
		os.Exit(1)
	}

	/* If the output format is not YUV420P, then a temporary YUV420P
	 * picture is needed too. It is then converted to the required
	 * output format. */
	ost.tmpFrame = nil
	if c.GetPixFmt() != avutil.AV_PIX_FMT_YUV420P {
		ost.tmpFrame = alloc_picture(avutil.AV_PIX_FMT_YUV420P, c.GetWidth(), c.GetHeight())
		if ost.tmpFrame == nil {
			fmt.Fprintf(os.Stderr, "Could not allocate temporary picture\n")
			os.Exit(1)
		}
	}

	/* copy the stream parameters to the muxer */
	ret = avcodec.AvcodecParametersFromContext(ost.st.GetCodecPar(), c)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not copy the stream parameters\n")
		os.Exit(1)
	}
}

/* Prepare a dummy image. */
func fill_yuv_image(pict *avutil.CAVFrame, frame_index int,
	width int, height int) {
	var x, y, i int

	i = frame_index

	/* Y */
	for y = 0; y < height; y++ {
		for x = 0; x < width; x++ {
			//pict.data[0][y*pict.linesize[0]+x] = x + y + i*3
			linesize := int(pict.GetLinesize()[0])
			data := unsafe.Slice((*uint8)(pict.GetData()[0]), height*linesize)
			data[y*linesize+x] = (uint8)(x + y + i*3)
		}
	}

	/* Cb and Cr */
	for y = 0; y < height/2; y++ {
		for x = 0; x < width/2; x++ {
			//pict.data[1][y*pict.linesize[1]+x] = 128 + y + i*2
			linesize := int(pict.GetLinesize()[1])
			data := unsafe.Slice((*uint8)(pict.GetData()[1]), height*linesize)
			data[y*linesize+x] = (uint8)(128 + y + i*2)
			//pict.data[2][y*pict.linesize[2]+x] = 64 + x + i*5
			linesize = int(pict.GetLinesize()[2])
			data = unsafe.Slice((*uint8)(pict.GetData()[2]), height*linesize)
			data[y*linesize+x] = (uint8)(64 + x + i*5)
		}
	}
}

func get_video_frame(ost *OutputStream) *avutil.CAVFrame {
	var c *avcodec.CAVCodecContext = ost.enc

	/* check if we want to generate more frames */
	rational := avutil.CAVRational{}
	rational.SetNum(1)
	rational.SetDen(1)
	if avutil.AvCompareTs(ost.nextPts, c.GetTimeBase(),
		STREAM_DURATION, rational) > 0 {
		return nil
	}

	/* when we pass a frame to the encoder, it may keep a reference to it
	 * internally; make sure we do not overwrite it here */
	if avutil.AvFrameMakeWritable(ost.frame) < 0 {
		os.Exit(1)
	}

	if c.GetPixFmt() != avutil.AV_PIX_FMT_YUV420P {
		/* as we only generate a YUV420P picture, we must convert it
		 * to the codec pixel format if needed */
		if ost.swsCtx == nil {
			ost.swsCtx = swscale.SwsGetContext(c.GetWidth(), c.GetHeight(),
				avutil.AV_PIX_FMT_YUV420P,
				c.GetWidth(), c.GetHeight(),
				c.GetPixFmt(),
				SCALE_FLAGS, nil, nil, nil)
			if ost.swsCtx == nil {
				fmt.Fprintf(os.Stderr,
					"Could not initialize the conversion context\n")
				os.Exit(1)
			}
		}
		fill_yuv_image(ost.tmpFrame, int(ost.nextPts), c.GetWidth(), c.GetHeight())
		//  sws_scale(ost.sws_ctx, (const uint8_t * const *) ost.tmp_frame.data,
		// 		   ost.tmp_frame.linesize, 0, c->height, ost->frame->data,
		// 		   ost.frame->linesize);
	} else {
		fill_yuv_image(ost.frame, int(ost.nextPts), c.GetWidth(), c.GetHeight())
	}

	ost.frame.SetPts(ost.nextPts)
	ost.nextPts++

	return ost.frame
}

/*
 * encode one video frame and send it to the muxer
 * return 1 when encoding is finished, 0 otherwise
 */
func write_video_frame(oc *avformat.CAVFormatContext, ost *OutputStream) int {
	return write_frame(oc, ost.enc, ost.st, get_video_frame(ost), ost.tmpPkt.CAVPacket)
}

func close_stream(oc *avformat.CAVFormatContext, ost *OutputStream) {
	ostEnc := ost.enc
	avcodec.AvcodecFreeContext(&ostEnc)
	ostFrame := ost.frame
	avutil.AvFrameFree(&ostFrame)
	ostTmpFrame := ost.tmpFrame
	avutil.AvFrameFree(&ostTmpFrame)
	ostTmpPkt := ost.tmpPkt.CAVPacket
	avcodec.AvPacketFree(&ostTmpPkt)
	swscale.SwsFreeContext(ost.swsCtx)
	ostSwrCtx := ost.swrCtx
	swresample.SwrFree(&ostSwrCtx)
}

/**************************************************************/
/* media file output */

func main() {
	var video_st, audio_st OutputStream = OutputStream{}, OutputStream{}
	var avfmt *avformat.CAVOutputFormat
	var filename string
	var oc *avformat.CAVFormatContext
	var audio_codec, video_codec *avcodec.CAVCodec
	var ret int
	have_video, have_audio := false, false
	encode_video, encode_audio := false, false
	var opt *avutil.CAVDictionary = nil
	var i int

	if len(os.Args) < 2 {
		fmt.Printf("usage: %s output_file\n"+
			"API example program to output a media file with libavformat.\n"+
			"This program generates a synthetic audio and video stream, encodes and\n"+
			"muxes them into a file named output_file.\n"+
			"The output format is automatically guessed according to the file extension.\n"+
			"Raw images can also be output by using '%%d' in the filename.\n"+
			"\n", os.Args[0])
		os.Exit(1)
	}

	filename = os.Args[1]
	for i = 2; i+1 < len(os.Args); i += 2 {
		if !(os.Args[i] == "-flags") || !(os.Args[i] == "-fflags") {
			avutil.AvDictSet(&opt, os.Args[i][1:], os.Args[i+1], 0)
		}
	}

	/* allocate the output media context */
	avformat.AvformatAllocOutputContext2(&oc, nil, "", filename)
	if oc == nil {
		fmt.Printf("Could not deduce output format from file extension: using MPEG.\n")
		avformat.AvformatAllocOutputContext2(&oc, nil, "mpeg", filename)
	}
	if oc == nil {
		os.Exit(1)
	}

	avfmt = oc.GetOFormat()

	/* Add the audio and video streams using the default format codecs
	 * and initialize the codecs. */
	if avfmt.GetVideoCodec() != avcodec.AV_CODEC_ID_NONE {
		add_stream(&video_st, oc, &video_codec, avfmt.GetVideoCodec())
		have_video = true
		encode_video = true
	}
	if avfmt.GetAudioCodec() != avcodec.AV_CODEC_ID_NONE {
		add_stream(&audio_st, oc, &audio_codec, avfmt.GetAudioCodec())
		have_audio = true
		encode_audio = true
	}

	/* Now that all the parameters are set, we can open the audio and
	 * video codecs and allocate the necessary encode buffers. */
	if have_video {
		open_video(oc, video_codec, &video_st, opt)
	}

	if have_audio {
		open_audio(oc, audio_codec, &audio_st, opt)
	}

	avformat.AvDumpFormat(oc, 0, filename, 1)

	/* open the output file, if needed */
	if (avfmt.GetFlags() & avformat.AVFMT_NOFILE) == 0 {
		ret = avformat.AvioOpen(oc.GetPBPtr(), filename, avformat.AVIO_FLAG_WRITE)
		if ret < 0 {
			fmt.Fprintf(os.Stderr, "Could not open '%s': %s\n", filename,
				avutil.AvErr2str(ret))
			os.Exit(1)
		}
	}

	/* Write the stream header, if any. */
	ret = avformat.AvformatWriteHeader(oc, &opt)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Error occurred when opening output file: %s\n",
			avutil.AvErr2str(ret))
		os.Exit(1)
	}

	for encode_video || encode_audio {
		/* select the stream to encode */
		if encode_video && (!encode_audio || avutil.AvCompareTs(video_st.nextPts, video_st.enc.GetTimeBase(),
			audio_st.nextPts, audio_st.enc.GetTimeBase()) <= 0) {
			encode_video = !(write_video_frame(oc, &video_st) != 0)
		} else {
			encode_audio = !(write_audio_frame(oc, &audio_st) != 0)
		}
	}

	avformat.AvWriteTrailer(oc)

	/* Close each codec. */
	if have_video {
		close_stream(oc, &video_st)
	}
	if have_audio {
		close_stream(oc, &audio_st)
	}

	if (avfmt.GetFlags() & avformat.AVFMT_NOFILE) == 0 {
		/* Close the output file. */
		avformat.AvioClosep(oc.GetPBPtr())
	}

	/* free the stream */
	avformat.AvformatFreeContext(oc)

	os.Exit(0)
}
