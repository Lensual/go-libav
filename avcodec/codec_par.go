package avcodec

/*
#cgo pkg-config: libavcodec

#include "libavcodec/codec_par.h"
*/
import "C"
import (
	"unsafe"

	"github.com/Lensual/go-libav/avutil"
)

/*
 * Codec parameters public API
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

//  #ifndef AVCODEC_CODEC_PAR_H
//  #define AVCODEC_CODEC_PAR_H

//  #include <stdint.h>

//  #include "libavutil/avutil.h"
//  #include "libavutil/channel_layout.h"
//  #include "libavutil/rational.h"
//  #include "libavutil/pixfmt.h"

//  #include "codec_id.h"
//  #include "defs.h"
//  #include "packet.h"

/**
 * @addtogroup lavc_core
 * @{
 */

/**
 * This struct describes the properties of an encoded stream.
 *
 * sizeof(AVCodecParameters) is not a part of the public ABI, this struct must
 * be allocated with avcodec_parameters_alloc() and freed with
 * avcodec_parameters_free().
 */
type CAVCodecParameters C.AVCodecParameters

//#region CAVCodecParameters

/**
 * General type of the encoded data.
 */
func (p *CAVCodecParameters) GetCodecType() avutil.CAVMediaType {
	return avutil.CAVMediaType(p.codec_type)
}

/**
 * General type of the encoded data.
 */
func (p *CAVCodecParameters) SetCodecType(codecType avutil.CAVMediaType) {
	p.codec_type = C.enum_AVMediaType(codecType)
}

/**
 * Specific type of the encoded data (the codec used).
 */
func (p *CAVCodecParameters) GetCodecId() CAVCodecID {
	return CAVCodecID(p.codec_id)
}

/**
 * Specific type of the encoded data (the codec used).
 */
func (p *CAVCodecParameters) SetCodecId(codecId CAVCodecID) {
	p.codec_id = C.enum_AVCodecID(codecId)
}

/**
 * Additional information about the codec (corresponds to the AVI FOURCC).
 */
func (p *CAVCodecParameters) GetCodecTag() uint32 {
	return uint32(p.codec_tag)
}

/**
 * Additional information about the codec (corresponds to the AVI FOURCC).
 */
func (p *CAVCodecParameters) SetCodecTag(codecTag uint32) {
	p.codec_tag = C.uint32_t(codecTag)
}

/**
 * Extra binary data needed for initializing the decoder, codec-dependent.
 *
 * Must be allocated with av_malloc() and will be freed by
 * avcodec_parameters_free(). The allocated size of extradata must be at
 * least extradata_size + AV_INPUT_BUFFER_PADDING_SIZE, with the padding
 * bytes zeroed.
 */
func (p *CAVCodecParameters) GetExtradata() unsafe.Pointer {
	return unsafe.Pointer(p.extradata)
}

/**
 * Extra binary data needed for initializing the decoder, codec-dependent.
 *
 * Must be allocated with av_malloc() and will be freed by
 * avcodec_parameters_free(). The allocated size of extradata must be at
 * least extradata_size + AV_INPUT_BUFFER_PADDING_SIZE, with the padding
 * bytes zeroed.
 */
func (p *CAVCodecParameters) SetExtradata(extradata unsafe.Pointer) {
	p.extradata = (*C.uint8_t)(extradata)
}

/**
 * Size of the extradata content in bytes.
 */
func (p *CAVCodecParameters) GetExtradataSize() int {
	return int(p.extradata_size)
}

/**
 * Size of the extradata content in bytes.
 */
func (p *CAVCodecParameters) SetExtradataSize(extradataSize int) {
	p.extradata_size = C.int(extradataSize)
}

/**
 * - video: the pixel format, the value corresponds to enum AVPixelFormat.
 * - audio: the sample format, the value corresponds to enum AVSampleFormat.
 */
func (p *CAVCodecParameters) GetFormat() int {
	return int(p.format)
}

/**
 * - video: the pixel format, the value corresponds to enum AVPixelFormat.
 * - audio: the sample format, the value corresponds to enum AVSampleFormat.
 */
func (p *CAVCodecParameters) SetFormat(format int) {
	p.format = C.int(format)
}

/**
 * The average bitrate of the encoded data (in bits per second).
 */
func (p *CAVCodecParameters) GetBitRate() int64 {
	return int64(p.bit_rate)
}

/**
 * The average bitrate of the encoded data (in bits per second).
 */
func (p *CAVCodecParameters) SetBitRate(bitRate int64) {
	p.bit_rate = C.int64_t(bitRate)
}

/**
 * The number of bits per sample in the codedwords.
 *
 * This is basically the bitrate per sample. It is mandatory for a bunch of
 * formats to actually decode them. It's the number of bits for one sample in
 * the actual coded bitstream.
 *
 * This could be for example 4 for ADPCM
 * For PCM formats this matches bits_per_raw_sample
 * Can be 0
 */
