package goavformat

import (
	"unsafe"

	"github.com/Lensual/go-libav/advance/goavcodec"
	"github.com/Lensual/go-libav/advance/goavutil"
	"github.com/Lensual/go-libav/avcodec"
	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/ctypes"
)

type AvformatContext struct {
	CAvformatContext *avformat.CAVFormatContext
}

//#region members

func (fmtCtx *AvformatContext) GetAvClass() *avutil.CAVClass {
	return fmtCtx.CAvformatContext.GetAvClass()
}

func (fmtCtx *AvformatContext) GetIformat() *avformat.CAVInputFormat {
	return fmtCtx.CAvformatContext.GetIformat()
}

func (fmtCtx *AvformatContext) GetOformat() *avformat.CAVOutputFormat {
	return fmtCtx.CAvformatContext.GetOformat()
}
func (fmtCtx *AvformatContext) SetOformat(oformat *avformat.CAVOutputFormat) {
	fmtCtx.CAvformatContext.SetOformat(oformat)
}

func (fmtCtx *AvformatContext) GetPrivData() unsafe.Pointer {
	return fmtCtx.CAvformatContext.GetPrivData()
}
func (fmtCtx *AvformatContext) SetPrivData(privData unsafe.Pointer) {
	fmtCtx.CAvformatContext.SetPrivData(privData)
}

func (fmtCtx *AvformatContext) GetPb() *AVIOContext {
	pb := fmtCtx.CAvformatContext.GetPb()
	if pb == nil {
		return nil
	}
	return &AVIOContext{
		CAVIOContext: pb,
	}
}
func (fmtCtx *AvformatContext) SetPb(avioCtx *AVIOContext) {
	fmtCtx.CAvformatContext.SetPb(avioCtx.CAVIOContext)
}

// SetIoContext alias SetPb
func (fmtCtx *AvformatContext) SetIoContext(avioCtx *AVIOContext) {
	fmtCtx.SetPb(avioCtx)
}

func (fmtCtx *AvformatContext) GetCtxFlags() int {
	return fmtCtx.CAvformatContext.GetCtxFlags()
}
func (fmtCtx *AvformatContext) SetCtxFlags(ctxFlags int) {
	fmtCtx.CAvformatContext.SetCtxFlags(ctxFlags)
}

func (fmtCtx *AvformatContext) GetNbStreams() uint {
	return fmtCtx.CAvformatContext.GetNbStreams()
}
func (fmtCtx *AvformatContext) SetNbStreams(nbStreams uint) {
	fmtCtx.CAvformatContext.SetNbStreams(nbStreams)
}

func (fmtCtx *AvformatContext) GetStreams() []*avformat.CAVStream {
	return unsafe.Slice(fmtCtx.CAvformatContext.GetStreams(), fmtCtx.CAvformatContext.GetNbStreams())
}
func (fmtCtx *AvformatContext) SetStreams(streams []*avformat.CAVStream) {
	fmtCtx.CAvformatContext.SetStreams(unsafe.SliceData(streams))
}

func (fmtCtx *AvformatContext) GetUrl() string {
	return fmtCtx.CAvformatContext.GetUrl()
}
func (fmtCtx *AvformatContext) SetUrl(url string) {
	fmtCtx.CAvformatContext.SetUrl(url)
}

func (fmtCtx *AvformatContext) GetStartTime() int64 {
	return fmtCtx.CAvformatContext.GetStartTime()
}
func (fmtCtx *AvformatContext) SetStartTime(startTime int64) {
	fmtCtx.CAvformatContext.SetStartTime(startTime)
}

func (fmtCtx *AvformatContext) GetDuration() int64 {
	return fmtCtx.CAvformatContext.GetDuration()
}
func (fmtCtx *AvformatContext) SetDuration(duration int64) {
	fmtCtx.CAvformatContext.SetDuration(duration)
}

func (fmtCtx *AvformatContext) GetBitRate() int64 {
	return fmtCtx.CAvformatContext.GetBitRate()
}
func (fmtCtx *AvformatContext) SetBitRate(bitRate int64) {
	fmtCtx.CAvformatContext.SetBitRate(bitRate)
}

