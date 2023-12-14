package goavfilter

import (
	"github.com/Lensual/go-libav/advance/goavutil"
	"github.com/Lensual/go-libav/avfilter"
)

type BufferSink struct {
	*AVFilterContext
}

func NewBufferSink(name string, args string, graphCtx *AVFilterGraph) (*BufferSink, int) {
	filt := GetByName("buffersink")
	filtCtx := BufferSink{&AVFilterContext{}}
	ret := avfilter.AvfilterGraphCreateFilter(&filtCtx.CAVFilterContext, filt.CAVFilter, name, args, nil, graphCtx.CAVFilterGraph)
	if filtCtx.CAVFilterContext == nil {
		return nil, ret
	}
	return &filtCtx, ret
}

func (filtCtx *BufferSink) GetFrame(frame *goavutil.AVFrame) int {
	return avfilter.AvBuffersinkGetFrame(filtCtx.CAVFilterContext, frame.CAVFrame)
}
