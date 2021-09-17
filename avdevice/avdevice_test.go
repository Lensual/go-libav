package avdevice_test

import (
	"testing"

	"github.com/Lensual/go-libav/avdevice"
)

func TestVersion(t *testing.T) {
	t.Log(avdevice.Version())
}
