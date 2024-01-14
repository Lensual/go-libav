package goavformat

/*
#cgo pkg-config: libavformat

#include "libavformat/avformat.h"

#include <stdint.h>

extern int go_read_packet(void *opaque, uint8_t *cbuf, int buf_size);
extern int go_write_packet(void *opaque, uint8_t *cbuf, int buf_size);
extern int64_t go_seek(void *opaque, int64_t offset, int whence);
*/
import "C"

import (
	"unsafe"

	"github.com/Lensual/go-libav/advance/goavcodec"
	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/ctypes"
)

type AVIOContext struct {
	CAVIOContext *avformat.CAVIOContext
}

//#region members

func (avioCtx *AVIOContext) GetAvClass() *avutil.CAVClass {
	return avioCtx.CAVIOContext.GetAvClass()
}

func (avioCtx *AVIOContext) GetBuffer() []byte {
	return unsafe.Slice((*byte)(avioCtx.CAVIOContext.GetBuffer()), avioCtx.CAVIOContext.GetBufferSize())
}
func (avioCtx *AVIOContext) SetBuffer(buffer []byte) {
	avioCtx.CAVIOContext.SetBuffer(C.CBytes(buffer))
}

func (avioCtx *AVIOContext) GetBufferSize() int {
	return avioCtx.CAVIOContext.GetBufferSize()
}
func (avioCtx *AVIOContext) SetBufferSize(bufferSize int) {
	avioCtx.CAVIOContext.SetBufferSize(bufferSize)
}

func (avioCtx *AVIOContext) GetBufPtr() unsafe.Pointer {
	return avioCtx.CAVIOContext.GetBufPtr()
}
func (avioCtx *AVIOContext) SetBufPtr(bufPtr unsafe.Pointer) {
	avioCtx.CAVIOContext.SetBufPtr(bufPtr)
}

func (avioCtx *AVIOContext) GetBufEnd() unsafe.Pointer {
	return avioCtx.CAVIOContext.GetBufEnd()
}
func (avioCtx *AVIOContext) SetBufEnd(bufPtr unsafe.Pointer) {
	avioCtx.CAVIOContext.SetBufEnd(bufPtr)
}

func (avioCtx *AVIOContext) GetOpaque() unsafe.Pointer {
	return avioCtx.CAVIOContext.GetOpaque()
}
func (avioCtx *AVIOContext) SetOpaque(opaque unsafe.Pointer) {
	avioCtx.CAVIOContext.SetOpaque(opaque)
}

func (avioCtx *AVIOContext) GetReadPacket() ctypes.CFunc {
	return avioCtx.CAVIOContext.GetReadPacket()
}
func (avioCtx *AVIOContext) SetReadPacket(readPacket ctypes.CFunc) {
	avioCtx.CAVIOContext.SetReadPacket(readPacket)
}

func (avioCtx *AVIOContext) GetWritePacket() ctypes.CFunc {
	return avioCtx.CAVIOContext.GetWritePacket()
}
func (avioCtx *AVIOContext) SetWritePacket(writePacket ctypes.CFunc) {
	avioCtx.CAVIOContext.SetWritePacket(writePacket)
}

func (avioCtx *AVIOContext) GetSeek() ctypes.CFunc {
	return avioCtx.CAVIOContext.GetSeek()
}
func (avioCtx *AVIOContext) SetSeek(seek ctypes.CFunc) {
	avioCtx.CAVIOContext.SetSeek(seek)
}

func (avioCtx *AVIOContext) GetPos() int64 {
	return avioCtx.CAVIOContext.GetPos()
}
func (avioCtx *AVIOContext) SetPos(pos int64) {
	avioCtx.CAVIOContext.SetPos(pos)
}

func (avioCtx *AVIOContext) GetEofReached() int {
	return avioCtx.CAVIOContext.GetEofReached()
}
func (avioCtx *AVIOContext) SetEofReached(eofReached int) {
	avioCtx.CAVIOContext.SetEofReached(eofReached)
}

func (avioCtx *AVIOContext) GetError() int {
	return avioCtx.CAVIOContext.GetError()
}
func (avioCtx *AVIOContext) SetError(_error int) {
	avioCtx.CAVIOContext.SetError(_error)
}

func (avioCtx *AVIOContext) GetWriteFlag() int {
	return avioCtx.CAVIOContext.GetWriteFlag()
}
func (avioCtx *AVIOContext) SetWriteFlag(writeFlag int) {
	avioCtx.CAVIOContext.SetWriteFlag(writeFlag)
}

