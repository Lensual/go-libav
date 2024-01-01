package avcodec

/*
#cgo pkg-config: libavcodec

#include "libavcodec/defs.h"
*/
import "C"
import (
	"unsafe"

	"github.com/Lensual/go-libav/ctypes"
)

/*
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

//  #ifndef AVCODEC_DEFS_H
//  #define AVCODEC_DEFS_H

/**
 * @file
 * @ingroup libavc
 * Misc types and constants that do not belong anywhere else.
 */

//  #include <stdint.h>
//  #include <stdlib.h>

/**
 * @ingroup lavc_decoding
 * Required number of additionally allocated bytes at the end of the input bitstream for decoding.
 * This is mainly needed because some optimized bitstream readers read
 * 32 or 64 bit at once and could read over the end.<br>
 * Note: If the first 23 bits of the additional bytes are not 0, then damaged
 * MPEG bitstreams could cause overread and segfault.
 */
const AV_INPUT_BUFFER_PADDING_SIZE = C.AV_INPUT_BUFFER_PADDING_SIZE

/**
 * Verify checksums embedded in the bitstream (could be of either encoded or
 * decoded data, depending on the format) and print an error message on mismatch.
 * If AV_EF_EXPLODE is also set, a mismatching checksum will result in the
 * decoder/demuxer returning an error.
 */
