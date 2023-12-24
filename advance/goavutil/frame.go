package goavutil

import (
	"unsafe"

	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/ctypes"
)

type AVFrame struct {
	CAVFrame *avutil.CAVFrame
}

// #region members

func (frame *AVFrame) GetData() [avutil.AV_NUM_DATA_POINTERS]unsafe.Pointer {
	return frame.CAVFrame.GetData()
}

func (frame *AVFrame) SetData(data [avutil.AV_NUM_DATA_POINTERS]unsafe.Pointer) {
	frame.CAVFrame.SetData(data)
}

func (frame *AVFrame) GetLineSize() [avutil.AV_NUM_DATA_POINTERS]int {
	ret := [avutil.AV_NUM_DATA_POINTERS]int{}
	cLinesize := frame.CAVFrame.GetLinesize()
	for i := range ret {
		ret[i] = int(cLinesize[i])
	}
	return ret
}

func (frame *AVFrame) SetLinesize(linesize [avutil.AV_NUM_DATA_POINTERS]int) {
	cLinesize := [avutil.AV_NUM_DATA_POINTERS]ctypes.Int{}
	for i := range linesize {
		cLinesize[i] = ctypes.Int(linesize[i])
	}
}

func (frame *AVFrame) GetExtendedData() *unsafe.Pointer {
	return frame.CAVFrame.GetExtendedData()
}

func (frame *AVFrame) SetExtendedData(extendedData *unsafe.Pointer) {
	frame.CAVFrame.SetExtendedData(extendedData)
}

func (frame *AVFrame) GetWidth() int {
	return frame.CAVFrame.GetWidth()
}

func (frame *AVFrame) SetWidth(width int) {
	frame.CAVFrame.SetWidth(width)
}

func (frame *AVFrame) GetHeight() int {
	return frame.CAVFrame.GetHeight()
}

func (frame *AVFrame) SetHeight(height int) {
	frame.CAVFrame.SetHeight(height)
}

func (frame *AVFrame) GetNbSamples() int {
	return frame.CAVFrame.GetNbSamples()
}

func (frame *AVFrame) SetNbSamples(nbSamples int) {
	frame.CAVFrame.SetNbSamples(nbSamples)
}

func (frame *AVFrame) GetFormat() int {
	return frame.CAVFrame.GetFormat()
}

func (frame *AVFrame) SetFormat(format int) {
	frame.CAVFrame.SetFormat(format)
}

func (frame *AVFrame) GetPictType() avutil.CAVPictureType {
	return frame.CAVFrame.GetPictType()
}

func (frame *AVFrame) SetPictType(pictType avutil.CAVPictureType) {
	frame.CAVFrame.SetPictType(pictType)
}

func (frame *AVFrame) GetSampleAspectRatio() avutil.CAVRational {
	return frame.CAVFrame.GetSampleAspectRatio()
}

