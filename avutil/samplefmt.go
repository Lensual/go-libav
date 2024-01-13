package avutil

/*
#cgo pkg-config: libavutil

#include <stdlib.h>
#include "libavutil/samplefmt.h"
*/
import "C"
import (
	"unsafe"

	"github.com/Lensual/go-libav/ctypes"
)

/*
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

//  #ifndef AVUTIL_SAMPLEFMT_H
//  #define AVUTIL_SAMPLEFMT_H

//  #include <stdint.h>

/**
 * @addtogroup lavu_audio
 * @{
 *
 * @defgroup lavu_sampfmts Audio sample formats
 *
 * Audio sample format enumeration and related convenience functions.
 * @{
 */

/**
 * Audio sample formats
 *
 * - The data described by the sample format is always in native-endian order.
 *   Sample values can be expressed by native C types, hence the lack of a signed
 *   24-bit sample format even though it is a common raw audio data format.
 *
 * - The floating-point formats are based on full volume being in the range
 *   [-1.0, 1.0]. Any values outside this range are beyond full volume level.
 *
 * - The data layout as used in av_samples_fill_arrays() and elsewhere in FFmpeg
 *   (such as AVFrame in libavcodec) is as follows:
 *
 * @par
 * For planar sample formats, each audio channel is in a separate data plane,
 * and linesize is the buffer size, in bytes, for a single plane. All data
 * planes must be the same size. For packed sample formats, only the first data
 * plane is used, and samples for each channel are interleaved. In this case,
 * linesize is the buffer size, in bytes, for the 1 plane.
 *
 */
type CAVSampleFormat C.enum_AVSampleFormat

const (
	AV_SAMPLE_FMT_NONE CAVSampleFormat = C.AV_SAMPLE_FMT_NONE
	AV_SAMPLE_FMT_U8   CAVSampleFormat = C.AV_SAMPLE_FMT_U8  ///< unsigned 8 bits
	AV_SAMPLE_FMT_S16  CAVSampleFormat = C.AV_SAMPLE_FMT_S16 ///< signed 16 bits
	AV_SAMPLE_FMT_S32  CAVSampleFormat = C.AV_SAMPLE_FMT_S32 ///< signed 32 bits
	AV_SAMPLE_FMT_FLT  CAVSampleFormat = C.AV_SAMPLE_FMT_FLT ///< float
	AV_SAMPLE_FMT_DBL  CAVSampleFormat = C.AV_SAMPLE_FMT_DBL ///< double

	AV_SAMPLE_FMT_U8P  CAVSampleFormat = C.AV_SAMPLE_FMT_U8P  ///< unsigned 8 bits, planar
	AV_SAMPLE_FMT_S16P CAVSampleFormat = C.AV_SAMPLE_FMT_S16P ///< signed 16 bits, planar
	AV_SAMPLE_FMT_S32P CAVSampleFormat = C.AV_SAMPLE_FMT_S32P ///< signed 32 bits, planar
	AV_SAMPLE_FMT_FLTP CAVSampleFormat = C.AV_SAMPLE_FMT_FLTP ///< float, planar
	AV_SAMPLE_FMT_DBLP CAVSampleFormat = C.AV_SAMPLE_FMT_DBLP ///< double, planar
	AV_SAMPLE_FMT_S64  CAVSampleFormat = C.AV_SAMPLE_FMT_S64  ///< signed 64 bits
	AV_SAMPLE_FMT_S64P CAVSampleFormat = C.AV_SAMPLE_FMT_S64P ///< signed 64 bits, planar

	AV_SAMPLE_FMT_NB CAVSampleFormat = C.AV_SAMPLE_FMT_NB ///< Number of sample formats. DO NOT USE if linking dynamically
)

/**
 * Return the name of sample_fmt, or NULL if sample_fmt is not
 * recognized.
 */
func AvGetSampleFmtName(sampleFmt CAVSampleFormat) string {
	return C.GoString(C.av_get_sample_fmt_name(C.enum_AVSampleFormat(sampleFmt)))
}

/**
 * Return a sample format corresponding to name, or AV_SAMPLE_FMT_NONE
 * on error.
 */
func AvGetSampleFmt(name string) CAVSampleFormat {
	var cName *C.char = nil
	if len(name) > 0 {
		cName = C.CString(name)
		defer C.free(unsafe.Pointer(cName))
	}

	return CAVSampleFormat(C.av_get_sample_fmt(cName))
}

/**
 * Return the planar<->packed alternative form of the given sample format, or
 * AV_SAMPLE_FMT_NONE on error. If the passed sample_fmt is already in the
 * requested planar/packed format, the format returned is the same as the
 * input.
 */
