package swscale

/*
#cgo LDFLAGS: -lswscale

#include "libswscale/swscale.h"
*/
import "C"

const (
	LIBSWSCALE_IDENT       = C.LIBSWSCALE_IDENT
	LIBSWSCALE_VERSION_INT = C.LIBSWSCALE_VERSION_INT
)
