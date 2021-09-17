package avformat

/*
#cgo LDFLAGS: -lavformat

#include "libavformat/avformat.h"
*/
import "C"

func AvformatAllocContext() *AVFormatContext {
	return &AVFormatContext{
		CAVFormatContext: C.avformat_alloc_context(),
	}
}

func AvformatFreeContext(avf *AVFormatContext) {
	C.avformat_free_context(avf.CAVFormatContext)
	avf = nil
}
