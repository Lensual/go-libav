package swresample

/*
#cgo LDFLAGS: -lswresample

#include "libswresample/swresample.h"
*/
import "C"

const (
	LIBSWRESAMPLE_IDENT       = C.LIBSWRESAMPLE_IDENT
	LIBSWRESAMPLE_VERSION_INT = C.LIBSWRESAMPLE_VERSION_INT
)
