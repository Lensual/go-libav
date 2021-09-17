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

func Version() string {
	return C.GoString(C.GET_LIBAVDEVICE_IDENT())
}

func VersionInt() int {
	return int(C.LIBAVDEVICE_VERSION_INT)
}
