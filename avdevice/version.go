package avdevice

/*
#cgo LDFLAGS: -lavdevice

#include "libavdevice/avdevice.h"

char* GET_LIBAVDEVICE_IDENT()
{
    return LIBAVDEVICE_IDENT;
}
*/
import "C"

func init() {
	minVersion := 58<<16 | 3<<8 | 100
	if C.LIBAVDEVICE_VERSION_INT < minVersion {
		panic("最小支持版本 Lavd58.3.100")
	}
}

func Version() string {
	return C.GoString(C.GET_LIBAVDEVICE_IDENT())
}
