package goavcodec

import (
	"github.com/Lensual/go-libav/avcodec"
)

type AVPacket struct {
	CAVPacket *avcodec.CAVPacket
}

func (pkt *AVPacket) Free() {
	avcodec.AvPacketFree(&pkt.CAVPacket)
}

func (pkt *AVPacket) Unref() {
	avcodec.AvPacketUnref(pkt.CAVPacket)
}

func NewAvPacket(size int) (*AVPacket, int) {
	var cAVPacket *avcodec.CAVPacket
	code := avcodec.AvNewPacket(cAVPacket, size)
	if code != 0 {
		return nil, code
	}

	return &AVPacket{
		CAVPacket: cAVPacket,
	}, code
}

func AllocAvPacket() *AVPacket {
	return &AVPacket{
		CAVPacket: avcodec.AvPacketAlloc(),
	}
}