func (fmtCtx *AvformatContext) GetPacketSize() uint {
	return fmtCtx.CAvformatContext.GetPacketSize()
}
func (fmtCtx *AvformatContext) SetPacketSize(packetSize uint) {
	fmtCtx.CAvformatContext.SetPacketSize(packetSize)
}

func (fmtCtx *AvformatContext) GetMaxDelay() int {
	return fmtCtx.CAvformatContext.GetMaxDelay()
}
func (fmtCtx *AvformatContext) SetMaxDelay(maxDelay int) {
	fmtCtx.CAvformatContext.SetMaxDelay(maxDelay)
}

func (fmtCtx *AvformatContext) GetFlags() int {
	return fmtCtx.CAvformatContext.GetFlags()
}
func (fmtCtx *AvformatContext) SetFlags(flags int) {
	fmtCtx.CAvformatContext.SetFlags(flags)
}

func (fmtCtx *AvformatContext) GetProbesize() int64 {
	return fmtCtx.CAvformatContext.GetProbesize()
}
func (fmtCtx *AvformatContext) SetProbesize(probesize int64) {
	fmtCtx.CAvformatContext.SetProbesize(probesize)
}

func (fmtCtx *AvformatContext) GetMaxAnalyzeDuration() int64 {
	return fmtCtx.CAvformatContext.GetMaxAnalyzeDuration()
}
func (fmtCtx *AvformatContext) SetMaxAnalyzeDuration(maxAnalyzeDuration int64) {
	fmtCtx.CAvformatContext.SetMaxAnalyzeDuration(maxAnalyzeDuration)
}

func (fmtCtx *AvformatContext) GetKey() unsafe.Pointer {
	return fmtCtx.CAvformatContext.GetKey()
}
func (fmtCtx *AvformatContext) SetKey(key unsafe.Pointer) {
	fmtCtx.CAvformatContext.SetKey(key)
}

func (fmtCtx *AvformatContext) GetKeylen() int {
	return fmtCtx.CAvformatContext.GetKeylen()
}
func (fmtCtx *AvformatContext) SetKeylen(keylen int) {
	fmtCtx.CAvformatContext.SetKeylen(keylen)
}

func (fmtCtx *AvformatContext) GetNbPrograms() uint {
	return fmtCtx.CAvformatContext.GetNbPrograms()
}
func (fmtCtx *AvformatContext) SetNbPrograms(nbPrograms uint) {
	fmtCtx.CAvformatContext.SetNbPrograms(nbPrograms)
}

func (fmtCtx *AvformatContext) GetPrograms() **avformat.CAVProgram {
	return fmtCtx.CAvformatContext.GetPrograms()
}
func (fmtCtx *AvformatContext) SetPrograms(programs **avformat.CAVProgram) {
	fmtCtx.CAvformatContext.SetPrograms(programs)
}

func (fmtCtx *AvformatContext) GetVideoCodecId() avcodec.CAVCodecID {
	return fmtCtx.CAvformatContext.GetVideoCodecId()
}
func (fmtCtx *AvformatContext) SetVideoCodecId(videoCodecId avcodec.CAVCodecID) {
	fmtCtx.CAvformatContext.SetVideoCodecId(videoCodecId)
}

func (fmtCtx *AvformatContext) GetAudioCodecId() avcodec.CAVCodecID {
	return fmtCtx.CAvformatContext.GetAudioCodecId()
}
func (fmtCtx *AvformatContext) SetAudioCodecId(audioCodecId avcodec.CAVCodecID) {
	fmtCtx.CAvformatContext.SetAudioCodecId(audioCodecId)
}

func (fmtCtx *AvformatContext) GetSubtitleCodecId() avcodec.CAVCodecID {
	return fmtCtx.CAvformatContext.GetSubtitleCodecId()
}
func (fmtCtx *AvformatContext) SetSubtitleCodecId(subtitleCodecId avcodec.CAVCodecID) {
	fmtCtx.CAvformatContext.SetSubtitleCodecId(subtitleCodecId)
}

func (fmtCtx *AvformatContext) GetMaxIndexSize() uint {
	return fmtCtx.CAvformatContext.GetMaxIndexSize()
}
func (fmtCtx *AvformatContext) SetMaxIndexSize(maxIndexSize uint) {
	fmtCtx.CAvformatContext.SetMaxIndexSize(maxIndexSize)
}

