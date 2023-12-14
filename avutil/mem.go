package avutil

/*
#cgo pkg-config: libavutil

#include "libavutil/mem.h"
*/
import "C"
import "unsafe"

func AvMalloc(bufSize uint64) unsafe.Pointer {
	return C.av_malloc(C.ulong(bufSize))
}

func AvFreep(ptr unsafe.Pointer) {
	C.av_freep(ptr)
}
