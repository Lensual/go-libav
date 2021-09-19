package avutil

/*
#cgo LDFLAGS: -lavutil

#include "libavutil/avutil.h"
*/
import "C"

const (
	LIBAVUTIL_IDENT       = C.LIBAVUTIL_IDENT
	LIBAVUTIL_VERSION_INT = C.LIBAVUTIL_VERSION_INT
)
