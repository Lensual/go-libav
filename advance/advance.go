package advance

import (
	"github.com/Lensual/go-libav/avformat"
	"github.com/Lensual/go-libav/avutil"
)

func NewCustomAVIO(bufSize int, writeFlag bool) *avformat.AVIOContext {
	buf := avutil.AvMalloc(uint64(bufSize))
	var wf int
	if writeFlag {
		wf = 1
	}

	//TODO

	return avformat.AvioAllocContext(buf, bufSize, wf, nil, nil, nil, nil) //TODO

}
