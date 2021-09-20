package advance

import (
	"unsafe"

	"github.com/Lensual/go-libav/avcodec"
)

type AVCodec struct {
	CAVCodec *avcodec.CAVCodec
}

func FindDecoder(id avcodec.AVCodecID) *AVCodec {
	return &AVCodec{
		CAVCodec: avcodec.AvcodecFindDecoder(id),
	}
}

func FindDecoderByName(name string) *AVCodec {
	return &AVCodec{
		CAVCodec: avcodec.AvcodecFindDecoderByName(name),
	}
}

func FindEncoder(id avcodec.AVCodecID) *AVCodec {
	return &AVCodec{
		CAVCodec: avcodec.AvcodecFindEncoder(id),
	}
}

func FindEncoderByName(name string) *AVCodec {
	return &AVCodec{
		CAVCodec: avcodec.AvcodecFindEncoderByName(name),
	}
}

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

func NewAVCodecContextById(id int, is_encoder bool) *AVCodecContext {
	codecid := avcodec.AVCodecID(id)
	var codec *AVCodec
	if is_encoder {
		codec = FindEncoder(codecid)
	} else {
		codec = FindDecoder(codecid)
	}
	return NewAVCodecContext(codec)
}

func NewAVCodecContextByName(name string, is_encoder bool) *AVCodecContext {
	var codec *AVCodec
	if is_encoder {
		codec = FindEncoderByName(name)
	} else {
		codec = FindDecoderByName(name)
	}
	return NewAVCodecContext(codec)
}

func (avctx *AVCodecContext) Free() {
	avcodec.AvPacketFree((**avcodec.CAVPacket)((unsafe.Pointer)(&avctx.CAVCodecContext)))
	avctx.CAVCodecContext = nil
}
