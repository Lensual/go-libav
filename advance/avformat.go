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
