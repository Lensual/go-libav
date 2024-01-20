package goavcodec

import (
	"context"
	"syscall"
	"unsafe"

	"github.com/Lensual/go-libav/advance/goavutil"
	"github.com/Lensual/go-libav/avcodec"
	"github.com/Lensual/go-libav/avutil"
)

type AVCodecContext struct {
	CAVCodecContext *avcodec.CAVCodecContext
}

//#region members

func (avctx *AVCodecContext) GetWidth() int {
	return avctx.CAVCodecContext.GetWidth()
}

func (avctx *AVCodecContext) SetWidth(width int) {
	avctx.CAVCodecContext.SetWidth(width)
}

func (avctx *AVCodecContext) GetHeight() int {
	return avctx.CAVCodecContext.GetHeight()
}

func (avctx *AVCodecContext) SetHeight(height int) {
	avctx.CAVCodecContext.SetHeight(height)
}

func (avctx *AVCodecContext) GetPixFmt() avutil.CAVPixelFormat {
	return avctx.CAVCodecContext.GetPixFmt()
}

func (avctx *AVCodecContext) SetPixFmt(pixFmt avutil.CAVPixelFormat) {
	avctx.CAVCodecContext.SetPixFmt(pixFmt)
}

func (avctx *AVCodecContext) GetSampleAspectRatio() avutil.CAVRational {
	return avctx.CAVCodecContext.GetSampleAspectRatio()
}

func (avctx *AVCodecContext) SetSampleAspectRatio(sampleAspectRatio avutil.CAVRational) {
	avctx.CAVCodecContext.SetSampleAspectRatio(sampleAspectRatio)
}

func (avctx *AVCodecContext) GetChLayout() *goavutil.AVChannelLayout {
	return &goavutil.AVChannelLayout{
		CAVChannelLayout: avctx.CAVCodecContext.GetChLayoutPtr(),
	}
}

//#endregion members

func NewContext(codec *AVCodec) *AVCodecContext {
	c := avcodec.AvcodecAllocContext3(codec.CAVCodec)
	if c == nil {
		return nil
	}
	return &AVCodecContext{
		CAVCodecContext: c,
	}
}

func (avctx *AVCodecContext) Free() {
	avcodec.AvcodecFreeContext(&avctx.CAVCodecContext)
}

func (avctx *AVCodecContext) SendPacket(pkt *AVPacket) int {
	var cPkt *avcodec.CAVPacket
	if pkt != nil {
		cPkt = pkt.CAVPacket
	}
	return avcodec.AvcodecSendPacket(avctx.CAVCodecContext, cPkt)
}

func (avctx *AVCodecContext) ReceiveFrame(frame *goavutil.AVFrame) int {
	return avcodec.AvcodecReceiveFrame(avctx.CAVCodecContext, frame.CAVFrame)
}

func (avctx *AVCodecContext) SendFrame(frame *goavutil.AVFrame) int {
	var cFrame *avutil.CAVFrame
	if frame != nil {
		cFrame = frame.CAVFrame
	}
	return avcodec.AvcodecSendFrame(avctx.CAVCodecContext, cFrame)
}

func (avctx *AVCodecContext) ReceivePacket(pkt *AVPacket) int {
	return avcodec.AvcodecReceivePacket(avctx.CAVCodecContext, pkt.CAVPacket)
}

// Decode AVPacket.
// Return a slice of decoded AVFrame, May be nil or empty.
func (avctx *AVCodecContext) Decode(pkt *AVPacket) ([]*goavutil.AVFrame, int) {
	code := avctx.SendPacket(pkt)
	if code < 0 {
		return nil, code
	}

	frames := make([]*goavutil.AVFrame, 0)

	for {
		frame := goavutil.AllocAVFrame()
		if frame == nil {
			code = int(syscall.ENOMEM)
			break
		}

		code = avctx.ReceiveFrame(frame)
		if code < 0 {
			frame.Free()
			break
		}

		frames = append(frames, frame)
	}

	return frames, code
}

