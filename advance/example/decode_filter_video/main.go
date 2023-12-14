package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/Lensual/go-libav/advance/goavcodec"
	"github.com/Lensual/go-libav/advance/goavfilter"
	"github.com/Lensual/go-libav/advance/goavformat"
	"github.com/Lensual/go-libav/advance/goavutil"
	"github.com/Lensual/go-libav/avfilter"
	"github.com/Lensual/go-libav/avutil"
)

//  #define _XOPEN_SOURCE 600 /* for usleep */

var filterDescr string = "scale=78:24,transpose=cclock"

//  /* other way:
// 	scale=78:24 [scl]; [scl] transpose=cclock // assumes "[in]" and "[out]" to be input output pads respectively
//   */

var fmtCtx *goavformat.AvformatContext
var decCtx *goavcodec.AVCodecContext

var buffersinkCtx *goavfilter.BufferSink
var buffersrcCtx *goavfilter.BufferSrc
var filterGraph *goavfilter.AVFilterGraph
var videoStreamIndex int = -1
var lastPts int64 = avutil.AV_NOPTS_VALUE

func openInputFile(filename string) int {
	var dec *goavcodec.AVCodec
	var ret int

	if ret, fmtCtx = goavformat.OpenInput(filename, nil, nil); ret < 0 {
		avutil.AvLog(nil, avutil.AV_LOG_ERROR, "Cannot open input file\n")
		return ret
	}

	if ret = fmtCtx.FindStreamInfo(nil); ret < 0 {
		avutil.AvLog(nil, avutil.AV_LOG_ERROR, "Cannot find stream information\n")
		return ret
	}

	/* select the video stream */
	ret, dec = fmtCtx.FindBestStream(avutil.AVMEDIA_TYPE_VIDEO, -1, -1, 0)
	if ret < 0 {
		avutil.AvLog(nil, avutil.AV_LOG_ERROR, "Cannot find a video stream in the input file\n")
		return ret
	}
	videoStreamIndex = ret

	/* create decoding context */
	decCtx = dec.CreateContext()
	if decCtx == nil {
		return avutil.AVERROR(int(syscall.ENOMEM))
	}
	decCtx.ParametersFrom(fmtCtx.GetStreams()[videoStreamIndex].GetCodecPar())

	/* init the video decoder */
	if ret = decCtx.Open(nil); ret < 0 {
		avutil.AvLog(nil, avutil.AV_LOG_ERROR, "Cannot open video decoder\n")
		return ret
	}

	return 0
}

func initFilters(filtersDescr string) int {
	var args string
	var ret int = 0
	// var buffersrc *goavfilter.AVFilter = goavfilter.GetByName("buffer")
	// var buffersink *goavfilter.AVFilter = goavfilter.GetByName("buffersink")
	var outputs *goavfilter.AVFilterInOut = goavfilter.AllocAvFilterInOut()
	var inputs *goavfilter.AVFilterInOut = goavfilter.AllocAvFilterInOut()
	var timeBase avutil.CAVRational = fmtCtx.GetStreams()[videoStreamIndex].GetTimeBase()
	var pixFmts []avutil.CAVPixelFormat = []avutil.CAVPixelFormat{avutil.AV_PIX_FMT_GRAY8, avutil.AV_PIX_FMT_NONE}
	cpixFmts := unsafe.SliceData(pixFmts)
	AV_PIX_FMT_NONE := avutil.AV_PIX_FMT_NONE

	filterGraph = goavfilter.AllocAvFilterGraph()
	if outputs == nil || inputs == nil || filterGraph == nil {
		ret = avutil.AVERROR(int(syscall.ENOMEM))
		goto end
	}

	/* buffer video source: the decoded frames from the decoder will be inserted here. */
	args = fmt.Sprintf("video_size=%dx%d:pix_fmt=%d:time_base=%d/%d:pixel_aspect=%d/%d",
		decCtx.GetWidth(), decCtx.GetHeight(), decCtx.GetPixFmt(),
		timeBase.GetNum(), timeBase.GetDen(),
		decCtx.GetSampleAspectRatio().GetNum(), decCtx.GetSampleAspectRatio().GetDen())

	buffersrcCtx, ret = filterGraph.CreateBufferSrc("in", args)
	if ret < 0 {
		avutil.AvLog(nil, avutil.AV_LOG_ERROR, "Cannot create buffer source\n")
		goto end
	}

	/* buffer video sink: to terminate the filter chain. */
	buffersinkCtx, ret = filterGraph.CreateBufferSink("out", "")
	if ret < 0 {
		avutil.AvLog(nil, avutil.AV_LOG_ERROR, "Cannot create buffer sink\n")
		goto end
	}

	ret = avutil.AvOptSetIntList(unsafe.Pointer(buffersinkCtx.CAVFilterContext), "pix_fmts", unsafe.Pointer(cpixFmts),
		uint64(AV_PIX_FMT_NONE), avutil.AV_OPT_SEARCH_CHILDREN)
	if ret < 0 {
		avutil.AvLog(nil, avutil.AV_LOG_ERROR, "Cannot set output pixel format\n")
		goto end
	}

	/*
	 * Set the endpoints for the filter graph. The filter_graph will
	 * be linked to the graph described by filters_descr.
	 */

	/*
	 * The buffer source output must be connected to the input pad of
	 * the first filter described by filters_descr; since the first
	 * filter input label is not specified, it is set to "in" by
	 * default.
	 */
	outputs.SetName("in")
	outputs.SetFilterCtx(buffersrcCtx.AVFilterContext)
	outputs.SetPadIdx(0)
	outputs.SetNext(nil)

	/*
	 * The buffer sink input must be connected to the output pad of
	 * the last filter described by filters_descr; since the last
	 * filter output label is not specified, it is set to "out" by
	 * default.
	 */
	inputs.SetName("out")
	inputs.SetFilterCtx(buffersinkCtx.AVFilterContext)
	inputs.SetPadIdx(0)
	inputs.SetNext(nil)

	if ret = filterGraph.ParsePtr(filtersDescr, inputs, outputs, nil); ret < 0 {
		goto end
	}

	if ret = filterGraph.ValidConfig(nil); ret < 0 {
		goto end
	}

end:
	inputs.Free()
	outputs.Free()

	return ret
}

