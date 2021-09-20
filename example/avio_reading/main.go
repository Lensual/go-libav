package main

import (
	"fmt"
	"os"
	"reflect"
	"syscall"
	"unsafe"

	"github.com/Lensual/go-libav/advance"
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

	fmtCtx := advance.NewAvformatContext()
	if fmtCtx == nil {
		panic(syscall.ENOMEM)
	}
	defer fmtCtx.CloseInput()

	avioCtx := advance.NewAvioContext(4096, func(cbuf unsafe.Pointer, buf_size int) int {
		//read_packet

		//映射C的内存
		var buf []byte
		h := (*reflect.SliceHeader)((unsafe.Pointer(&buf)))
		h.Cap = buf_size
		h.Len = buf_size
		h.Data = uintptr(cbuf)

		count, err := fd.Read(buf)
		if err != nil {
			return avutil.AVERROR_EOF
		}
		return count

	}, nil, nil)
	if avioCtx == nil {
		panic(syscall.ENOMEM)
	}

	fmtCtx.SetIOContext(avioCtx)
	ret := fmtCtx.OpenInput("")
	if ret < 0 {
		panic("Could not open input\n")
	}
	defer avioCtx.Free()

	ret = fmtCtx.FindStreamInfo()
	if ret < 0 {
		panic("Could not find stream information\n")
	}

	fmtCtx.DumpFormat(0, input_filename, 0)
}
