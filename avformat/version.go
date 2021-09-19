package avformat

/*
#cgo LDFLAGS: -lavformat

#include "libavformat/avformat.h"
*/
import "C"

const (
	LIBAVFORMAT_IDENT       = C.LIBAVFORMAT_IDENT
	LIBAVFORMAT_VERSION_INT = C.LIBAVFORMAT_VERSION_INT
)