func displayFrame(frame *goavutil.AVFrame, timeBase avutil.CAVRational) {
	var x, y int
	var p0, p *uint8
	var delay int64

	if frame.GetPts() != avutil.AV_NOPTS_VALUE {
		if lastPts != avutil.AV_NOPTS_VALUE {
			/* sleep roughly the right amount of time;
			 * usleep is in microseconds, just like AV_TIME_BASE. */
			delay = avutil.AvRescaleQ(frame.GetPts()-lastPts, timeBase, avutil.AV_TIME_BASE_Q())
			if delay > 0 && delay < 1000000 {
				time.Sleep(time.Microsecond * time.Duration(delay))
			}
		}
		lastPts = frame.GetPts()
	}

	/* Trivial ASCII grayscale display. */
	p0 = (*uint8)(frame.GetData()[0])
	fmt.Print("\033c")
	for y = 0; y < frame.GetHeight(); y++ {
		p = p0
		for x = 0; x < frame.GetWidth(); x++ {
			fmt.Printf(string(" .-+#"[*(p)/52]))
			p = (*uint8)(unsafe.Add(unsafe.Pointer(p), 1))
		}
		fmt.Println()
		p0 = (*uint8)(unsafe.Add(unsafe.Pointer(p0), frame.GetLineSize()[0]))
	}
	os.Stdout.Sync()
}

func main() {
	var ret int
	var packet *goavcodec.AVPacket
	var frame *goavutil.AVFrame
	var filtFrame *goavutil.AVFrame

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s file\n", os.Args[0])
		os.Exit(1)
	}

	frame = goavutil.AllocAVFrame()
	filtFrame = goavutil.AllocAVFrame()
	packet = goavcodec.AllocAvPacket()
	if frame == nil || filtFrame == nil || packet == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate frame or packet\n")
		os.Exit(1)
	}

	if ret = openInputFile(os.Args[1]); ret < 0 {
		goto end
	}
	if ret = initFilters(filterDescr); ret < 0 {
		goto end
	}

	/* read all packets */
	for {
		if ret = fmtCtx.ReadFrame(packet); ret < 0 {
			break
		}

		if packet.CAVPacket.GetStreamIndex() == videoStreamIndex {
			ret = decCtx.SendPacket(packet)
			if ret < 0 {
				avutil.AvLog(nil, avutil.AV_LOG_ERROR, "Error while sending a packet to the decoder\n")
				break
			}

			for ret >= 0 {
				ret = decCtx.ReceiveFrame(frame)
				if ret == avutil.AVERROR(int(syscall.EAGAIN)) || ret == avutil.AVERROR_EOF {
					break
				} else if ret < 0 {
					avutil.AvLog(nil, avutil.AV_LOG_ERROR, "Error while receiving a frame from the decoder\n")
					goto end
				}

				frame.SetPts(frame.GetBestEffortTimestamp())

				/* push the decoded frame into the filtergraph */
				if buffersrcCtx.AddFrameFlags(frame, avfilter.AV_BUFFERSRC_FLAG_KEEP_REF) < 0 {
					avutil.AvLog(nil, avutil.AV_LOG_ERROR, "Error while feeding the filtergraph\n")
					break
				}

				/* pull filtered frames from the filtergraph */
				for {
					ret = buffersinkCtx.GetFrame(filtFrame)
					if ret == avutil.AVERROR(int(syscall.EAGAIN)) || ret == avutil.AVERROR_EOF {
						break
					}
					if ret < 0 {
						goto end
					}
					displayFrame(filtFrame, buffersinkCtx.GetInputs()[0].GetTimeBase())
					filtFrame.Unref()
				}
				frame.Unref()
			}
		}
		packet.Unref()
	}
end:
	filterGraph.Free()
	decCtx.Free()
	fmtCtx.CloseInput()
	frame.Free()
	filtFrame.Free()
	packet.Free()

	if ret < 0 && ret != avutil.AVERROR_EOF {
		fmt.Fprintf(os.Stderr, "Error occurred: %s\n", avutil.AvErr2str(ret))
		os.Exit(1)
	}

	os.Exit(0)
}