func AvGetAltSampleFmt(sampleFmt CAVSampleFormat, planar int) CAVSampleFormat {
	return CAVSampleFormat(C.av_get_alt_sample_fmt(C.enum_AVSampleFormat(sampleFmt), C.int(planar)))
}

/**
 * Get the packed alternative form of the given sample format.
 *
 * If the passed sample_fmt is already in packed format, the format returned is
 * the same as the input.
 *
 * @return  the packed alternative form of the given sample format or
 *          AV_SAMPLE_FMT_NONE on error.
 */
func AvGetPackedSampleFmt(sampleFmt CAVSampleFormat) CAVSampleFormat {
	return CAVSampleFormat(C.av_get_packed_sample_fmt(C.enum_AVSampleFormat(sampleFmt)))
}

/**
 * Get the planar alternative form of the given sample format.
 *
 * If the passed sample_fmt is already in planar format, the format returned is
 * the same as the input.
 *
 * @return  the planar alternative form of the given sample format or
 *		 AV_SAMPLE_FMT_NONE on error.
 */
func AvGetPlanarSampleFmt(sampleFmt CAVSampleFormat) CAVSampleFormat {
	return CAVSampleFormat(C.av_get_planar_sample_fmt(C.enum_AVSampleFormat(sampleFmt)))
}

/**
 * Generate a string corresponding to the sample format with
 * sample_fmt, or a header if sample_fmt is negative.
 *
 * @param buf the buffer where to write the string
 * @param buf_size the size of buf
 * @param sample_fmt the number of the sample format to print the
 * corresponding info string, or a negative value to print the
 * corresponding header.
 * @return the pointer to the filled buffer or NULL if sample_fmt is
 * unknown or in case of other errors
 */
func AvGetSampleFmtString(buf unsafe.Pointer, bufSize int, sampleFmt CAVSampleFormat) string {
	return C.GoString(C.av_get_sample_fmt_string((*C.char)(buf), C.int(bufSize), C.enum_AVSampleFormat(sampleFmt)))
}

/**
 * Return number of bytes per sample.
 *
 * @param sample_fmt the sample format
 * @return number of bytes per sample or zero if unknown for the given
 * sample format
 */
func AvGetBytesPerSample(sampleFmt CAVSampleFormat) int {
	return int(C.av_get_bytes_per_sample(C.enum_AVSampleFormat(sampleFmt)))
}

/**
 * Check if the sample format is planar.
 *
 * @param sample_fmt the sample format to inspect
 * @return 1 if the sample format is planar, 0 if it is interleaved
 */
func AvSampleFmtIsPlanar(sampleFmt CAVSampleFormat) int {
	return int(C.av_sample_fmt_is_planar(C.enum_AVSampleFormat(sampleFmt)))
}

/**
 * Get the required buffer size for the given audio parameters.
 *
 * @param[out] linesize calculated linesize, may be NULL
 * @param nb_channels   the number of channels
 * @param nb_samples    the number of samples in a single channel
 * @param sample_fmt    the sample format
 * @param align         buffer size alignment (0 = default, 1 = no alignment)
 * @return              required buffer size, or negative error code on failure
 */
func AvSamplesGetBufferSize(linesize []ctypes.Int, nbChannels int, nbSamples int,
	sampleFmt CAVSampleFormat, align int) int {
	return int(C.av_samples_get_buffer_size(
		(*C.int)(unsafe.SliceData(linesize)),
		C.int(nbChannels),
		C.int(nbSamples),
		C.enum_AVSampleFormat(sampleFmt),
		C.int(align)))
}

/**
 * @}
 *
 * @defgroup lavu_sampmanip Samples manipulation
 *
 * Functions that manipulate audio samples
 * @{
 */

/**
 * Fill plane data pointers and linesize for samples with sample
 * format sample_fmt.
 *
 * The audio_data array is filled with the pointers to the samples data planes:
 * for planar, set the start point of each channel's data within the buffer,
 * for packed, set the start point of the entire buffer only.
 *
 * The value pointed to by linesize is set to the aligned size of each
 * channel's data buffer for planar layout, or to the aligned size of the
 * buffer for all channels for packed layout.
 *
 * The buffer in buf must be big enough to contain all the samples
 * (use av_samples_get_buffer_size() to compute its minimum size),
 * otherwise the audio_data pointers will point to invalid data.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param[out] audio_data  array to be filled with the pointer for each channel
 * @param[out] linesize    calculated linesize, may be NULL
 * @param buf              the pointer to a buffer containing the samples
 * @param nb_channels      the number of channels
 * @param nb_samples       the number of samples in a single channel
 * @param sample_fmt       the sample format
 * @param align            buffer size alignment (0 = default, 1 = no alignment)
 * @return                 minimum size in bytes required for the buffer on success,
 *                         or a negative error code on failure
 */
