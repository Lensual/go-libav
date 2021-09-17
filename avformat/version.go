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

func init() {
	minVersion := 58<<16 | 12<<8 | 100
	if C.LIBAVFORMAT_VERSION_INT < minVersion {
		panic("最小支持版本 Lavf58.12.100")
	}
}

func Version() string {
	return C.GoString(C.GET_LIBAVFORMAT_IDENT())
}
