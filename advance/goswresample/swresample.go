package goswresample

import (
	"context"
	"unsafe"

	"github.com/Lensual/go-libav/advance/goavutil"
	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/ctypes"
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

func NewSwrContextWithOpts(outChLayout *goavutil.AVChannelLayout, outSampleFmt avutil.CAVSampleFormat, outSampleRate int,
	inChLayout *goavutil.AVChannelLayout, inSampleFmt avutil.CAVSampleFormat, inSampleRate int) (*SwrContext, int) {
	swrCtx := SwrContext{}
	ret := swresample.SwrAllocSetOpts2(&swrCtx.CSwrContext, outChLayout.CAVChannelLayout, outSampleFmt, outSampleRate,
		inChLayout.CAVChannelLayout, inSampleFmt, inSampleRate, 0, nil)
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

func (swrCtx *SwrContext) ConvertUnsafeToUnsafe(out *unsafe.Pointer, outCount int, in *unsafe.Pointer, inCount int) int {
	return swresample.SwrConvert(swrCtx.CSwrContext, out, outCount, in, inCount)
}

func (swrCtx *SwrContext) Convert(in []byte, inCount int) ([]byte, int) {
	outCount := inCount
	outSampleFmt, _ := swrCtx.GetOutSampleFmt()
	outChLayout, _ := swrCtx.GetOutChLayout()

	outBufSize := avutil.AvSamplesGetBufferSize(nil, outChLayout.GetNbChannels(), outCount, outSampleFmt, 0)
	cOut := avutil.AvMalloc(ctypes.SizeT(outBufSize))
	defer avutil.AvFree(cOut)
	//TODO wait ffmpeg av_samples_alloc return buffer size
	//outBufSize := avutil.AvSamplesAlloc(&outPtr, nil, outChLayout.GetNbChannels(), int(outCount), outSampleFmt, 0)

	inBufSize := len(in)
	cIn := avutil.AvMalloc(ctypes.SizeT(inBufSize))
	defer avutil.AvFree(cIn)
	copy(unsafe.Slice((*byte)(cIn), inBufSize), in)

	ret := swrCtx.ConvertUnsafeToUnsafe(&cOut, outCount, &cIn, inCount)
	if ret <= 0 {
		return nil, ret
	}
	outLength := avutil.AvSamplesGetBufferSize(nil, outChLayout.GetNbChannels(), ret, outSampleFmt, 1)
	out := make([]byte, outLength)
	copy(out, unsafe.Slice((*byte)(cOut), outLength))
	return out, ret
}

func (swrCtx *SwrContext) GetDelay(base int64) int64 {
	return swresample.SwrGetDelay(swrCtx.CSwrContext, base)
}

func (swrCtx *SwrContext) ConvertFrameTo(output *goavutil.AVFrame, input *goavutil.AVFrame) int {
	return swresample.SwrConvertFrame(swrCtx.CSwrContext, output.CAVFrame, input.CAVFrame)
}

func (swrCtx *SwrContext) ConvertFrame(input *goavutil.AVFrame) (*goavutil.AVFrame, int) {
	output := goavutil.AllocAVFrame()

	outChLayout, code := swrCtx.GetOutChLayout()
	if code != 0 {
		output.Free()
		return nil, code
	}
	output.SetChLayout(outChLayout)

	outSampleRate, code := swrCtx.GetOutSampleRate()
	if code != 0 {
		output.Free()
		return nil, code
	}
	output.SetSampleRate(int(outSampleRate))

	outSampleFmt, code := swrCtx.GetOutSampleFmt()
	if code != 0 {
		output.Free()
		return nil, code
	}
	output.SetFormat(int(outSampleFmt))

	var cInput *avutil.CAVFrame
	if input != nil {
		cInput = input.CAVFrame
	}

	code = swresample.SwrConvertFrame(swrCtx.CSwrContext, output.CAVFrame, cInput)
	if code != 0 {
		output.Free()
		return nil, code
	}
	return output, code
}

func ConvertFrames() {
	// TODO
}

func (swrCtx *SwrContext) ConvertFrameChan(ctx context.Context, inputChan <-chan *goavutil.AVFrame) (context.Context, <-chan *goavutil.AVFrame) {
	ctx, cancel := context.WithCancelCause(ctx)
	outputChan := make(chan *goavutil.AVFrame)

	go func() {
		defer close(outputChan)
	loop:
		for {
			select {
			case <-ctx.Done():
				return
			case input, ok := <-inputChan:
				if !ok {
					break loop
				}
				output, code := swrCtx.ConvertFrame(input)
				input.Unref()
				if output != nil {
					select {
					case <-ctx.Done():
						output.Unref()
						return
					case outputChan <- output:
					}
				}
				if code != 0 {
					cancel(goavutil.AvErr(code))
					return
				}
			}
		}

		//flush resample
		output, code := swrCtx.ConvertFrame(nil)
		if output != nil {
			select {
			case <-ctx.Done():
				output.Unref()
				return
			case outputChan <- output:
			}
		}
		if code != 0 {
			cancel(goavutil.AvErr(code))
			return
		}
	}()

	return ctx, outputChan
}

func (swrCtx *SwrContext) GetInChLayout() (*goavutil.AVChannelLayout, int) {
	var layout avutil.CAVChannelLayout
	ret := avutil.AvOptGetChlayout(unsafe.Pointer(swrCtx.CSwrContext), "in_chlayout", 0, &layout)
	return &goavutil.AVChannelLayout{
		CAVChannelLayout: &layout,
	}, ret
}
func (swrCtx *SwrContext) SetInChLayout(layout *goavutil.AVChannelLayout) int {
	return avutil.AvOptSetChlayout(unsafe.Pointer(swrCtx.CSwrContext), "in_chlayout", layout.CAVChannelLayout, 0)
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

func (swrCtx *SwrContext) GetOutChLayout() (*goavutil.AVChannelLayout, int) {
	var layout avutil.CAVChannelLayout
	ret := avutil.AvOptGetChlayout(unsafe.Pointer(swrCtx.CSwrContext), "out_chlayout", 0, &layout)
	return &goavutil.AVChannelLayout{
		CAVChannelLayout: &layout,
	}, ret
}
func (swrCtx *SwrContext) SetOutChLayout(layout *goavutil.AVChannelLayout) int {
	return avutil.AvOptSetChlayout(unsafe.Pointer(swrCtx.CSwrContext), "out_chlayout", layout.CAVChannelLayout, 0)
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
