package advance

/*
#cgo LDFLAGS: -lavformat

#include "libavformat/avformat.h"
*/
import "C"

import (
	"unsafe"

	"github.com/Lensual/go-libav/avformat"
)

type AvformatContext struct {
	CAvformatContext *avformat.CAVFormatContext
}

func NewAvformatContext() *AvformatContext {
	ctx := avformat.AvformatAllocContext()
	if ctx == nil {
		return nil
	}
	return &AvformatContext{
		CAvformatContext: ctx,
	}
}

//设置IO上下文
func (fmtCtx AvformatContext) SetIOContext(avioCtx *AVIOContext) {
	(*C.AVFormatContext)(unsafe.Pointer(fmtCtx.CAvformatContext)).pb =
		(*C.struct_AVIOContext)(unsafe.Pointer(avioCtx.CAVIOContext))
}

func (fmtCtx AvformatContext) OpenInput(url string) int {
	return avformat.AvformatOpenInput(&fmtCtx.CAvformatContext, url, nil, nil) //TODO fmt options
}

func (fmtCtx AvformatContext) FindStreamInfo() int {
	return avformat.AvformatFindStreamInfo(fmtCtx.CAvformatContext, nil) //TODO options
}

func (fmtCtx AvformatContext) DumpFormat(index int, url string, is_output int) {
	avformat.AvDumpFormat(fmtCtx.CAvformatContext, index, url, is_output)
}

func (fmtCtx AvformatContext) CloseInput() {
	avformat.AvformatCloseInput(&fmtCtx.CAvformatContext)
}
