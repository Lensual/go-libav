package avcodec

/*
#cgo LDFLAGS: -lavcodec

#include "libavcodec/avcodec.h"
*/
import "C"

const (
	LIBAVCODEC_IDENT       = C.LIBAVCODEC_IDENT
	LIBAVCODEC_VERSION_INT = C.LIBAVCODEC_VERSION_INT
)
