package goavcodec_test

import (
	"testing"

	"github.com/Lensual/go-libav/advance/goavcodec"
	"github.com/Lensual/go-libav/avcodec"
)

func TestAvailableCodecs(t *testing.T) {
	availableCodecs := goavcodec.GetAvailableCodecs()
	availableCodecsNames := []string{}
	for _, v := range availableCodecs {
		availableCodecsNames = append(availableCodecsNames, v.CAVCodec.GetName())
	}

	t.Logf("AvailableCodecs:\t%v", availableCodecsNames)
}

func TestEncoder(t *testing.T) {
	enc := goavcodec.FindEncoder(avcodec.AV_CODEC_ID_AAC)
	supportedSamplerates := enc.GetSupportedSamplerates()
	t.Logf("SupportedSamplerates:\t%v", supportedSamplerates)
}