func (avioCtx *AVIOContext) GetMaxPacketSize() int {
	return avioCtx.CAVIOContext.GetMaxPacketSize()
}
func (avioCtx *AVIOContext) SetMaxPacketSize(maxPacketSize int) {
	avioCtx.CAVIOContext.SetMaxPacketSize(maxPacketSize)
}

func (avioCtx *AVIOContext) GetMinPacketSize() int {
	return avioCtx.CAVIOContext.GetMinPacketSize()
}
func (avioCtx *AVIOContext) SetMinPacketSize(minPacketSize int) {
	avioCtx.CAVIOContext.SetMinPacketSize(minPacketSize)
}

func (avioCtx *AVIOContext) GetCheckSum() uint64 {
	return avioCtx.CAVIOContext.GetCheckSum()
}
func (avioCtx *AVIOContext) SetCheckSum(checkSum uint64) {
	avioCtx.CAVIOContext.SetCheckSum(checkSum)
}

func (avioCtx *AVIOContext) GetCheckSumPtr() unsafe.Pointer {
	return avioCtx.CAVIOContext.GetCheckSumPtr()
}
func (avioCtx *AVIOContext) SetCheckSumPtr(checkSumPtr unsafe.Pointer) {
	avioCtx.CAVIOContext.SetCheckSumPtr(checkSumPtr)
}

func (avioCtx *AVIOContext) GetUpdateCheckSum() ctypes.CFunc {
	return avioCtx.CAVIOContext.GetUpdateCheckSum()
}
func (avioCtx *AVIOContext) SetUpdateCheckSum(updateCheckSum ctypes.CFunc) {
	avioCtx.CAVIOContext.SetUpdateCheckSum(updateCheckSum)
}

func (avioCtx *AVIOContext) GetReadPause() ctypes.CFunc {
	return avioCtx.CAVIOContext.GetReadPause()
}
func (avioCtx *AVIOContext) SetReadPause(readPause ctypes.CFunc) {
	avioCtx.CAVIOContext.SetReadPause(readPause)
}

func (avioCtx *AVIOContext) GetReadSeek() ctypes.CFunc {
	return avioCtx.CAVIOContext.GetReadSeek()
}
func (avioCtx *AVIOContext) SetReadSeek(readSeek ctypes.CFunc) {
	avioCtx.CAVIOContext.SetReadSeek(readSeek)
}

func (avioCtx *AVIOContext) GetSeekable() int {
	return avioCtx.CAVIOContext.GetSeekable()
}
func (avioCtx *AVIOContext) SetSeekable(seekable int) {
	avioCtx.CAVIOContext.SetSeekable(seekable)
}

func (avioCtx *AVIOContext) GetDirect() int {
	return avioCtx.CAVIOContext.GetDirect()
}
func (avioCtx *AVIOContext) SetDirect(direct int) {
	avioCtx.CAVIOContext.SetDirect(direct)
}

func (avioCtx *AVIOContext) GetProtocolWhitelist() string {
	return avioCtx.CAVIOContext.GetProtocolWhitelist()
}
func (avioCtx *AVIOContext) SetProtocolWhitelist(protocolWhitelist string) {
	avioCtx.CAVIOContext.SetProtocolWhitelist(protocolWhitelist)
}

func (avioCtx *AVIOContext) GetProtocolBlacklist() string {
	return avioCtx.CAVIOContext.GetProtocolBlacklist()
}
func (avioCtx *AVIOContext) SetProtocolBlacklist(protocolBlacklist string) {
	avioCtx.CAVIOContext.SetProtocolBlacklist(protocolBlacklist)
}

func (avioCtx *AVIOContext) GetWriteDataType() ctypes.CFunc {
	return avioCtx.CAVIOContext.GetWriteDataType()
}
func (avioCtx *AVIOContext) SetWriteDataType(writeDataType ctypes.CFunc) {
	avioCtx.CAVIOContext.SetWriteDataType(writeDataType)
}

func (avioCtx *AVIOContext) GetIgnoreBoundaryPoint() int {
	return avioCtx.CAVIOContext.GetIgnoreBoundaryPoint()
}
func (avioCtx *AVIOContext) SetIgnoreBoundaryPoint(ignoreBoundaryPoint int) {
	avioCtx.CAVIOContext.SetIgnoreBoundaryPoint(ignoreBoundaryPoint)
}

func (avioCtx *AVIOContext) GetBufPtrMax() unsafe.Pointer {
	return avioCtx.CAVIOContext.GetBufPtrMax()
}
func (avioCtx *AVIOContext) SetBufPtrMax(bufPtrMax unsafe.Pointer) {
	avioCtx.CAVIOContext.SetBufPtrMax(bufPtrMax)
}

