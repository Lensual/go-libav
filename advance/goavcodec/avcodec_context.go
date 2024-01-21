package goavcodec

import (
	"context"
	"syscall"
	"unsafe"

	"github.com/Lensual/go-libav/advance/goavutil"
	"github.com/Lensual/go-libav/avcodec"
	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/ctypes"
)

type AVCodecContext struct {
	CAVCodecContext *avcodec.CAVCodecContext
}

//#region members

func (avctx *AVCodecContext) GetAvClass() *avutil.CAVClass {
	return avctx.CAVCodecContext.GetAvClass()
}

func (avctx *AVCodecContext) GetLogLevelOffset() int {
	return avctx.CAVCodecContext.GetLogLevelOffset()
}
func (avctx *AVCodecContext) SetLogLevelOffset(logLevelOffset int) {
	avctx.CAVCodecContext.SetLogLevelOffset(logLevelOffset)
}

func (avctx *AVCodecContext) GetCodecType() avutil.CAVMediaType {
	return avctx.CAVCodecContext.GetCodecType()
}
func (avctx *AVCodecContext) SetCodecType(codecType avutil.CAVMediaType) {
	avctx.CAVCodecContext.SetCodecType(codecType)
}

func (avctx *AVCodecContext) GetCodec() *AVCodec {
	c := avctx.CAVCodecContext.GetCodec()
	if c == nil {
		return nil
	}
	return &AVCodec{
		CAVCodec: c,
	}
}

func (avctx *AVCodecContext) GetCodecId() avcodec.CAVCodecID {
	return avctx.CAVCodecContext.GetCodecId()
}
func (avctx *AVCodecContext) SetCodecId(codecId avcodec.CAVCodecID) {
	avctx.CAVCodecContext.SetCodecId(codecId)
}

func (avctx *AVCodecContext) GetCodecTag() uint {
	return avctx.CAVCodecContext.GetCodecTag()
}
func (avctx *AVCodecContext) SetCodecTag(codecTag uint) {
	avctx.CAVCodecContext.SetCodecTag(codecTag)
}

func (avctx *AVCodecContext) GetPrivData() unsafe.Pointer {
	return avctx.CAVCodecContext.GetPrivData()
}
func (avctx *AVCodecContext) SetPrivData(privData unsafe.Pointer) {
	avctx.CAVCodecContext.SetPrivData(privData)
}

// func (avctx *AVCodecContext) GetInternal() *C.struct_AVCodecInternal {
// 	return avctx.CAVCodecContext.GetInternal()
// }
// func (avctx *AVCodecContext) SetInternal(internal *C.struct_AVCodecInternal) {
// 	avctx.CAVCodecContext.SetInternal(internal)
// }

func (avctx *AVCodecContext) GetOpaque() unsafe.Pointer {
	return avctx.CAVCodecContext.GetOpaque()
}
func (avctx *AVCodecContext) SetOpaque(opaque unsafe.Pointer) {
	avctx.CAVCodecContext.SetOpaque(opaque)
}

func (avctx *AVCodecContext) GetBitRate() int64 {
	return avctx.CAVCodecContext.GetBitRate()
}
func (avctx *AVCodecContext) SetBitRate(bitRate int64) {
	avctx.CAVCodecContext.SetBitRate(bitRate)
}

func (avctx *AVCodecContext) GetBitRateTolerance() int {
	return avctx.CAVCodecContext.GetBitRateTolerance()
}
func (avctx *AVCodecContext) SetBitRateTolerance(bitRateTolerance int) {
	avctx.CAVCodecContext.SetBitRateTolerance(bitRateTolerance)
}

func (avctx *AVCodecContext) GetGlobalQuality() int {
	return avctx.CAVCodecContext.GetGlobalQuality()
}
func (avctx *AVCodecContext) SetGlobalQuality(globalQuality int) {
	avctx.CAVCodecContext.SetGlobalQuality(globalQuality)
}

func (avctx *AVCodecContext) GetCompressionLevel() int {
	return avctx.CAVCodecContext.GetCompressionLevel()
}
func (avctx *AVCodecContext) SetCompressionLevel(compressionLevel int) {
	avctx.CAVCodecContext.SetCompressionLevel(compressionLevel)
}

func (avctx *AVCodecContext) GetFlags() int {
	return avctx.CAVCodecContext.GetFlags()
}
func (avctx *AVCodecContext) SetFlags(flags int) {
	avctx.CAVCodecContext.SetFlags(flags)
}

func (avctx *AVCodecContext) GetFlags2() int {
	return avctx.CAVCodecContext.GetFlags2()
}
func (avctx *AVCodecContext) SetFlags2(flags2 int) {
	avctx.CAVCodecContext.SetFlags2(flags2)
}

func (avctx *AVCodecContext) GetExtradata() unsafe.Pointer {
	return avctx.CAVCodecContext.GetExtradata()
}
func (avctx *AVCodecContext) SetExtradata(extradata unsafe.Pointer) {
	avctx.CAVCodecContext.SetExtradata(extradata)
}

func (avctx *AVCodecContext) GetExtradataSize() int {
	return avctx.CAVCodecContext.GetExtradataSize()
}
func (avctx *AVCodecContext) SetExtradataSize(extradataSize int) {
	avctx.CAVCodecContext.SetExtradataSize(extradataSize)
}

func (avctx *AVCodecContext) GetTimeBase() avutil.CAVRational {
	return avctx.CAVCodecContext.GetTimeBase()
}
func (avctx *AVCodecContext) SetTimeBase(timeBase avutil.CAVRational) {
	avctx.CAVCodecContext.SetTimeBase(timeBase)
}

func (avctx *AVCodecContext) GetDelay() int {
	return avctx.CAVCodecContext.GetDelay()
}
func (avctx *AVCodecContext) SetDelay(delay int) {
	avctx.CAVCodecContext.SetDelay(delay)
}

func (avctx *AVCodecContext) GetWidth() int {
	return avctx.CAVCodecContext.GetWidth()
}
func (avctx *AVCodecContext) SetWidth(width int) {
	avctx.CAVCodecContext.SetWidth(width)
}

func (avctx *AVCodecContext) GetHeight() int {
	return avctx.CAVCodecContext.GetHeight()
}
func (avctx *AVCodecContext) SetHeight(height int) {
	avctx.CAVCodecContext.SetHeight(height)
}

func (avctx *AVCodecContext) GetCodedWidth() int {
	return avctx.CAVCodecContext.GetCodedWidth()
}
func (avctx *AVCodecContext) SetCodedWidth(codedWidth int) {
	avctx.CAVCodecContext.SetCodedWidth(codedWidth)
}

func (avctx *AVCodecContext) GetCodedHeight() int {
	return avctx.CAVCodecContext.GetCodedHeight()
}
func (avctx *AVCodecContext) SetCodedHeight(codedHeight int) {
	avctx.CAVCodecContext.SetCodedHeight(codedHeight)
}

func (avctx *AVCodecContext) GetGopSize() int {
	return avctx.CAVCodecContext.GetGopSize()
}
func (avctx *AVCodecContext) SetGopSize(gopSize int) {
	avctx.CAVCodecContext.SetGopSize(gopSize)
}