const (
	AV_EF_CRCCHECK  = C.AV_EF_CRCCHECK
	AV_EF_BITSTREAM = C.AV_EF_BITSTREAM ///< detect bitstream specification deviations
	AV_EF_BUFFER    = C.AV_EF_BUFFER    ///< detect improper bitstream length
	AV_EF_EXPLODE   = C.AV_EF_EXPLODE   ///< abort decoding on minor error detection

	AV_EF_IGNORE_ERR = C.AV_EF_IGNORE_ERR ///< ignore errors and continue
	AV_EF_CAREFUL    = C.AV_EF_CAREFUL    ///< consider things that violate the spec, are fast to calculate and have not been seen in the wild as errors
	AV_EF_COMPLIANT  = C.AV_EF_COMPLIANT  ///< consider all spec non compliances as errors
	AV_EF_AGGRESSIVE = C.AV_EF_AGGRESSIVE ///< consider things that a sane encoder/muxer should not do as an error

	FF_COMPLIANCE_VERY_STRICT  = C.FF_COMPLIANCE_VERY_STRICT ///< Strictly conform to an older more strict version of the spec or reference software.
	FF_COMPLIANCE_STRICT       = C.FF_COMPLIANCE_STRICT      ///< Strictly conform to all the things in the spec no matter what consequences.
	FF_COMPLIANCE_NORMAL       = C.FF_COMPLIANCE_NORMAL
	FF_COMPLIANCE_UNOFFICIAL   = C.FF_COMPLIANCE_UNOFFICIAL   ///< Allow unofficial extensions
	FF_COMPLIANCE_EXPERIMENTAL = C.FF_COMPLIANCE_EXPERIMENTAL ///< Allow nonstandardized experimental things.

	AV_PROFILE_UNKNOWN  = C.AV_PROFILE_UNKNOWN
	AV_PROFILE_RESERVED = C.AV_PROFILE_RESERVED

	AV_PROFILE_AAC_MAIN      = C.AV_PROFILE_AAC_MAIN
	AV_PROFILE_AAC_LOW       = C.AV_PROFILE_AAC_LOW
	AV_PROFILE_AAC_SSR       = C.AV_PROFILE_AAC_SSR
	AV_PROFILE_AAC_LTP       = C.AV_PROFILE_AAC_LTP
	AV_PROFILE_AAC_HE        = C.AV_PROFILE_AAC_HE
	AV_PROFILE_AAC_HE_V2     = C.AV_PROFILE_AAC_HE_V2
	AV_PROFILE_AAC_LD        = C.AV_PROFILE_AAC_LD
	AV_PROFILE_AAC_ELD       = C.AV_PROFILE_AAC_ELD
	AV_PROFILE_MPEG2_AAC_LOW = C.AV_PROFILE_MPEG2_AAC_LOW
	AV_PROFILE_MPEG2_AAC_HE  = C.AV_PROFILE_MPEG2_AAC_HE

	AV_PROFILE_DNXHD     = C.AV_PROFILE_DNXHD
	AV_PROFILE_DNXHR_LB  = C.AV_PROFILE_DNXHR_LB
	AV_PROFILE_DNXHR_SQ  = C.AV_PROFILE_DNXHR_SQ
	AV_PROFILE_DNXHR_HQ  = C.AV_PROFILE_DNXHR_HQ
	AV_PROFILE_DNXHR_HQX = C.AV_PROFILE_DNXHR_HQX
	AV_PROFILE_DNXHR_444 = C.AV_PROFILE_DNXHR_444

	AV_PROFILE_DTS              = C.AV_PROFILE_DTS
	AV_PROFILE_DTS_ES           = C.AV_PROFILE_DTS_ES
	AV_PROFILE_DTS_96_24        = C.AV_PROFILE_DTS_96_24
	AV_PROFILE_DTS_HD_HRA       = C.AV_PROFILE_DTS_HD_HRA
	AV_PROFILE_DTS_HD_MA        = C.AV_PROFILE_DTS_HD_MA
	AV_PROFILE_DTS_EXPRESS      = C.AV_PROFILE_DTS_EXPRESS
	AV_PROFILE_DTS_HD_MA_X      = C.AV_PROFILE_DTS_HD_MA_X
	AV_PROFILE_DTS_HD_MA_X_IMAX = C.AV_PROFILE_DTS_HD_MA_X_IMAX

	AV_PROFILE_EAC3_DDP_ATMOS = C.AV_PROFILE_EAC3_DDP_ATMOS

	AV_PROFILE_TRUEHD_ATMOS = C.AV_PROFILE_TRUEHD_ATMOS

	AV_PROFILE_MPEG2_422          = C.AV_PROFILE_MPEG2_422
	AV_PROFILE_MPEG2_HIGH         = C.AV_PROFILE_MPEG2_HIGH
	AV_PROFILE_MPEG2_SS           = C.AV_PROFILE_MPEG2_SS
	AV_PROFILE_MPEG2_SNR_SCALABLE = C.AV_PROFILE_MPEG2_SNR_SCALABLE
	AV_PROFILE_MPEG2_MAIN         = C.AV_PROFILE_MPEG2_MAIN
	AV_PROFILE_MPEG2_SIMPLE       = C.AV_PROFILE_MPEG2_SIMPLE

	AV_PROFILE_H264_CONSTRAINED = C.AV_PROFILE_H264_CONSTRAINED // 8+1; constraint_set1_flag
	AV_PROFILE_H264_INTRA       = C.AV_PROFILE_H264_INTRA       // 8+3; constraint_set3_flag

	AV_PROFILE_H264_BASELINE             = C.AV_PROFILE_H264_BASELINE
	AV_PROFILE_H264_CONSTRAINED_BASELINE = C.AV_PROFILE_H264_CONSTRAINED_BASELINE
	AV_PROFILE_H264_MAIN                 = C.AV_PROFILE_H264_MAIN
	AV_PROFILE_H264_EXTENDED             = C.AV_PROFILE_H264_EXTENDED
	AV_PROFILE_H264_HIGH                 = C.AV_PROFILE_H264_HIGH
	AV_PROFILE_H264_HIGH_10              = C.AV_PROFILE_H264_HIGH_10
	AV_PROFILE_H264_HIGH_10_INTRA        = C.AV_PROFILE_H264_HIGH_10_INTRA
	AV_PROFILE_H264_MULTIVIEW_HIGH       = C.AV_PROFILE_H264_MULTIVIEW_HIGH
	AV_PROFILE_H264_HIGH_422             = C.AV_PROFILE_H264_HIGH_422
	AV_PROFILE_H264_HIGH_422_INTRA       = C.AV_PROFILE_H264_HIGH_422_INTRA
	AV_PROFILE_H264_STEREO_HIGH          = C.AV_PROFILE_H264_STEREO_HIGH
	AV_PROFILE_H264_HIGH_444             = C.AV_PROFILE_H264_HIGH_444
	AV_PROFILE_H264_HIGH_444_PREDICTIVE  = C.AV_PROFILE_H264_HIGH_444_PREDICTIVE
	AV_PROFILE_H264_HIGH_444_INTRA       = C.AV_PROFILE_H264_HIGH_444_INTRA
	AV_PROFILE_H264_CAVLC_444            = C.AV_PROFILE_H264_CAVLC_444

	AV_PROFILE_VC1_SIMPLE   = C.AV_PROFILE_VC1_SIMPLE
	AV_PROFILE_VC1_MAIN     = C.AV_PROFILE_VC1_MAIN
	AV_PROFILE_VC1_COMPLEX  = C.AV_PROFILE_VC1_COMPLEX
	AV_PROFILE_VC1_ADVANCED = C.AV_PROFILE_VC1_ADVANCED

	AV_PROFILE_MPEG4_SIMPLE                    = C.AV_PROFILE_MPEG4_SIMPLE
	AV_PROFILE_MPEG4_SIMPLE_SCALABLE           = C.AV_PROFILE_MPEG4_SIMPLE_SCALABLE
	AV_PROFILE_MPEG4_CORE                      = C.AV_PROFILE_MPEG4_CORE
	AV_PROFILE_MPEG4_MAIN                      = C.AV_PROFILE_MPEG4_MAIN
	AV_PROFILE_MPEG4_N_BIT                     = C.AV_PROFILE_MPEG4_N_BIT
	AV_PROFILE_MPEG4_SCALABLE_TEXTURE          = C.AV_PROFILE_MPEG4_SCALABLE_TEXTURE
	AV_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION     = C.AV_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION
	AV_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE    = C.AV_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE
	AV_PROFILE_MPEG4_HYBRID                    = C.AV_PROFILE_MPEG4_HYBRID
	AV_PROFILE_MPEG4_ADVANCED_REAL_TIME        = C.AV_PROFILE_MPEG4_ADVANCED_REAL_TIME
	AV_PROFILE_MPEG4_CORE_SCALABLE             = C.AV_PROFILE_MPEG4_CORE_SCALABLE
	AV_PROFILE_MPEG4_ADVANCED_CODING           = C.AV_PROFILE_MPEG4_ADVANCED_CODING
	AV_PROFILE_MPEG4_ADVANCED_CORE             = C.AV_PROFILE_MPEG4_ADVANCED_CORE
	AV_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE = C.AV_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE
	AV_PROFILE_MPEG4_SIMPLE_STUDIO             = C.AV_PROFILE_MPEG4_SIMPLE_STUDIO
	AV_PROFILE_MPEG4_ADVANCED_SIMPLE           = C.AV_PROFILE_MPEG4_ADVANCED_SIMPLE

	AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0  = C.AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0
	AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1  = C.AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1
	AV_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION = C.AV_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION
	AV_PROFILE_JPEG2000_DCINEMA_2K             = C.AV_PROFILE_JPEG2000_DCINEMA_2K
	AV_PROFILE_JPEG2000_DCINEMA_4K             = C.AV_PROFILE_JPEG2000_DCINEMA_4K

	AV_PROFILE_VP9_0 = C.AV_PROFILE_VP9_0
	AV_PROFILE_VP9_1 = C.AV_PROFILE_VP9_1
	AV_PROFILE_VP9_2 = C.AV_PROFILE_VP9_2
	AV_PROFILE_VP9_3 = C.AV_PROFILE_VP9_3

	AV_PROFILE_HEVC_MAIN               = C.AV_PROFILE_HEVC_MAIN
	AV_PROFILE_HEVC_MAIN_10            = C.AV_PROFILE_HEVC_MAIN_10
	AV_PROFILE_HEVC_MAIN_STILL_PICTURE = C.AV_PROFILE_HEVC_MAIN_STILL_PICTURE
	AV_PROFILE_HEVC_REXT               = C.AV_PROFILE_HEVC_REXT
	AV_PROFILE_HEVC_SCC                = C.AV_PROFILE_HEVC_SCC

	AV_PROFILE_VVC_MAIN_10     = C.AV_PROFILE_VVC_MAIN_10
	AV_PROFILE_VVC_MAIN_10_444 = C.AV_PROFILE_VVC_MAIN_10_444

	AV_PROFILE_AV1_MAIN         = C.AV_PROFILE_AV1_MAIN
	AV_PROFILE_AV1_HIGH         = C.AV_PROFILE_AV1_HIGH
	AV_PROFILE_AV1_PROFESSIONAL = C.AV_PROFILE_AV1_PROFESSIONAL

	AV_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT            = C.AV_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT
	AV_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT = C.AV_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT
	AV_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT         = C.AV_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT
	AV_PROFILE_MJPEG_HUFFMAN_LOSSLESS                = C.AV_PROFILE_MJPEG_HUFFMAN_LOSSLESS
	AV_PROFILE_MJPEG_JPEG_LS                         = C.AV_PROFILE_MJPEG_JPEG_LS

	AV_PROFILE_SBC_MSBC = C.AV_PROFILE_SBC_MSBC

	AV_PROFILE_PRORES_PROXY    = C.AV_PROFILE_PRORES_PROXY
	AV_PROFILE_PRORES_LT       = C.AV_PROFILE_PRORES_LT
	AV_PROFILE_PRORES_STANDARD = C.AV_PROFILE_PRORES_STANDARD
	AV_PROFILE_PRORES_HQ       = C.AV_PROFILE_PRORES_HQ
	AV_PROFILE_PRORES_4444     = C.AV_PROFILE_PRORES_4444
	AV_PROFILE_PRORES_XQ       = C.AV_PROFILE_PRORES_XQ

	AV_PROFILE_ARIB_PROFILE_A = C.AV_PROFILE_ARIB_PROFILE_A
	AV_PROFILE_ARIB_PROFILE_C = C.AV_PROFILE_ARIB_PROFILE_C

	AV_PROFILE_KLVA_SYNC  = C.AV_PROFILE_KLVA_SYNC
	AV_PROFILE_KLVA_ASYNC = C.AV_PROFILE_KLVA_ASYNC

	AV_PROFILE_EVC_BASELINE = C.AV_PROFILE_EVC_BASELINE
	AV_PROFILE_EVC_MAIN     = C.AV_PROFILE_EVC_MAIN

	AV_LEVEL_UNKNOWN = C.AV_LEVEL_UNKNOWN
)

