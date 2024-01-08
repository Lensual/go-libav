package goavcodec

/*
#include <stdlib.h>
*/
import "C"

import (
	"context"
	"syscall"
	"unsafe"

	"github.com/Lensual/go-libav/advance/goavutil"
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
	poutbuf *unsafe.Pointer, poutbufSize *ctypes.Int,
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
		pkt.CAVPacket.GetDataPtr(), pkt.CAVPacket.GetSizePtr(),
		buf, bufSize,
		pts, dts, pos)
}

// buf: do not need padding
func (parser *AVCodecParserContext) ParseToUnsafe(avctx *AVCodecContext,
	poutbuf *unsafe.Pointer, poutbufSize *ctypes.Int,
	buf []byte,
	pts int64, dts int64, pos int64) int {

	var cBuf unsafe.Pointer = nil
	if buf != nil {
		cBuf = avutil.AvMallocz(ctypes.SizeT((len(buf) + avcodec.AV_INPUT_BUFFER_PADDING_SIZE)))
		copy(unsafe.Slice((*byte)(cBuf), len(buf)), buf)
		defer C.free(cBuf)
	}

	return parser.ParseUnsafeToUnsafe(avctx,
		poutbuf, poutbufSize,
		cBuf, len(buf),
		pts, dts, pos)
}

// buf: do not need padding
func (parser *AVCodecParserContext) ParseTo(avctx *AVCodecContext,
	pkt *AVPacket, buf []byte,
	pts int64, dts int64, pos int64) int {

	return parser.ParseToUnsafe(avctx,
		pkt.CAVPacket.GetDataPtr(), pkt.CAVPacket.GetSizePtr(),
		buf,
		pts, dts, pos)
}

// buf: do not need padding
func (parser *AVCodecParserContext) Parse(avctx *AVCodecContext, buf []byte,
	pts int64, dts int64, pos int64) (*AVPacket, int) {

	var outbuf unsafe.Pointer
	var outbufSize ctypes.Int
	code := parser.ParseToUnsafe(avctx,
		&outbuf, &outbufSize,
		buf,
		pts, dts, pos)

	if code < 0 {
		//has error
		if outbuf != nil {
			//has output
			C.free(outbuf)
		}
		return nil, code
	}

	var pkt *AVPacket
	if outbufSize > 0 {
		//has output
		pkt = AllocAvPacket()
		if pkt == nil {
			C.free(outbuf)
			return nil, int(syscall.ENOMEM)
		}
		pkt.CAVPacket.SetData(outbuf)
		pkt.CAVPacket.SetSize(int(outbufSize))
	}

	return pkt, code
}

// buf: do not need padding
func (parser *AVCodecParserContext) ParseChan(ctx context.Context, avctx *AVCodecContext, bufCh <-chan []byte) (context.Context, <-chan *AVPacket) {
	ctx, cancel := context.WithCancelCause(ctx)
	pktCh := make(chan *AVPacket)

	go func() {
		defer close(pktCh)
	loop:
		for {
			select {
			case <-ctx.Done():
				return
			case buf, ok := <-bufCh:
				if !ok {
					break loop
				}
				inLen := len(buf)
				inBuf := buf
				for inLen > 0 {
					pkt, code := parser.Parse(avctx, inBuf, avutil.AV_NOPTS_VALUE, avutil.AV_NOPTS_VALUE, 0)
					if code < 0 {
						if pkt != nil {
							pkt.Unref()
						}
						cancel(goavutil.AvErr(code))
						return
					}

					bufUsed := code
					inBuf = inBuf[bufUsed:]
					inLen -= bufUsed

					if pkt != nil {
						select {
						case <-ctx.Done():
							pkt.Unref()
							return
						case pktCh <- pkt:
						}
					}
				}
			}
		}
	}()

	return ctx, pktCh
}

// Free() is alias Close()
func (parser *AVCodecParserContext) Free() {
	parser.Close()
}

func (parser *AVCodecParserContext) Close() {
	if parser == nil || parser.CAVCodecParserContext == nil {
		return
	}
	avcodec.AvParserClose(parser.CAVCodecParserContext)
	parser.CAVCodecParserContext = nil
}