func (fmtCtx *AvformatContext) GetMaxPictureBuffer() uint {
	return fmtCtx.CAvformatContext.GetMaxPictureBuffer()
}
func (fmtCtx *AvformatContext) SetMaxPictureBuffer(maxPictureBuffer uint) {
	fmtCtx.CAvformatContext.SetMaxPictureBuffer(maxPictureBuffer)
}

func (fmtCtx *AvformatContext) GetNbChapters() uint {
	return fmtCtx.CAvformatContext.GetNbChapters()
}
func (fmtCtx *AvformatContext) SetNbChapters(nbChapter uint) {
	fmtCtx.CAvformatContext.SetNbChapters(nbChapter)
}

func (fmtCtx *AvformatContext) GetChapters() **avformat.CAVChapter {
	return fmtCtx.CAvformatContext.GetChapters()
}
func (fmtCtx *AvformatContext) SetChapters(chapters **avformat.CAVChapter) {
	fmtCtx.CAvformatContext.SetChapters(chapters)
}

func (fmtCtx *AvformatContext) GetMetadata() *avutil.CAVDictionary {
	return fmtCtx.CAvformatContext.GetMetadata()
}
func (fmtCtx *AvformatContext) SetMetadata(metadata *avutil.CAVDictionary) {
	fmtCtx.CAvformatContext.SetMetadata(metadata)
}

func (fmtCtx *AvformatContext) GetStartTimeRealtime() int64 {
	return fmtCtx.CAvformatContext.GetStartTimeRealtime()
}
func (fmtCtx *AvformatContext) SetStartTimeRealtime(startTimeRealtime int64) {
	fmtCtx.CAvformatContext.SetStartTimeRealtime(startTimeRealtime)
}

func (fmtCtx *AvformatContext) GetFpsProbeSize() int {
	return fmtCtx.CAvformatContext.GetFpsProbeSize()
}
func (fmtCtx *AvformatContext) SetFpsProbeSize(fpsProbeSize int) {
	fmtCtx.CAvformatContext.SetFpsProbeSize(fpsProbeSize)
}

func (fmtCtx *AvformatContext) GetErrorRecognition() int {
	return fmtCtx.CAvformatContext.GetErrorRecognition()
}
func (fmtCtx *AvformatContext) SetErrorRecognition(errorRecognition int) {
	fmtCtx.CAvformatContext.SetErrorRecognition(errorRecognition)
}

func (fmtCtx *AvformatContext) GetInterruptCallback() avformat.CAVIOInterruptCB {
	return fmtCtx.CAvformatContext.GetInterruptCallback()
}
func (fmtCtx *AvformatContext) SetInterruptCallback(errorRecognition avformat.CAVIOInterruptCB) {
	fmtCtx.CAvformatContext.SetInterruptCallback(errorRecognition)
}

func (fmtCtx *AvformatContext) GetDebug() int {
	return fmtCtx.CAvformatContext.GetDebug()
}
func (fmtCtx *AvformatContext) SetDebug(debug int) {
	fmtCtx.CAvformatContext.SetDebug(debug)
}

func (fmtCtx *AvformatContext) GetMaxInterleaveDelta() int64 {
	return fmtCtx.CAvformatContext.GetMaxInterleaveDelta()
}
func (fmtCtx *AvformatContext) SetMaxInterleaveDelta(maxInterleaveDelta int64) {
	fmtCtx.CAvformatContext.SetMaxInterleaveDelta(maxInterleaveDelta)
}

func (fmtCtx *AvformatContext) GetStrictStdCompliance() int {
	return fmtCtx.CAvformatContext.GetStrictStdCompliance()
}
func (fmtCtx *AvformatContext) SetStrictStdCompliance(strictStdCompliance int) {
	fmtCtx.CAvformatContext.SetStrictStdCompliance(strictStdCompliance)
}

func (fmtCtx *AvformatContext) GetEventFlags() int {
	return fmtCtx.CAvformatContext.GetEventFlags()
}
func (fmtCtx *AvformatContext) SetEventFlags(eventFlags int) {
	fmtCtx.CAvformatContext.SetEventFlags(eventFlags)
}

func (fmtCtx *AvformatContext) GetMaxTsProbe() int {
	return fmtCtx.CAvformatContext.GetMaxTsProbe()
}
func (fmtCtx *AvformatContext) SetMaxTsProbe(maxTsProbe int) {
	fmtCtx.CAvformatContext.SetMaxTsProbe(maxTsProbe)
}

