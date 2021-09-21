package advance

/*
#cgo LDFLAGS: -lavformat

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

type AVIOCallback func(buf []byte, len int) int
type AVIOSeekCallback func(offset int64, whence int) int64

//用于c的callback绑定avio对象
var avioBindingC []*AVIOContext

func init() {
	avioBindingC = make([]*AVIOContext, unsafe.Sizeof(uint8(1)))
}

type AVIOContext struct {
	id           uint8
	CAVIOContext *avformat.CAVIOContext
	BufferSize   int
	readFunc     AVIOCallback
	writeFunc    AVIOCallback
	seekFunc     AVIOSeekCallback
}

func NewAvioContext(bufSize int, readFunc AVIOCallback, writeFunc AVIOCallback, seekFunc AVIOSeekCallback) *AVIOContext {
	//TODO lock
	//分配可用id
	var id int = -1
	for i, v := range avioBindingC {
		if v == nil {
			id = i
		}
	}
	if id == -1 {
		return nil
	}

	//初始化C缓冲区
	cbuf := avutil.AvMalloc(uint64(bufSize))
	if cbuf == nil {
		return nil
	}

	//判断write flag
	var wf int
	if writeFunc != nil {
		wf = 1
	}

	var readCallback avformat.CFunc
	if readFunc != nil {
		readCallback = avformat.CFunc(C.go_read_packet)
	}
	var writeCallback avformat.CFunc
	if writeFunc != nil {
		writeCallback = avformat.CFunc(C.go_write_packet)
	}
	var seekCallback avformat.CFunc
	if seekFunc != nil {
		seekCallback = avformat.CFunc(C.go_seek)
	}

	//初始化AVIO
	cavio := avformat.AvioAllocContext(cbuf, bufSize, wf, unsafe.Pointer(uintptr(id)), //hack 这里把指针直接当int用
		readCallback, writeCallback, seekCallback) //TODO
	if cavio == nil {
		avutil.AvFreep(cbuf)
		return nil
	}

	avio := &AVIOContext{
		CAVIOContext: cavio,
		BufferSize:   bufSize,
		readFunc:     readFunc,
		writeFunc:    writeFunc,
		seekFunc:     seekFunc,
	}

	//AVIO绑定
	avioBindingC[avio.id] = avio

	return avio
}

func (avio *AVIOContext) GetBuffer() []byte {
	return UnsafePtr2ByteSlice(avio.CAVIOContext.GetCBuffer(), avio.BufferSize)
}

func (avio *AVIOContext) Free() {
	/* note: the internal buffer could have changed, and be != avio_ctx_buffer */
	bufptr := avio.CAVIOContext.GetCBuffer()
	avutil.AvFreep(unsafe.Pointer(&bufptr))
	avformat.AvioContextFree(avio.CAVIOContext)

	//TODO 释放ID

	avio.readFunc = nil
	avio.writeFunc = nil
	avio.seekFunc = nil
}

//export go_read_packet
func go_read_packet(opaque unsafe.Pointer, cbuf *C.uint8_t, buf_size C.int) C.int {
	id := uintptr(opaque) //hack 把指针当int用
	avio := avioBindingC[id]
	if avio != nil {
		buf := UnsafePtr2ByteSlice(unsafe.Pointer(cbuf), int(buf_size))
		return C.int(avio.readFunc(buf, int(buf_size)))
	}
	return 0
}

//export go_write_packet
func go_write_packet(opaque unsafe.Pointer, cbuf *C.uint8_t, buf_size C.int) C.int {
	id := uintptr(opaque) //hack 把指针当int用
	avio := avioBindingC[id]
	if avio != nil {
		buf := UnsafePtr2ByteSlice(unsafe.Pointer(cbuf), int(buf_size))
		return C.int(avio.writeFunc(buf, int(buf_size)))
	}
	return 0
}

//export go_seek
func go_seek(opaque unsafe.Pointer, offset C.int64_t, whence C.int) C.int64_t {
	id := uintptr(opaque) //hack 把指针当int用
	avio := avioBindingC[id]
	if avio != nil {
		return C.int64_t(avio.seekFunc(int64(offset), int(whence)))
	}
	return 0
}
