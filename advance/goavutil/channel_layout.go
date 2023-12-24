package goavutil

import (
	"github.com/Lensual/go-libav/avutil"
)

type AVChannelLayout struct {
	CAVChannelLayout *avutil.CAVChannelLayout
}

// #region member

func (layout *AVChannelLayout) GetNbChannels() int {
	return layout.CAVChannelLayout.GetNbChannels()
}

//#endregion member

func GetAvChannelLayoutDefault(nbChannels int) *AVChannelLayout {
	var c avutil.CAVChannelLayout
	avutil.AvChannelLayoutDefault(&c, nbChannels)
	return &AVChannelLayout{
		CAVChannelLayout: &c,
	}
}