func (avctx *AVCodecContext) GetPixFmt() avutil.CAVPixelFormat {
	return avctx.CAVCodecContext.GetPixFmt()
}
func (avctx *AVCodecContext) SetPixFmt(pixFmt avutil.CAVPixelFormat) {
	avctx.CAVCodecContext.SetPixFmt(pixFmt)
}

func (avctx *AVCodecContext) GetDrawHorizBand() ctypes.CFunc {
	return avctx.CAVCodecContext.GetDrawHorizBand()
}

func (avctx *AVCodecContext) SetDrawHorizBand(drawHorizBand ctypes.CFunc) {
	avctx.CAVCodecContext.SetDrawHorizBand(drawHorizBand)
}

func (avctx *AVCodecContext) GetGetFormat() ctypes.CFunc {
	return avctx.CAVCodecContext.GetGetFormat()
}

func (avctx *AVCodecContext) SetGetFormat(getFormat ctypes.CFunc) {
	avctx.CAVCodecContext.SetGetFormat(getFormat)
}

func (avctx *AVCodecContext) GetMaxBFrames() int {
	return avctx.CAVCodecContext.GetMaxBFrames()
}
func (avctx *AVCodecContext) SetMaxBFrames(maxBFrames int) {
	avctx.CAVCodecContext.SetMaxBFrames(maxBFrames)
}

func (avctx *AVCodecContext) GetBQuantFactor() float32 {
	return avctx.CAVCodecContext.GetBQuantFactor()
}
func (avctx *AVCodecContext) SetBQuantFactor(bQuantFactor float32) {
	avctx.CAVCodecContext.SetBQuantFactor(bQuantFactor)
}

func (avctx *AVCodecContext) GetBQuantOffset() float32 {
	return avctx.CAVCodecContext.GetBQuantOffset()
}
func (avctx *AVCodecContext) SetBQuantOffset(bQuantOffset float32) {
	avctx.CAVCodecContext.SetBQuantOffset(bQuantOffset)
}

func (avctx *AVCodecContext) GetHasBFrames() int {
	return avctx.CAVCodecContext.GetHasBFrames()
}
func (avctx *AVCodecContext) SetHasBFrames(hasBFrames int) {
	avctx.CAVCodecContext.SetHasBFrames(hasBFrames)
}

func (avctx *AVCodecContext) GetIQuantFactor() float32 {
	return avctx.CAVCodecContext.GetIQuantFactor()
}
func (avctx *AVCodecContext) SetIQuantFactor(iQuantFactor float32) {
	avctx.CAVCodecContext.SetIQuantFactor(iQuantFactor)
}

func (avctx *AVCodecContext) GetIQuantOffset() float32 {
	return avctx.CAVCodecContext.GetIQuantOffset()
}
func (avctx *AVCodecContext) SetIQuantOffset(iQuantOffset float32) {
	avctx.CAVCodecContext.SetIQuantOffset(iQuantOffset)
}

func (avctx *AVCodecContext) GetLumiMasking() float32 {
	return avctx.CAVCodecContext.GetLumiMasking()
}
func (avctx *AVCodecContext) SetLumiMasking(lumiMasking float32) {
	avctx.CAVCodecContext.SetLumiMasking(lumiMasking)
}

func (avctx *AVCodecContext) GetTemporalCplxMasking() float32 {
	return avctx.CAVCodecContext.GetTemporalCplxMasking()
}
func (avctx *AVCodecContext) SetTemporalCplxMasking(temporalCplxMasking float32) {
	avctx.CAVCodecContext.SetTemporalCplxMasking(temporalCplxMasking)
}

func (avctx *AVCodecContext) GetSpatialCplxMasking() float32 {
	return avctx.CAVCodecContext.GetSpatialCplxMasking()
}
func (avctx *AVCodecContext) SetSpatialCplxMasking(spatialCplxMasking float32) {
	avctx.CAVCodecContext.SetSpatialCplxMasking(spatialCplxMasking)
}

func (avctx *AVCodecContext) GetPMasking() float32 {
	return avctx.CAVCodecContext.GetPMasking()
}
func (avctx *AVCodecContext) SetPMasking(pMasking float32) {
	avctx.CAVCodecContext.SetPMasking(pMasking)
}

func (avctx *AVCodecContext) GetDarkMasking() float32 {
	return avctx.CAVCodecContext.GetDarkMasking()
}
func (avctx *AVCodecContext) SetDarkMasking(darkMasking float32) {
	avctx.CAVCodecContext.SetDarkMasking(darkMasking)
}

func (avctx *AVCodecContext) GetSampleAspectRatio() avutil.CAVRational {
	return avctx.CAVCodecContext.GetSampleAspectRatio()
}
func (avctx *AVCodecContext) SetSampleAspectRatio(sampleAspectRatio avutil.CAVRational) {
	avctx.CAVCodecContext.SetSampleAspectRatio(sampleAspectRatio)
}

func (avctx *AVCodecContext) GetMeCmp() int {
	return avctx.CAVCodecContext.GetMeCmp()
}
func (avctx *AVCodecContext) SetMeCmp(meCmp int) {
	avctx.CAVCodecContext.SetMeCmp(meCmp)
}

func (avctx *AVCodecContext) GetMeSubCmp() int {
	return avctx.CAVCodecContext.GetMeSubCmp()
}
func (avctx *AVCodecContext) SetMeSubCmp(meSubCmp int) {
	avctx.CAVCodecContext.SetMeSubCmp(meSubCmp)
}

func (avctx *AVCodecContext) GetMbCmp() int {
	return avctx.CAVCodecContext.GetMbCmp()
}
func (avctx *AVCodecContext) SetMbCmp(mbCmp int) {
	avctx.CAVCodecContext.SetMbCmp(mbCmp)
}

func (avctx *AVCodecContext) GetIldctCmp() int {
	return avctx.CAVCodecContext.GetIldctCmp()
}
func (avctx *AVCodecContext) SetIldctCmp(ildctCmp int) {
	avctx.CAVCodecContext.SetIldctCmp(ildctCmp)
}

func (avctx *AVCodecContext) GetDiaSize() int {
	return avctx.CAVCodecContext.GetDiaSize()
}
func (avctx *AVCodecContext) SetDiaSize(diaSize int) {
	avctx.CAVCodecContext.SetDiaSize(diaSize)
}

func (avctx *AVCodecContext) GetLastPredictorCount() int {
	return avctx.CAVCodecContext.GetLastPredictorCount()
}
func (avctx *AVCodecContext) SetLastPredictorCount(lastPredictorCount int) {
	avctx.CAVCodecContext.SetLastPredictorCount(lastPredictorCount)
}

func (avctx *AVCodecContext) GetMePreCmp() int {
	return avctx.CAVCodecContext.GetMePreCmp()
}
func (avctx *AVCodecContext) SetMePreCmp(mePreCmp int) {
	avctx.CAVCodecContext.SetMePreCmp(mePreCmp)
}

func (avctx *AVCodecContext) GetPreDiaSize() int {
	return avctx.CAVCodecContext.GetPreDiaSize()
}
func (avctx *AVCodecContext) SetPreDiaSize(preDiaSize int) {
	avctx.CAVCodecContext.SetPreDiaSize(preDiaSize)
}

