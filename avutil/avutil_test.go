package avutil_test

import (
	"testing"

	"github.com/Lensual/go-libav/avutil"
)

func TestVersion(t *testing.T) {
	t.Log(avutil.Version())
}