func (p *CAVCodecParameters) GetBitsPerCodedSample() int {
	return int(p.bits_per_coded_sample)
}

/**
 * The number of bits per sample in the codedwords.
 *
 * This is basically the bitrate per sample. It is mandatory for a bunch of
 * formats to actually decode them. It's the number of bits for one sample in
 * the actual coded bitstream.
 *
 * This could be for example 4 for ADPCM
 * For PCM formats this matches bits_per_raw_sample
 * Can be 0
 */
func (p *CAVCodecParameters) SetBitsPerCodedSample(bitsPerCodedSample int) {
	p.bits_per_coded_sample = C.int(bitsPerCodedSample)
}

/**
 * This is the number of valid bits in each output sample. If the
 * sample format has more bits, the least significant bits are additional
 * padding bits, which are always 0. Use right shifts to reduce the sample
 * to its actual size. For example, audio formats with 24 bit samples will
 * have bits_per_raw_sample set to 24, and format set to AV_SAMPLE_FMT_S32.
 * To get the original sample use "(int32_t)sample >> 8"."
 *
 * For ADPCM this might be 12 or 16 or similar
 * Can be 0
 */
func (p *CAVCodecParameters) GetBitsPerRawSample() int {
	return int(p.bits_per_raw_sample)
}

/**
 * This is the number of valid bits in each output sample. If the
 * sample format has more bits, the least significant bits are additional
 * padding bits, which are always 0. Use right shifts to reduce the sample
 * to its actual size. For example, audio formats with 24 bit samples will
 * have bits_per_raw_sample set to 24, and format set to AV_SAMPLE_FMT_S32.
 * To get the original sample use "(int32_t)sample >> 8"."
 *
 * For ADPCM this might be 12 or 16 or similar
 * Can be 0
 */
func (p *CAVCodecParameters) SetBitsPerRawSample(bitsPerRawSample int) {
	p.bits_per_raw_sample = C.int(bitsPerRawSample)
}

/**
 * Codec-specific bitstream restrictions that the stream conforms to.
 */

func (p *CAVCodecParameters) GetProfile() int {
	return int(p.profile)
}

func (p *CAVCodecParameters) SetProfile(profile int) {
	p.profile = C.int(profile)
}

func (p *CAVCodecParameters) GetLevel() int {
	return int(p.level)
}
func (p *CAVCodecParameters) SetLevel(level int) {
	p.level = C.int(level)
}

/**
 * Video only. The dimensions of the video frame in pixels.
 */

func (p *CAVCodecParameters) GetWidth() int {
	return int(p.width)
}
func (p *CAVCodecParameters) SetWidth(width int) {
	p.width = C.int(width)
}

func (p *CAVCodecParameters) GetHeight() int {
	return int(p.height)
}
func (p *CAVCodecParameters) SetHeight(height int) {
	p.height = C.int(height)
}

/**
 * Video only. The aspect ratio (width / height) which a single pixel
 * should have when displayed.
 *
 * When the aspect ratio is unknown / undefined, the numerator should be
 * set to 0 (the denominator may have any value).
 */
func (p *CAVCodecParameters) GetSampleAspectRatio() avutil.CAVRational {
	return *(*avutil.CAVRational)(unsafe.Pointer(&p.sample_aspect_ratio))
}

/**
 * Video only. The aspect ratio (width / height) which a single pixel
 * should have when displayed.
 *
 * When the aspect ratio is unknown / undefined, the numerator should be
 * set to 0 (the denominator may have any value).
 */
func (p *CAVCodecParameters) SetSampleAspectRatio(sampleAspectRatio avutil.CAVRational) {
	p.sample_aspect_ratio = *(*C.AVRational)(unsafe.Pointer(&sampleAspectRatio))
}

/**
 * Video only. The order of the fields in interlaced video.
 */
func (p *CAVCodecParameters) GetFieldOrder() CAVFieldOrder {
	return CAVFieldOrder(p.field_order)
}

/**
 * Video only. The order of the fields in interlaced video.
 */
func (p *CAVCodecParameters) SetFieldOrder(fieldOrder CAVFieldOrder) {
	p.field_order = C.enum_AVFieldOrder(fieldOrder)
}

/**
 * Video only. Additional colorspace characteristics.
 */

func (p *CAVCodecParameters) GetColorRange() avutil.CAVColorRange {
	return avutil.CAVColorRange(p.color_range)
}
func (p *CAVCodecParameters) SetColorRange(colorRange avutil.CAVColorRange) {
	p.color_range = C.enum_AVColorRange(colorRange)
}