func (avctx *AVCodecContext) GetMeSubpelQuality() int {
	return avctx.CAVCodecContext.GetMeSubpelQuality()
}
func (avctx *AVCodecContext) SetMeSubpelQuality(meSubpelQuality int) {
	avctx.CAVCodecContext.SetMeSubpelQuality(meSubpelQuality)
}

func (avctx *AVCodecContext) GetMeRange() int {
	return avctx.CAVCodecContext.GetMeRange()
}
func (avctx *AVCodecContext) SetMeRange(meRange int) {
	avctx.CAVCodecContext.SetMeRange(meRange)
}

func (avctx *AVCodecContext) GetSliceFlags() int {
	return avctx.CAVCodecContext.GetSliceFlags()
}
func (avctx *AVCodecContext) SetSliceFlags(sliceFlags int) {
	avctx.CAVCodecContext.SetSliceFlags(sliceFlags)
}

func (avctx *AVCodecContext) GetMbDecision() int {
	return avctx.CAVCodecContext.GetMbDecision()
}
func (avctx *AVCodecContext) SetMbDecision(mbDecision int) {
	avctx.CAVCodecContext.SetMbDecision(mbDecision)
}

func (avctx *AVCodecContext) GetIntraMatrix() *ctypes.UInt16 {
	return avctx.CAVCodecContext.GetIntraMatrix()
}
func (avctx *AVCodecContext) SetIntraMatrix(intraMatrix *ctypes.UInt16) {
	avctx.CAVCodecContext.SetIntraMatrix(intraMatrix)
}

func (avctx *AVCodecContext) GetInterMatrix() *ctypes.UInt16 {
	return avctx.CAVCodecContext.GetInterMatrix()
}
func (avctx *AVCodecContext) SetInterMatrix(interMatrix *ctypes.UInt16) {
	avctx.CAVCodecContext.SetInterMatrix(interMatrix)
}

func (avctx *AVCodecContext) GetIntraDcPrecision() int {
	return avctx.CAVCodecContext.GetIntraDcPrecision()
}
func (avctx *AVCodecContext) SetIntraDcPrecision(intraDcPrecision int) {
	avctx.CAVCodecContext.SetIntraDcPrecision(intraDcPrecision)
}

func (avctx *AVCodecContext) GetSkipTop() int {
	return avctx.CAVCodecContext.GetSkipTop()
}
func (avctx *AVCodecContext) SetSkipTop(skipTop int) {
	avctx.CAVCodecContext.SetSkipTop(skipTop)
}

func (avctx *AVCodecContext) GetSkipBottom() int {
	return avctx.CAVCodecContext.GetSkipBottom()
}
func (avctx *AVCodecContext) SetSkipBottom(skipBottom int) {
	avctx.CAVCodecContext.SetSkipBottom(skipBottom)
}

func (avctx *AVCodecContext) GetMbLmin() int {
	return avctx.CAVCodecContext.GetMbLmin()
}
func (avctx *AVCodecContext) SetMbLmin(mbLmin int) {
	avctx.CAVCodecContext.SetMbLmin(mbLmin)
}

func (avctx *AVCodecContext) GetMbLmax() int {
	return avctx.CAVCodecContext.GetMbLmax()
}
func (avctx *AVCodecContext) SetMbLmax(mbLmax int) {
	avctx.CAVCodecContext.SetMbLmax(mbLmax)
}

func (avctx *AVCodecContext) GetBidirRefine() int {
	return avctx.CAVCodecContext.GetBidirRefine()
}
func (avctx *AVCodecContext) SetBidirRefine(bidirRefine int) {
	avctx.CAVCodecContext.SetBidirRefine(bidirRefine)
}

func (avctx *AVCodecContext) GetKeyintMin() int {
	return avctx.CAVCodecContext.GetKeyintMin()
}
func (avctx *AVCodecContext) SetKeyintMin(keyintMin int) {
	avctx.CAVCodecContext.SetKeyintMin(keyintMin)
}

func (avctx *AVCodecContext) GetRefs() int {
	return avctx.CAVCodecContext.GetRefs()
}
func (avctx *AVCodecContext) SetRefs(refs int) {
	avctx.CAVCodecContext.SetRefs(refs)
}

func (avctx *AVCodecContext) GetMv0Threshold() int {
	return avctx.CAVCodecContext.GetMv0Threshold()
}
func (avctx *AVCodecContext) SetMv0Threshold(mv0Threshold int) {
	avctx.CAVCodecContext.SetMv0Threshold(mv0Threshold)
}

func (avctx *AVCodecContext) GetColorPrimaries() avutil.CAVColorPrimaries {
	return avctx.CAVCodecContext.GetColorPrimaries()
}
func (avctx *AVCodecContext) SetColorPrimaries(colorPrimaries avutil.CAVColorPrimaries) {
	avctx.CAVCodecContext.SetColorPrimaries(colorPrimaries)
}

func (avctx *AVCodecContext) GetColorTrc() avutil.CAVColorTransferCharacteristic {
	return avctx.CAVCodecContext.GetColorTrc()
}
func (avctx *AVCodecContext) SetColorTrc(colorTrc avutil.CAVColorTransferCharacteristic) {
	avctx.CAVCodecContext.SetColorTrc(colorTrc)
}

func (avctx *AVCodecContext) GetColorspace() avutil.CAVColorSpace {
	return avctx.CAVCodecContext.GetColorspace()
}
func (avctx *AVCodecContext) SetColorspace(colorspace avutil.CAVColorSpace) {
	avctx.CAVCodecContext.SetColorspace(colorspace)
}

func (avctx *AVCodecContext) GetColorRange() avutil.CAVColorRange {
	return avctx.CAVCodecContext.GetColorRange()
}
func (avctx *AVCodecContext) SetColorRange(colorRange avutil.CAVColorRange) {
	avctx.CAVCodecContext.SetColorRange(colorRange)
}

func (avctx *AVCodecContext) GetSampleChromaLocation() avutil.CAVChromaLocation {
	return avctx.CAVCodecContext.GetSampleChromaLocation()
}
func (avctx *AVCodecContext) SetSampleChromaLocation(chromaLocation avutil.CAVChromaLocation) {
	avctx.CAVCodecContext.SetSampleChromaLocation(chromaLocation)
}

func (avctx *AVCodecContext) GetSlices() int {
	return avctx.CAVCodecContext.GetSlices()
}
func (avctx *AVCodecContext) SetSlices(slices int) {
	avctx.CAVCodecContext.SetSlices(slices)
}

func (avctx *AVCodecContext) GetFieldOrder() avcodec.CAVFieldOrder {
	return avctx.CAVCodecContext.GetFieldOrder()
}
func (avctx *AVCodecContext) SetFieldOrder(fieldOrder avcodec.CAVFieldOrder) {
	avctx.CAVCodecContext.SetFieldOrder(fieldOrder)
}

func (avctx *AVCodecContext) GetSampleRate() int {
	return avctx.CAVCodecContext.GetSampleRate()
}
func (avctx *AVCodecContext) SetSampleRate(sampleRate int) {
	avctx.CAVCodecContext.SetSampleRate(sampleRate)
}

func (avctx *AVCodecContext) GetSampleFmt() avutil.CAVSampleFormat {
	return avctx.CAVCodecContext.GetSampleFmt()
}
func (avctx *AVCodecContext) SetSampleFmt(sampleFmt avutil.CAVSampleFormat) {
	avctx.CAVCodecContext.SetSampleFmt(sampleFmt)
}

