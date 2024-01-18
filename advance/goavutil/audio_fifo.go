package goavutil

import (
	"unsafe"

	"github.com/Lensual/go-libav/avutil"
)

type AVAudioFifo struct {
	CAVAudioFifo *avutil.CAVAudioFifo
}

func (af *AVAudioFifo) Free() {
	avutil.AvAudioFifoFree(af.CAVAudioFifo)
}

func (af *AVAudioFifo) Realloc(nbSamples int) int {
	return avutil.AvAudioFifoRealloc(af.CAVAudioFifo, nbSamples)
}

func (af *AVAudioFifo) Write(data *unsafe.Pointer, nbSamples int) int {
	return avutil.AvAudioFifoWrite(af.CAVAudioFifo, data, nbSamples)
}

func (af *AVAudioFifo) Peek(data *unsafe.Pointer, nbSamples int) int {
	return avutil.AvAudioFifoPeek(af.CAVAudioFifo, data, nbSamples)
}

func (af *AVAudioFifo) PeekAt(data *unsafe.Pointer, nbSamples int, offset int) int {
	return avutil.AvAudioFifoPeekAt(af.CAVAudioFifo, data, nbSamples, offset)
}

func (af *AVAudioFifo) Read(data *unsafe.Pointer, nbSamples int) int {
	return avutil.AvAudioFifoRead(af.CAVAudioFifo, data, nbSamples)
}

func (af *AVAudioFifo) Drain(nbSamples int) int {
	return avutil.AvAudioFifoDrain(af.CAVAudioFifo, nbSamples)
}

func (af *AVAudioFifo) Reset() {
	avutil.AvAudioFifoReset(af.CAVAudioFifo)
}

func (af *AVAudioFifo) Size() int {
	return avutil.AvAudioFifoSize(af.CAVAudioFifo)
}

func (af *AVAudioFifo) Space() int {
	return avutil.AvAudioFifoSpace(af.CAVAudioFifo)
}

func NewAVAudioFifo(sampleFmt avutil.CAVSampleFormat, channels int, nbSamples int) *AVAudioFifo {
	cAf := avutil.AvAudioFifoAlloc(sampleFmt, channels, nbSamples)
	if cAf == nil {
		return nil
	}
	return &AVAudioFifo{
		CAVAudioFifo: cAf,
	}
}
