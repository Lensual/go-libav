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
	"sync"
	"unsafe"

	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
)

type AVIOCallback func(buf []byte, len int) int
type AVIOSeekCallback func(offset int64, whence int) int64

//用于c的callback绑定avio对象
var avioBindingC []*AVIOContext
var avioBindingC_lock sync.RWMutex

func init() {
	avioBindingC = make([]*AVIOContext, unsafe.Sizeof(uint8(1)))
}

type AVIOContext struct {
	id           int
	CAVIOContext *avformat.CAVIOContext
	BufferSize   int
	readFunc     AVIOCallback
	writeFunc    AVIOCallback
	seekFunc     AVIOSeekCallback
}

func NewAvioContext(bufSize int, readFunc AVIOCallback, writeFunc AVIOCallback, seekFunc AVIOSeekCallback) *AVIOContext {

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
	cavioCtx := avformat.AvioAllocContext(cbuf, bufSize, wf, unsafe.Pointer(uintptr(id)), //hack 这里把指针直接当int用
		readCallback, writeCallback, seekCallback) //TODO
	if cavioCtx == nil {
		avutil.AvFreep(cbuf)
		return nil
	}

	avioCtx := &AVIOContext{
		id:           -1,
		CAVIOContext: cavioCtx,
		BufferSize:   bufSize,
		readFunc:     readFunc,
		writeFunc:    writeFunc,
		seekFunc:     seekFunc,
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
			//clean
			avioCtx.Free()
			return nil
		}

		//AVIO绑定
		avioBindingC[avioCtx.id] = avioCtx
	}
	avioBindingC_lock.Unlock()

	return avioCtx
}

func (avioCtx *AVIOContext) GetBuffer() []byte {
	return UnsafePtr2ByteSlice(avioCtx.CAVIOContext.GetCBuffer(), avioCtx.BufferSize)
}

func (avioCtx *AVIOContext) Free() {
	/* note: the internal buffer could have changed, and be != avio_ctx_buffer */
	bufptr := avioCtx.CAVIOContext.GetCBuffer()
	avutil.AvFreep(unsafe.Pointer(&bufptr))
	avformat.AvioContextFree(avioCtx.CAVIOContext)

	//释放AVIO绑定
	if avioCtx.id >= 0 {
		avioBindingC_lock.Lock()
		{
			avioBindingC[avioCtx.id] = nil
		}
		avioBindingC_lock.Unlock()
		avioCtx.id = -1
	}

	avioCtx.readFunc = nil
	avioCtx.writeFunc = nil
	avioCtx.seekFunc = nil
}

//export go_read_packet
func go_read_packet(opaque unsafe.Pointer, cbuf *C.uint8_t, buf_size C.int) C.int {
	id := uintptr(opaque) //hack 把指针当int用
	avioCtx := avioBindingC[id]
	if avioCtx != nil {
		buf := UnsafePtr2ByteSlice(unsafe.Pointer(cbuf), int(buf_size))
		return C.int(avioCtx.readFunc(buf, int(buf_size)))
	}
	return 0
}

//export go_write_packet
func go_write_packet(opaque unsafe.Pointer, cbuf *C.uint8_t, buf_size C.int) C.int {
	id := uintptr(opaque) //hack 把指针当int用
	avioCtx := avioBindingC[id]
	if avioCtx != nil {
		buf := UnsafePtr2ByteSlice(unsafe.Pointer(cbuf), int(buf_size))
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
