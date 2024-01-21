package goavutil

import (
	"unsafe"

	"github.com/Lensual/go-libav/avutil"
)

type AVChannelLayout struct {
	CAVChannelLayout *avutil.CAVChannelLayout
}

// #region member

func (layout *AVChannelLayout) GetOrder() avutil.CAVChannelOrder {
	return layout.CAVChannelLayout.GetOrder()
}
func (layout *AVChannelLayout) SetOrder(order avutil.CAVChannelOrder) {
	layout.CAVChannelLayout.SetOrder(order)
}

func (layout *AVChannelLayout) GetNbChannels() int {
	return layout.CAVChannelLayout.GetNbChannels()
}
func (layout *AVChannelLayout) SetNbChannels(nbChannels int) {
	layout.CAVChannelLayout.SetNbChannels(nbChannels)
}

func (layout *AVChannelLayout) GetMask() uint64 {
	return layout.CAVChannelLayout.GetMask()
}
func (layout *AVChannelLayout) SetMask(mask uint64) {
	layout.CAVChannelLayout.SetMask(mask)
}

func (layout *AVChannelLayout) GetMap() *avutil.CAVChannelCustom {
	return layout.CAVChannelLayout.GetMap()
}
func (layout *AVChannelLayout) SetMap(_map *avutil.CAVChannelCustom) {
	layout.CAVChannelLayout.SetMap(_map)
}

func (layout *AVChannelLayout) GetOpaque() unsafe.Pointer {
	return layout.CAVChannelLayout.GetOpaque()
}
func (layout *AVChannelLayout) SetOpaque(opaque unsafe.Pointer) {
	layout.CAVChannelLayout.SetOpaque(unsafe.Pointer(&opaque))
}

//#endregion member

func (layout *AVChannelLayout) Uninit() {
	avutil.AvChannelLayoutUninit(layout.CAVChannelLayout)
}

func (layout *AVChannelLayout) CopyTo(dst *AVChannelLayout) int {
	return avutil.AvChannelLayoutCopy(dst.CAVChannelLayout, layout.CAVChannelLayout)
}

func (layout *AVChannelLayout) Check() bool {
	return avutil.AvChannelLayoutCheck(layout.CAVChannelLayout) == 1
}

func (layout *AVChannelLayout) EqualsErr(b *AVChannelLayout) (bool, error) {
	if layout.CAVChannelLayout == b.CAVChannelLayout {
		return true, nil
	}
	ret := avutil.AvChannelLayoutCompare(layout.CAVChannelLayout, b.CAVChannelLayout)
	if ret != 0 && ret != 1 {
		return false, AvErr(ret)
	}
	return ret == 0, nil
}

func (layout *AVChannelLayout) Equals(b *AVChannelLayout) bool {
	ret, err := layout.EqualsErr(b)
	if err != nil {
		return false
	}
	return ret
}

func NewAvChannelLayoutFromMask(mask uint64) (*AVChannelLayout, int) {
	var c *avutil.CAVChannelLayout
	ret := avutil.AvChannelLayoutFromMask(c, mask)
	if ret != 0 {
		return nil, ret
	}
	return &AVChannelLayout{
		CAVChannelLayout: c,
	}, ret
}

func NewAvChannelLayoutFromString(str string) (*AVChannelLayout, int) {
	var c *avutil.CAVChannelLayout
	ret := avutil.AvChannelLayoutFromString(c, str)
	if ret != 0 {
		return nil, ret
	}
	return &AVChannelLayout{
		CAVChannelLayout: c,
	}, ret
}

func NewAvChannelLayoutDefault(nbChannels int) *AVChannelLayout {
	var c avutil.CAVChannelLayout
	avutil.AvChannelLayoutDefault(&c, nbChannels)
	return &AVChannelLayout{
		CAVChannelLayout: &c,
	}
}

func NewAvChannelLayout(opaque *unsafe.Pointer) *AVChannelLayout {
	c := avutil.AvChannelLayoutStandard(opaque)
	if c == nil {
		return nil
	}
	return &AVChannelLayout{
		CAVChannelLayout: c,
	}
}
