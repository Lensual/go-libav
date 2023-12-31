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
	"sync"
	"unsafe"

	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/ctypes"
)

type AVIOReadCallback func(buf []byte, len int) int
type AVIOWriteCallback func(buf []byte, len int) int
type AVIOSeekCallback func(offset int64, whence int) int64

// 用于c的callback绑定avio对象
var avioBindingC []*GoAVIOContext
var avioBindingC_lock sync.RWMutex

func init() {
	avioBindingC = make([]*GoAVIOContext, unsafe.Sizeof(uint8(1)))
}

// HACK
// 使用导出的go方法实现AVIO
// 由于go对象指针不可传递给C，迫不得已生成唯一id存入C.AVIOContext.opaque来实现Go对象成员方法的访问
// 除此之外还可以尝试使用fd协议配合os.pipe实现
type GoAVIOContext struct {
	*AVIOContext
	id        int
	readFunc  AVIOReadCallback
	writeFunc AVIOWriteCallback
	seekFunc  AVIOSeekCallback
}

func NewGoAvioContext(bufSize int, readFunc AVIOReadCallback, writeFunc AVIOWriteCallback, seekFunc AVIOSeekCallback) *GoAVIOContext {
	avioCtx := &GoAVIOContext{
		AVIOContext: &AVIOContext{},
		id:          -1,
		readFunc:    readFunc,
		writeFunc:   writeFunc,
		seekFunc:    seekFunc,
	}

	//分配可用id
	avioBindingC_lock.Lock()
	{
		for i, v := range avioBindingC {
			if v == nil {
				avioCtx.id = i
			}
		}
		if avioCtx.id < 0 {
			avioCtx.Free()
			return nil
		}

		//AVIO绑定
		avioBindingC[avioCtx.id] = avioCtx
	}
	avioBindingC_lock.Unlock()

	//初始化C缓冲区
	cbuf := avutil.AvMalloc(ctypes.SizeT(bufSize))
	if cbuf == nil {
		avioCtx.Free()
		return nil
	}

	//判断write flag
	var wf int
	if writeFunc != nil {
		wf = 1
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

	//初始化AVIO
	cavioCtx := avformat.AvioAllocContext(cbuf, bufSize, wf, unsafe.Pointer(uintptr(avioCtx.id)), //hack 这里把指针直接当int用
		readCallback, writeCallback, seekCallback) //TODO
	if cavioCtx == nil {
		avutil.AvFreep(cbuf)
		avioCtx.Free()
		return nil
	}
	avioCtx.AVIOContext.CAVIOContext = cavioCtx

	return avioCtx
}

func (avioCtx *GoAVIOContext) GetBuffer() []byte {
	return avioCtx.AVIOContext.GetBuffer()
}

func (avioCtx *GoAVIOContext) Free() {
	avioCtx.AVIOContext.Free()

	//释放AVIO绑定
	if avioCtx.id >= 0 {
		avioBindingC_lock.Lock()
		{
			avioBindingC[avioCtx.id] = nil
		}
		avioBindingC_lock.Unlock()
		avioCtx.id = -1
	}

	avioCtx.AVIOContext = nil
	avioCtx.readFunc = nil
	avioCtx.writeFunc = nil
	avioCtx.seekFunc = nil

	avioCtx = nil
}

//export go_read_packet
func go_read_packet(opaque unsafe.Pointer, cbuf *C.uint8_t, buf_size C.int) C.int {
	id := uintptr(opaque) //hack 把指针当int用
	avioCtx := avioBindingC[id]
	if avioCtx != nil {
		buf := unsafe.Slice((*byte)(unsafe.Pointer(cbuf)), int(buf_size))
		return C.int(avioCtx.readFunc(buf, int(buf_size)))
	}
	return 0
}

//export go_write_packet
func go_write_packet(opaque unsafe.Pointer, cbuf *C.uint8_t, buf_size C.int) C.int {
	id := uintptr(opaque) //hack 把指针当int用
	avioCtx := avioBindingC[id]
	if avioCtx != nil {
		buf := unsafe.Slice((*byte)(unsafe.Pointer(cbuf)), int(buf_size))
		return C.int(avioCtx.writeFunc(buf, int(buf_size)))
	}
	return 0
}

//export go_seek
func go_seek(opaque unsafe.Pointer, offset C.int64_t, whence C.int) C.int64_t {
	id := uintptr(opaque) //hack 把指针当int用
	avioCtx := avioBindingC[id]
	if avioCtx != nil {
		return C.int64_t(avioCtx.seekFunc(int64(offset), int(whence)))
	}
	return 0
}