type CAVFieldOrder C.enum_AVFieldOrder

const (
	AV_FIELD_UNKNOWN     CAVFieldOrder = C.AV_FIELD_UNKNOWN
	AV_FIELD_PROGRESSIVE CAVFieldOrder = C.AV_FIELD_PROGRESSIVE
	AV_FIELD_TT          CAVFieldOrder = C.AV_FIELD_TT ///< Top coded_first, top displayed first
	AV_FIELD_BB          CAVFieldOrder = C.AV_FIELD_BB ///< Bottom coded first, bottom displayed first
	AV_FIELD_TB          CAVFieldOrder = C.AV_FIELD_TB ///< Top coded first, bottom displayed first
	AV_FIELD_BT          CAVFieldOrder = C.AV_FIELD_BT ///< Bottom coded first, top displayed first
)

/**
 * @ingroup lavc_decoding
 */

type CAVDiscard C.enum_AVDiscard

const (
	/* We leave some space between them for extensions (drop some
	 * keyframes for intra-only or drop just some bidir frames). */
	AVDISCARD_NONE     CAVDiscard = C.AVDISCARD_NONE     ///< discard nothing
	AVDISCARD_DEFAULT  CAVDiscard = C.AVDISCARD_DEFAULT  ///< discard useless packets like 0 size packets in avi
	AVDISCARD_NONREF   CAVDiscard = C.AVDISCARD_NONREF   ///< discard all non reference
	AVDISCARD_BIDIR    CAVDiscard = C.AVDISCARD_BIDIR    ///< discard all bidirectional frames
	AVDISCARD_NONINTRA CAVDiscard = C.AVDISCARD_NONINTRA ///< discard all non intra frames
	AVDISCARD_NONKEY   CAVDiscard = C.AVDISCARD_NONKEY   ///< discard all frames except keyframes
	AVDISCARD_ALL      CAVDiscard = C.AVDISCARD_ALL      ///< discard all
)