func (avctx *AVCodecContext) GetFrameSize() int {
	return avctx.CAVCodecContext.GetFrameSize()
}
func (avctx *AVCodecContext) SetFrameSize(frameSize int) {
	avctx.CAVCodecContext.SetFrameSize(frameSize)
}

func (avctx *AVCodecContext) GetBlockAlign() int {
	return avctx.CAVCodecContext.GetBlockAlign()
}
func (avctx *AVCodecContext) SetBlockAlign(blockAlign int) {
	avctx.CAVCodecContext.SetBlockAlign(blockAlign)
}

func (avctx *AVCodecContext) GetCutoff() int {
	return avctx.CAVCodecContext.GetCutoff()
}
func (avctx *AVCodecContext) SetCutoff(cutoff int) {
	avctx.CAVCodecContext.SetCutoff(cutoff)
}

func (avctx *AVCodecContext) GetAudioServiceType() avcodec.CAVAudioServiceType {
	return avctx.CAVCodecContext.GetAudioServiceType()
}
func (avctx *AVCodecContext) SetAudioServiceType(audioServiceType avcodec.CAVAudioServiceType) {
	avctx.CAVCodecContext.SetAudioServiceType(audioServiceType)
}

func (avctx *AVCodecContext) GetRequestSampleFmt() avutil.CAVSampleFormat {
	return avctx.CAVCodecContext.GetRequestSampleFmt()
}
func (avctx *AVCodecContext) SetRequestSampleFmt(requestSampleFmt avutil.CAVSampleFormat) {
	avctx.CAVCodecContext.SetRequestSampleFmt(requestSampleFmt)
}

func (avctx *AVCodecContext) GetGetBuffer2() ctypes.CFunc {
	return avctx.CAVCodecContext.GetGetBuffer2()
}
func (avctx *AVCodecContext) SetGetBuffer2(getBuffer2 ctypes.CFunc) {
	avctx.CAVCodecContext.SetGetBuffer2(getBuffer2)
}

func (avctx *AVCodecContext) GetQcompress() float32 {
	return avctx.CAVCodecContext.GetQcompress()
}
func (avctx *AVCodecContext) SetQcompress(qcompress float32) {
	avctx.CAVCodecContext.SetQcompress(qcompress)
}

func (avctx *AVCodecContext) GetQblur() float32 {
	return avctx.CAVCodecContext.GetQblur()
}
func (avctx *AVCodecContext) SetQblur(qblur float32) {
	avctx.CAVCodecContext.SetQblur(qblur)
}

func (avctx *AVCodecContext) GetQmin() int {
	return avctx.CAVCodecContext.GetQmin()
}
func (avctx *AVCodecContext) SetQmin(qmin int) {
	avctx.CAVCodecContext.SetQmin(qmin)
}

func (avctx *AVCodecContext) GetQmax() int {
	return avctx.CAVCodecContext.GetQmax()
}
func (avctx *AVCodecContext) SetQmax(qmax int) {
	avctx.CAVCodecContext.SetQmax(qmax)
}

func (avctx *AVCodecContext) GetMaxQdiff() int {
	return avctx.CAVCodecContext.GetMaxQdiff()
}
func (avctx *AVCodecContext) SetMaxQdiff(maxQdiff int) {
	avctx.CAVCodecContext.SetMaxQdiff(maxQdiff)
}

func (avctx *AVCodecContext) GetRcBufferSize() int {
	return avctx.CAVCodecContext.GetRcBufferSize()
}
func (avctx *AVCodecContext) SetRcBufferSize(rcBufferSize int) {
	avctx.CAVCodecContext.SetRcBufferSize(rcBufferSize)
}

func (avctx *AVCodecContext) GetRcOverrideCount() int {
	return avctx.CAVCodecContext.GetRcOverrideCount()
}
func (avctx *AVCodecContext) SetRcOverrideCount(rcOverrideCount int) {
	avctx.CAVCodecContext.SetRcOverrideCount(rcOverrideCount)
}

func (avctx *AVCodecContext) GetRcOverride() *avcodec.CRcOverride {
	return avctx.CAVCodecContext.GetRcOverride()
}
func (avctx *AVCodecContext) SetRcOverride(rcOverride *avcodec.CRcOverride) {
	avctx.CAVCodecContext.SetRcOverride(rcOverride)
}

func (avctx *AVCodecContext) GetRcMaxRate() int64 {
	return avctx.CAVCodecContext.GetRcMaxRate()
}
func (avctx *AVCodecContext) SetRcMaxRate(rcMaxRate int64) {
	avctx.CAVCodecContext.SetRcMaxRate(rcMaxRate)
}

func (avctx *AVCodecContext) GetRcMinRate() int64 {
	return avctx.CAVCodecContext.GetRcMinRate()
}
func (avctx *AVCodecContext) SetRcMinRate(rcMinRate int64) {
	avctx.CAVCodecContext.SetRcMinRate(rcMinRate)
}

func (avctx *AVCodecContext) GetRcMaxAvailableVbvUse() float32 {
	return avctx.CAVCodecContext.GetRcMaxAvailableVbvUse()
}
func (avctx *AVCodecContext) SetRcMaxAvailableVbvUse(rcMaxAvailableVbvUse float32) {
	avctx.CAVCodecContext.SetRcMaxAvailableVbvUse(rcMaxAvailableVbvUse)
}

func (avctx *AVCodecContext) GetRcMinVbvOverflowUse() float32 {
	return avctx.CAVCodecContext.GetRcMinVbvOverflowUse()
}
func (avctx *AVCodecContext) SetRcMinVbvOverflowUse(rcMinVbvOverflowUse float32) {
	avctx.CAVCodecContext.SetRcMinVbvOverflowUse(rcMinVbvOverflowUse)
}

func (avctx *AVCodecContext) GetRcInitialBufferOccupancy() int {
	return avctx.CAVCodecContext.GetRcInitialBufferOccupancy()
}
func (avctx *AVCodecContext) SetRcInitialBufferOccupancy(rcInitialBufferOccupancy int) {
	avctx.CAVCodecContext.SetRcInitialBufferOccupancy(rcInitialBufferOccupancy)
}

func (avctx *AVCodecContext) GetTrellis() int {
	return avctx.CAVCodecContext.GetTrellis()
}
func (avctx *AVCodecContext) SetTrellis(trellis int) {
	avctx.CAVCodecContext.SetTrellis(trellis)
}

func (avctx *AVCodecContext) GetStatsOut() string {
	return avctx.CAVCodecContext.GetStatsOut()
}
func (avctx *AVCodecContext) SetStatsOut(statsOut string) {
	avctx.CAVCodecContext.SetStatsOut(statsOut)
}

func (avctx *AVCodecContext) GetStatsIn() string {
	return avctx.CAVCodecContext.GetStatsIn()
}
func (avctx *AVCodecContext) SetStatsIn(statsIn string) {
	avctx.CAVCodecContext.SetStatsIn(statsIn)
}

func (avctx *AVCodecContext) GetWorkaroundBugs() int {
	return avctx.CAVCodecContext.GetWorkaroundBugs()
}
func (avctx *AVCodecContext) SetWorkaroundBugs(workaroundBugs int) {
	avctx.CAVCodecContext.SetWorkaroundBugs(workaroundBugs)
}