func (fmtCtx *AvformatContext) GetAvoidNegativeTs() int {
	return fmtCtx.CAvformatContext.GetAvoidNegativeTs()
}
func (fmtCtx *AvformatContext) SetAvoidNegativeTs(avoidNegativeTs int) {
	fmtCtx.CAvformatContext.SetAvoidNegativeTs(avoidNegativeTs)
}

func (fmtCtx *AvformatContext) GetTsId() int {
	return fmtCtx.CAvformatContext.GetTsId()
}
func (fmtCtx *AvformatContext) SetTsId(tsId int) {
	fmtCtx.CAvformatContext.SetTsId(tsId)
}

func (fmtCtx *AvformatContext) GetAudioPreload() int {
	return fmtCtx.CAvformatContext.GetAudioPreload()
}
func (fmtCtx *AvformatContext) SetAudioPreload(audioPreload int) {
	fmtCtx.CAvformatContext.SetAudioPreload(audioPreload)
}

func (fmtCtx *AvformatContext) GetMaxChunkDuration() int {
	return fmtCtx.CAvformatContext.GetMaxChunkDuration()
}
func (fmtCtx *AvformatContext) SetMaxChunkDuration(maxChunkDuration int) {
	fmtCtx.CAvformatContext.SetMaxChunkDuration(maxChunkDuration)
}

func (fmtCtx *AvformatContext) GetMaxChunkSize() int {
	return fmtCtx.CAvformatContext.GetMaxChunkSize()
}
func (fmtCtx *AvformatContext) SetMaxChunkSize(maxChunkSize int) {
	fmtCtx.CAvformatContext.SetMaxChunkSize(maxChunkSize)
}

func (fmtCtx *AvformatContext) GetUseWallclockAsTimestamps() int {
	return fmtCtx.CAvformatContext.GetUseWallclockAsTimestamps()
}
func (fmtCtx *AvformatContext) SetUseWallclockAsTimestamps(useWallclockAsTimestamps int) {
	fmtCtx.CAvformatContext.SetUseWallclockAsTimestamps(useWallclockAsTimestamps)
}

func (fmtCtx *AvformatContext) GetAvioFlags() int {
	return fmtCtx.CAvformatContext.GetAvioFlags()
}
func (fmtCtx *AvformatContext) SetAvioFlags(avioFlags int) {
	fmtCtx.CAvformatContext.SetAvioFlags(avioFlags)
}

func (fmtCtx *AvformatContext) GetDurationEstimationMethod() avformat.CAVDurationEstimationMethod {
	return fmtCtx.CAvformatContext.GetDurationEstimationMethod()
}
func (fmtCtx *AvformatContext) SetDurationEstimationMethod(durationEstimationMethod avformat.CAVDurationEstimationMethod) {
	fmtCtx.CAvformatContext.SetDurationEstimationMethod(durationEstimationMethod)
}

func (fmtCtx *AvformatContext) GetSkipInitialBytes() int64 {
	return fmtCtx.CAvformatContext.GetSkipInitialBytes()
}
func (fmtCtx *AvformatContext) SetSkipInitialBytes(skipInitialBytes int64) {
	fmtCtx.CAvformatContext.SetSkipInitialBytes(skipInitialBytes)
}

func (fmtCtx *AvformatContext) GetCorrectTsOverflow() uint {
	return fmtCtx.CAvformatContext.GetCorrectTsOverflow()
}
func (fmtCtx *AvformatContext) SetCorrectTsOverflow(correctTsOverflow uint) {
	fmtCtx.CAvformatContext.SetCorrectTsOverflow(correctTsOverflow)
}

func (fmtCtx *AvformatContext) GetSeek2Any() int {
	return fmtCtx.CAvformatContext.GetSeek2Any()
}
func (fmtCtx *AvformatContext) SetSeek2Any(seek2Any int) {
	fmtCtx.CAvformatContext.SetSeek2Any(seek2Any)
}

func (fmtCtx *AvformatContext) GetFlushPackets() int {
	return fmtCtx.CAvformatContext.GetFlushPackets()
}
func (fmtCtx *AvformatContext) SetFlushPackets(flushPackets int) {
	fmtCtx.CAvformatContext.SetFlushPackets(flushPackets)
}

