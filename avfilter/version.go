package avfilter

/*
#cgo LDFLAGS: -lavfilter

#include "libavfilter/avfilter.h"
*/
import "C"

const (
	LIBAVFILTER_IDENT       = C.LIBAVFILTER_IDENT
	LIBAVFILTER_VERSION_INT = C.LIBAVFILTER_VERSION_INT
)
