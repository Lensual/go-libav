package avfilter

/*
#cgo LDFLAGS: -lavfilter

#include "libavfilter/avfilter.h"

char* GET_LIBAVFILTER_IDENT()
{
    return LIBAVFILTER_IDENT;
}
*/
import "C"

func init() {
	minVersion := 7<<16 | 16<<8 | 100
	if C.LIBAVFILTER_VERSION_INT < minVersion {
		panic("最小支持版本 Lavfi7.16.100")
	}
}

func Version() string {
	return C.GoString(C.GET_LIBAVFILTER_IDENT())
}
