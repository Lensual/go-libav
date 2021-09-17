package avformat

/*
#cgo LDFLAGS: -lavformat

#include "libavformat/avformat.h"

char* GET_LIBAVFORMAT_IDENT()
{
    return LIBAVFORMAT_IDENT;
}
*/
import "C"

func Version() string {
	return C.GoString(C.GET_LIBAVFORMAT_IDENT())
}

func VersionInt() int {
	return int(C.LIBAVFORMAT_VERSION_INT)
}