// Decode AVPacket from channel.
//
// @param ctx The context.Context to cancel this goroutine.
// @param pktChan The AVPacket channel.
//
// @return context.Context The context.Context to get the cause of this gorouine.
// @return <-chan *goavutil.AVFrame The channel to read decoded AVFrame.
func (avctx *AVCodecContext) DecodeChan(ctx context.Context, pktChan <-chan *AVPacket) (context.Context, <-chan *goavutil.AVFrame) {
	ctx, cancel := context.WithCancelCause(ctx)
	frameChan := make(chan *goavutil.AVFrame)

	go func() {
		defer close(frameChan)
	loop:
		for {
			select {
			case <-ctx.Done():
				return
			case pkt, ok := <-pktChan:
				if !ok {
					break loop
				}
				frames, code := avctx.Decode(pkt)
				pkt.Unref()
				for _, frame := range frames {
					select {
					case <-ctx.Done():
						for _, frame := range frames {
							frame.Unref()
						}
						return
					case frameChan <- frame:
					}
				}
				if code == avutil.AVERROR(int(syscall.EAGAIN)) || code == avutil.AVERROR_EOF {
					continue
				}
				if code < 0 {
					cancel(goavutil.AvErr(code))
					return
				}
			}
		}

		//flush codec
		frames, code := avctx.Decode(nil)
		for _, frame := range frames {
			select {
			case <-ctx.Done():
				for _, frame := range frames {
					frame.Unref()
				}
				return
			case frameChan <- frame:
			}
		}
		if code == avutil.AVERROR(int(syscall.EAGAIN)) || code == avutil.AVERROR_EOF {
			return
		}
		if code < 0 {
			cancel(goavutil.AvErr(code))
			return
		}
	}()

	return ctx, frameChan
}

// ChunkingFrameChan chunking read frame by AVAudioFifo.
//
// Example:
//
//	if avctx.CAVCodecContext.GetCodec().GetCapabilities()&avcodec.AV_CODEC_CAP_VARIABLE_FRAME_SIZE == 0 {
//		ctx, frameChan = avctx.chunkingFrameChan(ctx, frameChan)
//	}
//	avctx.EncodeChan(ctx, frameChan)
func (avctx *AVCodecContext) ChunkingFrameChan(ctx context.Context, frameInChan <-chan *goavutil.AVFrame) (context.Context, <-chan *goavutil.AVFrame) {
	ctx, cancel := context.WithCancelCause(ctx)
	frameOutChan := make(chan *goavutil.AVFrame)

	af := goavutil.NewAVAudioFifo(avctx.CAVCodecContext.GetSampleFmt(), avctx.GetChLayout().GetNbChannels(), 1)
	if af == nil {
		cancel(goavutil.AvErr(avutil.AVERROR(int(syscall.ENOMEM))))
		return ctx, nil
	}

	go func() {
		defer af.Free()
		defer close(frameOutChan)
	loop:
		for {
			select {
			case <-ctx.Done():
				return
			case frame, ok := <-frameInChan:
				// write fifo
				if !ok {
					break loop
				}
				if code := af.Realloc(af.Size() + frame.GetNbSamples()); code < 0 {
					frame.Unref()
					cancel(goavutil.AvErr(code))
					return
				}
				frameData := frame.GetData()
				code := af.Write(unsafe.SliceData(frameData[:]), frame.GetNbSamples())
				frame.Unref()
				if code < 0 {
					cancel(goavutil.AvErr(code))
					return
				}

				// read fifo
				frameSize := avctx.CAVCodecContext.GetFrameSize()
				for af.Size() >= frameSize {
					frame = goavutil.AllocAVFrame()
					frame.SetNbSamples(frameSize)
					avctx.GetChLayout().CopyTo(frame.GetChLayout())
					frame.SetFormat(int(avctx.CAVCodecContext.GetSampleFmt()))
					frame.SetSampleRate(avctx.CAVCodecContext.GetSampleRate())
					code = frame.AllocBuffer(0)
					if code < 0 {
						frame.Unref()
						cancel(goavutil.AvErr(code))
						return
					}
					frameData = frame.GetData()
					code = af.Read(unsafe.SliceData(frameData[:]), frameSize)
					if code < 0 {
						frame.Unref()
						cancel(goavutil.AvErr(code))
						return
					}
					select {
					case <-ctx.Done():
						frame.Unref()
						return
					case frameOutChan <- frame:
					}
				}
			}
		}

		// flush fifo
		if af.Size() >= 0 {
			frame := goavutil.AllocAVFrame()
			frame.SetNbSamples(af.Size())
			avctx.GetChLayout().CopyTo(frame.GetChLayout())
			frame.SetFormat(int(avctx.CAVCodecContext.GetSampleFmt()))
			frame.SetSampleRate(avctx.CAVCodecContext.GetSampleRate())
			code := frame.AllocBuffer(0)
			if code < 0 {
				frame.Unref()
				cancel(goavutil.AvErr(code))
				return
			}
			frameData := frame.GetData()
			code = af.Read(unsafe.SliceData(frameData[:]), af.Size())
			if code < 0 {
				frame.Unref()
				cancel(goavutil.AvErr(code))
				return
			}
			select {
			case <-ctx.Done():
				frame.Unref()
				return
			case frameOutChan <- frame:
			}
		}
	}()

	return ctx, frameOutChan
}

