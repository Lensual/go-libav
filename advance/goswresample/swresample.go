package goswresample

import (
	"context"
	"math"
	"unsafe"

	"github.com/pkg/errors"

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

func (swrCtx *SwrContext) IsInitialized() bool {
	return swresample.SwrIsInitialized(swrCtx.CSwrContext) > 0
}

func (swrCtx *SwrContext) Free() {
	if swrCtx == nil || swrCtx.CSwrContext == nil {
		return
	}
	swresample.SwrFree(&swrCtx.CSwrContext)
}

func (swrCtx *SwrContext) Close() {
	swresample.SwrClose(swrCtx.CSwrContext)
}

// Convert pointer to pointer.
func (swrCtx *SwrContext) ConvertUnsafeToUnsafe(out *unsafe.Pointer, outCount int, in *unsafe.Pointer, inCount int) int {
	return swresample.SwrConvert(swrCtx.CSwrContext, out, outCount, in, inCount)
}

// Convert slice to new slice.
func (swrCtx *SwrContext) Convert(in []byte, inCount int) ([]byte, int) {
	inSampleRate, code := swrCtx.GetInSampleRate()
	if code < 0 {
		return nil, code
	}
	outSampleRate, code := swrCtx.GetOutSampleRate()
	if code < 0 {
		return nil, code
	}
	outCount := int(avutil.AvRescaleRnd(swresample.SwrGetDelay(swrCtx.CSwrContext, inSampleRate)+
		int64(inCount), outSampleRate, inSampleRate, avutil.AV_ROUND_UP))
	outSampleFmt, code := swrCtx.GetOutSampleFmt()
	if code < 0 {
		return nil, code
	}
	outChLayout, code := swrCtx.GetOutChLayout()
	if code < 0 {
		return nil, code
	}

	outBufSize := avutil.AvSamplesGetBufferSize(nil, outChLayout.GetNbChannels(), outCount, outSampleFmt, 0)
	cOut := avutil.AvMalloc(ctypes.SizeT(outBufSize))
	defer avutil.AvFree(cOut)
	//TODO wait ffmpeg implement av_samples_alloc return buffer size
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

// Convert frame data from channel.
//
// Parameter:
//
//	ctx: The context.Context to cancel this goroutine.
//	inChan: The []byte channel.
//
// Return:
//
//	context.Context: The context.Context to get the cause of this gorouine.
//	<-chan []byte: The channel to read converted frame data.
//
// Example:
//
//	ctx, outDataCh = trans.swrCtx.ConvertChan(parentCtx, inDataCh)
//	for {
//		select {
//		case <-ctx.Done():
//			break loop
//		case outData, ok := <-outDataCh:
//			if !ok {
//				break loop
//			}
//			// do something.
//		}
//	}
//	err := context.Cause(ctx)
//	if err != nil {
//		panic(err)
//	}
func (swrCtx *SwrContext) ConvertChan(ctx context.Context, inChan <-chan []byte) (context.Context, chan<- []byte) {
	ctx, cancel := context.WithCancelCause(ctx)
	outChan := make(chan []byte)

	go func() {
		defer close(outChan)

		inSampleFmt, code := swrCtx.GetOutSampleFmt()
		if code < 0 {
			cancel(goavutil.AvErr(code))
			return
		}
		inChLayout, code := swrCtx.GetOutChLayout()
		if code < 0 {
			cancel(goavutil.AvErr(code))
			return
		}

	loop:
		for {
			select {
			case <-ctx.Done():
				return
			case in, ok := <-inChan:
				if !ok {
					break loop
				}
				// compute inCount by sampleFmt and channelNum
				inCount := float64(len(in)) / float64(avutil.AvGetBytesPerSample(inSampleFmt)*inChLayout.GetNbChannels())
				if math.Trunc(inCount) != inCount {
					cancel(errors.New("not integer"))
					return
				}
				out, code := swrCtx.Convert(in, int(inCount))
				if out != nil {
					outChan <- out
				}
				if code < 0 {
					cancel(goavutil.AvErr(code))
					return
				}
			}
		}

		//flush resample
		out, code := swrCtx.Convert(nil, 0)
		if out != nil {
			select {
			case <-ctx.Done():
				return
			case outChan <- out:
			}
		}
		if code < 0 {
			cancel(goavutil.AvErr(code))
			return
		}
	}()

	return ctx, outChan
}

func (swrCtx *SwrContext) NextPts(pts int64) int64 {
	return swresample.SwrNextPts(swrCtx.CSwrContext, pts)
}

func (swrCtx *SwrContext) SetCompensation(pts int64, sampleDelta int, compensationDistance int) int {
	return swresample.SwrSetCompensation(swrCtx.CSwrContext, sampleDelta, compensationDistance)
}

func (swrCtx *SwrContext) SetChannelMapping() int {
	//TODO
	panic("not implemented")
}

func (swrCtx *SwrContext) SetMatrix() int {
	//TODO
	panic("not implemented")
}

func (swrCtx *SwrContext) DropOutput(count int) int {
	return swresample.SwrDropOutput(swrCtx.CSwrContext, count)
}

func (swrCtx *SwrContext) InjectSilence(count int) int {
	return swresample.SwrInjectSilence(swrCtx.CSwrContext, count)
}

func (swrCtx *SwrContext) GetDelay(base int64) int64 {
	return swresample.SwrGetDelay(swrCtx.CSwrContext, base)
}

func (swrCtx *SwrContext) SwrGetOutSamples(inSamples int) int {
	return swresample.SwrGetOutSamples(swrCtx.CSwrContext, inSamples)
}

// Convert AVFrame to existing AVFrame.
//
// Example:
//
//	output := goavutil.AllocAVFrame()
//	output.SetChLayout(outChLayout)
//	output.SetSampleRate(int(outSampleRate))
//	output.SetFormat(int(outSampleFmt))
//	code := swrCtx.ConvertFrameToFrame(output, input)
func (swrCtx *SwrContext) ConvertFrameToFrame(output *goavutil.AVFrame, input *goavutil.AVFrame) int {
	var cInput *avutil.CAVFrame
	if input != nil {
		cInput = input.CAVFrame
	}

	return swresample.SwrConvertFrame(swrCtx.CSwrContext, output.CAVFrame, cInput)
}

// Convert AVFrame to new AVFrame.
//
// Example:
//
//	outFrame, code := ConvertFrame(inFrame)
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

	code = swrCtx.ConvertFrameToFrame(output, input)
	if code != 0 {
		output.Free()
		return nil, code
	}
	return output, code
}

// Convert AVFrame from channel.
//
// Parameter:
//
//	ctx: The context.Context to cancel this goroutine.
//	inputChan: The AVFrame channel.
//
// Return:
//
//	context.Context: The context.Context to get the cause of this gorouine.
//	<-chan *goavutil.AVFrame: The channel to read converted AVFrame.
//
// Example:
//
//	ctx, outFrameCh = trans.swrCtx.ConvertFrameChan(parentCtx, inFrameCh)
//	for {
//		select {
//		case <-ctx.Done():
//			break loop
//		case frame, ok := <-outFrameCh:
//			if !ok {
//				break loop
//			}
//			// do something.
//		}
//	}
//	err := context.Cause(ctx)
//	if err != nil {
//		panic(err)
//	}
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

func (swrCtx *SwrContext) SwrConfigFrame(output *goavutil.AVFrame, input *goavutil.AVFrame) int {
	return swresample.SwrConfigFrame(swrCtx.CSwrContext, output.CAVFrame, input.CAVFrame)
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
	var sampleRate ctypes.Int64
	return int64(sampleRate), avutil.AvOptGetInt(unsafe.Pointer(swrCtx.CSwrContext), "in_sample_rate", 0, &sampleRate)
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
	var sampleRate ctypes.Int64
	return int64(sampleRate), avutil.AvOptGetInt(unsafe.Pointer(swrCtx.CSwrContext), "out_sample_rate", 0, &sampleRate)
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