func (avctx *AVCodecContext) GetStrictStdCompliance() int {
	return avctx.CAVCodecContext.GetStrictStdCompliance()
}
func (avctx *AVCodecContext) SetStrictStdCompliance(strictStdCompliance int) {
	avctx.CAVCodecContext.SetStrictStdCompliance(strictStdCompliance)
}

func (avctx *AVCodecContext) GetErrorConcealment() int {
	return avctx.CAVCodecContext.GetErrorConcealment()
}
func (avctx *AVCodecContext) SetErrorConcealment(errorConcealment int) {
	avctx.CAVCodecContext.SetErrorConcealment(errorConcealment)
}

func (avctx *AVCodecContext) GetDebug() int {
	return avctx.CAVCodecContext.GetDebug()
}
func (avctx *AVCodecContext) SetDebug(debug int) {
	avctx.CAVCodecContext.SetDebug(debug)
}

func (avctx *AVCodecContext) GetErrRecognition() int {
	return avctx.CAVCodecContext.GetErrRecognition()
}
func (avctx *AVCodecContext) SetErrRecognition(errRecognition int) {
	avctx.CAVCodecContext.SetErrRecognition(errRecognition)
}

func (avctx *AVCodecContext) GetHwaccel() *avcodec.CAVHWAccel {
	return avctx.CAVCodecContext.GetHwaccel()
}

func (avctx *AVCodecContext) GetHwaccelContext() unsafe.Pointer {
	return avctx.CAVCodecContext.GetHwaccelContext()
}
func (avctx *AVCodecContext) SetHwaccelContext(hwaccelContext unsafe.Pointer) {
	avctx.CAVCodecContext.SetHwaccelContext(hwaccelContext)
}

func (avctx *AVCodecContext) GetError() [avutil.AV_NUM_DATA_POINTERS]ctypes.UInt64 {
	return avctx.CAVCodecContext.GetError()
}
func (avctx *AVCodecContext) SetError(_error [avutil.AV_NUM_DATA_POINTERS]ctypes.UInt64) {
	avctx.CAVCodecContext.SetError(_error)
}

func (avctx *AVCodecContext) GetDctAlgo() int {
	return avctx.CAVCodecContext.GetDctAlgo()
}
func (avctx *AVCodecContext) SetDctAlgo(dctAlgo int) {
	avctx.CAVCodecContext.SetDctAlgo(dctAlgo)
}

func (avctx *AVCodecContext) GetIdctAlgo() int {
	return avctx.CAVCodecContext.GetIdctAlgo()
}
func (avctx *AVCodecContext) SetIdctAlgo(idctAlgo int) {
	avctx.CAVCodecContext.SetIdctAlgo(idctAlgo)
}

func (avctx *AVCodecContext) GetBitsPerCodedSample() int {
	return avctx.CAVCodecContext.GetBitsPerCodedSample()
}
func (avctx *AVCodecContext) SetBitsPerCodedSample(bitsPerCodedSample int) {
	avctx.CAVCodecContext.SetBitsPerCodedSample(bitsPerCodedSample)
}

func (avctx *AVCodecContext) GetBitsPerRawSample() int {
	return avctx.CAVCodecContext.GetBitsPerRawSample()
}
func (avctx *AVCodecContext) SetBitsPerRawSample(bitsPerRawSample int) {
	avctx.CAVCodecContext.SetBitsPerRawSample(bitsPerRawSample)
}

func (avctx *AVCodecContext) GetLowres() int {
	return avctx.CAVCodecContext.GetLowres()
}
func (avctx *AVCodecContext) SetLowres(lowres int) {
	avctx.CAVCodecContext.SetLowres(lowres)
}

func (avctx *AVCodecContext) GetThreadCount() int {
	return avctx.CAVCodecContext.GetThreadCount()
}
func (avctx *AVCodecContext) SetThreadCount(threadCount int) {
	avctx.CAVCodecContext.SetThreadCount(threadCount)
}

func (avctx *AVCodecContext) GetThreadType() int {
	return avctx.CAVCodecContext.GetThreadType()
}
func (avctx *AVCodecContext) SetThreadType(threadType int) {
	avctx.CAVCodecContext.SetThreadType(threadType)
}

func (avctx *AVCodecContext) GetActiveThreadType() int {
	return avctx.CAVCodecContext.GetActiveThreadType()
}
func (avctx *AVCodecContext) SetActiveThreadType(activeThreadType int) {
	avctx.CAVCodecContext.SetActiveThreadType(activeThreadType)
}

func (avctx *AVCodecContext) GetExecute() ctypes.CFunc {
	return avctx.CAVCodecContext.GetExecute()
}
func (avctx *AVCodecContext) SetExecute(execute ctypes.CFunc) {
	avctx.CAVCodecContext.SetExecute(execute)
}

func (avctx *AVCodecContext) GetExecute2() ctypes.CFunc {
	return avctx.CAVCodecContext.GetExecute2()
}
func (avctx *AVCodecContext) SetExecute2(execute2 ctypes.CFunc) {
	avctx.CAVCodecContext.SetExecute2(execute2)
}

func (avctx *AVCodecContext) GetNsseWeight() int {
	return avctx.CAVCodecContext.GetNsseWeight()
}
func (avctx *AVCodecContext) SetNsseWeight(nsseWeight int) {
	avctx.CAVCodecContext.SetNsseWeight(nsseWeight)
}

func (avctx *AVCodecContext) GetProfile() int {
	return avctx.CAVCodecContext.GetProfile()
}
func (avctx *AVCodecContext) SetProfile(profile int) {
	avctx.CAVCodecContext.SetProfile(profile)
}

func (avctx *AVCodecContext) GetLevel() int {
	return avctx.CAVCodecContext.GetLevel()
}
func (avctx *AVCodecContext) SetLevel(level int) {
	avctx.CAVCodecContext.SetLevel(level)
}

func (avctx *AVCodecContext) GetSkipLoopFilter() avcodec.CAVDiscard {
	return avctx.CAVCodecContext.GetSkipLoopFilter()
}
func (avctx *AVCodecContext) SetSkipLoopFilter(skipLoopFilter avcodec.CAVDiscard) {
	avctx.CAVCodecContext.SetSkipLoopFilter(skipLoopFilter)
}

func (avctx *AVCodecContext) GetSkipIdct() avcodec.CAVDiscard {
	return avctx.CAVCodecContext.GetSkipIdct()
}
func (avctx *AVCodecContext) SetSkipIdct(skipIdct avcodec.CAVDiscard) {
	avctx.CAVCodecContext.SetSkipIdct(skipIdct)
}

func (avctx *AVCodecContext) GetSkipFrame() avcodec.CAVDiscard {
	return avctx.CAVCodecContext.GetSkipFrame()
}
func (avctx *AVCodecContext) SetSkipFrame(skipFrame avcodec.CAVDiscard) {
	avctx.CAVCodecContext.SetSkipFrame(skipFrame)
}

func (avctx *AVCodecContext) GetSubtitleHeader() unsafe.Pointer {
	return avctx.CAVCodecContext.GetSubtitleHeader()
}
func (avctx *AVCodecContext) SetSubtitleHeader(subtitleHeader unsafe.Pointer) {
	avctx.CAVCodecContext.SetSubtitleHeader(subtitleHeader)
}

