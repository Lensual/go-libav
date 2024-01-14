package goavfilter

import (
	"unsafe"

	"github.com/Lensual/go-libav/avfilter"
)

type AVFilterGraph struct {
	CAVFilterGraph *avfilter.CAVFilterGraph
}

func (g *AVFilterGraph) CreateFilter(filt *AVFilter, name string, args string) (*AVFilterContext, int) {
	return NewContext(filt, name, args, g)
}

func (g *AVFilterGraph) CreateBufferSrc(name string, args string) (*BufferSrc, int) {
	return NewBufferSrc(name, args, g)
}

func (g *AVFilterGraph) CreateBufferSink(name string, args string) (*BufferSink, int) {
	return NewBufferSink(name, args, g)
}

func (g *AVFilterGraph) ValidConfig(logCtx unsafe.Pointer) int {
	return avfilter.AvfilterGraphConfig(g.CAVFilterGraph, logCtx)
}

func (g *AVFilterGraph) Free() {
	avfilter.AvfilterGraphFree(&g.CAVFilterGraph)
}

func (g *AVFilterGraph) ParsePtr(filters string, inputs *AVFilterInOut, outputs *AVFilterInOut, logCtx unsafe.Pointer) int {
	return avfilter.AvfilterGraphParsePtr(g.CAVFilterGraph, filters, &inputs.CAVFilterInOut, &outputs.CAVFilterInOut, logCtx)
}

func AllocAvFilterGraph() *AVFilterGraph {
	g := avfilter.AvfilterGraphAlloc()
	if g == nil {
		return nil
	}
	return &AVFilterGraph{
		CAVFilterGraph: g,
	}
}