func (avioCtx *AVIOContext) GetBytesRead() int64 {
	return avioCtx.CAVIOContext.GetBytesRead()
}

func (avioCtx *AVIOContext) GetBytesWritten() int64 {
	return avioCtx.CAVIOContext.GetBytesWritten()
}

//#endregion members

func (avioCtx *AVIOContext) Free() {
	if avioCtx.CAVIOContext != nil {
		/* note: the internal buffer could have changed, and be != avio_ctx_buffer */
		bufptr := avioCtx.CAVIOContext.GetBuffer()
		avutil.AvFreep(unsafe.Pointer(&bufptr))
		avformat.AvioContextFree(&avioCtx.CAVIOContext)
	}
}

func (avioCtx *AVIOContext) W8(b int) {
	avformat.AvioW8(avioCtx.CAVIOContext, b)
}
func (avioCtx *AVIOContext) Write(buf unsafe.Pointer, size int) {
	avformat.AvioWrite(avioCtx.CAVIOContext, buf, size)
}
func (avioCtx *AVIOContext) Wl64(val uint64) {
	avformat.AvioWl64(avioCtx.CAVIOContext, val)
}
func (avioCtx *AVIOContext) Wb64(val uint64) {
	avformat.AvioWb64(avioCtx.CAVIOContext, val)
}
func (avioCtx *AVIOContext) Wl32(val uint) {
	avformat.AvioWl32(avioCtx.CAVIOContext, val)
}
func (avioCtx *AVIOContext) Wb32(val uint) {
	avformat.AvioWb32(avioCtx.CAVIOContext, val)
}
func (avioCtx *AVIOContext) Wl24(val uint) {
	avformat.AvioWl24(avioCtx.CAVIOContext, val)
}
func (avioCtx *AVIOContext) Wb24(val uint) {
	avformat.AvioWb24(avioCtx.CAVIOContext, val)
}
func (avioCtx *AVIOContext) Wl16(val uint) {
	avformat.AvioWl16(avioCtx.CAVIOContext, val)
}
func (avioCtx *AVIOContext) Wb16(val uint) {
	avformat.AvioWb16(avioCtx.CAVIOContext, val)
}

func (avioCtx *AVIOContext) PutStr(str string) {
	avformat.AvioPutStr(avioCtx.CAVIOContext, str)
}

func (avioCtx *AVIOContext) PutStr16le(str string) {
	avformat.AvioPutStr16le(avioCtx.CAVIOContext, str)
}

func (avioCtx *AVIOContext) PutStr16be(str string) {
	avformat.AvioPutStr16be(avioCtx.CAVIOContext, str)
}

func (avioCtx *AVIOContext) WriteMarker(time int64, _type avformat.CAVIODataMarkerType) {
	avformat.AvioWriteMarker(avioCtx.CAVIOContext, time, _type)
}

func (avioCtx *AVIOContext) Seek(offset int64, whence int) int64 {
	return avformat.AvioSeek(avioCtx.CAVIOContext, offset, whence)
}

func (avioCtx *AVIOContext) Skip(offset int64) int64 {
	return avformat.AvioSkip(avioCtx.CAVIOContext, offset)
}

func (avioCtx *AVIOContext) Tell() int64 {
	return avformat.AvioTell(avioCtx.CAVIOContext)
}

func (avioCtx *AVIOContext) Size() int64 {
	return avformat.AvioSize(avioCtx.CAVIOContext)
}

func (avioCtx *AVIOContext) Feof() int {
	return avformat.AvioFeof(avioCtx.CAVIOContext)
}

func (avioCtx *AVIOContext) Flush() {
	avformat.AvioFlush(avioCtx.CAVIOContext)
}

func (avioCtx *AVIOContext) Read(buf unsafe.Pointer, size int) int {
	return avformat.AvioRead(avioCtx.CAVIOContext, buf, size)
}

func (avioCtx *AVIOContext) ReadPartial(buf unsafe.Pointer, size int) int {
	return avformat.AvioReadPartial(avioCtx.CAVIOContext, buf, size)
}