func (fmtCtx *AvformatContext) GetProbeScore() int {
	return fmtCtx.CAvformatContext.GetProbeScore()
}
func (fmtCtx *AvformatContext) SetProbeScore(probeScore int) {
	fmtCtx.CAvformatContext.SetProbeScore(probeScore)
}

func (fmtCtx *AvformatContext) GetFormatProbesize() int {
	return fmtCtx.CAvformatContext.GetFormatProbesize()
}
func (fmtCtx *AvformatContext) SetFormatProbesize(formatProbeSize int) {
	fmtCtx.CAvformatContext.SetFormatProbesize(formatProbeSize)
}

func (fmtCtx *AvformatContext) GetCodecWhitelist() string {
	return fmtCtx.CAvformatContext.GetCodecWhitelist()
}
func (fmtCtx *AvformatContext) SetCodecWhitelist(codecWhitelist string) {
	fmtCtx.CAvformatContext.SetCodecWhitelist(codecWhitelist)
}

func (fmtCtx *AvformatContext) GetFormatWhitelist() string {
	return fmtCtx.CAvformatContext.GetFormatWhitelist()
}
func (fmtCtx *AvformatContext) SetFormatWhitelist(formatWhitelist string) {
	fmtCtx.CAvformatContext.SetFormatWhitelist(formatWhitelist)
}

func (fmtCtx *AvformatContext) GetIoRespositioned() int {
	return fmtCtx.CAvformatContext.GetIoRespositioned()
}
func (fmtCtx *AvformatContext) SetIoRespositioned(ioRespositioned int) {
	fmtCtx.CAvformatContext.SetIoRespositioned(ioRespositioned)
}

func (fmtCtx *AvformatContext) GetVideoCodec() *goavcodec.AVCodec {
	codec := fmtCtx.CAvformatContext.GetVideoCodec()
	if codec == nil {
		return nil
	}
	return &goavcodec.AVCodec{
		CAVCodec: codec,
	}
}
func (fmtCtx *AvformatContext) SetVideoCodec(videoCodec *goavcodec.AVCodec) {
	fmtCtx.CAvformatContext.SetVideoCodec(videoCodec.CAVCodec)
}

func (fmtCtx *AvformatContext) GetAudioCodec() *goavcodec.AVCodec {
	codec := fmtCtx.CAvformatContext.GetAudioCodec()
	if codec == nil {
		return nil
	}
	return &goavcodec.AVCodec{
		CAVCodec: codec,
	}
}
func (fmtCtx *AvformatContext) SetAudioCodec(audioCodec *goavcodec.AVCodec) {
	fmtCtx.CAvformatContext.SetAudioCodec(audioCodec.CAVCodec)
}

func (fmtCtx *AvformatContext) GetSubtitleCodec() *goavcodec.AVCodec {
	codec := fmtCtx.CAvformatContext.GetSubtitleCodec()
	if codec == nil {
		return nil
	}
	return &goavcodec.AVCodec{
		CAVCodec: codec,
	}
}
func (fmtCtx *AvformatContext) SetSubtitleCodec(subtitleCodec *goavcodec.AVCodec) {
	fmtCtx.CAvformatContext.SetSubtitleCodec(subtitleCodec.CAVCodec)
}

func (fmtCtx *AvformatContext) GetDataCodec() *goavcodec.AVCodec {
	codec := fmtCtx.CAvformatContext.GetDataCodec()
	if codec == nil {
		return nil
	}
	return &goavcodec.AVCodec{
		CAVCodec: codec,
	}
}
func (fmtCtx *AvformatContext) SetDataCodec(dataCodec *goavcodec.AVCodec) {
	fmtCtx.CAvformatContext.SetDataCodec(dataCodec.CAVCodec)
}

func (fmtCtx *AvformatContext) GetMetadataHeaderPadding() int {
	return fmtCtx.CAvformatContext.GetMetadataHeaderPadding()
}
func (fmtCtx *AvformatContext) SetMetadataHeaderPadding(metadataHeaderPadding int) {
	fmtCtx.CAvformatContext.SetMetadataHeaderPadding(metadataHeaderPadding)
}

