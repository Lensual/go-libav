package avutil

/*
#cgo pkg-config: libavutil

#include "libavutil/version.h"
*/
import "C"

const (
	LIBAVUTIL_IDENT       = C.LIBAVUTIL_IDENT
	LIBAVUTIL_VERSION_INT = C.LIBAVUTIL_VERSION_INT
)
