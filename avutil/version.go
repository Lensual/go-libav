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

func Version() string {
	return C.GoString(C.GET_LIBAVUTIL_IDENT())
}

func VersionInt() int {
	return int(C.LIBAVUTIL_VERSION_INT)
}