// Encode AVFrame.
// Return a slice of encoded AVPacket, May be nil or empty.
func (avctx *AVCodecContext) Encode(frame *goavutil.AVFrame) ([]*AVPacket, int) {
	code := avctx.SendFrame(frame)
	if code < 0 {
		return nil, code
	}

	pkts := make([]*AVPacket, 0)

	for {
		pkt := AllocAvPacket()
		if pkt == nil {
			code = int(syscall.ENOMEM)
			break
		}
		code = avctx.ReceivePacket(pkt)
		if code < 0 {
			pkt.Free()
			break
		}

		pkts = append(pkts, pkt)
	}

	return pkts, code
}

// Encode AVPacket from channel.
//
// @param ctx The context.Context to cancel this goroutine.
// @param frameChan The AVFrame channel.
//
// @return context.Context The context.Context to get the cause of this gorouine.
// @return <-chan *AVPacket The channel to read encoded AVPacket.
func (avctx *AVCodecContext) EncodeChan(ctx context.Context, frameChan <-chan *goavutil.AVFrame) (context.Context, <-chan *AVPacket) {
	ctx, cancel := context.WithCancelCause(ctx)
	pktChan := make(chan *AVPacket)

	go func() {
		defer close(pktChan)
	loop:
		for {
			select {
			case <-ctx.Done():
				return
			case frame, ok := <-frameChan:
				if !ok {
					break loop
				}
				pkts, code := avctx.Encode(frame)
				frame.Unref()
				for _, pkt := range pkts {
					select {
					case <-ctx.Done():
						for _, pkt := range pkts {
							pkt.Unref()
						}
						return
					case pktChan <- pkt:
					}
				}
				if code == avutil.AVERROR(int(syscall.EAGAIN)) || code == avutil.AVERROR_EOF {
					continue
				}
				if code < 0 {
					cancel(goavutil.AvErr(code))
					return
				}
			}
		}

		//flush codec
		pkts, code := avctx.Encode(nil)
		for _, pkt := range pkts {
			select {
			case <-ctx.Done():
				for _, pkt := range pkts {
					pkt.Unref()
				}
				return
			case pktChan <- pkt:
			}
		}
		if code == avutil.AVERROR(int(syscall.EAGAIN)) || code == avutil.AVERROR_EOF {
			return
		}
		if code < 0 {
			cancel(goavutil.AvErr(code))
			return
		}
	}()

	return ctx, pktChan
}

func (avctx *AVCodecContext) ParametersFrom(par *avcodec.CAVCodecParameters) int {
	return avcodec.AvcodecParametersToContext(avctx.CAVCodecContext, par)
}

func (avctx *AVCodecContext) ParametersTo(par *avcodec.CAVCodecParameters) int {
	return avcodec.AvcodecParametersFromContext(par, avctx.CAVCodecContext)
}

func (avctx *AVCodecContext) Open(options **avutil.CAVDictionary) int {
	return avcodec.AvcodecOpen2(avctx.CAVCodecContext, avctx.CAVCodecContext.GetCodec(), options)
}

func (avctx *AVCodecContext) Close() int {
	return avcodec.AvcodecClose(avctx.CAVCodecContext)
}

func (avctx *AVCodecContext) IsOpen() bool {
	return avcodec.AvcodecIsOpen(avctx.CAVCodecContext) > 0
}
