package goswresample

import (
	"unsafe"

	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/swresample"
)

type SwrContext struct {
	CSwrContext *swresample.CSwrContext
}

func NewSwrContext() *SwrContext {
	c := swresample.SwrAlloc()
	if c == nil {
		return nil
	}
	return &SwrContext{
		CSwrContext: c,
	}
}

func NewSwrContextWithOpts(outChLayout *avutil.CAVChannelLayout, outSampleFmt avutil.CAVSampleFormat, outSampleRate int,
	inChLayout *avutil.CAVChannelLayout, inSampleFmt avutil.CAVSampleFormat, inSampleRate int) (*SwrContext, int) {
	swrCtx := SwrContext{}
	ret := swresample.SwrAllocSetOpts2(&swrCtx.CSwrContext, outChLayout, outSampleFmt, outSampleRate,
		inChLayout, inSampleFmt, inSampleRate, 0, nil)
	if ret != 0 {
		return nil, ret
	}
	return &swrCtx, ret
}

func (swrCtx *SwrContext) Init() int {
	return swresample.SwrInit(swrCtx.CSwrContext)
}

func (swrCtx *SwrContext) Free() {
	if swrCtx == nil || swrCtx.CSwrContext == nil {
		return
	}
	swresample.SwrFree(&swrCtx.CSwrContext)
}

func (swrCtx *SwrContext) ConvertUnsafe(in *unsafe.Pointer, inCount int, out *unsafe.Pointer, outCount int) int {
	return swresample.SwrConvert(swrCtx.CSwrContext, out, outCount, in, inCount)
}

// func (swrCtx *SwrContext) Convert(in []byte, inCount int, outCount int) ([]byte, int) {
// 	outSampleFmt, _ := swrCtx.GetOutSampleFmt()
// 	outChLayout, _ := swrCtx.GetOutChLayout()
// 	outPtr := unsafe.Pointer(nil)
// 	avutil.AvSamplesAlloc(&outPtr, nil, outChLayout.GetNbChannels(), int(outCount), outSampleFmt, 0)

// 	inPtr := unsafe.Pointer(unsafe.SliceData(in))
// 	ret := swresample.SwrConvert(swrCtx.CSwrContext, &outPtr, int(outCount), &inPtr, int(inCount))
// 	return
// }

func (swrCtx *SwrContext) GetDelay(base int64) int64 {
	return swresample.SwrGetDelay(swrCtx.CSwrContext, base)
}

// func (swrCtx *SwrContext) ConvertFrame() {
// 	swresample.SwrConvertFrame(&swrCtx.CSwrContext)
// }

func (swrCtx *SwrContext) GetInChLayout() (*avutil.CAVChannelLayout, int) {
	var layout avutil.CAVChannelLayout
	return &layout, avutil.AvOptGetChlayout(unsafe.Pointer(swrCtx.CSwrContext), "in_chlayout", 0, &layout)
}
func (swrCtx *SwrContext) SetInChLayout(layout *avutil.CAVChannelLayout) int {
	return avutil.AvOptSetChlayout(unsafe.Pointer(swrCtx.CSwrContext), "in_chlayout", layout, 0)
}

func (swrCtx *SwrContext) GetInSampleRate() (int64, int) {
	var sampleRate int64
	return sampleRate, avutil.AvOptGetInt(unsafe.Pointer(swrCtx.CSwrContext), "in_sample_rate", 0, &sampleRate)
}
func (swrCtx *SwrContext) SetInSampleRate(sampleRate int64) int {
	return avutil.AvOptSetInt(unsafe.Pointer(swrCtx.CSwrContext), "in_sample_rate", sampleRate, 0)
}

func (swrCtx *SwrContext) GetInSampleFmt() (avutil.CAVSampleFormat, int) {
	var fmt avutil.CAVSampleFormat
	return fmt, avutil.AvOptGetSampleFmt(unsafe.Pointer(swrCtx.CSwrContext), "in_sample_fmt", 0, &fmt)
}
func (swrCtx *SwrContext) SetInSampleFmt(fmt avutil.CAVSampleFormat) int {
	return avutil.AvOptSetSampleFmt(unsafe.Pointer(swrCtx.CSwrContext), "in_sample_fmt", fmt, 0)
}

func (swrCtx *SwrContext) GetOutChLayout() (*avutil.CAVChannelLayout, int) {
	var layout avutil.CAVChannelLayout
	return &layout, avutil.AvOptGetChlayout(unsafe.Pointer(swrCtx.CSwrContext), "out_chlayout", 0, &layout)
}
func (swrCtx *SwrContext) SetOutChLayout(layout *avutil.CAVChannelLayout) int {
	return avutil.AvOptSetChlayout(unsafe.Pointer(swrCtx.CSwrContext), "out_chlayout", layout, 0)
}

func (swrCtx *SwrContext) GetOutSampleRate() (int64, int) {
	var sampleRate int64
	return sampleRate, avutil.AvOptGetInt(unsafe.Pointer(swrCtx.CSwrContext), "out_sample_rate", 0, &sampleRate)
}
func (swrCtx *SwrContext) SetOutSampleRate(sampleRate int64) int {
	return avutil.AvOptSetInt(unsafe.Pointer(swrCtx.CSwrContext), "out_sample_rate", sampleRate, 0)
}

func (swrCtx *SwrContext) GetOutSampleFmt() (avutil.CAVSampleFormat, int) {
	var fmt avutil.CAVSampleFormat
	return fmt, avutil.AvOptGetSampleFmt(unsafe.Pointer(swrCtx.CSwrContext), "out_sample_fmt", 0, &fmt)
}
func (swrCtx *SwrContext) SetOutSampleFmt(fmt avutil.CAVSampleFormat) int {
	return avutil.AvOptSetSampleFmt(unsafe.Pointer(swrCtx.CSwrContext), "out_sample_fmt", fmt, 0)
}