func (frame *AVFrame) SetSampleAspectRatio(sampleAspectRatio avutil.CAVRational) {
	frame.CAVFrame.SetSampleAspectRatio(sampleAspectRatio)
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

func (frame *AVFrame) SetPktDts(pktDts int64) {
	frame.CAVFrame.SetPktDts(pktDts)
}

func (frame *AVFrame) GetTimeBase() avutil.CAVRational {
	return frame.CAVFrame.GetTimeBase()
}

func (frame *AVFrame) SetTimeBase(timeBase avutil.CAVRational) {
	frame.CAVFrame.SetTimeBase(timeBase)
}

func (frame *AVFrame) GetQuality() int {
	return frame.CAVFrame.GetQuality()
}

func (frame *AVFrame) SetQuality(quality int) {
	frame.CAVFrame.SetQuality(quality)
}

func (frame *AVFrame) GetOpaque() unsafe.Pointer {
	return frame.CAVFrame.GetOpaque()
}

func (frame *AVFrame) SetOpaque(opaque unsafe.Pointer) {
	frame.CAVFrame.SetOpaque(opaque)
}

func (frame *AVFrame) GetRepeatPict() int {
	return frame.CAVFrame.GetRepeatPict()
}

func (frame *AVFrame) SetRepeatPict(repeatPict int) {
	frame.CAVFrame.SetRepeatPict(repeatPict)
}

func (frame *AVFrame) GetSampleRate() int {
	return frame.CAVFrame.GetSampleRate()
}

func (frame *AVFrame) SetSampleRate(sampleRate int) {
	frame.CAVFrame.SetSampleRate(sampleRate)
}

func (frame *AVFrame) GetBuf() [avutil.AV_NUM_DATA_POINTERS]*avutil.CAVBufferRef {
	return frame.CAVFrame.GetBuf()
}

func (frame *AVFrame) SetBuf(buf [avutil.AV_NUM_DATA_POINTERS]*avutil.CAVBufferRef) {
	frame.CAVFrame.SetBuf(buf)
}

func (frame *AVFrame) GetExtendedBuf() **avutil.CAVBufferRef {
	return frame.CAVFrame.GetExtendedBuf()
}

func (frame *AVFrame) SetExtendedBuf(extendedBuf **avutil.CAVBufferRef) {
	frame.CAVFrame.SetExtendedBuf(extendedBuf)
}

func (frame *AVFrame) GetNbExtendedBuf() int {
	return frame.CAVFrame.GetNbExtendedBuf()
}

func (frame *AVFrame) SetNbExtendedBuf(nbExtendedBuf int) {
	frame.CAVFrame.SetNbExtendedBuf(nbExtendedBuf)
}

func (frame *AVFrame) GetSideData() **avutil.CAVFrameSideData {
	return frame.CAVFrame.GetSideData()
}

func (frame *AVFrame) SetSideData(sideData **avutil.CAVFrameSideData) {
	frame.CAVFrame.SetSideData(sideData)
}

func (frame *AVFrame) GetNbSideData() int {
	return frame.CAVFrame.GetNbSideData()
}

func (frame *AVFrame) SetNbSideData(nbSideData int) {
	frame.CAVFrame.SetNbSideData(nbSideData)
}

func (frame *AVFrame) GetFlags() int {
	return frame.CAVFrame.GetFlags()
}

func (frame *AVFrame) SetFlags(flags int) {
	frame.CAVFrame.SetFlags(flags)
}

func (frame *AVFrame) GetColorRange() avutil.CAVColorRange {
	return frame.CAVFrame.GetColorRange()
}

func (frame *AVFrame) SetColorRange(colorRange avutil.CAVColorRange) {
	frame.CAVFrame.SetColorRange(colorRange)
}

func (frame *AVFrame) GetColorPrimaries() avutil.CAVColorPrimaries {
	return frame.CAVFrame.GetColorPrimaries()
}

func (frame *AVFrame) SetColorPrimaries(colorPrimaries avutil.CAVColorPrimaries) {
	frame.CAVFrame.SetColorPrimaries(colorPrimaries)
}

func (frame *AVFrame) GetColorTrc() avutil.CAVColorTransferCharacteristic {
	return frame.CAVFrame.GetColorTrc()
}

func (frame *AVFrame) SetColorTrc(colorTrc avutil.CAVColorTransferCharacteristic) {
	frame.CAVFrame.SetColorTrc(colorTrc)
}

func (frame *AVFrame) GetColorspace() avutil.CAVColorSpace {
	return frame.CAVFrame.GetColorspace()
}

func (frame *AVFrame) SetColorspace(colorspace avutil.CAVColorSpace) {
	frame.CAVFrame.SetColorspace(colorspace)
}

func (frame *AVFrame) GetChromaLocation() avutil.CAVChromaLocation {
	return frame.CAVFrame.GetChromaLocation()
}

func (frame *AVFrame) SetChromaLocation(chromaLocation avutil.CAVChromaLocation) {
	frame.CAVFrame.SetChromaLocation(chromaLocation)
}

func (frame *AVFrame) GetBestEffortTimestamp() int64 {
	return frame.CAVFrame.GetBestEffortTimestamp()
}

func (frame *AVFrame) SetBestEffortTimestamp(bestEffortTimestamp int64) {
	frame.CAVFrame.SetBestEffortTimestamp(bestEffortTimestamp)
}

func (frame *AVFrame) GetMetadata() *avutil.CAVDictionary {
	return frame.CAVFrame.GetMetadata()
}

func (frame *AVFrame) SetMetadata(metadata *avutil.CAVDictionary) {
	frame.CAVFrame.SetMetadata(metadata)
}

func (frame *AVFrame) GetDecodeErrorFlags() int {
	return frame.CAVFrame.GetDecodeErrorFlags()
}

func (frame *AVFrame) SetDecodeErrorFlags(decodeErrorFlags int) {
	frame.CAVFrame.SetDecodeErrorFlags(decodeErrorFlags)
}

func (frame *AVFrame) GetHwFramesCtx() *avutil.CAVBufferRef {
	return frame.CAVFrame.GetHwFramesCtx()
}

func (frame *AVFrame) SetHwFramesCtx(hwFramesCtx *avutil.CAVBufferRef) {
	frame.CAVFrame.SetHwFramesCtx(hwFramesCtx)
}

func (frame *AVFrame) GetOpaqueRef() *avutil.CAVBufferRef {
	return frame.CAVFrame.GetOpaqueRef()
}

func (frame *AVFrame) SetOpaqueRef(opaqueRef *avutil.CAVBufferRef) {
	frame.CAVFrame.SetOpaqueRef(opaqueRef)
}

func (frame *AVFrame) GetCropTop() ctypes.SizeT {
	return frame.CAVFrame.GetCropTop()
}

func (frame *AVFrame) SetCropTop(cropTop ctypes.SizeT) {
	frame.CAVFrame.SetCropTop(cropTop)
}

func (frame *AVFrame) GetCropBottom() ctypes.SizeT {
	return frame.CAVFrame.GetCropBottom()
}

func (frame *AVFrame) SetCropBottom(cropBottom ctypes.SizeT) {
	frame.CAVFrame.SetCropBottom(cropBottom)
}

func (frame *AVFrame) GetCropLeft() ctypes.SizeT {
	return frame.CAVFrame.GetCropLeft()
}

func (frame *AVFrame) SetCropLeft(cropLeft ctypes.SizeT) {
	frame.CAVFrame.SetCropLeft(cropLeft)
}

func (frame *AVFrame) GetCropRight() ctypes.SizeT {
	return frame.CAVFrame.GetCropRight()
}

func (frame *AVFrame) SetCropRight(cropRight ctypes.SizeT) {
	frame.CAVFrame.SetCropRight(cropRight)
}

func (frame *AVFrame) GetPrivateRef() *avutil.CAVBufferRef {
	return frame.CAVFrame.GetPrivateRef()
}

func (frame *AVFrame) SetPrivateRef(privateRef *avutil.CAVBufferRef) {
	frame.CAVFrame.SetPrivateRef(privateRef)
}

func (frame *AVFrame) GetChLayout() *AVChannelLayout {
	return &AVChannelLayout{
		CAVChannelLayout: frame.CAVFrame.GetChLayoutPtr(),
	}
}

func (frame *AVFrame) SetChLayout(chLayout *AVChannelLayout) {
	frame.CAVFrame.SetChLayout(*chLayout.CAVChannelLayout)
}

func (frame *AVFrame) GetDuration() int64 {
	return frame.CAVFrame.GetDuration()
}

func (frame *AVFrame) SetDuration(duration int64) {
	frame.CAVFrame.SetDuration(duration)
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

func (frame *AVFrame) RefTo(dst *AVFrame) int {
	return avutil.AvFrameRef(dst.CAVFrame, frame.CAVFrame)
}

func (frame *AVFrame) ReplaceTo(dst *AVFrame) int {
	return avutil.AvFrameReplace(dst.CAVFrame, frame.CAVFrame)
}

func (frame *AVFrame) Clone() *AVFrame {
	newFrame := avutil.AvFrameClone(frame.CAVFrame)
	if newFrame == nil {
		return nil
	}
	return &AVFrame{
		CAVFrame: newFrame,
	}
}

func (frame *AVFrame) Unref() {
	avutil.AvFrameUnref(frame.CAVFrame)
}

func (frame *AVFrame) MoveRefTo(dst *AVFrame) {
	avutil.AvFrameMoveRef(dst.CAVFrame, frame.CAVFrame)
}

func (frame *AVFrame) AllocBuffer(align int) int {
	return avutil.AvFrameGetBuffer(frame.CAVFrame, align)
}

func (frame *AVFrame) IsWritable() int {
	return avutil.AvFrameIsWritable(frame.CAVFrame)
}

func (frame *AVFrame) MakeWritable() int {
	return avutil.AvFrameMakeWritable(frame.CAVFrame)
}

func (frame *AVFrame) CopyTo(dst *AVFrame) int {
	return avutil.AvFrameCopy(dst.CAVFrame, frame.CAVFrame)
}

func (frame *AVFrame) CopyPropsTo(dst *AVFrame) int {
	return avutil.AvFrameCopyProps(dst.CAVFrame, frame.CAVFrame)
}

func (frame *AVFrame) GetPlaneBuffer(plane int) *avutil.CAVBufferRef {
	return avutil.AvFrameGetPlaneBuffer(frame.CAVFrame, plane)
}

func (frame *AVFrame) NewSideData(_type avutil.CAVFrameSideDataType, size int) *avutil.CAVFrameSideData {
	return avutil.AvFrameNewSideData(frame.CAVFrame, _type, size)
}

func (frame *AVFrame) NewSideDataFromBuf(_type avutil.CAVFrameSideDataType, buf *avutil.CAVBufferRef) *avutil.CAVFrameSideData {
	return avutil.AvFrameNewSideDataFromBuf(frame.CAVFrame, _type, buf)
}

func (frame *AVFrame) GetSideDataByType(_type avutil.CAVFrameSideDataType) *avutil.CAVFrameSideData {
	return avutil.AvFrameGetSideData(frame.CAVFrame, _type)
}

func (frame *AVFrame) RemoveSideData(_type avutil.CAVFrameSideDataType) {
	avutil.AvFrameRemoveSideData(frame.CAVFrame, _type)
}

func (frame *AVFrame) ApplyCropping(flags int) int {
	return avutil.AvFrameApplyCropping(frame.CAVFrame, flags)
}
