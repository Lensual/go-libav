// Example for goswresample
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Lensual/go-libav/advance/goavutil"
	"github.com/Lensual/go-libav/advance/goswresample"
	"github.com/Lensual/go-libav/avutil"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s input_file output_file\n"+
			"Input 44100hz 2ch pcm_s16le file.\n"+
			"Output 8000hz 1ch pcm_s16le file.\n",
			os.Args[0])
		os.Exit(1)
	}

	inFileName := os.Args[1]
	outFileName := os.Args[2]

	inFile, err := os.OpenFile(inFileName, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	outChLayout := goavutil.GetAvChannelLayoutDefault(1)
	outSampleFmt := avutil.AV_SAMPLE_FMT_S16
	outSampleRate := 8000

	inChLayout := goavutil.GetAvChannelLayoutDefault(2)
	inSampleFmt := avutil.AV_SAMPLE_FMT_S16
	inSampleRate := 44100

	swrCtx, _ := goswresample.NewSwrContextWithOpts(outChLayout, outSampleFmt, outSampleRate, inChLayout, inSampleFmt, inSampleRate)
	defer swrCtx.Free()

	code := swrCtx.Init()
	if code != 0 {
		panic(goavutil.AvErr(code))
	}

	inCount := 1024
	inBufSize := avutil.AvSamplesGetBufferSize(nil, inChLayout.GetNbChannels(), inCount, inSampleFmt, 0)
	in := make([]byte, inBufSize)
	for {
		_, err := inFile.Read(in)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		out, outCount := swrCtx.Convert(in, inCount)
		if outCount < 0 {
			panic(goavutil.AvErr(outCount))
		}
		if outCount > 0 {
			_, err = outFile.Write(out)
			if err != nil {
				panic(err)
			}
		}
	}

	out, outCount := swrCtx.Convert(nil, 0)
	if outCount < 0 {
		panic(goavutil.AvErr(outCount))
	}
	if outCount > 0 {
		_, err = outFile.Write(out)
		if err != nil {
			panic(err)
		}
	}
}