func (fmtCtx *AvformatContext) GetOpaque() unsafe.Pointer {
	return fmtCtx.CAvformatContext.GetOpaque()
}
func (fmtCtx *AvformatContext) SetOpaque(opaque unsafe.Pointer) {
	fmtCtx.CAvformatContext.SetOpaque(opaque)
}

func (fmtCtx *AvformatContext) GetControlMessageCb() avformat.CAvFormatControlMessage {
	return fmtCtx.CAvformatContext.GetControlMessageCb()
}
func (fmtCtx *AvformatContext) SetControlMessageCb(controlMessageCb avformat.CAvFormatControlMessage) {
	fmtCtx.CAvformatContext.SetControlMessageCb(controlMessageCb)
}

func (fmtCtx *AvformatContext) GetOutputTsOffset() int64 {
	return fmtCtx.CAvformatContext.GetOutputTsOffset()
}
func (fmtCtx *AvformatContext) SetOutputTsOffset(outputTsOffset int64) {
	fmtCtx.CAvformatContext.SetOutputTsOffset(outputTsOffset)
}

func (fmtCtx *AvformatContext) GetDumpSeparator() string {
	return fmtCtx.CAvformatContext.GetDumpSeparator()
}
func (fmtCtx *AvformatContext) SetDumpSeparator(dumpSeparator string) {
	fmtCtx.CAvformatContext.SetDumpSeparator(dumpSeparator)
}

func (fmtCtx *AvformatContext) GetDataCodecId() avcodec.CAVCodecID {
	return fmtCtx.CAvformatContext.GetDataCodecId()
}
func (fmtCtx *AvformatContext) SetDataCodecId(dataCodecid avcodec.CAVCodecID) {
	fmtCtx.CAvformatContext.SetDataCodecId(dataCodecid)
}

func (fmtCtx *AvformatContext) GetProtocolWhitelist() string {
	return fmtCtx.CAvformatContext.GetProtocolWhitelist()
}
func (fmtCtx *AvformatContext) SetProtocolWhitelist(protocolWhitelist string) {
	fmtCtx.CAvformatContext.SetProtocolWhitelist(protocolWhitelist)
}

func (fmtCtx *AvformatContext) GetIoOpen() ctypes.CFunc {
	return fmtCtx.CAvformatContext.GetIoOpen()
}
func (fmtCtx *AvformatContext) SetIoOpen(ioOpen ctypes.CFunc) {
	fmtCtx.CAvformatContext.SetIoOpen(ioOpen)
}

func (fmtCtx *AvformatContext) GetProtocolBlacklist() string {
	return fmtCtx.CAvformatContext.GetProtocolBlacklist()
}
func (fmtCtx *AvformatContext) SetProtocolBlacklist(protocolBlackList string) {
	fmtCtx.CAvformatContext.SetProtocolBlacklist(protocolBlackList)
}

func (fmtCtx *AvformatContext) GetMaxStreams() int {
	return fmtCtx.CAvformatContext.GetMaxStreams()
}
func (fmtCtx *AvformatContext) SetMaxStreams(maxStreams int) {
	fmtCtx.CAvformatContext.SetMaxStreams(maxStreams)
}

func (fmtCtx *AvformatContext) GetSkipEstimateDurationFromPts() int {
	return fmtCtx.CAvformatContext.GetSkipEstimateDurationFromPts()
}
func (fmtCtx *AvformatContext) SetSkipEstimateDurationFromPts(skipEstimateDurationFromPts int) {
	fmtCtx.CAvformatContext.SetSkipEstimateDurationFromPts(skipEstimateDurationFromPts)
}

func (fmtCtx *AvformatContext) GetMaxProbePackets() int {
	return fmtCtx.CAvformatContext.GetMaxProbePackets()
}
func (fmtCtx *AvformatContext) SetMaxProbePackets(maxProbePackets int) {
	fmtCtx.CAvformatContext.SetMaxProbePackets(maxProbePackets)
}

func (fmtCtx *AvformatContext) GetIoClose2() ctypes.CFunc {
	return fmtCtx.CAvformatContext.GetIoClose2()
}
func (fmtCtx *AvformatContext) SetIoClose2(ioClose2 ctypes.CFunc) {
	fmtCtx.CAvformatContext.SetIoClose2(ioClose2)
}

//#endregion members