type CAVAudioServiceType C.enum_AVAudioServiceType

const (
	AV_AUDIO_SERVICE_TYPE_MAIN              CAVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_MAIN
	AV_AUDIO_SERVICE_TYPE_EFFECTS           CAVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_EFFECTS
	AV_AUDIO_SERVICE_TYPE_VISUALLY_IMPAIRED CAVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_VISUALLY_IMPAIRED
	AV_AUDIO_SERVICE_TYPE_HEARING_IMPAIRED  CAVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_HEARING_IMPAIRED
	AV_AUDIO_SERVICE_TYPE_DIALOGUE          CAVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_DIALOGUE
	AV_AUDIO_SERVICE_TYPE_COMMENTARY        CAVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_COMMENTARY
	AV_AUDIO_SERVICE_TYPE_EMERGENCY         CAVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_EMERGENCY
	AV_AUDIO_SERVICE_TYPE_VOICE_OVER        CAVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_VOICE_OVER
	AV_AUDIO_SERVICE_TYPE_KARAOKE           CAVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_KARAOKE
	AV_AUDIO_SERVICE_TYPE_NB                CAVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_NB ///< Not part of ABI
)

/**
 * Pan Scan area.
 * This specifies the area which should be displayed.
 * Note there may be multiple such areas for one frame.
 */
