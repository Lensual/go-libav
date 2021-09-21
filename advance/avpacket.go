package advance

import (
	"errors"

	"github.com/Lensual/go-libav/avcodec"
	"github.com/Lensual/go-libav/avutil"
)

type AvPacket struct {
	CAvPacket *avcodec.CAVPacket
}

func NewAVPacket(size int) (*AvPacket, error) {
	var cAVPacket *avcodec.CAVPacket
	code := avcodec.AvNewPacket(cAVPacket, size)
	if code != 0 {
		return nil, errors.New(avutil.Err2str(code))
	}

	return &AvPacket{
		CAvPacket: cAVPacket,
	}, nil
}