func (fmtCtx *AvformatContext) InjectGlobalSideData() {
	avformat.AvFormatInjectGlobalSideData(fmtCtx.CAvformatContext)
}

func (fmtCtx *AvformatContext) Free() {
	avformat.AvformatFreeContext(fmtCtx.CAvformatContext)
}

func (fmtCtx *AvformatContext) NewStream(c *goavcodec.AVCodec) *avformat.CAVStream {
	return avformat.AvformatNewStream(fmtCtx.CAvformatContext, c.CAVCodec)
}

func (fmtCtx *AvformatContext) NewProgram(id int) *avformat.CAVProgram {
	return avformat.AvNewProgram(fmtCtx.CAvformatContext, id)
}

func (fmtCtx *AvformatContext) OpenInput(url string, fmt *avformat.CAVInputFormat, options **avutil.CAVDictionary) int {
	return avformat.AvformatOpenInput(&fmtCtx.CAvformatContext, url, fmt, options)
}

func (fmtCtx *AvformatContext) FindStreamInfo(options **avutil.CAVDictionary) int {
	return avformat.AvformatFindStreamInfo(fmtCtx.CAvformatContext, options)
}

func (fmtCtx *AvformatContext) FindProgramFromStream(last *avformat.CAVProgram, s int) *avformat.CAVProgram {
	return avformat.AvFindProgramFromStream(fmtCtx.CAvformatContext, last, s)
}

func (fmtCtx *AvformatContext) ProgramAddStreamIndex(progid int, idx uint) {
	avformat.AvProgramAddStreamIndex(fmtCtx.CAvformatContext, progid, idx)
}

func (fmtCtx *AvformatContext) FindBestStream(_type avutil.CAVMediaType, wanted_stream_nb int,
	related_stream int, flags int) (int, *goavcodec.AVCodec) {
	dec := goavcodec.AVCodec{}
	ret := avformat.AvFindBestStream(fmtCtx.CAvformatContext, _type, wanted_stream_nb, related_stream, &dec.CAVCodec, flags)
	return ret, &dec
}

func (fmtCtx *AvformatContext) ReadFrame(pkt *goavcodec.AVPacket) int {
	return avformat.AvReadFrame(fmtCtx.CAvformatContext, pkt.CAVPacket)
}

func (fmtCtx *AvformatContext) SeekFrame(streamIndex int, timestamp int64, flags int) int {
	return avformat.AvSeekFrame(fmtCtx.CAvformatContext, streamIndex, timestamp, flags)
}

func (fmtCtx *AvformatContext) SeekFile(streamIndex int, minTs int64, ts int64, maxTs int64, flags int) int {
	return avformat.AvformatSeekFile(fmtCtx.CAvformatContext, streamIndex, minTs, ts, maxTs, flags)
}

func (fmtCtx *AvformatContext) Flush() int {
	return avformat.AvformatFlush(fmtCtx.CAvformatContext)
}

func (fmtCtx *AvformatContext) ReadPlay() int {
	return avformat.AvReadPlay(fmtCtx.CAvformatContext)
}

func (fmtCtx *AvformatContext) ReadPause() int {
	return avformat.AvReadPause(fmtCtx.CAvformatContext)
}

func (fmtCtx *AvformatContext) CloseInput() {
	avformat.AvformatCloseInput(&fmtCtx.CAvformatContext)
	fmtCtx.CAvformatContext = nil
}

func (fmtCtx *AvformatContext) WriteHeader(options **avutil.CAVDictionary) int {
	return avformat.AvformatWriteHeader(fmtCtx.CAvformatContext, options)
}

func (fmtCtx *AvformatContext) InitOutput(options **avutil.CAVDictionary) int {
	return avformat.AvformatInitOutput(fmtCtx.CAvformatContext, options)
}

func (fmtCtx *AvformatContext) WriteFrame(pkt *goavcodec.AVPacket) int {
	return avformat.AvWriteFrame(fmtCtx.CAvformatContext, pkt.CAVPacket)
}

func (fmtCtx *AvformatContext) InterleavedWriteFrame(pkt *goavcodec.AVPacket) int {
	return avformat.AvInterleavedWriteFrame(fmtCtx.CAvformatContext, pkt.CAVPacket)
}

