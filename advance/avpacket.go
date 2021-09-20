package advance

import (
	"errors"
	"syscall"

	"github.com/Lensual/go-libav/avcodec"
	"github.com/Lensual/go-libav/avutil"
)

type AvPacket struct {
	CAvPacket *avcodec.CAVPacket
}

func NewAVPacket(size int) (*AvPacket, error) {
	var cAVPacket *avcodec.CAVPacket
	if size > 0 {
		code := avcodec.AvNewPacket(cAVPacket, size)
		if code != 0 {
			return nil, errors.New(avutil.Err2str(code))
		}
	} else {
		cAVPacket = avcodec.AvPacketAlloc()
		if cAVPacket == nil {
			return nil, errors.New(syscall.ENOMEM.Error())
		}
	}

	return &AvPacket{
		CAvPacket: cAVPacket,
	}, nil
}
