package avutil

/*
#cgo LDFLAGS: -lavutil

#include "libavutil/avutil.h"

char* GET_LIBAVUTIL_IDENT()
{
    return LIBAVUTIL_IDENT;
}
*/
import "C"

func init() {
	minVersion := 56<<16 | 14<<8 | 100
	if C.LIBAVUTIL_VERSION_INT < minVersion {
		panic("最小支持版本 Lavu56.14.100")
	}
}

func Version() string {
	return C.GoString(C.GET_LIBAVUTIL_IDENT())
}
