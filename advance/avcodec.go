package advance

import (
	"github.com/Lensual/go-libav/avcodec"
)

type AVCodec struct {
	CAVCodec *avcodec.CAVCodec
}

func (codec *AVCodec) CreateContext() *AVCodecContext {
	return NewAVCodecContext(codec)
}

func (codec *AVCodec) CreateParserContext() *AVCodecParserContext {
	return NewAVCodecParserContext(codec.CAVCodec.ID())
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
