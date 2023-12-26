package goavcodec

import (
	"syscall"
	"unsafe"

	"github.com/Lensual/go-libav/avcodec"
	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/ctypes"
)

type AVCodecParserContext struct {
	CAVCodecParserContext *avcodec.CAVCodecParserContext
}

func NewAVCodecParserContext(codecId int) *AVCodecParserContext {
	parser := avcodec.AvParserInit(codecId)
	if parser == nil {
		return nil
	}
	return &AVCodecParserContext{
		CAVCodecParserContext: parser,
	}
}

func (parser *AVCodecParserContext) ParseUnsafeToUnsafe(avctx *AVCodecContext,
	poutbuf unsafe.Pointer, poutbufSize *ctypes.Int,
	buf unsafe.Pointer, bufSize int,
	pts int64, dts int64, pos int64) int {

	return avcodec.AvParserParse2(parser.CAVCodecParserContext, avctx.CAVCodecContext,
		poutbuf, poutbufSize,
		buf, bufSize,
		pts, dts, pos)
}

func (parser *AVCodecParserContext) ParseUnsafeTo(avctx *AVCodecContext,
	pkt *AVPacket,
	buf unsafe.Pointer, bufSize int,
	pts int64, dts int64, pos int64) int {

	return parser.ParseUnsafeToUnsafe(avctx,
		pkt.CAVPacket.GetData(), pkt.CAVPacket.GetSizePtr(),
		buf, bufSize,
		pts, dts, pos)
}

func (parser *AVCodecParserContext) ParseToUnsafe(avctx *AVCodecContext,
	poutbuf unsafe.Pointer, poutbufSize *ctypes.Int,
	buf []byte,
	pts int64, dts int64, pos int64) int {

	length := len(buf)
	cBuf := avutil.AvMalloc(ctypes.SizeT(length))
	defer avutil.AvFree(cBuf)
	copy(unsafe.Slice((*byte)(cBuf), length), buf)

	return parser.ParseUnsafeToUnsafe(avctx,
		poutbuf, poutbufSize,
		cBuf, len(buf),
		pts, dts, pos)
}

func (parser *AVCodecParserContext) ParseTo(avctx *AVCodecContext,
	pkt *AVPacket, buf []byte,
	pts int64, dts int64, pos int64) int {

	return parser.ParseToUnsafe(avctx,
		pkt.CAVPacket.GetData(), pkt.CAVPacket.GetSizePtr(),
		buf,
		pts, dts, pos)
}

func (parser *AVCodecParserContext) Parse(avctx *AVCodecContext, data []byte,
	pts int64, dts int64, pos int64) (*AVPacket, int) {

	pkt := AllocAvPacket()
	if pkt == nil {
		return nil, int(syscall.ENOMEM)
	}

	ret := parser.ParseTo(avctx,
		pkt,
		data,
		pts, dts, pos)

	if ret != 0 {
		pkt.Free()
		return nil, ret
	}

	if pkt.CAVPacket.GetData() == nil {
		pkt = nil
	}

	return pkt, ret
}

func (parser *AVCodecParserContext) Free() {
	avcodec.AvParserClose(parser.CAVCodecParserContext)
}
