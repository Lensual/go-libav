package avformat

/*
#cgo LDFLAGS: -lavformat

#include "libavformat/avformat.h"
*/
import "C"

type AVFormatContext struct {
	CAVFormatContext *C.AVFormatContext
}
