package goavfilter

import (
	"unsafe"

	"github.com/Lensual/go-libav/avfilter"
)

type AVFilterContext struct {
	CAVFilterContext *avfilter.CAVFilterContext
}

//#region members

func (filtCtx *AVFilterContext) GetInputs() []*avfilter.CAVFilterLink {
	return unsafe.Slice(filtCtx.CAVFilterContext.GetInputs(), filtCtx.CAVFilterContext.GetNbInputs())
}

func (filtCtx *AVFilterContext) GetOutputs() []*avfilter.CAVFilterLink {
	return unsafe.Slice(filtCtx.CAVFilterContext.GetOutputs(), filtCtx.CAVFilterContext.GetNbOutputs())
}

//#endregion members

func NewAvFilterContext(filt *AVFilter, name string, args string, graphCtx *AVFilterGraph) (*AVFilterContext, int) {
	filtCtx := AVFilterContext{}
	ret := avfilter.AvfilterGraphCreateFilter(&filtCtx.CAVFilterContext, filt.CAVFilter, name, args, nil, graphCtx.CAVFilterGraph)
	if filtCtx.CAVFilterContext == nil {
		return nil, ret
	}
	return &filtCtx, ret
}
