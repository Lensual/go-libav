package goavfilter

import (
	"github.com/Lensual/go-libav/advance/goavutil"
	"github.com/Lensual/go-libav/avfilter"
)

type BufferSrc struct {
	*AVFilterContext
}

func NewBufferSrc(name string, args string, graphCtx *AVFilterGraph) (*BufferSrc, int) {
	filt := GetByName("buffer")
	filtCtx := BufferSrc{&AVFilterContext{}}
	ret := avfilter.AvfilterGraphCreateFilter(&filtCtx.CAVFilterContext, filt.CAVFilter, name, args, nil, graphCtx.CAVFilterGraph)
	if filtCtx.CAVFilterContext == nil {
		return nil, ret
	}
	return &filtCtx, ret
}

func (filtCtx *BufferSrc) AddFrameFlags(frame *goavutil.AVFrame, flags int) int {
	return avfilter.AvBuffersrcAddFrameFlags(filtCtx.CAVFilterContext, frame.CAVFrame, flags)
}
