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

func (parser *AVCodecParserContext) ParseUnsafe(avctx *AVCodecContext,
	poutbuf unsafe.Pointer, poutbufSize *ctypes.Int,
	buf unsafe.Pointer, bufSize int,
	pts int64, dts int64, pos int64) int {

	return avcodec.AvParserParse2(parser.CAVCodecParserContext, avctx.CAVCodecContext,
		poutbuf, poutbufSize,
		buf, bufSize,
		pts, dts, pos)
}

func (parser *AVCodecParserContext) ParsePktUnsafe(avctx *AVCodecContext, pkt *AVPacket,
	data unsafe.Pointer, length int,
	pts int64, dts int64, pos int64) int {

	return parser.ParseUnsafe(avctx,
		pkt.CAVPacket.GetData(), pkt.CAVPacket.GetSizePtr(),
		data, length, pts, dts, pos)
}

func (parser *AVCodecParserContext) ParsePkt(avctx *AVCodecContext, pkt *AVPacket,
	data []byte,
	pts int64, dts int64, pos int64) int {

	length := len(data)
	cData := avutil.AvMalloc(ctypes.SizeT(length))
	defer avutil.AvFree(cData)
	copy(unsafe.Slice((*byte)(cData), length), data)

	return parser.ParsePktUnsafe(avctx, pkt, cData, length, pts, dts, pos)
}

func (parser *AVCodecParserContext) Parse(avctx *AVCodecContext, data []byte,
	pts int64, dts int64, pos int64) (*AVPacket, int) {

	pkt := AllocAvPacket()
	if pkt == nil {
		return nil, int(syscall.ENOMEM)
	}

	return pkt, parser.ParsePkt(avctx, pkt, data, pts, dts, pos)
}

func (parser *AVCodecParserContext) Free() {
	avcodec.AvParserClose(parser.CAVCodecParserContext)
}
