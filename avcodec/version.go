package avcodec

/*
#cgo LDFLAGS: -lavcodec

#include "libavcodec/avcodec.h"

char* GET_LIBAVCODEC_IDENT()
{
    return LIBAVCODEC_IDENT;
}
*/
import "C"

func Version() string {
	return C.GoString(C.GET_LIBAVCODEC_IDENT())
}

func VersionInt() int {
	return int(C.LIBAVCODEC_VERSION_INT)
}