func (p *CAVCodecParameters) GetColorPrimaries() avutil.CAVColorPrimaries {
	return avutil.CAVColorPrimaries(p.color_primaries)
}
func (p *CAVCodecParameters) SetColorPrimaries(colorPrimaries avutil.CAVColorPrimaries) {
	p.color_primaries = C.enum_AVColorPrimaries(colorPrimaries)
}

func (p *CAVCodecParameters) GetColorTrc() avutil.CAVColorTransferCharacteristic {
	return avutil.CAVColorTransferCharacteristic(p.color_trc)
}
func (p *CAVCodecParameters) SetColorTrc(colorTrc avutil.CAVColorTransferCharacteristic) {
	p.color_trc = C.enum_AVColorTransferCharacteristic(colorTrc)
}

func (p *CAVCodecParameters) GetColorspace() avutil.CAVColorSpace {
	return avutil.CAVColorSpace(p.color_space)
}
func (p *CAVCodecParameters) SetColorspace(colorspace avutil.CAVColorSpace) {
	p.color_space = C.enum_AVColorSpace(colorspace)
}

func (p *CAVCodecParameters) GetChromaLocation() avutil.CAVChromaLocation {
	return avutil.CAVChromaLocation(p.chroma_location)
}
func (p *CAVCodecParameters) SetChromaLocation(chromaLocation avutil.CAVChromaLocation) {
	p.chroma_location = C.enum_AVChromaLocation(chromaLocation)
}

/**
 * Video only. Number of delayed frames.
 */
func (p *CAVCodecParameters) GetVideoDelay() int {
	return int(p.video_delay)
}

/**
 * Video only. Number of delayed frames.
 */
func (p *CAVCodecParameters) SetVideoDelay(videoDelay int) {
	p.video_delay = C.int(videoDelay)
}

//  #if FF_API_OLD_CHANNEL_LAYOUT
// 	 /**
// 	  * Audio only. The channel layout bitmask. May be 0 if the channel layout is
// 	  * unknown or unspecified, otherwise the number of bits set must be equal to
// 	  * the channels field.
// 	  * @deprecated use ch_layout
// 	  */
// 	 attribute_deprecated
// 	 uint64_t channel_layout;
// 	 /**
// 	  * Audio only. The number of audio channels.
// 	  * @deprecated use ch_layout.nb_channels
// 	  */
// 	 attribute_deprecated
// 	 int      channels;
//  #endif

/**
 * Audio only. The number of audio samples per second.
 */
func (p *CAVCodecParameters) GetSampleRate() int {
	return int(p.sample_rate)
}

/**
 * Audio only. The number of audio samples per second.
 */
func (p *CAVCodecParameters) SetSampleRate(sampleRate int) {
	p.sample_rate = C.int(sampleRate)
}

/**
 * Audio only. The number of bytes per coded audio frame, required by some
 * formats.
 *
 * Corresponds to nBlockAlign in WAVEFORMATEX.
 */
func (p *CAVCodecParameters) GetBlockAlign() int {
	return int(p.block_align)
}

/**
 * Audio only. The number of bytes per coded audio frame, required by some
 * formats.
 *
 * Corresponds to nBlockAlign in WAVEFORMATEX.
 */
func (p *CAVCodecParameters) SetBlockAlign(blockAlign int) {
	p.block_align = C.int(blockAlign)
}

/**
 * Audio only. Audio frame size, if known. Required by some formats to be static.
 */
func (p *CAVCodecParameters) GetFrameSize() int {
	return int(p.frame_size)
}

/**
 * Audio only. Audio frame size, if known. Required by some formats to be static.
 */
func (p *CAVCodecParameters) SetFrameSize(frameSize int) {
	p.frame_size = C.int(frameSize)
}

/**
 * Audio only. The amount of padding (in samples) inserted by the encoder at
 * the beginning of the audio. I.e. this number of leading decoded samples
 * must be discarded by the caller to get the original audio without leading
 * padding.
 */
func (p *CAVCodecParameters) GetInitialPadding() int {
	return int(p.initial_padding)
}

/**
 * Audio only. The amount of padding (in samples) inserted by the encoder at
 * the beginning of the audio. I.e. this number of leading decoded samples
 * must be discarded by the caller to get the original audio without leading
 * padding.
 */
func (p *CAVCodecParameters) SetInitialPadding(initialPadding int) {
	p.initial_padding = C.int(initialPadding)
}

/**
 * Audio only. The amount of padding (in samples) appended by the encoder to
 * the end of the audio. I.e. this number of decoded samples must be
 * discarded by the caller from the end of the stream to get the original
 * audio without any trailing padding.
 */
func (p *CAVCodecParameters) GetTrailingPadding() int {
	return int(p.trailing_padding)
}