func (avioCtx *AVIOContext) R8() int {
	return avformat.AvioR8(avioCtx.CAVIOContext)
}
func (avioCtx *AVIOContext) Rl16() uint {
	return avformat.AvioRl16(avioCtx.CAVIOContext)
}
func (avioCtx *AVIOContext) Rl24() uint {
	return avformat.AvioRl24(avioCtx.CAVIOContext)
}
func (avioCtx *AVIOContext) Rl32() uint {
	return avformat.AvioRl32(avioCtx.CAVIOContext)
}
func (avioCtx *AVIOContext) Rl64() uint64 {
	return avformat.AvioRl64(avioCtx.CAVIOContext)
}
func (avioCtx *AVIOContext) Rb16() uint {
	return avformat.AvioRb16(avioCtx.CAVIOContext)
}
func (avioCtx *AVIOContext) Rb24() uint {
	return avformat.AvioRb24(avioCtx.CAVIOContext)
}
func (avioCtx *AVIOContext) Rb32() uint {
	return avformat.AvioRb32(avioCtx.CAVIOContext)
}
func (avioCtx *AVIOContext) Rb64() uint64 {
	return avformat.AvioRb64(avioCtx.CAVIOContext)
}

func (avioCtx *AVIOContext) GetStr(maxlen int, buf unsafe.Pointer, buflen int) int {
	return avformat.AvioGetStr(avioCtx.CAVIOContext, maxlen, buf, buflen)
}

func (avioCtx *AVIOContext) GetStr16le(maxlen int, buf unsafe.Pointer, buflen int) int {
	return avformat.AvioGetStr16le(avioCtx.CAVIOContext, maxlen, buf, buflen)
}

func (avioCtx *AVIOContext) GetStr16be(maxlen int, buf unsafe.Pointer, buflen int) int {
	return avformat.AvioGetStr16be(avioCtx.CAVIOContext, maxlen, buf, buflen)
}

func (avioCtx *AVIOContext) Close() int {
	ret := avformat.AvioClose(avioCtx.CAVIOContext)
	if ret == 0 {
		avioCtx.CAVIOContext = nil
	}
	return ret
}

func (avioCtx *AVIOContext) GetDynBuf(pbuffer unsafe.Pointer) int {
	return avformat.AvioGetDynBuf(avioCtx.CAVIOContext, pbuffer)
}

func (avioCtx *AVIOContext) CloseDynBuf(pbuffer unsafe.Pointer) int {
	return avformat.AvioCloseDynBuf(avioCtx.CAVIOContext, pbuffer)
}

func (avioCtx *AVIOContext) Pause(pause int) int {
	return avformat.AvioPause(avioCtx.CAVIOContext, pause)
}

func (avioCtx *AVIOContext) SeekTime(streamIndex int, timestamp int64, flags int) int64 {
	return avformat.AvioSeekTime(avioCtx.CAVIOContext, streamIndex, timestamp, flags)
}

func (avioCtx *AVIOContext) Accept() (*AVIOContext, int) {
	client := AVIOContext{}
	ret := avformat.AvioAccept(avioCtx.CAVIOContext, &client.CAVIOContext)
	if ret < 0 {
		return nil, ret
	}
	return &client, ret
}

func (avioCtx *AVIOContext) Handshake() int {
	return avformat.AvioHandshake(avioCtx.CAVIOContext)
}

func (avioCtx *AVIOContext) GetPacket(pkt *goavcodec.AVPacket, size int) int {
	return avformat.AvGetPacket(avioCtx.CAVIOContext, pkt.CAVPacket, size)
}

func (avioCtx *AVIOContext) AppendPacket(pkt *goavcodec.AVPacket, size int) int {
	return avformat.AvAppendPacket(avioCtx.CAVIOContext, pkt.CAVPacket, size)
}

func AvioOpen(url string, flags int) (*AVIOContext, int) {
	avioCtx := AVIOContext{}
	ret := avformat.AvioOpen(&avioCtx.CAVIOContext, url, flags)
	if ret < 0 {
		return nil, ret
	}
	return &avioCtx, ret
}

func AvioOpen2(url string, flags int, intCb *avformat.CAVIOInterruptCB, options **avutil.CAVDictionary) (*AVIOContext, int) {
	avioCtx := AVIOContext{}
	ret := avformat.AvioOpen2(&avioCtx.CAVIOContext, url, flags, intCb, options)
	if ret < 0 {
		return nil, ret
	}
	return &avioCtx, ret
}

func OpenDynBuf() (*AVIOContext, int) {
	avioCtx := AVIOContext{}
	ret := avformat.AvioOpenDynBuf(&avioCtx.CAVIOContext)
	if ret != 0 {
		return nil, ret
	}
	return &avioCtx, ret
}

func GetAvailableProtocols(output bool) []string {
	p := unsafe.Pointer(nil)
	arr := []string{}
	outputInt := 0
	if output {
		outputInt = 1
	}
	for {
		name := avformat.AvioEnumProtocols(&p, outputInt)
		if len(name) <= 0 {
			break
		}
		arr = append(arr, name)
	}
	return arr
}
