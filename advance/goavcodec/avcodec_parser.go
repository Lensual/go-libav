package goavcodec

import (
	"unsafe"

	"github.com/Lensual/go-libav/avcodec"
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

func (parser *AVCodecParserContext) Parse(avctx *AVCodecContext, pkt *AVPacket,
	data unsafe.Pointer, length int,
	pts int64, dts int64, pos int64) int {

	return avcodec.AvParserParse2(parser.CAVCodecParserContext, avctx.CAVCodecContext,
		pkt.CAVPacket.GetData(), pkt.CAVPacket.GetSize(),
		data, length, pts, dts, pos)
}

func (parser *AVCodecParserContext) Free() {
	avcodec.AvParserClose(parser.CAVCodecParserContext)
}