type CAVPanScan C.AVPanScan

//#region CAVPanScan

/**
 * id
 * - encoding: Set by user.
 * - decoding: Set by libavcodec.
 */
func (ps *CAVPanScan) GetId() int {
	return int(ps.id)
}

/**
 * id
 * - encoding: Set by user.
 * - decoding: Set by libavcodec.
 */
func (ps *CAVPanScan) SetId(id int) {
	ps.id = C.int(id)
}

/**
 * width and height in 1/16 pel
 * - encoding: Set by user.
 * - decoding: Set by libavcodec.
 */

func (ps *CAVPanScan) GetWidth() int {
	return int(ps.width)
}
func (ps *CAVPanScan) SetWidth(width int) {
	ps.width = C.int(width)
}

func (ps *CAVPanScan) GetHeight() int {
	return int(ps.height)
}
func (ps *CAVPanScan) SetHeight(height int) {
	ps.height = C.int(height)
}

/**
 * position of the top left corner in 1/16 pel for up to 3 fields/frames
 * - encoding: Set by user.
 * - decoding: Set by libavcodec.
 */
func (ps *CAVPanScan) GetPosition() [3][2]ctypes.Int16 {
	cArr := (*[2]ctypes.Int16)(unsafe.Pointer(unsafe.SliceData(ps.position[:])))
	return ([3][2]ctypes.Int16)(unsafe.Slice(cArr, 3))
}

/**
 * position of the top left corner in 1/16 pel for up to 3 fields/frames
 * - encoding: Set by user.
 * - decoding: Set by libavcodec.
 */
func (ps *CAVPanScan) SetPosition(position [3][2]ctypes.Int16) {
	cArr := (*[2]C.int16_t)(unsafe.Pointer(unsafe.SliceData(position[:])))
	ps.position = ([3][2]C.int16_t)(unsafe.Slice(cArr, 3))
}

//#endregion CAVPanScan

/**
 * This structure describes the bitrate properties of an encoded bitstream. It
 * roughly corresponds to a subset the VBV parameters for MPEG-2 or HRD
 * parameters for H.264/HEVC.
 */
type CAVCPBProperties = C.AVCPBProperties

//#region CAVCPBProperties

/**
 * Maximum bitrate of the stream, in bits per second.
 * Zero if unknown or unspecified.
 */
func (p *CAVCPBProperties) GetMaxBitrate() int64 {
	return int64(p.max_bitrate)
}

/**
 * Maximum bitrate of the stream, in bits per second.
 * Zero if unknown or unspecified.
 */
func (p *CAVCPBProperties) SetMaxBitrate(maxBitrate int64) {
	p.max_bitrate = C.int64_t(maxBitrate)
}