/**
 * Audio only. The amount of padding (in samples) appended by the encoder to
 * the end of the audio. I.e. this number of decoded samples must be
 * discarded by the caller from the end of the stream to get the original
 * audio without any trailing padding.
 */
func (p *CAVCodecParameters) SetTrailingPadding(trailingPadding int) {
	p.trailing_padding = C.int(trailingPadding)
}

/**
 * Audio only. Number of samples to skip after a discontinuity.
 */
func (p *CAVCodecParameters) GetSeekPreroll() int {
	return int(p.seek_preroll)
}

/**
 * Audio only. Number of samples to skip after a discontinuity.
 */
func (p *CAVCodecParameters) SetSeekPreroll(seekPreroll int) {
	p.seek_preroll = C.int(seekPreroll)
}

/**
 * Audio only. The channel layout and number of channels.
 */
func (p *CAVCodecParameters) GetChLayout() avutil.CAVChannelLayout {
	return *(*avutil.CAVChannelLayout)(unsafe.Pointer(&p.ch_layout))
}

/**
 * Audio only. The channel layout and number of channels.
 */
func (p *CAVCodecParameters) SetChLayout(chLayout avutil.CAVChannelLayout) {
	p.ch_layout = *(*C.AVChannelLayout)(unsafe.Pointer(&chLayout))
}

/**
 * Video only. Number of frames per second, for streams with constant frame
 * durations. Should be set to { 0, 1 } when some frames have differing
 * durations or if the value is not known.
 *
 * @note This field correponds to values that are stored in codec-level
 * headers and is typically overridden by container/transport-layer
 * timestamps, when available. It should thus be used only as a last resort,
 * when no higher-level timing information is available.
 */
func (p *CAVCodecParameters) GetFramerate() avutil.CAVRational {
	return *(*avutil.CAVRational)(unsafe.Pointer(&p.framerate))
}

/**
 * Video only. Number of frames per second, for streams with constant frame
 * durations. Should be set to { 0, 1 } when some frames have differing
 * durations or if the value is not known.
 *
 * @note This field correponds to values that are stored in codec-level
 * headers and is typically overridden by container/transport-layer
 * timestamps, when available. It should thus be used only as a last resort,
 * when no higher-level timing information is available.
 */
func (p *CAVCodecParameters) SetFramerate(framerate avutil.CAVRational) {
	p.framerate = *(*C.AVRational)(unsafe.Pointer(&framerate))
}

/**
 * Additional data associated with the entire stream.
 */
func (p *CAVCodecParameters) GetCodedSideData() *CAVPacketSideData {
	return (*CAVPacketSideData)(p.coded_side_data)
}

/**
 * Additional data associated with the entire stream.
 */
func (p *CAVCodecParameters) SetCodedSideData(codedSideData *CAVPacketSideData) {
	p.coded_side_data = (*C.AVPacketSideData)(codedSideData)
}

/**
 * Amount of entries in @ref coded_side_data.
 */
func (p *CAVCodecParameters) GetNbCodedSideData() int {
	return int(p.nb_coded_side_data)
}

/**
 * Amount of entries in @ref coded_side_data.
 */
func (p *CAVCodecParameters) SetNbCodedSideData(nbCodedSideData int) {
	p.nb_coded_side_data = C.int(nbCodedSideData)
}

//#endregion CAVCodecParameters

/**
 * Allocate a new AVCodecParameters and set its fields to default values
 * (unknown/invalid/0). The returned struct must be freed with
 * avcodec_parameters_free().
 */
func AvcodecParametersAlloc() *CAVCodecParameters {
	return (*CAVCodecParameters)(C.avcodec_parameters_alloc())
}

/**
 * Free an AVCodecParameters instance and everything associated with it and
 * write NULL to the supplied pointer.
 */
func AvcodecParametersFree(par **CAVCodecParameters) {
	C.avcodec_parameters_free((**C.AVCodecParameters)(unsafe.Pointer(par)))
}

/**
 * Copy the contents of src to dst. Any allocated fields in dst are freed and
 * replaced with newly allocated duplicates of the corresponding fields in src.
 *
 * @return >= 0 on success, a negative AVERROR code on failure.
 */
func AvcodecParametersCopy(dst *CAVCodecParameters, src *CAVCodecParameters) int {
	return int(C.avcodec_parameters_copy((*C.AVCodecParameters)(dst), (*C.AVCodecParameters)(src)))
}

/**
 * This function is the same as av_get_audio_frame_duration(), except it works
 * with AVCodecParameters instead of an AVCodecContext.
 */
func AvGetAudioFrameDuration2(par *CAVCodecParameters, frameBytes int) int {
	return int(C.av_get_audio_frame_duration2((*C.AVCodecParameters)(par), C.int(frameBytes)))
}

/**
 * @}
 */

//  #endif // AVCODEC_CODEC_PAR_H
