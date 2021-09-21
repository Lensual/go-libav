package advance

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/Lensual/go-libav/avcodec"
	"github.com/Lensual/go-libav/avutil"
)

type AVCodecContext struct {
	CAVCodecContext *avcodec.CAVCodecContext
}

func NewAVCodecContext(codec *AVCodec) *AVCodecContext {
	cavctx := avcodec.AvcodecAllocContext(codec.CAVCodec)
	avcodec.AvcodecOpen(cavctx, codec.CAVCodec, nil) //TODO options
	return &AVCodecContext{
		CAVCodecContext: cavctx,
	}
}

func (avctx *AVCodecContext) Free() {
	avcodec.AvPacketFree((**avcodec.CAVPacket)((unsafe.Pointer)(&avctx.CAVCodecContext)))
	avctx.CAVCodecContext = nil
}

//解码
func (avctx *AVCodecContext) Decode(pkt *AvPacket) ([]AvFrame, error) {
	code := avcodec.AvcodecSendPacket(avctx.CAVCodecContext, pkt.CAvPacket)
	if code < 0 {
		return nil, errors.New(avutil.Err2str(code))
	}

	frames := make([]AvFrame, 0)

	var err error

	for {
		frame := NewAVFrame(0)
		if frame == nil {
			err = syscall.ENOMEM
			break
		}

		code := avcodec.AvcodecReceiveFrame(avctx.CAVCodecContext, frame.CAvFrame)
		if code < 0 {
			frame.Free()
			break
		}

		frames = append(frames, *frame)
	}

	return frames, err
}

//编码
func (avctx *AVCodecContext) Encode() {
}
