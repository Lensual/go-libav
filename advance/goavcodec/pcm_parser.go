package goavcodec

import (
	"context"

	"github.com/Lensual/go-libav/advance/goavutil"
)

func ParsePcmChan(ctx context.Context, dataInCh <-chan []byte) (context.Context, <-chan *AVPacket) {
	ctx, cancel := context.WithCancelCause(ctx)
	parsedPktCh := make(chan *AVPacket)
	go func() {
		defer close(parsedPktCh)
	loop:
		for {
			select {
			case <-ctx.Done():
				return
			case data, ok := <-dataInCh:
				if !ok {
					break loop
				}
				pkt, code := NewAvPacketFromData(data)
				if code < 0 {
					cancel(goavutil.AvErr(code))
					return
				}
				select {
				case <-ctx.Done():
					pkt.Unref()
					return
				case parsedPktCh <- pkt:
				}
			}
		}
	}()
	return ctx, parsedPktCh
}
