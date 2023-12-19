package main

import (
	"fmt"
	"syscall"

	"github.com/Lensual/go-libav/advance/goavcodec"
	"github.com/Lensual/go-libav/advance/goavformat"
	"github.com/Lensual/go-libav/avutil"
)

func main() {
	fmtCtx, _ := goavformat.OpenInput("/home/lensual/test_decode.mp4", nil, nil)
	defer fmtCtx.CloseInput()

	fmtCtx.DumpFormat(0, "", 0)

	streamIdx, dec := fmtCtx.FindBestStream(avutil.AVMEDIA_TYPE_VIDEO, -1, -1, 0)
	decCtx := dec.CreateContext()
	decCtx.ParametersFrom(fmtCtx.GetStreams()[streamIdx].GetCodecPar())
	decCtx.Open(nil)

	for {
		pkt := goavcodec.AllocAvPacket()
		ret := fmtCtx.ReadFrame(pkt)

		if pkt.CAVPacket.GetStreamIndex() == streamIdx {
			if ret != 0 {
				pkt.Unref()
				break
			}
			frames, ret := decCtx.Decode(pkt)
			if ret != 0 {
				if ret != avutil.AVERROR(int(syscall.EAGAIN)) {
					panic(ret)
				}
			}
			for _, frame := range frames {
				fmt.Println(frame.GetPts())
				frame.Unref()
			}
		}
		pkt.Unref()
	}

}
