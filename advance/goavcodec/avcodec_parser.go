package goavcodec

import (
	"unsafe"

	"github.com/Lensual/go-libav/avcodec"
)

type AVCodecParser struct {
	CAVCodecParser *avcodec.CAVCodecParser
}

// TODO
func GetAvailableParsers() []*AVCodecParser {
	p := unsafe.Pointer(nil)
	arr := []*AVCodecParser{}
	for {
		parser := avcodec.AvParserIterate(&p)
		if parser == nil {
			break
		}
		arr = append(arr, &AVCodecParser{
			CAVCodecParser: parser,
		})
	}
	return arr
}
