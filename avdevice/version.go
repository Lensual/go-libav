package avdevice

/*
#cgo LDFLAGS: -lavdevice

#include "libavdevice/avdevice.h"
*/
import "C"

const (
	LIBAVDEVICE_IDENT       = C.LIBAVDEVICE_IDENT
	LIBAVDEVICE_VERSION_INT = C.LIBAVDEVICE_VERSION_INT
)
