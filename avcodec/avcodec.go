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

func init() {
	minVersion := 58<<16 | 18<<8 | 100
	if C.LIBAVCODEC_VERSION_INT < minVersion {
		panic("最小支持版本 Lavc58.18.100")
	}
}

func Version() string {
	return C.GoString(C.GET_LIBAVCODEC_IDENT())
}