func AvSamplesFillArrays(audioData *unsafe.Pointer, linesize *ctypes.Int,
	buf unsafe.Pointer,
	nbChannels int, nbSamples int,
	sampleFmt CAVSampleFormat, align int) int {
	return int(C.av_samples_fill_arrays((**C.uint8_t)(unsafe.Pointer(audioData)), (*C.int)(linesize),
		(*C.uint8_t)(buf),
		C.int(nbChannels), C.int(nbSamples),
		C.enum_AVSampleFormat(sampleFmt), C.int(align)))
}

/**
 * Allocate a samples buffer for nb_samples samples, and fill data pointers and
 * linesize accordingly.
 * The allocated samples buffer can be freed by using av_freep(&audio_data[0])
 * Allocated data will be initialized to silence.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param[out] audio_data  array to be filled with the pointer for each channel
 * @param[out] linesize    aligned size for audio buffer(s), may be NULL
 * @param nb_channels      number of audio channels
 * @param nb_samples       number of samples per channel
 * @param sample_fmt       the sample format
 * @param align            buffer size alignment (0 = default, 1 = no alignment)
 * @return                 >=0 on success or a negative error code on failure
 * @todo return the size of the allocated buffer in case of success at the next bump
 * @see av_samples_fill_arrays()
 * @see av_samples_alloc_array_and_samples()
 */
func AvSamplesAlloc(audioData *unsafe.Pointer, linesize *ctypes.Int, nbChannels int,
	nbSamples int, sampleFmt CAVSampleFormat, align int) int {
	return int(C.av_samples_alloc((**C.uint8_t)(unsafe.Pointer(audioData)), (*C.int)(linesize),
		C.int(nbChannels), C.int(nbSamples), C.enum_AVSampleFormat(sampleFmt), C.int(align)))
}

/**
 * Allocate a data pointers array, samples buffer for nb_samples
 * samples, and fill data pointers and linesize accordingly.
 *
 * This is the same as av_samples_alloc(), but also allocates the data
 * pointers array.
 *
 * @see av_samples_alloc()
 */
func AvSamplesAllocArrayAndSamples(audioData **unsafe.Pointer, linesize *ctypes.Int, nbChannels int,
	nbSamples int, sampleFmt CAVSampleFormat, align int) int {
	return int(C.av_samples_alloc_array_and_samples((***C.uint8_t)(unsafe.Pointer(audioData)), (*C.int)(linesize),
		C.int(nbChannels), C.int(nbSamples), C.enum_AVSampleFormat(sampleFmt), C.int(align)))
}

/**
 * Copy samples from src to dst.
 *
 * @param dst destination array of pointers to data planes
 * @param src source array of pointers to data planes
 * @param dst_offset offset in samples at which the data will be written to dst
 * @param src_offset offset in samples at which the data will be read from src
 * @param nb_samples number of samples to be copied
 * @param nb_channels number of audio channels
 * @param sample_fmt audio sample format
 */
func AvSamplesCopy(dst *unsafe.Pointer, src *unsafe.Pointer, dstOffset int,
	srcOffset int, nbSamples int, nbChannels int,
	sampleFmt CAVSampleFormat) int {
	return int(C.av_samples_copy((**C.uint8_t)(unsafe.Pointer(dst)), (**C.uint8_t)(unsafe.Pointer(src)), C.int(dstOffset),
		C.int(srcOffset), C.int(nbSamples), C.int(nbChannels),
		C.enum_AVSampleFormat(sampleFmt)))
}

/**
 * Fill an audio buffer with silence.
 *
 * @param audio_data  array of pointers to data planes
 * @param offset      offset in samples at which to start filling
 * @param nb_samples  number of samples to fill
 * @param nb_channels number of audio channels
 * @param sample_fmt  audio sample format
 */
func AvSamplesSetSilence(audioData *unsafe.Pointer, offset int, nbSamples int,
	nbChannels int, sampleFmt CAVSampleFormat) int {
	return int(C.av_samples_set_silence((**C.uint8_t)(unsafe.Pointer(audioData)), C.int(offset), C.int(nbSamples),
		C.int(nbChannels), C.enum_AVSampleFormat(sampleFmt)))
}

/**
 * @}
 * @}
 */
//  #endif /* AVUTIL_SAMPLEFMT_H */
