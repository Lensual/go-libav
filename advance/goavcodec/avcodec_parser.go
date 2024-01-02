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

func (parser *AVCodecParserContext) ParseToUnsafe(avctx *AVCodecContext,
	poutbuf *unsafe.Pointer, poutbufSize *ctypes.Int,
	buf []byte,
	pts int64, dts int64, pos int64) int {

	var cBuf unsafe.Pointer = nil
	if buf != nil {
		cBuf = C.CBytes(buf)
		defer C.free(cBuf)
	}

	code := parser.ParseUnsafeToUnsafe(avctx,
		poutbuf, poutbufSize,
		cBuf, len(buf),
		pts, dts, pos)

	if code < 0 {
		//has error
		return code
	}

	if *poutbufSize > 0 {
		//has output
		if *poutbuf == cBuf {
			*poutbuf = avutil.AvMemdup(*poutbuf, ctypes.SizeT(*poutbufSize))
		}
	}

	return code
}

func (parser *AVCodecParserContext) ParseTo(avctx *AVCodecContext,
	pkt *AVPacket, buf []byte,
	pts int64, dts int64, pos int64) int {

	return parser.ParseToUnsafe(avctx,
		pkt.CAVPacket.GetDataPtr(), pkt.CAVPacket.GetSizePtr(),
		buf,
		pts, dts, pos)
}

func (parser *AVCodecParserContext) Parse(avctx *AVCodecContext, buf []byte,
	pts int64, dts int64, pos int64) (*AVPacket, int) {

	var dataPtr unsafe.Pointer
	var sizePtr ctypes.Int
	code := parser.ParseToUnsafe(avctx,
		&dataPtr, &sizePtr,
		buf,
		pts, dts, pos)

	if code < 0 {
		//has error
		if dataPtr != nil {
			//has output
			C.free(dataPtr)
		}
		return nil, code
	}

	var pkt *AVPacket
	if sizePtr > 0 {
		//has output
		pkt = AllocAvPacket()
		if pkt == nil {
			C.free(dataPtr)
			return nil, int(syscall.ENOMEM)
		}
		pkt.CAVPacket.SetData(dataPtr)
		pkt.CAVPacket.SetSize(int(sizePtr))
	}

	return pkt, code
}

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

					if code < 0 {
						cancel(goavutil.AvErr(code))
						return
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
