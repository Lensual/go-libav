package advance_test

import (
	"testing"

	"github.com/Lensual/go-libav/avcodec"
	"github.com/Lensual/go-libav/avdevice"
	"github.com/Lensual/go-libav/avfilter"
	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/swresample"
	"github.com/Lensual/go-libav/swscale"
)

func TestVersion(t *testing.T) {
	t.Log(avutil.Version())
	t.Log(avcodec.Version())
	t.Log(avformat.Version())
	t.Log(avdevice.Version())
	t.Log(avfilter.Version())
	t.Log(swscale.Version())
	t.Log(swresample.Version())
}
