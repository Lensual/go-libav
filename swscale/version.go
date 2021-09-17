package swscale

/*
#cgo LDFLAGS: -lswscale

#include "libswscale/swscale.h"

char* GET_LIBSWSCALE_IDENT()
{
    return LIBSWSCALE_IDENT;
}
*/
import "C"

func Version() string {
	return C.GoString(C.GET_LIBSWSCALE_IDENT())
}

func VersionInt() int {
	return int(C.LIBSWSCALE_VERSION_INT)
}
