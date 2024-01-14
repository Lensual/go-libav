package goavformat

import (
	"errors"
	"io"

	"github.com/Lensual/go-libav/avutil"
)

// GoAVIOContext2 is Go API for AVIOContext.
// The field opaque of AVIOContext is occupied by cgo.Handle.
//
// Based from GoAVIOContext.
type GoAVIOContext2 struct {
	*GoAVIOContext
	reader io.Reader
	writer io.Writer
	seeker io.Seeker
}

// NewGoAvioContext2
//
// unsafeBuf: True if callback with unsafe buffer, False if callback with copy memory to Go heap.
//
// return: Allocated GoAVIOContext2 or nil on failure.
func NewGoAvioContext2(bufSize int, reader io.Reader, writer io.Writer, seeker io.Seeker, unsafeBuf bool) *GoAVIOContext2 {
	avioCtx := &GoAVIOContext2{
		reader: reader,
		writer: writer,
		seeker: seeker,
	}

	var readFunc AVIOReadCallback = nil
	if reader != nil {
		readFunc = func(avioCtx *GoAVIOContext, buf []byte) int {
			ret, err := reader.Read(buf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					return avutil.AVERROR_EOF
				}
				return avutil.AVERROR_UNKNOWN
			}
			return ret
		}
	}

	var writeFunc AVIOWriteCallback = nil
	if writer != nil {
		writeFunc = func(avioCtx *GoAVIOContext, buf []byte) int {
			ret, err := writer.Write(buf)
			if err != nil {
				return avutil.AVERROR_UNKNOWN
			}
			return ret
		}
	}

	var seekFunc AVIOSeekCallback = nil
	if seeker != nil {
		seekFunc = func(avioCtx *GoAVIOContext, offset int64, whence int) int64 {
			ret, err := seeker.Seek(offset, whence)
			if err != nil {
				return avutil.AVERROR_UNKNOWN
			}
			return ret
		}
	}

	avioCtx.GoAVIOContext = NewGoAvioContext(bufSize, readFunc, writeFunc, seekFunc, unsafeBuf)
	if avioCtx.GoAVIOContext == nil {
		return nil
	}

	return avioCtx
}

// Free all resources allocated by GoAVIOContext2.
func (avioCtx *GoAVIOContext2) Free() {
	avioCtx.GoAVIOContext.Free()
	avioCtx.GoAVIOContext = nil

	avioCtx.reader = nil
	avioCtx.writer = nil
	avioCtx.seeker = nil
}
