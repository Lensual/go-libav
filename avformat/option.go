package avformat

/*
#cgo LDFLAGS: -lavformat

#include "libavformat/avformat.h"
*/
import "C"

type CAVFormatContext C.AVFormatContext

/*
Allocate an AVFormatContext.

avformat_free_context() can be used to free the context and everything
allocated by the framework within it.
*/
func AvformatAllocContext() *CAVFormatContext {
	return (*CAVFormatContext)(C.avformat_alloc_context())
}

/*
Free an AVFormatContext and all its streams.

@param s context to free
*/
func AvformatFreeContext(s *CAVFormatContext) {
	C.avformat_free_context((*C.AVFormatContext)(s))
}
