package avutil

/*
#cgo pkg-config: libavutil

#include "libavutil/audio_fifo.h"
*/
import "C"
import "unsafe"

/*
 * Audio FIFO
 * Copyright (c) 2012 Justin Ruggles <justin.ruggles@gmail.com>
 *
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

/**
 * @file
 * Audio FIFO Buffer
 */

//  #ifndef AVUTIL_AUDIO_FIFO_H
//  #define AVUTIL_AUDIO_FIFO_H

//  #include "attributes.h"
//  #include "samplefmt.h"

/**
 * @addtogroup lavu_audio
 * @{
 *
 * @defgroup lavu_audiofifo Audio FIFO Buffer
 * @{
 */

/**
 * Context for an Audio FIFO Buffer.
 *
 * - Operates at the sample level rather than the byte level.
 * - Supports multiple channels with either planar or packed sample format.
 * - Automatic reallocation when writing to a full buffer.
 */
type CAVAudioFifo C.AVAudioFifo

/**
 * Free an AVAudioFifo.
 *
 * @param af  AVAudioFifo to free
 */
func AvAudioFifoFree(af *CAVAudioFifo) {
	C.av_audio_fifo_free((*C.AVAudioFifo)(af))
}

/**
 * Allocate an AVAudioFifo.
 *
 * @param sample_fmt  sample format
 * @param channels    number of channels
 * @param nb_samples  initial allocation size, in samples
 * @return            newly allocated AVAudioFifo, or NULL on error
 */
func AvAudioFifoAlloc(sampleFmt CAVSampleFormat, channels int, nbSamples int) *CAVAudioFifo {
	return (*CAVAudioFifo)(C.av_audio_fifo_alloc(C.enum_AVSampleFormat(sampleFmt), C.int(channels), C.int(nbSamples)))
}

/**
 * Reallocate an AVAudioFifo.
 *
 * @param af          AVAudioFifo to reallocate
 * @param nb_samples  new allocation size, in samples
 * @return            0 if OK, or negative AVERROR code on failure
 */
func AvAudioFifoRealloc(af *CAVAudioFifo, nbSamples int) int {
	return int(C.av_audio_fifo_realloc((*C.AVAudioFifo)(af), C.int(nbSamples)))
}

/**
 * Write data to an AVAudioFifo.
 *
 * The AVAudioFifo will be reallocated automatically if the available space
 * is less than nb_samples.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to write to
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to write
 * @return            number of samples actually written, or negative AVERROR
 *                    code on failure. If successful, the number of samples
 *                    actually written will always be nb_samples.
 */
func AvAudioFifoWrite(af *CAVAudioFifo, data *unsafe.Pointer, nbSamples int) int {
	return int(C.av_audio_fifo_write((*C.AVAudioFifo)(af), data, C.int(nbSamples)))
}

/**
 * Peek data from an AVAudioFifo.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to read from
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to peek
 * @return            number of samples actually peek, or negative AVERROR code
 *                    on failure. The number of samples actually peek will not
 *                    be greater than nb_samples, and will only be less than
 *                    nb_samples if av_audio_fifo_size is less than nb_samples.
 */
func AvAudioFifoPeek(af *CAVAudioFifo, data *unsafe.Pointer, nbSamples int) int {
	return int(C.av_audio_fifo_peek((*C.AVAudioFifo)(af), data, C.int(nbSamples)))
}

/**
 * Peek data from an AVAudioFifo.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to read from
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to peek
 * @param offset      offset from current read position
 * @return            number of samples actually peek, or negative AVERROR code
 *                    on failure. The number of samples actually peek will not
 *                    be greater than nb_samples, and will only be less than
 *                    nb_samples if av_audio_fifo_size is less than nb_samples.
 */
func AvAudioFifoPeekAt(af *CAVAudioFifo, data *unsafe.Pointer,
	nbSamples int, offset int) int {
	return int(C.av_audio_fifo_peek_at((*C.AVAudioFifo)(af), data,
		C.int(nbSamples), C.int(offset)))
}

/**
 * Read data from an AVAudioFifo.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to read from
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to read
 * @return            number of samples actually read, or negative AVERROR code
 *                    on failure. The number of samples actually read will not
 *                    be greater than nb_samples, and will only be less than
 *                    nb_samples if av_audio_fifo_size is less than nb_samples.
 */
func AvAudioFifoRead(af *CAVAudioFifo, data *unsafe.Pointer, nbSamples int) int {
	return int(C.av_audio_fifo_read((*C.AVAudioFifo)(af), data, C.int(nbSamples)))
}

/**
 * Drain data from an AVAudioFifo.
 *
 * Removes the data without reading it.
 *
 * @param af          AVAudioFifo to drain
 * @param nb_samples  number of samples to drain
 * @return            0 if OK, or negative AVERROR code on failure
 */
func AvAudioFifoDrain(af *CAVAudioFifo, nbSamples int) int {
	return int(C.av_audio_fifo_drain((*C.AVAudioFifo)(af), C.int(nbSamples)))
}

/**
 * Reset the AVAudioFifo buffer.
 *
 * This empties all data in the buffer.
 *
 * @param af  AVAudioFifo to reset
 */
func AvAudioFifoReset(af *CAVAudioFifo) {
	C.av_audio_fifo_reset((*C.AVAudioFifo)(af))
}

/**
 * Get the current number of samples in the AVAudioFifo available for reading.
 *
 * @param af  the AVAudioFifo to query
 * @return    number of samples available for reading
 */
func AvAudioFifoSize(af *CAVAudioFifo) int {
	return int(C.av_audio_fifo_size((*C.AVAudioFifo)(af)))
}

/**
 * Get the current number of samples in the AVAudioFifo available for writing.
 *
 * @param af  the AVAudioFifo to query
 * @return    number of samples available for writing
 */
func AvAudioFifoSpace(af *CAVAudioFifo) int {
	return int(C.av_audio_fifo_space((*C.AVAudioFifo)(af)))
}

/**
 * @}
 * @}
 */

//  #endif /* AVUTIL_AUDIO_FIFO_H */
