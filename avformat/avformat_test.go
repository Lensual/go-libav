package avformat_test

import (
	"testing"

	"github.com/Lensual/go-libav/avformat"
)

func TestVersion(t *testing.T) {
	t.Log(avformat.Version())
}
