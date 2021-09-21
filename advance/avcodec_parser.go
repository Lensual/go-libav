package advance

import "github.com/Lensual/go-libav/avcodec"

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
