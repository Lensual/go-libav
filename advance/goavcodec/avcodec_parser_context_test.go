package goavcodec_test

import (
	"context"
	"io"
	"os"
	"syscall"
	"testing"

	"github.com/Lensual/go-libav/advance/goavcodec"
	"github.com/Lensual/go-libav/avcodec"
)

func TestParseChan(t *testing.T) {
	codec := goavcodec.FindEncoder(avcodec.AV_CODEC_ID_MP2)
	codecCtx := codec.CreateContext()
	if codecCtx == nil {
		panic(syscall.ENOMEM)
	}
	defer codecCtx.Free()

	parserCtx := codec.CreateParserContext()
	if parserCtx == nil {
		panic(syscall.ENOMEM)
	}
	defer parserCtx.Free()

	inFile, err := os.OpenFile("/data/lensual/git.bjxctec.com/paas/asr/test/test_sin_48000_2ch_s16p.mp2", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	bufCh := make(chan []byte)
	go func() {
		defer inFile.Close()
		defer close(bufCh)
		for {
			buf := make([]byte, 1) //size 1 for test
			count, err := inFile.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			bufCh <- buf[:count]
		}
		println("read done")
	}()

	_, pktCh := parserCtx.ParseChan(context.Background(), codecCtx, bufCh)

	for pkt := range pktCh {
		pkt.CAVPacket.GetDataPtr()
	}

}