/**
 * Minimum bitrate of the stream, in bits per second.
 * Zero if unknown or unspecified.
 */
func (p *CAVCPBProperties) GetMinBitrate() int64 {
	return int64(p.min_bitrate)
}

/**
 * Minimum bitrate of the stream, in bits per second.
 * Zero if unknown or unspecified.
 */
func (p *CAVCPBProperties) SetMinBitrate(minBitrate int64) {
	p.min_bitrate = C.int64_t(minBitrate)
}

/**
 * Average bitrate of the stream, in bits per second.
 * Zero if unknown or unspecified.
 */
func (p *CAVCPBProperties) GetAvgBitrate() int64 {
	return int64(p.avg_bitrate)
}

/**
 * Average bitrate of the stream, in bits per second.
 * Zero if unknown or unspecified.
 */
func (p *CAVCPBProperties) SetAvgBitrate(avgBitrate int64) {
	p.avg_bitrate = C.int64_t(avgBitrate)
}

/**
 * The size of the buffer to which the ratecontrol is applied, in bits.
 * Zero if unknown or unspecified.
 */
func (p *CAVCPBProperties) GetBufferSize() int64 {
	return int64(p.buffer_size)
}

/**
 * The size of the buffer to which the ratecontrol is applied, in bits.
 * Zero if unknown or unspecified.
 */
func (p *CAVCPBProperties) SetBufferSize(bufferSize int64) {
	p.buffer_size = C.int64_t(bufferSize)
}

/**
 * The delay between the time the packet this structure is associated with
 * is received and the time when it should be decoded, in periods of a 27MHz
 * clock.
 *
 * UINT64_MAX when unknown or unspecified.
 */
func (p *CAVCPBProperties) GetVbvDelay() uint64 {
	return uint64(p.vbv_delay)
}

/**
 * The delay between the time the packet this structure is associated with
 * is received and the time when it should be decoded, in periods of a 27MHz
 * clock.
 *
 * UINT64_MAX when unknown or unspecified.
 */
func (p *CAVCPBProperties) SetVbvDelay(vbvDelay uint64) {
	p.vbv_delay = C.uint64_t(vbvDelay)
}

//#endregion CAVCPBProperties

/**
 * Allocate a CPB properties structure and initialize its fields to default
 * values.
 *
 * @param size if non-NULL, the size of the allocated struct will be written
 *             here. This is useful for embedding it in side data.
 *
 * @return the newly allocated struct or NULL on failure
 */
func AvCpbPropertiesAlloc(size *ctypes.SizeT) *CAVCPBProperties {
	return (*CAVCPBProperties)(C.av_cpb_properties_alloc((*C.size_t)(size)))
}

/**
 * This structure supplies correlation between a packet timestamp and a wall clock
 * production time. The definition follows the Producer Reference Time ('prft')
 * as defined in ISO/IEC 14496-12
 */
type CAVProducerReferenceTime C.AVProducerReferenceTime

//#region CAVProducerReferenceTime

/**
 * A UTC timestamp, in microseconds, since Unix epoch (e.g, av_gettime()).
 */
func (prt *CAVProducerReferenceTime) GetWallclock() int64 {
	return int64(prt.wallclock)
}

/**
 * A UTC timestamp, in microseconds, since Unix epoch (e.g, av_gettime()).
 */
func (prt *CAVProducerReferenceTime) SetWallclock(wallclock int64) {
	prt.wallclock = C.int64_t(wallclock)
}

func (prt *CAVProducerReferenceTime) GetFlags() int {
	return int(prt.flags)
}
func (prt *CAVProducerReferenceTime) SetFlags(flags int) {
	prt.flags = C.int(flags)
}

//#endregion CAVProducerReferenceTime

/**
 * Encode extradata length to a buffer. Used by xiph codecs.
 *
 * @param s buffer to write to; must be at least (v/255+1) bytes long
 * @param v size of extradata in bytes
 * @return number of bytes written to the buffer.
 */
func AvXiphlacing(s unsafe.Pointer, v uint) uint {
	return uint(C.av_xiphlacing((*C.uchar)(s), C.uint(v)))
}

//  #endif // AVCODEC_DEFS_H
