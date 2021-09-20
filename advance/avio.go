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
	"reflect"
	"unsafe"

	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
)

type AVIOCallback func(cbuf unsafe.Pointer, buf_size int) int
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

	//初始化AVIO
	// C.read_packet(nil, nil, 0)
	// unsafe.Pointer(C._Cfunc_read_packet)
	cavio := avformat.AvioAllocContext(cbuf, bufSize, wf, unsafe.Pointer(uintptr(id)), //hack 这里把指针直接当int用
		avformat.CFunc(C.go_read_packet), avformat.CFunc(C.go_write_packet), avformat.CFunc(C.go_seek)) //TODO
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

func (avio *AVIOContext) GetCBuffer() unsafe.Pointer {
	return unsafe.Pointer((*C.AVIOContext)((unsafe.Pointer)(avio.CAVIOContext)).buffer)
}

func (avio *AVIOContext) GetBuffer() []byte {
	//映射C的内存
	var buf []byte
	h := (*reflect.SliceHeader)((unsafe.Pointer(&buf)))
	h.Cap = avio.BufferSize
	h.Len = avio.BufferSize
	h.Data = uintptr(avio.GetCBuffer())

	return buf
}

func (avio *AVIOContext) Free() {
	//TODO
	/* note: the internal buffer could have changed, and be != avio_ctx_buffer */
	bufptr := avio.GetCBuffer()
	avutil.AvFreep(unsafe.Pointer(&bufptr))
	avformat.AvioContextFree(avio.CAVIOContext)
}

//export go_read_packet
func go_read_packet(opaque unsafe.Pointer, cbuf *C.uint8_t, buf_size C.int) C.int {
	id := uintptr(opaque) //hack 把指针当int用
	avio := avioBindingC[id]
	if avio != nil {
		return C.int(avio.readFunc(unsafe.Pointer(cbuf), int(buf_size)))
	}
	return 0
}

//export go_write_packet
func go_write_packet(opaque unsafe.Pointer, cbuf *C.uint8_t, buf_size C.int) C.int {
	id := uintptr(opaque) //hack 把指针当int用
	avio := avioBindingC[id]
	if avio != nil {
		return C.int(avio.writeFunc(unsafe.Pointer(cbuf), int(buf_size)))
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
