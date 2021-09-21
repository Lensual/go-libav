package advance

import (
	"github.com/Lensual/go-libav/avutil"
)

type AvFrame struct {
	CAvFrame *avutil.CAVFrame
}

func NewAVFrame(size int) *AvFrame {
	return &AvFrame{
		CAvFrame: avutil.AvFrameAlloc(),
	}
}

func (frame *AvFrame) Free() {
	avutil.AvFrameFree(&frame.CAvFrame)
}