func (fmtCtx *AvformatContext) WriteUncodedFrame(streamIndex int, frame *goavutil.AVFrame) int {
	return avformat.AvWriteUncodedFrame(fmtCtx.CAvformatContext, streamIndex, frame.CAVFrame)
}

func (fmtCtx *AvformatContext) InterleavedWriteUncodedFrame(streamIndex int, frame *goavutil.AVFrame) int {
	return avformat.AvInterleavedWriteUncodedFrame(fmtCtx.CAvformatContext, streamIndex, frame.CAVFrame)
}

func (fmtCtx *AvformatContext) WriteUncodedFrameQuery(streamIndex int) int {
	return avformat.AvWriteUncodedFrameQuery(fmtCtx.CAvformatContext, streamIndex)
}

func (fmtCtx *AvformatContext) WriteTrailer() int {
	return avformat.AvWriteTrailer(fmtCtx.CAvformatContext)
}

func (fmtCtx *AvformatContext) GetOutputTimestamp(stream int) (dts int64, wall int64, ret int) {
	var cDts *ctypes.Int64
	var cWall *ctypes.Int64
	ret = avformat.AvGetOutputTimestamp(fmtCtx.CAvformatContext, stream, cDts, cWall)
	if ret != 0 {
		return 0, 0, ret
	}
	return int64(*cDts), int64(*cWall), ret
}

func (fmtCtx *AvformatContext) DumpFormat(index int, url string, is_output int) {
	avformat.AvDumpFormat(fmtCtx.CAvformatContext, index, url, is_output)
}

func (fmtCtx *AvformatContext) FindDefaultStreamIndex() int {
	return avformat.AvFindDefaultStreamIndex(fmtCtx.CAvformatContext)
}

func (fmtCtx *AvformatContext) GuessSampleAspectRatio(stream *avformat.CAVStream, frame *goavutil.AVFrame) avutil.CAVRational {
	return avformat.AvGuessSampleAspectRatio(fmtCtx.CAvformatContext, stream, frame.CAVFrame)
}

func (fmtCtx *AvformatContext) GuessFrameRate(stream *avformat.CAVStream, frame *goavutil.AVFrame) avutil.CAVRational {
	return avformat.AvGuessFrameRate(fmtCtx.CAvformatContext, stream, frame.CAVFrame)
}

func (fmtCtx *AvformatContext) MatchStreamSpecifier(st *avformat.CAVStream, spec string) int {
	return avformat.AvformatMatchStreamSpecifier(fmtCtx.CAvformatContext, st, spec)
}

func (fmtCtx *AvformatContext) QueueAttachedPictures() int {
	return avformat.AvformatQueueAttachedPictures(fmtCtx.CAvformatContext)
}

func GetAvailableMuxer() []*avformat.CAVOutputFormat {
	p := unsafe.Pointer(nil)
	arr := []*avformat.CAVOutputFormat{}
	for {
		muxer := avformat.AvMuxerIterate(&p)
		if muxer == nil {
			break
		}
		arr = append(arr, muxer)
	}
	return arr
}

func GetAvailableDeuxer() []*avformat.CAVInputFormat {
	p := unsafe.Pointer(nil)
	arr := []*avformat.CAVInputFormat{}
	for {
		demuxer := avformat.AvDeuxerIterate(&p)
		if demuxer == nil {
			break
		}
		arr = append(arr, demuxer)
	}
	return arr
}

func AllocContext() *AvformatContext {
	ctx := avformat.AvformatAllocContext()
	if ctx == nil {
		return nil
	}
	return &AvformatContext{
		CAvformatContext: ctx,
	}
}

func AllocOutputContext(oformat *avformat.CAVOutputFormat, formatName string, filename string) (*AvformatContext, int) {
	fmtCtx := AvformatContext{}
	ret := avformat.AvformatAllocOutputContext2(&fmtCtx.CAvformatContext, oformat, formatName, filename)
	if ret >= 0 {
		return nil, ret
	}
	return &fmtCtx, ret
}

func OpenInput(url string, fmt *avformat.CAVInputFormat, options **avutil.CAVDictionary) (*AvformatContext, int) {
	fmtCtx := AvformatContext{}
	ret := avformat.AvformatOpenInput(&fmtCtx.CAvformatContext, url, fmt, options)
	if ret != 0 {
		return nil, ret
	}
	return &fmtCtx, ret
}
