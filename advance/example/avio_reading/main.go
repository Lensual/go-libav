package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/Lensual/go-libav/advance/goavformat"
	"github.com/Lensual/go-libav/avutil"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s input_file\n"+
			"API example program to show how to read from a custom buffer "+
			"accessed through AVIOContext.\n", os.Args[0])
		os.Exit(1)
	}

	input_filename := os.Args[1]

	fd, err := os.Open(input_filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	fmtCtx := goavformat.AllocAvformatContext()
	if fmtCtx == nil {
		panic(syscall.ENOMEM)
	}
	defer fmtCtx.CloseInput()

	avioCtx := goavformat.NewGoAvioContext(4096, func(buf []byte, buf_size int) int {
		//read_packet
		count, err := fd.Read(buf)
		if err != nil {
			return avutil.AVERROR_EOF
		}
		return count

	}, nil, nil)
	if avioCtx == nil {
		panic(syscall.ENOMEM)
	}

	fmtCtx.SetIOContext(avioCtx.AVIOContext)
	ret := fmtCtx.OpenInput("", nil, nil)
	if ret < 0 {
		panic("Could not open input\n")
	}
	defer avioCtx.Free()

	ret = fmtCtx.FindStreamInfo(nil)
	if ret < 0 {
		panic("Could not find stream information\n")
	}

	fmtCtx.DumpFormat(0, input_filename, 0)
}
