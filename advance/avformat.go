package advance

import (
	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
)

type CustomAVIO struct {
	CAVIOContext *avformat.CAVIOContext
}

func NewCustomAVIO(bufSize int, writeFlag bool) *CustomAVIO {
	buf := avutil.AvMalloc(uint64(bufSize))
	var wf int
	if writeFlag {
		wf = 1
	}

	//TODO

	return &CustomAVIO{
		CAVIOContext: avformat.AvioAllocContext(buf, bufSize, wf, nil, nil, nil, nil), //TODO
	}

}

type AvformatContext struct {
	CAvformatContext *avformat.CAVFormatContext
}

func NewAvformatContext() *AvformatContext {
	return &AvformatContext{
		CAvformatContext: avformat.AvformatAllocContext(),
	}
}