func (avctx *AVCodecContext) GetSubtitleHeaderSize() int {
	return avctx.CAVCodecContext.GetSubtitleHeaderSize()
}
func (avctx *AVCodecContext) SetSubtitleHeaderSize(subtitleHeaderSize int) {
	avctx.CAVCodecContext.SetSubtitleHeaderSize(subtitleHeaderSize)
}

func (avctx *AVCodecContext) GetInitialPadding() int {
	return avctx.CAVCodecContext.GetInitialPadding()
}
func (avctx *AVCodecContext) SetInitialPadding(initialPadding int) {
	avctx.CAVCodecContext.SetInitialPadding(initialPadding)
}

func (avctx *AVCodecContext) GetFramerate() avutil.CAVRational {
	return avctx.CAVCodecContext.GetFramerate()
}
func (avctx *AVCodecContext) SetFramerate(framerate avutil.CAVRational) {
	avctx.CAVCodecContext.SetFramerate(framerate)
}

func (avctx *AVCodecContext) GetSwPixFmt() avutil.CAVPixelFormat {
	return avctx.CAVCodecContext.GetSwPixFmt()
}
func (avctx *AVCodecContext) SetSwPixFmt(swPixFmt avutil.CAVPixelFormat) {
	avctx.CAVCodecContext.SetSwPixFmt(swPixFmt)
}

func (avctx *AVCodecContext) GetPktTimebase() avutil.CAVRational {
	return avctx.CAVCodecContext.GetPktTimebase()
}
func (avctx *AVCodecContext) SetPktTimebase(pktTimebase avutil.CAVRational) {
	avctx.CAVCodecContext.SetPktTimebase(pktTimebase)
}

func (avctx *AVCodecContext) GetCodecDescriptor() *avcodec.CAVCodecDescriptor {
	return avctx.CAVCodecContext.GetCodecDescriptor()
}

func (avctx *AVCodecContext) GetPtsCorrectionNumFaultyPts() int64 {
	return avctx.CAVCodecContext.GetPtsCorrectionNumFaultyPts()
}
func (avctx *AVCodecContext) SetPtsCorrectionNumFaultyPts(ptsCorrectionNumFaultyPts int64) {
	avctx.CAVCodecContext.SetPtsCorrectionNumFaultyPts(ptsCorrectionNumFaultyPts)
}

func (avctx *AVCodecContext) GetPtsCorrectionNumFaultyDts() int64 {
	return avctx.CAVCodecContext.GetPtsCorrectionNumFaultyDts()
}
func (avctx *AVCodecContext) SetPtsCorrectionNumFaultyDts(ptsCorrectionNumFaultyDts int64) {
	avctx.CAVCodecContext.SetPtsCorrectionNumFaultyDts(ptsCorrectionNumFaultyDts)
}

func (avctx *AVCodecContext) GetPtsCorrectionLastPts() int64 {
	return avctx.CAVCodecContext.GetPtsCorrectionLastPts()
}
func (avctx *AVCodecContext) SetPtsCorrectionLastPts(ptsCorrectionLastPts int64) {
	avctx.CAVCodecContext.SetPtsCorrectionLastPts(ptsCorrectionLastPts)
}

func (avctx *AVCodecContext) GetPtsCorrectionLastDts() int64 {
	return avctx.CAVCodecContext.GetPtsCorrectionLastDts()
}
func (avctx *AVCodecContext) SetPtsCorrectionLastDts(ptsCorrectionLastDts int64) {
	avctx.CAVCodecContext.SetPtsCorrectionLastDts(ptsCorrectionLastDts)
}

func (avctx *AVCodecContext) GetSubCharenc() string {
	return avctx.CAVCodecContext.GetSubCharenc()
}
func (avctx *AVCodecContext) SetSubCharenc(subCharenc string) {
	avctx.CAVCodecContext.SetSubCharenc(subCharenc)
}

func (avctx *AVCodecContext) GetSubCharencMode() int {
	return avctx.CAVCodecContext.GetSubCharencMode()
}
func (avctx *AVCodecContext) SetSubCharencMode(subCharencMode int) {
	avctx.CAVCodecContext.SetSubCharencMode(subCharencMode)
}

func (avctx *AVCodecContext) GetSkipAlpha() int {
	return avctx.CAVCodecContext.GetSkipAlpha()
}
func (avctx *AVCodecContext) SetSkipAlpha(skipAlpha int) {
	avctx.CAVCodecContext.SetSkipAlpha(skipAlpha)
}

func (avctx *AVCodecContext) GetSeekPreroll() int {
	return avctx.CAVCodecContext.GetSeekPreroll()
}
func (avctx *AVCodecContext) SetSeekPreroll(seekPreroll int) {
	avctx.CAVCodecContext.SetSeekPreroll(seekPreroll)
}

func (avctx *AVCodecContext) GetChromaIntraMatrix() *ctypes.UInt16 {
	return avctx.CAVCodecContext.GetChromaIntraMatrix()
}
func (avctx *AVCodecContext) SetChromaIntraMatrix(chromaIntraMatrix *ctypes.UInt16) {
	avctx.CAVCodecContext.SetChromaIntraMatrix(chromaIntraMatrix)
}

func (avctx *AVCodecContext) GetDumpSeparator() string {
	return avctx.CAVCodecContext.GetDumpSeparator()
}
func (avctx *AVCodecContext) SetDumpSeparator(dumpSeparator string) {
	avctx.CAVCodecContext.SetDumpSeparator(dumpSeparator)
}

func (avctx *AVCodecContext) GetCodecWhitelist() string {
	return avctx.CAVCodecContext.GetCodecWhitelist()
}
func (avctx *AVCodecContext) SetCodecWhitelist(codecWhitelist string) {
	avctx.CAVCodecContext.SetCodecWhitelist(codecWhitelist)
}

func (avctx *AVCodecContext) GetProperties() uint {
	return avctx.CAVCodecContext.GetProperties()
}
func (avctx *AVCodecContext) SetProperties(properties uint) {
	avctx.CAVCodecContext.SetProperties(properties)
}

func (avctx *AVCodecContext) GetCodedSideData() *avcodec.CAVPacketSideData {
	return avctx.CAVCodecContext.GetCodedSideData()
}
func (avctx *AVCodecContext) SetCodedSideData(codedSideData *avcodec.CAVPacketSideData) {
	avctx.CAVCodecContext.SetCodedSideData(codedSideData)
}

func (avctx *AVCodecContext) GetNbCodedSideData() int {
	return avctx.CAVCodecContext.GetNbCodedSideData()
}
func (avctx *AVCodecContext) SetNbCodedSideData(nbCodedSideData int) {
	avctx.CAVCodecContext.SetNbCodedSideData(nbCodedSideData)
}

func (avctx *AVCodecContext) GetHwFramesCtx() *avutil.CAVBufferRef {
	return avctx.CAVCodecContext.GetHwFramesCtx()
}
func (avctx *AVCodecContext) SetHwFramesCtx(hwFramesCtx *avutil.CAVBufferRef) {
	avctx.CAVCodecContext.SetHwFramesCtx(hwFramesCtx)
}

