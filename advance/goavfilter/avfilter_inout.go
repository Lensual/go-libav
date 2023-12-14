package goavfilter

import (
	"github.com/Lensual/go-libav/avfilter"
)

type AVFilterInOut struct {
	CAVFilterInOut *avfilter.CAVFilterInOut
}

//#region members

func (inout *AVFilterInOut) GetName() string {
	return inout.CAVFilterInOut.GetName()
}

func (inout *AVFilterInOut) SetName(name string) {
	inout.CAVFilterInOut.SetName(name)
}

func (inout *AVFilterInOut) GetFilterCtx() *AVFilterContext {
	filtCtx := inout.CAVFilterInOut.GetFilterCtx()
	if filtCtx == nil {
		return nil
	}
	return &AVFilterContext{
		CAVFilterContext: filtCtx,
	}
}

func (inout *AVFilterInOut) SetFilterCtx(filterCtx *AVFilterContext) {
	inout.CAVFilterInOut.SetFilterCtx(filterCtx.CAVFilterContext)
}

func (inout *AVFilterInOut) GetPadIdx() int {
	return inout.CAVFilterInOut.GetPadIdx()
}

func (inout *AVFilterInOut) SetPadIdx(padIdx int) {
	inout.CAVFilterInOut.SetPadIdx(padIdx)
}

func (inout *AVFilterInOut) GetNext() *AVFilterInOut {
	next := inout.CAVFilterInOut.GetNext()
	if next == nil {
		return nil
	}
	return &AVFilterInOut{
		CAVFilterInOut: next,
	}
}

func (inout *AVFilterInOut) SetNext(next *AVFilterInOut) {
	var v *avfilter.CAVFilterInOut = nil
	if next != nil {
		v = next.CAVFilterInOut
	}
	inout.CAVFilterInOut.SetNext(v)
}

//#endregion members

func (inout *AVFilterInOut) Free() {
	avfilter.AvfilterInoutFree(&inout.CAVFilterInOut)
}

func AllocAvFilterInOut() *AVFilterInOut {
	io := avfilter.AvfilterInoutAlloc()
	if io == nil {
		return nil
	}
	return &AVFilterInOut{
		CAVFilterInOut: io,
	}
}
