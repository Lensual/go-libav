package avcodec_test

import (
	"testing"

	"github.com/Lensual/go-libav/avcodec"
)

func TestVersion(t *testing.T) {
	t.Log(avcodec.Version())
}