func (avctx *AVCodecContext) GetTrailingPadding() int {
	return avctx.CAVCodecContext.GetTrailingPadding()
}
func (avctx *AVCodecContext) SetTrailingPadding(trailingPadding int) {
	avctx.CAVCodecContext.SetTrailingPadding(trailingPadding)
}

func (avctx *AVCodecContext) GetMaxPixels() int64 {
	return avctx.CAVCodecContext.GetMaxPixels()
}
func (avctx *AVCodecContext) SetMaxPixels(maxPixels int64) {
	avctx.CAVCodecContext.SetMaxPixels(maxPixels)
}

func (avctx *AVCodecContext) GetHwDeviceCtx() *avutil.CAVBufferRef {
	return avctx.CAVCodecContext.GetHwDeviceCtx()
}
func (avctx *AVCodecContext) SetHwDeviceCtx(hwDeviceCtx *avutil.CAVBufferRef) {
	avctx.CAVCodecContext.SetHwDeviceCtx(hwDeviceCtx)
}

func (avctx *AVCodecContext) GetHwaccelFlags() int {
	return avctx.CAVCodecContext.GetHwaccelFlags()
}
func (avctx *AVCodecContext) SetHwaccelFlags(hwaccelFlags int) {
	avctx.CAVCodecContext.SetHwaccelFlags(hwaccelFlags)
}

func (avctx *AVCodecContext) GetApplyCropping() int {
	return avctx.CAVCodecContext.GetApplyCropping()
}
func (avctx *AVCodecContext) SetApplyCropping(applyCropping int) {
	avctx.CAVCodecContext.SetApplyCropping(applyCropping)
}

func (avctx *AVCodecContext) GetExtraHwFrames() int {
	return avctx.CAVCodecContext.GetExtraHwFrames()
}
func (avctx *AVCodecContext) SetExtraHwFrames(extraHwFrames int) {
	avctx.CAVCodecContext.SetExtraHwFrames(extraHwFrames)
}

func (avctx *AVCodecContext) GetDiscardDamagedPercentage() int {
	return avctx.CAVCodecContext.GetDiscardDamagedPercentage()
}
func (avctx *AVCodecContext) SetDiscardDamagedPercentage(discardDamagedPercentage int) {
	avctx.CAVCodecContext.SetDiscardDamagedPercentage(discardDamagedPercentage)
}

func (avctx *AVCodecContext) GetMaxSamples() int64 {
	return avctx.CAVCodecContext.GetMaxSamples()
}
func (avctx *AVCodecContext) SetMaxSamples(maxSamples int64) {
	avctx.CAVCodecContext.SetMaxSamples(maxSamples)
}

func (avctx *AVCodecContext) GetExportSideData() int {
	return avctx.CAVCodecContext.GetExportSideData()
}
func (avctx *AVCodecContext) SetExportSideData(exportSideData int) {
	avctx.CAVCodecContext.SetExportSideData(exportSideData)
}

func (avctx *AVCodecContext) GetGetEncodeBuffer() ctypes.CFunc {
	return avctx.CAVCodecContext.GetGetEncodeBuffer()
}
func (avctx *AVCodecContext) SetGetEncodeBuffer(getEncodeBuffer ctypes.CFunc) {
	avctx.CAVCodecContext.SetGetEncodeBuffer(getEncodeBuffer)
}

func (avctx *AVCodecContext) GetChLayout() *goavutil.AVChannelLayout {
	return &goavutil.AVChannelLayout{
		CAVChannelLayout: avctx.CAVCodecContext.GetChLayoutPtr(),
	}
}
func (avctx *AVCodecContext) SetChLayout(chLayout *goavutil.AVChannelLayout) {
	avctx.CAVCodecContext.SetChLayout(*chLayout.CAVChannelLayout)
}

func (avctx *AVCodecContext) GetFrameNum() int64 {
	return avctx.CAVCodecContext.GetFrameNum()
}
func (avctx *AVCodecContext) SetFrameNum(frameNum int64) {
	avctx.CAVCodecContext.SetFrameNum(frameNum)
}

//#endregion members

func NewContext(codec *AVCodec) *AVCodecContext {
	c := avcodec.AvcodecAllocContext3(codec.CAVCodec)
	if c == nil {
		return nil
	}
	return &AVCodecContext{
		CAVCodecContext: c,
	}
}

func (avctx *AVCodecContext) Free() {
	avcodec.AvcodecFreeContext(&avctx.CAVCodecContext)
}

func (avctx *AVCodecContext) SendPacket(pkt *AVPacket) int {
	var cPkt *avcodec.CAVPacket
	if pkt != nil {
		cPkt = pkt.CAVPacket
	}
	return avcodec.AvcodecSendPacket(avctx.CAVCodecContext, cPkt)
}

func (avctx *AVCodecContext) ReceiveFrame(frame *goavutil.AVFrame) int {
	return avcodec.AvcodecReceiveFrame(avctx.CAVCodecContext, frame.CAVFrame)
}

func (avctx *AVCodecContext) SendFrame(frame *goavutil.AVFrame) int {
	var cFrame *avutil.CAVFrame
	if frame != nil {
		cFrame = frame.CAVFrame
	}
	return avcodec.AvcodecSendFrame(avctx.CAVCodecContext, cFrame)
}

func (avctx *AVCodecContext) ReceivePacket(pkt *AVPacket) int {
	return avcodec.AvcodecReceivePacket(avctx.CAVCodecContext, pkt.CAVPacket)
}

// Decode AVPacket.
// Return a slice of decoded AVFrame, May be nil or empty.
func (avctx *AVCodecContext) Decode(pkt *AVPacket) ([]*goavutil.AVFrame, int) {
	code := avctx.SendPacket(pkt)
	if code < 0 {
		return nil, code
	}

	frames := make([]*goavutil.AVFrame, 0)

	for {
		frame := goavutil.AllocAVFrame()
		if frame == nil {
			code = int(syscall.ENOMEM)
			break
		}

		code = avctx.ReceiveFrame(frame)
		if code < 0 {
			frame.Free()
			break
		}

		frames = append(frames, frame)
	}

	return frames, code
}

// Decode AVPacket from channel.
//
// @param ctx The context.Context to cancel this goroutine.
// @param pktChan The AVPacket channel.
//
// @return context.Context The context.Context to get the cause of this gorouine.
// @return <-chan *goavutil.AVFrame The channel to read decoded AVFrame.
func (avctx *AVCodecContext) DecodeChan(ctx context.Context, pktChan <-chan *AVPacket) (context.Context, <-chan *goavutil.AVFrame) {
	ctx, cancel := context.WithCancelCause(ctx)
	frameChan := make(chan *goavutil.AVFrame)

	go func() {
		defer close(frameChan)
	loop:
		for {
			select {
			case <-ctx.Done():
				return
			case pkt, ok := <-pktChan:
				if !ok {
					break loop
				}
				frames, code := avctx.Decode(pkt)
				pkt.Unref()
				for _, frame := range frames {
					select {
					case <-ctx.Done():
						for _, frame := range frames {
							frame.Unref()
						}
						return
					case frameChan <- frame:
					}
				}
				if code == avutil.AVERROR(int(syscall.EAGAIN)) || code == avutil.AVERROR_EOF {
					continue
				}
				if code < 0 {
					cancel(goavutil.AvErr(code))
					return
				}
			}
		}

		//flush codec
		frames, code := avctx.Decode(nil)
		for _, frame := range frames {
			select {
			case <-ctx.Done():
				for _, frame := range frames {
					frame.Unref()
				}
				return
			case frameChan <- frame:
			}
		}
		if code == avutil.AVERROR(int(syscall.EAGAIN)) || code == avutil.AVERROR_EOF {
			return
		}
		if code < 0 {
			cancel(goavutil.AvErr(code))
			return
		}
	}()

	return ctx, frameChan
}

