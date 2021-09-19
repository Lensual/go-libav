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
	t.Log(avutil.LIBAVUTIL_IDENT)
	t.Log(avcodec.LIBAVCODEC_IDENT)
	t.Log(avformat.LIBAVFORMAT_IDENT)
	t.Log(avdevice.LIBAVDEVICE_IDENT)
	t.Log(avfilter.LIBAVFILTER_IDENT)
	t.Log(swscale.LIBSWSCALE_IDENT)
	t.Log(swresample.LIBSWRESAMPLE_IDENT)
}
