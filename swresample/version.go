package swresample

/*
#cgo LDFLAGS: -lswresample

#include "libswresample/swresample.h"

char* GET_LIBSWRESAMPLE_IDENT()
{
    return LIBSWRESAMPLE_IDENT;
}
*/
import "C"

func Version() string {
	return C.GoString(C.GET_LIBSWRESAMPLE_IDENT())
}

func VersionInt() int {
	return int(C.LIBSWRESAMPLE_VERSION_INT)
}
