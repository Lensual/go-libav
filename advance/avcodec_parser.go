package advance

import (
	"unsafe"

	"github.com/Lensual/go-libav/avcodec"
)

type AVCodecParserContext struct {
	CAVCodecParserContext *avcodec.CAVCodecParserContext
}

func NewAVCodecParserContext(codecId uint32) *AVCodecParserContext {
	parser := avcodec.AvParserInit(int(codecId))
	if parser == nil {
		return nil
	}
	return &AVCodecParserContext{
		CAVCodecParserContext: parser,
	}
}

func (parser *AVCodecParserContext) Parse(avctx *AVCodecContext, pkt *AvPacket,
	data unsafe.Pointer, length int64,
	pts int64, dts int64, pos int64) {

	avcodec.AvParserParse(parser.CAVCodecParserContext, avctx.CAVCodecContext,
		pkt.CAvPacket.GetData(), pkt.CAvPacket.GetSize(),
		data, int(length), pts, dts, pos)
}

func (parser *AVCodecParserContext) Free() {
	avcodec.AvParserClose(parser.CAVCodecParserContext)
}
