package avutil

/*
#cgo LDFLAGS: -lavutil

#include "libavutil/avutil.h"
*/
import "C"
import "unsafe"

func AvMalloc(bufSize uint64) unsafe.Pointer {
	return C.av_malloc(C.ulong(bufSize))
}
