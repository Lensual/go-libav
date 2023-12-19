package goavformat

import (
	"unsafe"

	"github.com/Lensual/go-libav/advance/goavcodec"
	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
)

type AvformatContext struct {
	CAvformatContext *avformat.CAVFormatContext
}

//#region members

func (fmtCtx *AvformatContext) SetIOContext(avioCtx *AVIOContext) {
	fmtCtx.CAvformatContext.SetPB(avioCtx.CAVIOContext)
}

func (fmtCtx *AvformatContext) GetStreams() []*avformat.CAVStream {
	return unsafe.Slice(fmtCtx.CAvformatContext.GetStreams(), fmtCtx.CAvformatContext.GetNbStreams())
}

//#endregion members

func (fmtCtx *AvformatContext) Free() {
	avformat.AvformatFreeContext(fmtCtx.CAvformatContext)
}

func (fmtCtx *AvformatContext) OpenInput(url string, fmt *avformat.CAVInputFormat, options **avutil.CAVDictionary) int {
	return avformat.AvformatOpenInput(&fmtCtx.CAvformatContext, url, fmt, options)
}

func (fmtCtx *AvformatContext) FindStreamInfo(options **avutil.CAVDictionary) int {
	return avformat.AvformatFindStreamInfo(fmtCtx.CAvformatContext, options)
}

func (fmtCtx *AvformatContext) FindBestStream(_type avutil.CAVMediaType, wanted_stream_nb int,
	related_stream int, flags int) (int, *goavcodec.AVCodec) {
	dec := goavcodec.AVCodec{}
	ret := avformat.AvFindBestStream(fmtCtx.CAvformatContext, _type, wanted_stream_nb, related_stream, &dec.CAVCodec, flags)
	return ret, &dec
}

func (fmtCtx *AvformatContext) ReadFrame(pkt *goavcodec.AVPacket) int {
	return avformat.AvReadFrame(fmtCtx.CAvformatContext, pkt.CAVPacket)
}

func (fmtCtx *AvformatContext) CloseInput() {
	avformat.AvformatCloseInput(&fmtCtx.CAvformatContext)
	fmtCtx.CAvformatContext = nil
}

func (fmtCtx *AvformatContext) DumpFormat(index int, url string, is_output int) {
	avformat.AvDumpFormat(fmtCtx.CAvformatContext, index, url, is_output)
}

func GetAvailableMuxer() []*avformat.CAVOutputFormat {
	p := unsafe.Pointer(nil)
	arr := []*avformat.CAVOutputFormat{}
	for {
		muxer := avformat.AvMuxerIterate(&p)
		if muxer == nil {
			break
		}
		arr = append(arr, muxer)
	}
	return arr
}

func GetAvailableDeuxer() []*avformat.CAVInputFormat {
	p := unsafe.Pointer(nil)
	arr := []*avformat.CAVInputFormat{}
	for {
		demuxer := avformat.AvDeuxerIterate(&p)
		if demuxer == nil {
			break
		}
		arr = append(arr, demuxer)
	}
	return arr
}

func AllocAvformatContext() *AvformatContext {
	ctx := avformat.AvformatAllocContext()
	if ctx == nil {
		return nil
	}
	return &AvformatContext{
		CAvformatContext: ctx,
	}
}

func OpenInput(url string, fmt *avformat.CAVInputFormat, options **avutil.CAVDictionary) (*AvformatContext, int) {
	fmtCtx := AvformatContext{}
	ret := avformat.AvformatOpenInput(&fmtCtx.CAvformatContext, url, fmt, options)
	if ret != 0 {
		return nil, ret
	}
	return &fmtCtx, ret
}
