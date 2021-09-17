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

func Version() string {
	return C.GoString(C.GET_LIBAVFILTER_IDENT())
}

func VersionInt() int {
	return int(C.LIBAVFILTER_VERSION_INT)
}
