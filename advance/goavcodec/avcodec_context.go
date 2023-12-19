package goavcodec

import (
	"syscall"

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

//#endregion members

func NewAVCodecContext(codec *AVCodec) *AVCodecContext {
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
	return avcodec.AvcodecSendPacket(avctx.CAVCodecContext, pkt.CAVPacket)
}

func (avctx *AVCodecContext) ReceiveFrame(frame *goavutil.AVFrame) int {
	return avcodec.AvcodecReceiveFrame(avctx.CAVCodecContext, frame.CAVFrame)
}

func (avctx *AVCodecContext) SendFrame(frame *goavutil.AVFrame) int {
	return avcodec.AvcodecSendFrame(avctx.CAVCodecContext, frame.CAVFrame)
}

func (avctx *AVCodecContext) ReceivePacket(pkt *AVPacket) int {
	return avcodec.AvcodecReceivePacket(avctx.CAVCodecContext, pkt.CAVPacket)
}

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

func (avctx *AVCodecContext) Encode(frame *goavutil.AVFrame) ([]AVPacket, int) {
	code := avctx.SendFrame(frame)
	if code < 0 {
		return nil, code
	}

	pkts := make([]AVPacket, 0)

	for {
		var pkt *AVPacket
		pkt, code = NewAvPacket(0)
		if code != 0 {
			break
		}

		code = avctx.ReceivePacket(pkt)
		if code < 0 {
			pkt.Free()
			break
		}

		pkts = append(pkts, *pkt)
	}

	return pkts, code
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
