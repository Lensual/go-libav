package goavformat

/*
#cgo pkg-config: libavformat

#include "libavformat/avformat.h"

#include <stdint.h>

extern int go_read_packet(void *opaque, uint8_t *cbuf, int buf_size);
extern int go_write_packet(void *opaque, uint8_t *cbuf, int buf_size);
extern int64_t go_seek(void *opaque, int64_t offset, int whence);
*/
import "C"

import (
	"unsafe"

	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
)

type AVIOContext struct {
	CAVIOContext *avformat.CAVIOContext
}

func (avioCtx *AVIOContext) GetBuffer() []byte {
	return unsafe.Slice((*byte)(avioCtx.CAVIOContext.GetBuffer()), avioCtx.CAVIOContext.GetBufferSize())
}

func (avioCtx *AVIOContext) Free() {
	if avioCtx.CAVIOContext != nil {
		/* note: the internal buffer could have changed, and be != avio_ctx_buffer */
		bufptr := avioCtx.CAVIOContext.GetBuffer()
		avutil.AvFreep(unsafe.Pointer(&bufptr))
		avformat.AvioContextFree(&avioCtx.CAVIOContext)
	}
}

func GetAvailableProtocols(output bool) []string {
	p := unsafe.Pointer(nil)
	arr := []string{}
	outputInt := 0
	if output {
		outputInt = 1
	}
	for {
		name := avformat.AvioEnumProtocols(&p, outputInt)
		if len(name) <= 0 {
			break
		}
		arr = append(arr, name)
	}
	return arr
}
