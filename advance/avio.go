package advance

/*
#include <stdint.h>

static int read_packet(void *opaque, uint8_t *buf, int buf_size)
{
    int go_read_packet(void *opaque, uint8_t *buf, int buf_size);
	return go_read_packet(opaque,buf,buf_size);
}

static int write_packet(void *opaque, uint8_t *buf, int buf_size)
{
    int go_write_packet(void *opaque, uint8_t *buf, int buf_size);
	return go_write_packet(opaque,buf,buf_size);
}

static int64_t seek(void *opaque, int64_t offset, int whence)
{
    int64_t go_seek(void *opaque, int64_t offset, int whence);
	return go_seek(opaque,offset,whence);
}
*/
import "C"

import (
	"reflect"
	"unsafe"

	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
)

type AVIOCallback func(buf unsafe.Pointer, buf_size int) int
type AVIOSeekCallback func(offset int64, whence int) int64

//用于c的callback绑定avio对象
var avioBindingC []*AVIOContext

func init() {
	avioBindingC = make([]*AVIOContext, unsafe.Sizeof(uint8(1)))
}

type AVIOContext struct {
	id           uint8
	CAVIOContext *avformat.CAVIOContext
	CBuffer      unsafe.Pointer
	Buffer       []byte
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

	//映射C的内存
	var buf []byte
	h := (*reflect.SliceHeader)((unsafe.Pointer(&buf)))
	h.Cap = bufSize
	h.Len = bufSize
	h.Data = uintptr(cbuf)

	//判断write flag
	var wf int
	if writeFunc != nil {
		wf = 1
	}

	//初始化AVIO
	cavio := avformat.AvioAllocContext(cbuf, bufSize, wf, unsafe.Pointer(uintptr(id)), //hack 这里把指针直接当int用
		avformat.CFunc(C.read_packet), avformat.CFunc(C.write_packet), nil) //TODO
	if cavio == nil {
		avutil.AvFreep(cbuf)
		buf = nil
		return nil
	}

	avio := &AVIOContext{
		CAVIOContext: cavio,
		CBuffer:      cbuf,
		BufferSize:   bufSize,
		Buffer:       buf,
		readFunc:     readFunc,
		writeFunc:    writeFunc,
		seekFunc:     seekFunc,
	}

	//AVIO绑定
	avioBindingC[avio.id] = avio

	return avio
}

func (avio *AVIOContext) Free() {
	//TODO
	avformat.AvioContextFree(&avio.CAVIOContext)
}

//export go_read_packet
func go_read_packet(opaque unsafe.Pointer, buf *C.uint8_t, buf_size C.int) C.int {
	id := uintptr(opaque) //hack 把指针当int用
	avio := avioBindingC[id]
	return C.int(avio.readFunc(unsafe.Pointer(buf), int(buf_size)))
}

//export go_write_packet
func go_write_packet(opaque unsafe.Pointer, buf *C.uint8_t, buf_size C.int) C.int {
	id := uintptr(opaque) //hack 把指针当int用
	avio := avioBindingC[id]
	return C.int(avio.writeFunc(unsafe.Pointer(buf), int(buf_size)))
}

//export go_seek
func go_seek(opaque unsafe.Pointer, offset C.int64_t, whence C.int) C.int64_t {
	id := uintptr(opaque) //hack 把指针当int用
	avio := avioBindingC[id]
	return C.int64_t(avio.seekFunc(int64(offset), int(whence)))
}
