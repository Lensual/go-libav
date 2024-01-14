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
	"runtime/cgo"
	"unsafe"

	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/ctypes"
)

type AVIOReadCallback func(avioCtx *GoAVIOContext, buf []byte) int
type AVIOWriteCallback func(avioCtx *GoAVIOContext, buf []byte) int
type AVIOSeekCallback func(avioCtx *GoAVIOContext, offset int64, whence int) int64

// GoAVIOContext is Go API for AVIOContext.
// The field opaque of AVIOContext is occupied by cgo.Handle.
type GoAVIOContext struct {
	*AVIOContext
	cgoHandle cgo.Handle
	readFunc  AVIOReadCallback
	writeFunc AVIOWriteCallback
	seekFunc  AVIOSeekCallback
	unsafeBuf bool
}

// NewGoAvioContext
//
// unsafeBuf: True if callback with unsafe buffer, False if callback with copy memory to Go heap.
//
// return: Allocated GoAVIOContext or nil on failure.
func NewGoAvioContext(bufSize int, readFunc AVIOReadCallback, writeFunc AVIOWriteCallback, seekFunc AVIOSeekCallback, unsafeBuf bool) *GoAVIOContext {
	avioCtx := &GoAVIOContext{
		AVIOContext: &AVIOContext{},
		readFunc:    readFunc,
		writeFunc:   writeFunc,
		seekFunc:    seekFunc,
		unsafeBuf:   unsafeBuf,
	}

	// Pass avioCtx pointer to C.
	avioCtx.cgoHandle = cgo.NewHandle(avioCtx)

	// alloc C buffer
	cbuf := avutil.AvMalloc(ctypes.SizeT(bufSize))
	if cbuf == nil {
		avioCtx.Free()
		return nil
	}

	var writeFlag int = 0
	if writeFunc != nil {
		writeFlag = 1
	}

	var readCallback ctypes.CFunc
	if readFunc != nil {
		readCallback = ctypes.CFunc(C.go_read_packet)
	}
	var writeCallback ctypes.CFunc
	if writeFunc != nil {
		writeCallback = ctypes.CFunc(C.go_write_packet)
	}
	var seekCallback ctypes.CFunc
	if seekFunc != nil {
		seekCallback = ctypes.CFunc(C.go_seek)
	}

	// alloc and initialize avio
	cavioCtx := avformat.AvioAllocContext(cbuf, bufSize, writeFlag, unsafe.Pointer(avioCtx.cgoHandle),
		readCallback, writeCallback, seekCallback)
	if cavioCtx == nil {
		avutil.AvFreep(&cbuf)
		return nil
	}
	avioCtx.AVIOContext.CAVIOContext = cavioCtx

	return avioCtx
}

func (avioCtx *GoAVIOContext) IsUnsafeBuffer() bool {
	return avioCtx.unsafeBuf
}

func (avioCtx *GoAVIOContext) Flush() {
	avioCtx.AVIOContext.Flush()
}

// Free all resources allocated by GoAVIOContext.
func (avioCtx *GoAVIOContext) Free() {
	if avioCtx.AVIOContext == nil {
		return
	}

	/* note: the internal buffer could have changed */
	avutil.AvFree(avioCtx.CAVIOContext.GetBuffer())

	avioCtx.AVIOContext.Free()
	avioCtx.AVIOContext = nil

	avioCtx.cgoHandle.Delete()

	avioCtx.readFunc = nil
	avioCtx.writeFunc = nil
	avioCtx.seekFunc = nil
}

//export go_read_packet
func go_read_packet(opaque unsafe.Pointer, cbuf *C.uint8_t, buf_size C.int) C.int {
	avioCtxCgoHandle := cgo.Handle(opaque)
	avioCtx := avioCtxCgoHandle.Value().(*GoAVIOContext)
	var buf []byte = nil
	if avioCtx.unsafeBuf {
		buf = unsafe.Slice((*byte)(unsafe.Pointer(cbuf)), int(buf_size))
	} else {
		buf = C.GoBytes(unsafe.Pointer(cbuf), buf_size)
	}
	return C.int(avioCtx.readFunc(avioCtx, buf))
}

//export go_write_packet
func go_write_packet(opaque unsafe.Pointer, cbuf *C.uint8_t, buf_size C.int) C.int {
	avioCtxCgoHandle := cgo.Handle(opaque)
	avioCtx := avioCtxCgoHandle.Value().(*GoAVIOContext)
	var buf []byte = nil
	if cbuf != nil {
		if avioCtx.unsafeBuf {
			buf = unsafe.Slice((*byte)(unsafe.Pointer(cbuf)), int(buf_size))
		} else {
			buf = C.GoBytes(unsafe.Pointer(cbuf), buf_size)
		}
	}
	return C.int(avioCtx.writeFunc(avioCtx, buf))
}

//export go_seek
func go_seek(opaque unsafe.Pointer, offset C.int64_t, whence C.int) C.int64_t {
	avioCtxCgoHandle := cgo.Handle(opaque)
	avioCtx := avioCtxCgoHandle.Value().(*GoAVIOContext)
	return C.int64_t(avioCtx.seekFunc(avioCtx, int64(offset), int(whence)))
}
