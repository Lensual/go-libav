package goavutil

import (
	"unsafe"

	"github.com/Lensual/go-libav/avutil"
)

type AVFrame struct {
	CAVFrame *avutil.CAVFrame
}

// #region members

func (frame *AVFrame) GetData() [avutil.AV_NUM_DATA_POINTERS]unsafe.Pointer {
	return frame.CAVFrame.GetData()
}

func (frame *AVFrame) GetLineSize() [avutil.AV_NUM_DATA_POINTERS]int {
	ret := [avutil.AV_NUM_DATA_POINTERS]int{}
	for i := range ret {
		ret[i] = int(frame.CAVFrame.GetLinesize()[i])
	}
	return ret
}

func (frame *AVFrame) GetExtendedData() *unsafe.Pointer {
	return frame.CAVFrame.GetExtendedData()
}

func (frame *AVFrame) GetWidth() int {
	return frame.CAVFrame.GetWidth()
}

func (frame *AVFrame) GetHeight() int {
	return frame.CAVFrame.GetHeight()
}

func (frame *AVFrame) GetNbSamples() int {
	return frame.CAVFrame.GetNbSamples()
}

func (frame *AVFrame) GetFormat() int {
	return frame.CAVFrame.GetFormat()
}

func (frame *AVFrame) GetSampleAspectRatio() avutil.CAVRational {
	return frame.CAVFrame.GetSampleAspectRatio()
}

func (frame *AVFrame) GetPts() int64 {
	return frame.CAVFrame.GetPts()
}

func (frame *AVFrame) SetPts(pts int64) {
	frame.CAVFrame.SetPts(pts)
}

func (frame *AVFrame) GetPktDts() int64 {
	return frame.CAVFrame.GetPktDts()
}

func (frame *AVFrame) GetTimeBase() avutil.CAVRational {
	return frame.CAVFrame.GetTimeBase()
}

func (frame *AVFrame) GetQuality() int {
	return frame.CAVFrame.GetQuality()
}

func (frame *AVFrame) GetSampleRate() int {
	return frame.CAVFrame.GetSampleRate()
}

func (frame *AVFrame) GetBestEffortTimestamp() int64 {
	return frame.CAVFrame.GetBestEffortTimestamp()
}

func (frame *AVFrame) GetChLayout() *AVChannelLayout {
	return &AVChannelLayout{
		CAVChannelLayout: frame.CAVFrame.GetChLayoutPtr(),
	}
}

func (frame *AVFrame) GetChLayoutPtr() *avutil.CAVChannelLayout {
	return frame.CAVFrame.GetChLayoutPtr()
}

func (frame *AVFrame) GetDuration() int64 {
	return frame.CAVFrame.GetDuration()
}

//#endregion members

func AllocAVFrame() *AVFrame {
	return &AVFrame{
		CAVFrame: avutil.AvFrameAlloc(),
	}
}

func (frame *AVFrame) Free() {
	avutil.AvFrameFree(&frame.CAVFrame)
}

func (frame *AVFrame) Unref() {
	avutil.AvFrameUnref(frame.CAVFrame)
}