// ChunkingFrameChan chunking read frame by AVAudioFifo.
//
// Example:
//
//	if avctx.CAVCodecContext.GetCodec().GetCapabilities()&avcodec.AV_CODEC_CAP_VARIABLE_FRAME_SIZE == 0 {
//		ctx, frameChan = avctx.chunkingFrameChan(ctx, frameChan)
//	}
//	avctx.EncodeChan(ctx, frameChan)
func (avctx *AVCodecContext) ChunkingFrameChan(ctx context.Context, frameInChan <-chan *goavutil.AVFrame) (context.Context, <-chan *goavutil.AVFrame) {
	ctx, cancel := context.WithCancelCause(ctx)
	frameOutChan := make(chan *goavutil.AVFrame)

	af := goavutil.NewAVAudioFifo(avctx.CAVCodecContext.GetSampleFmt(), avctx.GetChLayout().GetNbChannels(), 1)
	if af == nil {
		cancel(goavutil.AvErr(avutil.AVERROR(int(syscall.ENOMEM))))
		return ctx, nil
	}

	go func() {
		defer af.Free()
		defer close(frameOutChan)
	loop:
		for {
			select {
			case <-ctx.Done():
				return
			case frame, ok := <-frameInChan:
				// write fifo
				if !ok {
					break loop
				}
				if code := af.Realloc(af.Size() + frame.GetNbSamples()); code < 0 {
					frame.Unref()
					cancel(goavutil.AvErr(code))
					return
				}
				frameData := frame.GetData()
				code := af.Write(unsafe.SliceData(frameData[:]), frame.GetNbSamples())
				frame.Unref()
				if code < 0 {
					cancel(goavutil.AvErr(code))
					return
				}

				// read fifo
				frameSize := avctx.CAVCodecContext.GetFrameSize()
				for af.Size() >= frameSize {
					frame = goavutil.AllocAVFrame()
					frame.SetNbSamples(frameSize)
					avctx.GetChLayout().CopyTo(frame.GetChLayout())
					frame.SetFormat(int(avctx.CAVCodecContext.GetSampleFmt()))
					frame.SetSampleRate(avctx.CAVCodecContext.GetSampleRate())
					code = frame.AllocBuffer(0)
					if code < 0 {
						frame.Unref()
						cancel(goavutil.AvErr(code))
						return
					}
					frameData = frame.GetData()
					code = af.Read(unsafe.SliceData(frameData[:]), frameSize)
					if code < 0 {
						frame.Unref()
						cancel(goavutil.AvErr(code))
						return
					}
					select {
					case <-ctx.Done():
						frame.Unref()
						return
					case frameOutChan <- frame:
					}
				}
			}
		}

		// flush fifo
		if af.Size() >= 0 {
			frame := goavutil.AllocAVFrame()
			frame.SetNbSamples(af.Size())
			avctx.GetChLayout().CopyTo(frame.GetChLayout())
			frame.SetFormat(int(avctx.CAVCodecContext.GetSampleFmt()))
			frame.SetSampleRate(avctx.CAVCodecContext.GetSampleRate())
			code := frame.AllocBuffer(0)
			if code < 0 {
				frame.Unref()
				cancel(goavutil.AvErr(code))
				return
			}
			frameData := frame.GetData()
			code = af.Read(unsafe.SliceData(frameData[:]), af.Size())
			if code < 0 {
				frame.Unref()
				cancel(goavutil.AvErr(code))
				return
			}
			select {
			case <-ctx.Done():
				frame.Unref()
				return
			case frameOutChan <- frame:
			}
		}
	}()

	return ctx, frameOutChan
}

// Encode AVFrame.
// Return a slice of encoded AVPacket, May be nil or empty.
func (avctx *AVCodecContext) Encode(frame *goavutil.AVFrame) ([]*AVPacket, int) {
	code := avctx.SendFrame(frame)
	if code < 0 {
		return nil, code
	}

	pkts := make([]*AVPacket, 0)

	for {
		pkt := AllocAvPacket()
		if pkt == nil {
			code = int(syscall.ENOMEM)
			break
		}
		code = avctx.ReceivePacket(pkt)
		if code < 0 {
			pkt.Free()
			break
		}

		pkts = append(pkts, pkt)
	}

	return pkts, code
}

// Encode AVPacket from channel.
//
// @param ctx The context.Context to cancel this goroutine.
// @param frameChan The AVFrame channel.
//
// @return context.Context The context.Context to get the cause of this gorouine.
// @return <-chan *AVPacket The channel to read encoded AVPacket.
func (avctx *AVCodecContext) EncodeChan(ctx context.Context, frameChan <-chan *goavutil.AVFrame) (context.Context, <-chan *AVPacket) {
	ctx, cancel := context.WithCancelCause(ctx)
	pktChan := make(chan *AVPacket)

	go func() {
		defer close(pktChan)
	loop:
		for {
			select {
			case <-ctx.Done():
				return
			case frame, ok := <-frameChan:
				if !ok {
					break loop
				}
				pkts, code := avctx.Encode(frame)
				frame.Unref()
				for _, pkt := range pkts {
					select {
					case <-ctx.Done():
						for _, pkt := range pkts {
							pkt.Unref()
						}
						return
					case pktChan <- pkt:
					}
				}
				if code == avutil.AVERROR(int(syscall.EAGAIN)) || code == avutil.AVERROR_EOF {
					continue
				}
				if code < 0 {
					cancel(goavutil.AvErr(code))
					return
				}
			}
		}

		//flush codec
		pkts, code := avctx.Encode(nil)
		for _, pkt := range pkts {
			select {
			case <-ctx.Done():
				for _, pkt := range pkts {
					pkt.Unref()
				}
				return
			case pktChan <- pkt:
			}
		}
		if code == avutil.AVERROR(int(syscall.EAGAIN)) || code == avutil.AVERROR_EOF {
			return
		}
		if code < 0 {
			cancel(goavutil.AvErr(code))
			return
		}
	}()

	return ctx, pktChan
}

func (avctx *AVCodecContext) ParametersFrom(par *avcodec.CAVCodecParameters) int {
	return avcodec.AvcodecParametersToContext(avctx.CAVCodecContext, par)
}

func (avctx *AVCodecContext) ParametersTo(par *avcodec.CAVCodecParameters) int {
	return avcodec.AvcodecParametersFromContext(par, avctx.CAVCodecContext)
}

func (avctx *AVCodecContext) Open(options **avutil.CAVDictionary) int {
	return avcodec.AvcodecOpen2(avctx.CAVCodecContext, avctx.CAVCodecContext.GetCodec(), options)
}

func (avctx *AVCodecContext) Close() int {
	return avcodec.AvcodecClose(avctx.CAVCodecContext)
}

func (avctx *AVCodecContext) IsOpen() bool {
	return avcodec.AvcodecIsOpen(avctx.CAVCodecContext) > 0
}
