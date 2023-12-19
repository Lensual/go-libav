package advance_test

import (
	"testing"

	"github.com/Lensual/go-libav/advance/goavformat"
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

func TestConfiguration(t *testing.T) {
	t.Logf("AvformatConfiguration:\t%v", avformat.AvformatConfiguration())
	t.Logf("SwscaleConfiguration:\t%v", swscale.SwscaleConfiguration())
	t.Logf("SwresampleConfiguration:\t%v", swresample.SwresampleConfiguration())
}

func TestLicense(t *testing.T) {
	t.Logf("AvformatLicense:\t%v", avformat.AvformatLicense())
	t.Logf("SwscaleLicense:\t%v", swscale.SwscaleLicense())
	t.Logf("SwresampleLicense:\t%v", swresample.SwresampleLicense())
}

func TestAvailableProtocols(t *testing.T) {
	t.Logf("AvailableInputProtocols:\t\t%v", goavformat.GetAvailableProtocols(false))
	t.Logf("AvailableOutputProtocols:\t%v", goavformat.GetAvailableProtocols(true))
}

func TestAvailableMuxer(t *testing.T) {
	availableMuxer := goavformat.GetAvailableMuxer()
	availableMuxersNames := []string{}
	for _, v := range availableMuxer {
		availableMuxersNames = append(availableMuxersNames, v.GetName())
	}
	t.Logf("AvailableMuxer:\t%v", availableMuxersNames)

	availableDemuxer := goavformat.GetAvailableMuxer()
	availableDemuxersNames := []string{}
	for _, v := range availableDemuxer {
		availableDemuxersNames = append(availableDemuxersNames, v.GetName())
	}
	t.Logf("AvailablDemuxer:\t%v", availableDemuxersNames)
}
