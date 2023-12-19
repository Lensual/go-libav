package goavcodec

import "C"

import (
	"unsafe"

	"github.com/Lensual/go-libav/avcodec"
)

type AVCodec struct {
	CAVCodec *avcodec.CAVCodec
}

//region member

func (codec *AVCodec) GetSupportedSamplerates() []int {
	cArr := codec.CAVCodec.GetSupportedSamplerates()
	if cArr == nil {
		return nil
	}
	p := unsafe.Pointer(cArr)
	pSize := int(unsafe.Sizeof(C.int(0)))
	arr := []int{}
	for {
		p = unsafe.Add(p, pSize)
		val := int(*(*C.int)(p))
		if val == 0 {
			break
		}
		arr = append(arr, val)
	}
	return arr
}

//endregion member

func (codec *AVCodec) CreateContext() *AVCodecContext {
	return NewAVCodecContext(codec)
}

func (codec *AVCodec) CreateParserContext() *AVCodecParserContext {
	return NewAVCodecParserContext(int(codec.CAVCodec.GetId()))
}

func GetAvailableCodecs() []*AVCodec {
	p := unsafe.Pointer(nil)
	arr := []*AVCodec{}
	for {
		codec := avcodec.AvCodecIterate(&p)
		if codec == nil {
			break
		}
		arr = append(arr, &AVCodec{
			CAVCodec: codec,
		})
	}
	return arr
}

func FindDecoder(id avcodec.CAVCodecID) *AVCodec {
	c := avcodec.AvcodecFindDecoder(id)
	if c == nil {
		return nil
	}
	return &AVCodec{
		CAVCodec: c,
	}
}

func FindDecoderByName(name string) *AVCodec {
	c := avcodec.AvcodecFindDecoderByName(name)
	if c == nil {
		return nil
	}
	return &AVCodec{
		CAVCodec: c,
	}
}

func FindEncoder(id avcodec.CAVCodecID) *AVCodec {
	c := avcodec.AvcodecFindEncoder(id)
	if c == nil {
		return nil
	}
	return &AVCodec{
		CAVCodec: c,
	}
}

func FindEncoderByName(name string) *AVCodec {
	c := avcodec.AvcodecFindEncoderByName(name)
	if c == nil {
		return nil
	}
	return &AVCodec{
		CAVCodec: c,
	}
}
