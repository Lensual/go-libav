package avfilter_test

import (
	"testing"

	"github.com/Lensual/go-libav/avfilter"
)

func TestVersion(t *testing.T) {
	t.Log(avfilter.Version())
}
