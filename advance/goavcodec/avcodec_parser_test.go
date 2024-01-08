package goavcodec_test

import (
	"testing"

	"github.com/Lensual/go-libav/advance/goavcodec"
)

// TODO
func TestAvailableParsers(t *testing.T) {
	availableParsers := goavcodec.GetAvailableParsers()
	availableParsersNames := []any{}
	for _, v := range availableParsers {
		availableParsersNames = append(availableParsersNames, v.CAVCodecParser.GetCodecIds())
	}

	t.Logf("AvailableParsers:\t%v", availableParsersNames)
}
