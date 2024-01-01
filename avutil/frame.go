package avutil

/*
#cgo pkg-config: libavutil

#include "libavutil/frame.h"
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

/**
 * @file
 * @ingroup lavu_frame
 * reference-counted frame API
 */

//  #ifndef AVUTIL_FRAME_H
//  #define AVUTIL_FRAME_H

//  #include <stddef.h>
//  #include <stdint.h>

//  #include "avutil.h"
//  #include "buffer.h"
//  #include "channel_layout.h"
//  #include "dict.h"
//  #include "rational.h"
//  #include "samplefmt.h"
//  #include "pixfmt.h"
//  #include "version.h"

/**
 * @defgroup lavu_frame AVFrame
 * @ingroup lavu_data
 *
 * @{
 * AVFrame is an abstraction for reference-counted raw multimedia data.
 */

type CAVFrameSideDataType C.enum_AVFrameSideDataType

const (
	/**
	 * The data is the AVPanScan struct defined in libavcodec.
	 */
	AV_FRAME_DATA_PANSCAN CAVFrameSideDataType = C.AV_FRAME_DATA_PANSCAN
	/**
	 * ATSC A53 Part 4 Closed Captions.
	 * A53 CC bitstream is stored as uint8_t in AVFrameSideData.data.
	 * The number of bytes of CC data is AVFrameSideData.size.
	 */
	AV_FRAME_DATA_A53_CC CAVFrameSideDataType = C.AV_FRAME_DATA_A53_CC
	/**
	 * Stereoscopic 3d metadata.
	 * The data is the AVStereo3D struct defined in libavutil/stereo3d.h.
	 */
	AV_FRAME_DATA_STEREO3D CAVFrameSideDataType = C.AV_FRAME_DATA_STEREO3D
	/**
	 * The data is the AVMatrixEncoding enum defined in libavutil/channel_layout.h.
	 */
	AV_FRAME_DATA_MATRIXENCODING CAVFrameSideDataType = C.AV_FRAME_DATA_MATRIXENCODING
	/**
	 * Metadata relevant to a downmix procedure.
	 * The data is the AVDownmixInfo struct defined in libavutil/downmix_info.h.
	 */
	AV_FRAME_DATA_DOWNMIX_INFO CAVFrameSideDataType = C.AV_FRAME_DATA_DOWNMIX_INFO
	/**
	 * ReplayGain information in the form of the AVReplayGain struct.
	 */
	AV_FRAME_DATA_REPLAYGAIN CAVFrameSideDataType = C.AV_FRAME_DATA_REPLAYGAIN
	/**
	 * This side data contains a 3x3 transformation matrix describing an affine
	 * transformation that needs to be applied to the frame for correct
	 * presentation.
	 *
	 * See libavutil/display.h for a detailed description of the data.
	 */
	AV_FRAME_DATA_DISPLAYMATRIX CAVFrameSideDataType = C.AV_FRAME_DATA_DISPLAYMATRIX
	/**
	 * Active Format Description data consisting of a single byte as specified
	 * in ETSI TS 101 154 using AVActiveFormatDescription enum.
	 */
	AV_FRAME_DATA_AFD CAVFrameSideDataType = C.AV_FRAME_DATA_AFD
	/**
	 * Motion vectors exported by some codecs (on demand through the export_mvs
	 * flag set in the libavcodec AVCodecContext flags2 option).
	 * The data is the AVMotionVector struct defined in
	 * libavutil/motion_vector.h.
	 */
	AV_FRAME_DATA_MOTION_VECTORS CAVFrameSideDataType = C.AV_FRAME_DATA_MOTION_VECTORS
	/**
	 * Recommmends skipping the specified number of samples. This is exported
	 * only if the "skip_manual" AVOption is set in libavcodec.
	 * This has the same format as AV_PKT_DATA_SKIP_SAMPLES.
	 * @code
	 * u32le number of samples to skip from start of this packet
	 * u32le number of samples to skip from end of this packet
	 * u8    reason for start skip
	 * u8    reason for end   skip (0=padding silence, 1=convergence)
	 * @endcode
	 */
	AV_FRAME_DATA_SKIP_SAMPLES CAVFrameSideDataType = C.AV_FRAME_DATA_SKIP_SAMPLES
	/**
	 * This side data must be associated with an audio frame and corresponds to
	 * enum AVAudioServiceType defined in avcodec.h.
	 */
	AV_FRAME_DATA_AUDIO_SERVICE_TYPE CAVFrameSideDataType = C.AV_FRAME_DATA_AUDIO_SERVICE_TYPE
	/**
	 * Mastering display metadata associated with a video frame. The payload is
	 * an AVMasteringDisplayMetadata type and contains information about the
	 * mastering display color volume.
	 */
	AV_FRAME_DATA_MASTERING_DISPLAY_METADATA CAVFrameSideDataType = C.AV_FRAME_DATA_MASTERING_DISPLAY_METADATA
	/**
	 * The GOP timecode in 25 bit timecode format. Data format is 64-bit integer.
	 * This is set on the first frame of a GOP that has a temporal reference of 0.
	 */
	AV_FRAME_DATA_GOP_TIMECODE CAVFrameSideDataType = C.AV_FRAME_DATA_GOP_TIMECODE

	/**
	 * The data represents the AVSphericalMapping structure defined in
	 * libavutil/spherical.h.
	 */
	AV_FRAME_DATA_SPHERICAL CAVFrameSideDataType = C.AV_FRAME_DATA_SPHERICAL

	/**
	 * Content light level (based on CTA-861.3). This payload contains data in
	 * the form of the AVContentLightMetadata struct.
	 */
	AV_FRAME_DATA_CONTENT_LIGHT_LEVEL CAVFrameSideDataType = C.AV_FRAME_DATA_CONTENT_LIGHT_LEVEL

	/**
	 * The data contains an ICC profile as an opaque octet buffer following the
	 * format described by ISO 15076-1 with an optional name defined in the
	 * metadata key entry "name".
	 */
	AV_FRAME_DATA_ICC_PROFILE CAVFrameSideDataType = C.AV_FRAME_DATA_ICC_PROFILE

	/**
	 * Timecode which conforms to SMPTE ST 12-1. The data is an array of 4 uint32_t
	 * where the first uint32_t describes how many (1-3) of the other timecodes are used.
	 * The timecode format is described in the documentation of av_timecode_get_smpte_from_framenum()
	 * function in libavutil/timecode.h.
	 */
	AV_FRAME_DATA_S12M_TIMECODE CAVFrameSideDataType = C.AV_FRAME_DATA_S12M_TIMECODE

	/**
	 * HDR dynamic metadata associated with a video frame. The payload is
	 * an AVDynamicHDRPlus type and contains information for color
	 * volume transform - application 4 of SMPTE 2094-40:2016 standard.
	 */
	AV_FRAME_DATA_DYNAMIC_HDR_PLUS CAVFrameSideDataType = C.AV_FRAME_DATA_DYNAMIC_HDR_PLUS

	/**
	 * Regions Of Interest, the data is an array of AVRegionOfInterest type, the number of
	 * array element is implied by AVFrameSideData.size / AVRegionOfInterest.self_size.
	 */
	AV_FRAME_DATA_REGIONS_OF_INTEREST CAVFrameSideDataType = C.AV_FRAME_DATA_REGIONS_OF_INTEREST

	/**
	 * Encoding parameters for a video frame, as described by AVVideoEncParams.
	 */
	AV_FRAME_DATA_VIDEO_ENC_PARAMS CAVFrameSideDataType = C.AV_FRAME_DATA_VIDEO_ENC_PARAMS

	/**
	 * User data unregistered metadata associated with a video frame.
	 * This is the H.26[45] UDU SEI message, and shouldn't be used for any other purpose
	 * The data is stored as uint8_t in AVFrameSideData.data which is 16 bytes of
	 * uuid_iso_iec_11578 followed by AVFrameSideData.size - 16 bytes of user_data_payload_byte.
	 */
	AV_FRAME_DATA_SEI_UNREGISTERED CAVFrameSideDataType = C.AV_FRAME_DATA_SEI_UNREGISTERED

	/**
	 * Film grain parameters for a frame, described by AVFilmGrainParams.
	 * Must be present for every frame which should have film grain applied.
	 */
	AV_FRAME_DATA_FILM_GRAIN_PARAMS CAVFrameSideDataType = C.AV_FRAME_DATA_FILM_GRAIN_PARAMS

	/**
	 * Bounding boxes for object detection and classification,
	 * as described by AVDetectionBBoxHeader.
	 */
	AV_FRAME_DATA_DETECTION_BBOXES CAVFrameSideDataType = C.AV_FRAME_DATA_DETECTION_BBOXES

	/**
	 * Dolby Vision RPU raw data, suitable for passing to x265
	 * or other libraries. Array of uint8_t, with NAL emulation
	 * bytes intact.
	 */
	AV_FRAME_DATA_DOVI_RPU_BUFFER CAVFrameSideDataType = C.AV_FRAME_DATA_DOVI_RPU_BUFFER

	/**
	 * Parsed Dolby Vision metadata, suitable for passing to a software
	 * implementation. The payload is the AVDOVIMetadata struct defined in
	 * libavutil/dovi_meta.h.
	 */
	AV_FRAME_DATA_DOVI_METADATA CAVFrameSideDataType = C.AV_FRAME_DATA_DOVI_METADATA

	/**
	 * HDR Vivid dynamic metadata associated with a video frame. The payload is
	 * an AVDynamicHDRVivid type and contains information for color
	 * volume transform - CUVA 005.1-2021.
	 */
	AV_FRAME_DATA_DYNAMIC_HDR_VIVID CAVFrameSideDataType = C.AV_FRAME_DATA_DYNAMIC_HDR_VIVID

	/**
	 * Ambient viewing environment metadata, as defined by H.274.
	 */
	AV_FRAME_DATA_AMBIENT_VIEWING_ENVIRONMENT CAVFrameSideDataType = C.AV_FRAME_DATA_AMBIENT_VIEWING_ENVIRONMENT

	/**
	 * Provide encoder-specific hinting information about changed/unchanged
	 * portions of a frame.  It can be used to pass information about which
	 * macroblocks can be skipped because they didn't change from the
	 * corresponding ones in the previous frame. This could be useful for
	 * applications which know this information in advance to speed up
	 * encoding.
	 */
	AV_FRAME_DATA_VIDEO_HINT CAVFrameSideDataType = C.AV_FRAME_DATA_VIDEO_HINT
)

type CAVActiveFormatDescription C.enum_AVActiveFormatDescription

const (
	AV_AFD_SAME         CAVActiveFormatDescription = C.AV_AFD_SAME
	AV_AFD_4_3          CAVActiveFormatDescription = C.AV_AFD_4_3
	AV_AFD_16_9         CAVActiveFormatDescription = C.AV_AFD_16_9
	AV_AFD_14_9         CAVActiveFormatDescription = C.AV_AFD_14_9
	AV_AFD_4_3_SP_14_9  CAVActiveFormatDescription = C.AV_AFD_4_3_SP_14_9
	AV_AFD_16_9_SP_14_9 CAVActiveFormatDescription = C.AV_AFD_16_9_SP_14_9
	AV_AFD_SP_4_3       CAVActiveFormatDescription = C.AV_AFD_SP_4_3
)

/**
 * Structure to hold side data for an AVFrame.
 *
 * sizeof(AVFrameSideData) is not a part of the public ABI, so new fields may be added
 * to the end with a minor bump.
 */
type CAVFrameSideData C.AVFrameSideData

// #region CAVFrameSideData

func (sd *CAVFrameSideData) GetType() CAVFrameSideDataType {
	return CAVFrameSideDataType(sd._type)
}
func (sd *CAVFrameSideData) GetData() unsafe.Pointer {
	return unsafe.Pointer(sd.data)
}
func (sd *CAVFrameSideData) GetSize() ctypes.SizeT {
	return ctypes.SizeT(sd.size)
}
func (sd *CAVFrameSideData) GetMetadata() *CAVDictionary {
	return (*CAVDictionary)(sd.metadata)
}
func (sd *CAVFrameSideData) GetBuf() *CAVBufferRef {
	return (*CAVBufferRef)(sd.buf)
}

//#endregion CAVFrameSideData

/**
 * Structure describing a single Region Of Interest.
 *
 * When multiple regions are defined in a single side-data block, they
 * should be ordered from most to least important - some encoders are only
 * capable of supporting a limited number of distinct regions, so will have
 * to truncate the list.
 *
 * When overlapping regions are defined, the first region containing a given
 * area of the frame applies.
 */
type CAVRegionOfInterest C.AVRegionOfInterest

//#region CAVRegionOfInterest

/**
 * Must be set to the size of this data structure (that is,
 * sizeof(AVRegionOfInterest)).
 */
func (roi *CAVRegionOfInterest) GetSelfSize() uint32 {
	return uint32(roi.self_size)
}

/**
 * Distance in pixels from the top edge of the frame to the top and
 * bottom edges and from the left edge of the frame to the left and
 * right edges of the rectangle defining this region of interest.
 *
 * The constraints on a region are encoder dependent, so the region
 * actually affected may be slightly larger for alignment or other
 * reasons.
 */
func (roi *CAVRegionOfInterest) GetTop() int {
	return int(roi.top)
}
func (roi *CAVRegionOfInterest) GetBottom() int {
	return int(roi.bottom)
}
func (roi *CAVRegionOfInterest) GetLeft() int {
	return int(roi.left)
}
func (roi *CAVRegionOfInterest) GetRight() int {
	return int(roi.right)
}

/**
 * Quantisation offset.
 *
 * Must be in the range -1 to +1.  A value of zero indicates no quality
 * change.  A negative value asks for better quality (less quantisation),
 * while a positive value asks for worse quality (greater quantisation).
 *
 * The range is calibrated so that the extreme values indicate the
 * largest possible offset - if the rest of the frame is encoded with the
 * worst possible quality, an offset of -1 indicates that this region
 * should be encoded with the best possible quality anyway.  Intermediate
 * values are then interpolated in some codec-dependent way.
 *
 * For example, in 10-bit H.264 the quantisation parameter varies between
 * -12 and 51.  A typical qoffset value of -1/10 therefore indicates that
 * this region should be encoded with a QP around one-tenth of the full
 * range better than the rest of the frame.  So, if most of the frame
 * were to be encoded with a QP of around 30, this region would get a QP
 * of around 24 (an offset of approximately -1/10 * (51 - -12) = -6.3).
 * An extreme value of -1 would indicate that this region should be
 * encoded with the best possible quality regardless of the treatment of
 * the rest of the frame - that is, should be encoded at a QP of -12.
 */
func (roi *CAVRegionOfInterest) GetQoffset() CAVRational {
	return CAVRational(roi.qoffset)
}

//#endregion CAVRegionOfInterest

/**
 * This structure describes decoded (raw) audio or video data.
 *
 * AVFrame must be allocated using av_frame_alloc(). Note that this only
 * allocates the AVFrame itself, the buffers for the data must be managed
 * through other means (see below).
 * AVFrame must be freed with av_frame_free().
 *
 * AVFrame is typically allocated once and then reused multiple times to hold
 * different data (e.g. a single AVFrame to hold frames received from a
 * decoder). In such a case, av_frame_unref() will free any references held by
 * the frame and reset it to its original clean state before it
 * is reused again.
 *
 * The data described by an AVFrame is usually reference counted through the
 * AVBuffer API. The underlying buffer references are stored in AVFrame.buf /
 * AVFrame.extended_buf. An AVFrame is considered to be reference counted if at
 * least one reference is set, i.e. if AVFrame.buf[0] != NULL. In such a case,
 * every single data plane must be contained in one of the buffers in
 * AVFrame.buf or AVFrame.extended_buf.
 * There may be a single buffer for all the data, or one separate buffer for
 * each plane, or anything in between.
 *
 * sizeof(AVFrame) is not a part of the public ABI, so new fields may be added
 * to the end with a minor bump.
 *
 * Fields can be accessed through AVOptions, the name string used, matches the
 * C structure field name for fields accessible through AVOptions. The AVClass
 * for AVFrame can be obtained from avcodec_get_frame_class()
 */
type CAVFrame C.AVFrame

// #region CAVFrame
const AV_NUM_DATA_POINTERS = C.AV_NUM_DATA_POINTERS

/**
 * pointer to the picture/channel planes.
 * This might be different from the first allocated byte. For video,
 * it could even point to the end of the image data.
 *
 * All pointers in data and extended_data must point into one of the
 * AVBufferRef in buf or extended_buf.
 *
 * Some decoders access areas outside 0,0 - width,height, please
 * see avcodec_align_dimensions2(). Some filters and swscale can read
 * up to 16 bytes beyond the planes, if these filters are to be used,
 * then 16 extra bytes must be allocated.
 *
 * NOTE: Pointers not needed by the format MUST be set to NULL.
 *
 * @attention In case of video, the data[] pointers can point to the
 * end of image data in order to reverse line order, when used in
 * combination with negative values in the linesize[] array.
 */
func (f *CAVFrame) GetData() [AV_NUM_DATA_POINTERS]unsafe.Pointer {
	cArr := (*unsafe.Pointer)(unsafe.Pointer(unsafe.SliceData(f.data[:])))
	return ([AV_NUM_DATA_POINTERS]unsafe.Pointer)(unsafe.Slice(cArr, AV_NUM_DATA_POINTERS))
}

/**
 * pointer to the picture/channel planes.
 * This might be different from the first allocated byte. For video,
 * it could even point to the end of the image data.
 *
 * All pointers in data and extended_data must point into one of the
 * AVBufferRef in buf or extended_buf.
 *
 * Some decoders access areas outside 0,0 - width,height, please
 * see avcodec_align_dimensions2(). Some filters and swscale can read
 * up to 16 bytes beyond the planes, if these filters are to be used,
 * then 16 extra bytes must be allocated.
 *
 * NOTE: Pointers not needed by the format MUST be set to NULL.
 *
 * @attention In case of video, the data[] pointers can point to the
 * end of image data in order to reverse line order, when used in
 * combination with negative values in the linesize[] array.
 */
func (f *CAVFrame) SetData(data [AV_NUM_DATA_POINTERS]unsafe.Pointer) {
	cArr := (**C.uint8_t)(unsafe.Pointer(unsafe.SliceData(data[:])))
	f.data = ([AV_NUM_DATA_POINTERS]*C.uint8_t)(unsafe.Slice(cArr, AV_NUM_DATA_POINTERS))
}

/**
 * For video, a positive or negative value, which is typically indicating
 * the size in bytes of each picture line, but it can also be:
 * - the negative byte size of lines for vertical flipping
 *   (with data[n] pointing to the end of the data
 * - a positive or negative multiple of the byte size as for accessing
 *   even and odd fields of a frame (possibly flipped)
 *
 * For audio, only linesize[0] may be set. For planar audio, each channel
 * plane must be the same size.
 *
 * For video the linesizes should be multiples of the CPUs alignment
 * preference, this is 16 or 32 for modern desktop CPUs.
 * Some code requires such alignment other code can be slower without
 * correct alignment, for yet other it makes no difference.
 *
 * @note The linesize may be larger than the size of usable data -- there
 * may be extra padding present for performance reasons.
 *
 * @attention In case of video, line size values can be negative to achieve
 * a vertically inverted iteration over image lines.
 */
func (f *CAVFrame) GetLinesize() [AV_NUM_DATA_POINTERS]ctypes.Int {
	cArr := unsafe.SliceData(f.linesize[:])
	return ([AV_NUM_DATA_POINTERS]ctypes.Int)(unsafe.Slice((*ctypes.Int)(unsafe.Pointer(cArr)), AV_NUM_DATA_POINTERS))
}

/**
 * For video, a positive or negative value, which is typically indicating
 * the size in bytes of each picture line, but it can also be:
 * - the negative byte size of lines for vertical flipping
 *   (with data[n] pointing to the end of the data
 * - a positive or negative multiple of the byte size as for accessing
 *   even and odd fields of a frame (possibly flipped)
 *
 * For audio, only linesize[0] may be set. For planar audio, each channel
 * plane must be the same size.
 *
 * For video the linesizes should be multiples of the CPUs alignment
 * preference, this is 16 or 32 for modern desktop CPUs.
 * Some code requires such alignment other code can be slower without
 * correct alignment, for yet other it makes no difference.
 *
 * @note The linesize may be larger than the size of usable data -- there
 * may be extra padding present for performance reasons.
 *
 * @attention In case of video, line size values can be negative to achieve
 * a vertically inverted iteration over image lines.
 */
func (f *CAVFrame) SetLinesize(linesize [AV_NUM_DATA_POINTERS]ctypes.Int) {
	cArr := (*C.int)(unsafe.Pointer(unsafe.SliceData(linesize[:])))
	f.linesize = ([AV_NUM_DATA_POINTERS]C.int)(unsafe.Slice(cArr, AV_NUM_DATA_POINTERS))
}

/**
 * pointers to the data planes/channels.
 *
 * For video, this should simply point to data[].
 *
 * For planar audio, each channel has a separate data pointer, and
 * linesize[0] contains the size of each channel buffer.
 * For packed audio, there is just one data pointer, and linesize[0]
 * contains the total size of the buffer for all channels.
 *
 * Note: Both data and extended_data should always be set in a valid frame,
 * but for planar audio with more channels that can fit in data,
 * extended_data must be used in order to access all channels.
 */
func (f *CAVFrame) GetExtendedData() *unsafe.Pointer {
	return (*unsafe.Pointer)(unsafe.Pointer(f.extended_data))
}

/**
 * pointers to the data planes/channels.
 *
 * For video, this should simply point to data[].
 *
 * For planar audio, each channel has a separate data pointer, and
 * linesize[0] contains the size of each channel buffer.
 * For packed audio, there is just one data pointer, and linesize[0]
 * contains the total size of the buffer for all channels.
 *
 * Note: Both data and extended_data should always be set in a valid frame,
 * but for planar audio with more channels that can fit in data,
 * extended_data must be used in order to access all channels.
 */
func (f *CAVFrame) SetExtendedData(extendedData *unsafe.Pointer) {
	f.extended_data = (**C.uint8_t)(unsafe.Pointer(extendedData))
}

/**
 * @name Video dimensions
 * Video frames only. The coded dimensions (in pixels) of the video frame,
 * i.e. the size of the rectangle that contains some well-defined values.
 *
 * @note The part of the frame intended for display/presentation is further
 * restricted by the @ref cropping "Cropping rectangle".
 * @{
 */

func (f *CAVFrame) GetWidth() int {
	return int(f.width)
}

func (f *CAVFrame) SetWidth(width int) {
	f.width = C.int(width)
}

func (f *CAVFrame) GetHeight() int {
	return int(f.height)
}

func (f *CAVFrame) SetHeight(height int) {
	f.height = C.int(height)
}

/**
 * @}
 */

/**
 * number of audio samples (per channel) described by this frame
 */
func (f *CAVFrame) GetNbSamples() int {
	return int(f.nb_samples)
}

/**
 * number of audio samples (per channel) described by this frame
 */
func (f *CAVFrame) SetNbSamples(nbSamples int) {
	f.nb_samples = C.int(nbSamples)
}

/**
 * format of the frame, -1 if unknown or unset
 * Values correspond to enum AVPixelFormat for video frames,
 * enum AVSampleFormat for audio)
 */
func (f *CAVFrame) GetFormat() int {
	return int(f.format)
}

/**
 * format of the frame, -1 if unknown or unset
 * Values correspond to enum AVPixelFormat for video frames,
 * enum AVSampleFormat for audio)
 */
func (f *CAVFrame) SetFormat(format int) {
	f.format = C.int(format)
}

// #if FF_API_FRAME_KEY
// 	/**
// 	 * 1 -> keyframe, 0-> not
// 	 *
// 	 * @deprecated Use AV_FRAME_FLAG_KEY instead
// 	 */
// 	attribute_deprecated
// 	int key_frame;
// #endif

/**
 * Picture type of the frame.
 */
func (f *CAVFrame) GetPictType() CAVPictureType {
	return CAVPictureType(f.pict_type)
}

/**
 * Picture type of the frame.
 */
func (f *CAVFrame) SetPictType(pictType CAVPictureType) {
	f.pict_type = C.enum_AVPictureType(pictType)
}

/**
 * Sample aspect ratio for the video frame, 0/1 if unknown/unspecified.
 */
func (f *CAVFrame) GetSampleAspectRatio() CAVRational {
	return CAVRational(f.sample_aspect_ratio)
}

/**
 * Sample aspect ratio for the video frame, 0/1 if unknown/unspecified.
 */
func (f *CAVFrame) SetSampleAspectRatio(sampleAspectRatio CAVRational) {
	f.sample_aspect_ratio = C.AVRational(sampleAspectRatio)
}

/**
 * Presentation timestamp in time_base units (time when frame should be shown to user).
 */
func (f *CAVFrame) GetPts() int64 {
	return int64(f.pts)
}

/**
 * Presentation timestamp in time_base units (time when frame should be shown to user).
 */
func (f *CAVFrame) SetPts(pts int64) {
	f.pts = C.int64_t(pts)
}

/**
 * DTS copied from the AVPacket that triggered returning this frame. (if frame threading isn't used)
 * This is also the Presentation time of this AVFrame calculated from
 * only AVPacket.dts values without pts values.
 */
func (f *CAVFrame) GetPktDts() int64 {
	return int64(f.pkt_dts)
}

/**
 * DTS copied from the AVPacket that triggered returning this frame. (if frame threading isn't used)
 * This is also the Presentation time of this AVFrame calculated from
 * only AVPacket.dts values without pts values.
 */
func (f *CAVFrame) SetPktDts(pktDts int64) {
	f.pkt_dts = C.int64_t(pktDts)
}

/**
 * Time base for the timestamps in this frame.
 * In the future, this field may be set on frames output by decoders or
 * filters, but its value will be by default ignored on input to encoders
 * or filters.
 */
func (f *CAVFrame) GetTimeBase() CAVRational {
	return CAVRational(f.time_base)
}

/**
 * Time base for the timestamps in this frame.
 * In the future, this field may be set on frames output by decoders or
 * filters, but its value will be by default ignored on input to encoders
 * or filters.
 */
func (f *CAVFrame) SetTimeBase(timeBase CAVRational) {
	f.time_base = C.AVRational(timeBase)
}

// #if FF_API_FRAME_PICTURE_NUMBER
// 	/**
// 	 * picture number in bitstream order
// 	 */
// 	attribute_deprecated
// 	int coded_picture_number;
// 	/**
// 	 * picture number in display order
// 	 */
// 	attribute_deprecated
// 	int display_picture_number;
// #endif

/**
* quality (between 1 (good) and FF_LAMBDA_MAX (bad))
 */
func (f *CAVFrame) GetQuality() int {
	return int(f.quality)
}

/**
* quality (between 1 (good) and FF_LAMBDA_MAX (bad))
 */
func (f *CAVFrame) SetQuality(quality int) {
	f.quality = C.int(quality)
}

/**
 * Frame owner's private data.
 *
 * This field may be set by the code that allocates/owns the frame data.
 * It is then not touched by any library functions, except:
 * - it is copied to other references by av_frame_copy_props() (and hence by
 *   av_frame_ref());
 * - it is set to NULL when the frame is cleared by av_frame_unref()
 * - on the caller's explicit request. E.g. libavcodec encoders/decoders
 *   will copy this field to/from @ref AVPacket "AVPackets" if the caller sets
 *   @ref AV_CODEC_FLAG_COPY_OPAQUE.
 *
 * @see opaque_ref the reference-counted analogue
 */
func (f *CAVFrame) GetOpaque() unsafe.Pointer {
	return f.opaque
}

/**
 * Frame owner's private data.
 *
 * This field may be set by the code that allocates/owns the frame data.
 * It is then not touched by any library functions, except:
 * - it is copied to other references by av_frame_copy_props() (and hence by
 *   av_frame_ref());
 * - it is set to NULL when the frame is cleared by av_frame_unref()
 * - on the caller's explicit request. E.g. libavcodec encoders/decoders
 *   will copy this field to/from @ref AVPacket "AVPackets" if the caller sets
 *   @ref AV_CODEC_FLAG_COPY_OPAQUE.
 *
 * @see opaque_ref the reference-counted analogue
 */
func (f *CAVFrame) SetOpaque(opaque unsafe.Pointer) {
	f.opaque = opaque
}

/**
 * Number of fields in this frame which should be repeated, i.e. the total
 * duration of this frame should be repeat_pict + 2 normal field durations.
 *
 * For interlaced frames this field may be set to 1, which signals that this
 * frame should be presented as 3 fields: beginning with the first field (as
 * determined by AV_FRAME_FLAG_TOP_FIELD_FIRST being set or not), followed
 * by the second field, and then the first field again.
 *
 * For progressive frames this field may be set to a multiple of 2, which
 * signals that this frame's duration should be (repeat_pict + 2) / 2
 * normal frame durations.
 *
 * @note This field is computed from MPEG2 repeat_first_field flag and its
 * associated flags, H.264 pic_struct from picture timing SEI, and
 * their analogues in other codecs. Typically it should only be used when
 * higher-layer timing information is not available.
 */
func (f *CAVFrame) GetRepeatPict() int {
	return int(f.repeat_pict)
}

/**
 * Number of fields in this frame which should be repeated, i.e. the total
 * duration of this frame should be repeat_pict + 2 normal field durations.
 *
 * For interlaced frames this field may be set to 1, which signals that this
 * frame should be presented as 3 fields: beginning with the first field (as
 * determined by AV_FRAME_FLAG_TOP_FIELD_FIRST being set or not), followed
 * by the second field, and then the first field again.
 *
 * For progressive frames this field may be set to a multiple of 2, which
 * signals that this frame's duration should be (repeat_pict + 2) / 2
 * normal frame durations.
 *
 * @note This field is computed from MPEG2 repeat_first_field flag and its
 * associated flags, H.264 pic_struct from picture timing SEI, and
 * their analogues in other codecs. Typically it should only be used when
 * higher-layer timing information is not available.
 */
func (f *CAVFrame) SetRepeatPict(repeatPict int) {
	f.repeat_pict = C.int(repeatPict)
}

// #if FF_API_INTERLACED_FRAME
// 	/**
// 	 * The content of the picture is interlaced.
// 	 *
// 	 * @deprecated Use AV_FRAME_FLAG_INTERLACED instead
// 	 */
// 	attribute_deprecated
// 	int interlaced_frame;

// 	/**
// 	 * If the content is interlaced, is top field displayed first.
// 	 *
// 	 * @deprecated Use AV_FRAME_FLAG_TOP_FIELD_FIRST instead
// 	 */
// 	attribute_deprecated
// 	int top_field_first;
// #endif

// #if FF_API_PALETTE_HAS_CHANGED
// 	/**
// 	 * Tell user application that palette has changed from previous frame.
// 	 */
// 	attribute_deprecated
// 	int palette_has_changed;
// #endif

// #if FF_API_REORDERED_OPAQUE
// 	/**
// 	 * reordered opaque 64 bits (generally an integer or a double precision float
// 	 * PTS but can be anything).
// 	 * The user sets AVCodecContext.reordered_opaque to represent the input at
// 	 * that time,
// 	 * the decoder reorders values as needed and sets AVFrame.reordered_opaque
// 	 * to exactly one of the values provided by the user through AVCodecContext.reordered_opaque
// 	 *
// 	 * @deprecated Use AV_CODEC_FLAG_COPY_OPAQUE instead
// 	 */
// 	attribute_deprecated
// 	int64_t reordered_opaque;
// #endif

/**
 * Sample rate of the audio data.
 */
func (f *CAVFrame) GetSampleRate() int {
	return int(f.sample_rate)
}

/**
 * Sample rate of the audio data.
 */
func (f *CAVFrame) SetSampleRate(sampleRate int) {
	f.sample_rate = C.int(sampleRate)
}

// #if FF_API_OLD_CHANNEL_LAYOUT
// 	/**
// 	 * Channel layout of the audio data.
// 	 * @deprecated use ch_layout instead
// 	 */
// 	attribute_deprecated
// 	uint64_t channel_layout;
// #endif

/**
 * AVBuffer references backing the data for this frame. All the pointers in
 * data and extended_data must point inside one of the buffers in buf or
 * extended_buf. This array must be filled contiguously -- if buf[i] is
 * non-NULL then buf[j] must also be non-NULL for all j < i.
 *
 * There may be at most one AVBuffer per data plane, so for video this array
 * always contains all the references. For planar audio with more than
 * AV_NUM_DATA_POINTERS channels, there may be more buffers than can fit in
 * this array. Then the extra AVBufferRef pointers are stored in the
 * extended_buf array.
 */
func (f *CAVFrame) GetBuf() [AV_NUM_DATA_POINTERS]*CAVBufferRef {
	arrp := (**CAVBufferRef)(unsafe.Pointer(unsafe.SliceData(f.buf[:])))
	return ([AV_NUM_DATA_POINTERS]*CAVBufferRef)(unsafe.Slice(arrp, AV_NUM_DATA_POINTERS))
}

/**
 * AVBuffer references backing the data for this frame. All the pointers in
 * data and extended_data must point inside one of the buffers in buf or
 * extended_buf. This array must be filled contiguously -- if buf[i] is
 * non-NULL then buf[j] must also be non-NULL for all j < i.
 *
 * There may be at most one AVBuffer per data plane, so for video this array
 * always contains all the references. For planar audio with more than
 * AV_NUM_DATA_POINTERS channels, there may be more buffers than can fit in
 * this array. Then the extra AVBufferRef pointers are stored in the
 * extended_buf array.
 */
func (f *CAVFrame) SetBuf(buf [AV_NUM_DATA_POINTERS]*CAVBufferRef) {
	arrp := (**C.AVBufferRef)(unsafe.Pointer(unsafe.SliceData(buf[:])))
	f.buf = ([AV_NUM_DATA_POINTERS]*C.AVBufferRef)(unsafe.Slice(arrp, AV_NUM_DATA_POINTERS))
}

/**
 * For planar audio which requires more than AV_NUM_DATA_POINTERS
 * AVBufferRef pointers, this array will hold all the references which
 * cannot fit into AVFrame.buf.
 *
 * Note that this is different from AVFrame.extended_data, which always
 * contains all the pointers. This array only contains the extra pointers,
 * which cannot fit into AVFrame.buf.
 *
 * This array is always allocated using av_malloc() by whoever constructs
 * the frame. It is freed in av_frame_unref().
 */
func (f *CAVFrame) GetExtendedBuf() **CAVBufferRef {
	return (**CAVBufferRef)(unsafe.Pointer(f.extended_buf))
}

/**
 * For planar audio which requires more than AV_NUM_DATA_POINTERS
 * AVBufferRef pointers, this array will hold all the references which
 * cannot fit into AVFrame.buf.
 *
 * Note that this is different from AVFrame.extended_data, which always
 * contains all the pointers. This array only contains the extra pointers,
 * which cannot fit into AVFrame.buf.
 *
 * This array is always allocated using av_malloc() by whoever constructs
 * the frame. It is freed in av_frame_unref().
 */
func (f *CAVFrame) SetExtendedBuf(extendedBuf **CAVBufferRef) {
	f.extended_buf = (**C.AVBufferRef)(unsafe.Pointer(extendedBuf))
}

/**
 * Number of elements in extended_buf.
 */
func (f *CAVFrame) GetNbExtendedBuf() int {
	return int(f.nb_extended_buf)
}

/**
 * Number of elements in extended_buf.
 */
func (f *CAVFrame) SetNbExtendedBuf(nbExtendedBuf int) {
	f.nb_extended_buf = C.int(nbExtendedBuf)
}

func (f *CAVFrame) GetSideData() **CAVFrameSideData {
	return (**CAVFrameSideData)(unsafe.Pointer(f.side_data))
}

func (f *CAVFrame) SetSideData(sideData **CAVFrameSideData) {
	f.side_data = (**C.AVFrameSideData)(unsafe.Pointer(sideData))
}

func (f *CAVFrame) GetNbSideData() int {
	return int(f.nb_side_data)
}

func (f *CAVFrame) SetNbSideData(nbSideData int) {
	f.nb_side_data = C.int(nbSideData)
}

/**
 * @defgroup lavu_frame_flags AV_FRAME_FLAGS
 * @ingroup lavu_frame
 * Flags describing additional frame properties.
 *
 * @{
 */

/**
 * The frame data may be corrupted, e.g. due to decoding errors.
 */
const AV_FRAME_FLAG_CORRUPT = C.AV_FRAME_FLAG_CORRUPT

/**
 * A flag to mark frames that are keyframes.
 */
const AV_FRAME_FLAG_KEY = C.AV_FRAME_FLAG_KEY

/**
 * A flag to mark the frames which need to be decoded, but shouldn't be output.
 */
const AV_FRAME_FLAG_DISCARD = C.AV_FRAME_FLAG_DISCARD

/**
 * A flag to mark frames whose content is interlaced.
 */
const AV_FRAME_FLAG_INTERLACED = C.AV_FRAME_FLAG_INTERLACED

/**
 * A flag to mark frames where the top field is displayed first if the content
 * is interlaced.
 */
const AV_FRAME_FLAG_TOP_FIELD_FIRST = C.AV_FRAME_FLAG_TOP_FIELD_FIRST

/**
 * @}
 */

/**
 * Frame flags, a combination of @ref lavu_frame_flags
 */
func (f *CAVFrame) GetFlags() int {
	return int(f.flags)
}

/**
 * Frame flags, a combination of @ref lavu_frame_flags
 */
func (f *CAVFrame) SetFlags(flags int) {
	f.flags = C.int(flags)
}

/**
 * MPEG vs JPEG YUV range.
 * - encoding: Set by user
 * - decoding: Set by libavcodec
 */
func (f *CAVFrame) GetColorRange() CAVColorRange {
	return CAVColorRange(f.color_range)
}
func (f *CAVFrame) SetColorRange(colorRange CAVColorRange) {
	f.color_range = C.enum_AVColorRange(colorRange)
}

func (f *CAVFrame) GetColorPrimaries() CAVColorPrimaries {
	return CAVColorPrimaries(f.color_primaries)
}
func (f *CAVFrame) SetColorPrimaries(colorPrimaries CAVColorPrimaries) {
	f.color_primaries = C.enum_AVColorPrimaries(colorPrimaries)
}

func (f *CAVFrame) GetColorTrc() CAVColorTransferCharacteristic {
	return CAVColorTransferCharacteristic(f.color_trc)
}
func (f *CAVFrame) SetColorTrc(colorTrc CAVColorTransferCharacteristic) {
	f.color_trc = C.enum_AVColorTransferCharacteristic(colorTrc)
}

/**
 * YUV colorspace type.
 * - encoding: Set by user
 * - decoding: Set by libavcodec
 */
func (f *CAVFrame) GetColorspace() CAVColorSpace {
	return CAVColorSpace(f.colorspace)
}
func (f *CAVFrame) SetColorspace(colorspace CAVColorSpace) {
	f.colorspace = C.enum_AVColorSpace(colorspace)
}

func (f *CAVFrame) GetChromaLocation() CAVChromaLocation {
	return CAVChromaLocation(f.chroma_location)
}
func (f *CAVFrame) SetChromaLocation(chromaLocation CAVChromaLocation) {
	f.chroma_location = C.enum_AVChromaLocation(chromaLocation)
}

/**
 * frame timestamp estimated using various heuristics, in stream time base
 * - encoding: unused
 * - decoding: set by libavcodec, read by user.
 */
func (f *CAVFrame) GetBestEffortTimestamp() int64 {
	return int64(f.best_effort_timestamp)
}

/**
 * frame timestamp estimated using various heuristics, in stream time base
 * - encoding: unused
 * - decoding: set by libavcodec, read by user.
 */
func (f *CAVFrame) SetBestEffortTimestamp(bestEffortTimestamp int64) {
	f.best_effort_timestamp = C.int64_t(bestEffortTimestamp)
}

// #if FF_API_FRAME_PKT
// 	/**
// 	 * reordered pos from the last AVPacket that has been input into the decoder
// 	 * - encoding: unused
// 	 * - decoding: Read by user.
// 	 * @deprecated use AV_CODEC_FLAG_COPY_OPAQUE to pass through arbitrary user
// 	 *             data from packets to frames
// 	 */
// 	attribute_deprecated
// 	int64_t pkt_pos;
// #endif

// #if FF_API_PKT_DURATION
// 	/**
// 	 * duration of the corresponding packet, expressed in
// 	 * AVStream->time_base units, 0 if unknown.
// 	 * - encoding: unused
// 	 * - decoding: Read by user.
// 	 *
// 	 * @deprecated use duration instead
// 	 */
// 	attribute_deprecated
// 	int64_t pkt_duration;
// #endif

/**
 * metadata.
 * - encoding: Set by user.
 * - decoding: Set by libavcodec.
 */
func (f *CAVFrame) GetMetadata() *CAVDictionary {
	return (*CAVDictionary)(f.metadata)
}

/**
 * metadata.
 * - encoding: Set by user.
 * - decoding: Set by libavcodec.
 */
func (f *CAVFrame) SetMetadata(metadata *CAVDictionary) {
	f.metadata = (*C.AVDictionary)(metadata)
}

/**
 * decode error flags of the frame, set to a combination of
 * FF_DECODE_ERROR_xxx flags if the decoder produced a frame, but there
 * were errors during the decoding.
 * - encoding: unused
 * - decoding: set by libavcodec, read by user.
 */
func (f *CAVFrame) GetDecodeErrorFlags() int {
	return int(f.decode_error_flags)
}

/**
 * decode error flags of the frame, set to a combination of
 * FF_DECODE_ERROR_xxx flags if the decoder produced a frame, but there
 * were errors during the decoding.
 * - encoding: unused
 * - decoding: set by libavcodec, read by user.
 */
func (f *CAVFrame) SetDecodeErrorFlags(decodeErrorFlags int) {
	f.decode_error_flags = C.int(decodeErrorFlags)
}

const (
	FF_DECODE_ERROR_INVALID_BITSTREAM  = C.FF_DECODE_ERROR_INVALID_BITSTREAM
	FF_DECODE_ERROR_MISSING_REFERENCE  = C.FF_DECODE_ERROR_MISSING_REFERENCE
	FF_DECODE_ERROR_CONCEALMENT_ACTIVE = C.FF_DECODE_ERROR_CONCEALMENT_ACTIVE
	FF_DECODE_ERROR_DECODE_SLICES      = C.FF_DECODE_ERROR_DECODE_SLICES
)

// #if FF_API_OLD_CHANNEL_LAYOUT
// 	/**
// 	 * number of audio channels, only used for audio.
// 	 * - encoding: unused
// 	 * - decoding: Read by user.
// 	 * @deprecated use ch_layout instead
// 	 */
// 	attribute_deprecated
// 	int channels;
// #endif

// #if FF_API_FRAME_PKT
// 	/**
// 	 * size of the corresponding packet containing the compressed
// 	 * frame.
// 	 * It is set to a negative value if unknown.
// 	 * - encoding: unused
// 	 * - decoding: set by libavcodec, read by user.
// 	 * @deprecated use AV_CODEC_FLAG_COPY_OPAQUE to pass through arbitrary user
// 	 *             data from packets to frames
// 	 */
// 	attribute_deprecated
// 	int pkt_size;
// #endif

/**
 * For hwaccel-format frames, this should be a reference to the
 * AVHWFramesContext describing the frame.
 */
func (f *CAVFrame) GetHwFramesCtx() *CAVBufferRef {
	return (*CAVBufferRef)(f.hw_frames_ctx)
}

/**
 * For hwaccel-format frames, this should be a reference to the
 * AVHWFramesContext describing the frame.
 */
func (f *CAVFrame) SetHwFramesCtx(hwFramesCtx *CAVBufferRef) {
	f.hw_frames_ctx = (*C.AVBufferRef)(hwFramesCtx)
}

/**
 * Frame owner's private data.
 *
 * This field may be set by the code that allocates/owns the frame data.
 * It is then not touched by any library functions, except:
 * - a new reference to the underlying buffer is propagated by
 *   av_frame_copy_props() (and hence by av_frame_ref());
 * - it is unreferenced in av_frame_unref();
 * - on the caller's explicit request. E.g. libavcodec encoders/decoders
 *   will propagate a new reference to/from @ref AVPacket "AVPackets" if the
 *   caller sets @ref AV_CODEC_FLAG_COPY_OPAQUE.
 *
 * @see opaque the plain pointer analogue
 */
func (f *CAVFrame) GetOpaqueRef() *CAVBufferRef {
	return (*CAVBufferRef)(f.opaque_ref)
}

/**
 * Frame owner's private data.
 *
 * This field may be set by the code that allocates/owns the frame data.
 * It is then not touched by any library functions, except:
 * - a new reference to the underlying buffer is propagated by
 *   av_frame_copy_props() (and hence by av_frame_ref());
 * - it is unreferenced in av_frame_unref();
 * - on the caller's explicit request. E.g. libavcodec encoders/decoders
 *   will propagate a new reference to/from @ref AVPacket "AVPackets" if the
 *   caller sets @ref AV_CODEC_FLAG_COPY_OPAQUE.
 *
 * @see opaque the plain pointer analogue
 */
func (f *CAVFrame) SetOpaqueRef(opaqueRef *CAVBufferRef) {
	f.opaque_ref = (*C.AVBufferRef)(opaqueRef)
}

/**
 * @anchor cropping
 * @name Cropping
 * Video frames only. The number of pixels to discard from the the
 * top/bottom/left/right border of the frame to obtain the sub-rectangle of
 * the frame intended for presentation.
 * @{
 */
func (f *CAVFrame) GetCropTop() ctypes.SizeT {
	return ctypes.SizeT(f.crop_top)
}
func (f *CAVFrame) SetCropTop(cropTop ctypes.SizeT) {
	f.crop_top = C.size_t(cropTop)
}
func (f *CAVFrame) GetCropBottom() ctypes.SizeT {
	return ctypes.SizeT(f.crop_bottom)
}
func (f *CAVFrame) SetCropBottom(cropBottom ctypes.SizeT) {
	f.crop_bottom = C.size_t(cropBottom)
}
func (f *CAVFrame) GetCropLeft() ctypes.SizeT {
	return ctypes.SizeT(f.crop_left)
}
func (f *CAVFrame) SetCropLeft(cropLeft ctypes.SizeT) {
	f.crop_left = C.size_t(cropLeft)
}
func (f *CAVFrame) GetCropRight() ctypes.SizeT {
	return ctypes.SizeT(f.crop_right)
}
func (f *CAVFrame) SetCropRight(cropRight ctypes.SizeT) {
	f.crop_right = C.size_t(cropRight)
}

/**
 * @}
 */

/**
 * AVBufferRef for internal use by a single libav* library.
 * Must not be used to transfer data between libraries.
 * Has to be NULL when ownership of the frame leaves the respective library.
 *
 * Code outside the FFmpeg libs should never check or change the contents of the buffer ref.
 *
 * FFmpeg calls av_buffer_unref() on it when the frame is unreferenced.
 * av_frame_copy_props() calls create a new reference with av_buffer_ref()
 * for the target frame's private_ref field.
 */
func (f *CAVFrame) GetPrivateRef() *CAVBufferRef {
	return (*CAVBufferRef)(f.private_ref)
}

/**
 * AVBufferRef for internal use by a single libav* library.
 * Must not be used to transfer data between libraries.
 * Has to be NULL when ownership of the frame leaves the respective library.
 *
 * Code outside the FFmpeg libs should never check or change the contents of the buffer ref.
 *
 * FFmpeg calls av_buffer_unref() on it when the frame is unreferenced.
 * av_frame_copy_props() calls create a new reference with av_buffer_ref()
 * for the target frame's private_ref field.
 */
func (f *CAVFrame) SetPrivateRef(privateRef *CAVBufferRef) {
	f.private_ref = (*C.AVBufferRef)(privateRef)
}

/**
 * Channel layout of the audio data.
 */
func (f *CAVFrame) GetChLayout() CAVChannelLayout {
	return CAVChannelLayout(f.ch_layout)
}

/**
 * Channel layout of the audio data.
 */
func (f *CAVFrame) SetChLayout(chLayout CAVChannelLayout) {
	f.ch_layout = C.AVChannelLayout(chLayout)
}

/**
 * Channel layout of the audio data.
 */
func (f *CAVFrame) GetChLayoutPtr() *CAVChannelLayout {
	return (*CAVChannelLayout)(&f.ch_layout)
}

/**
 * Duration of the frame, in the same units as pts. 0 if unknown.
 */
func (f *CAVFrame) GetDuration() int64 {
	return int64(f.duration)
}

/**
 * Duration of the frame, in the same units as pts. 0 if unknown.
 */
func (f *CAVFrame) SetDuration(duration int64) {
	f.duration = C.int64_t(duration)
}

//#endregion CAVFrame

/**
 * Allocate an AVFrame and set its fields to default values.  The resulting
 * struct must be freed using av_frame_free().
 *
 * @return An AVFrame filled with default values or NULL on failure.
 *
 * @note this only allocates the AVFrame itself, not the data buffers. Those
 * must be allocated through other means, e.g. with av_frame_get_buffer() or
 * manually.
 */
func AvFrameAlloc() *CAVFrame {
	return (*CAVFrame)(C.av_frame_alloc())
}

/**
 * Free the frame and any dynamically allocated objects in it,
 * e.g. extended_data. If the frame is reference counted, it will be
 * unreferenced first.
 *
 * @param frame frame to be freed. The pointer will be set to NULL.
 */
func AvFrameFree(frame **CAVFrame) {
	C.av_frame_free((**C.AVFrame)(unsafe.Pointer(frame)))
}

/**
 * Set up a new reference to the data described by the source frame.
 *
 * Copy frame properties from src to dst and create a new reference for each
 * AVBufferRef from src.
 *
 * If src is not reference counted, new buffers are allocated and the data is
 * copied.
 *
 * @warning: dst MUST have been either unreferenced with av_frame_unref(dst),
 *           or newly allocated with av_frame_alloc() before calling this
 *           function, or undefined behavior will occur.
 *
 * @return 0 on success, a negative AVERROR on error
 */
func AvFrameRef(dst *CAVFrame, src *CAVFrame) int {
	return int(C.av_frame_ref((*C.AVFrame)(dst), (*C.AVFrame)(src)))
}

/**
 * Ensure the destination frame refers to the same data described by the source
 * frame, either by creating a new reference for each AVBufferRef from src if
 * they differ from those in dst, by allocating new buffers and copying data if
 * src is not reference counted, or by unrefencing it if src is empty.
 *
 * Frame properties on dst will be replaced by those from src.
 *
 * @return 0 on success, a negative AVERROR on error. On error, dst is
 *         unreferenced.
 */
func AvFrameReplace(dst *CAVFrame, src *CAVFrame) int {
	return int(C.av_frame_replace((*C.AVFrame)(dst), (*C.AVFrame)(src)))
}

/**
 * Create a new frame that references the same data as src.
 *
 * This is a shortcut for av_frame_alloc()+av_frame_ref().
 *
 * @return newly created AVFrame on success, NULL on error.
 */
func AvFrameClone(src *CAVFrame) *CAVFrame {
	return (*CAVFrame)(C.av_frame_clone((*C.AVFrame)(src)))
}

/**
 * Unreference all the buffers referenced by frame and reset the frame fields.
 */
func AvFrameUnref(frame *CAVFrame) {
	C.av_frame_unref((*C.AVFrame)(frame))
}

/**
 * Move everything contained in src to dst and reset src.
 *
 * @warning: dst is not unreferenced, but directly overwritten without reading
 *           or deallocating its contents. Call av_frame_unref(dst) manually
 *           before calling this function to ensure that no memory is leaked.
 */
func AvFrameMoveRef(dst *CAVFrame, src *CAVFrame) {
	C.av_frame_move_ref((*C.AVFrame)(dst), (*C.AVFrame)(src))
}

/**
 * Allocate new buffer(s) for audio or video data.
 *
 * The following fields must be set on frame before calling this function:
 * - format (pixel format for video, sample format for audio)
 * - width and height for video
 * - nb_samples and ch_layout for audio
 *
 * This function will fill AVFrame.data and AVFrame.buf arrays and, if
 * necessary, allocate and fill AVFrame.extended_data and AVFrame.extended_buf.
 * For planar formats, one buffer will be allocated for each plane.
 *
 * @warning: if frame already has been allocated, calling this function will
 *           leak memory. In addition, undefined behavior can occur in certain
 *           cases.
 *
 * @param frame frame in which to store the new buffers.
 * @param align Required buffer size alignment. If equal to 0, alignment will be
 *              chosen automatically for the current CPU. It is highly
 *              recommended to pass 0 here unless you know what you are doing.
 *
 * @return 0 on success, a negative AVERROR on error.
 */
func AvFrameGetBuffer(frame *CAVFrame, align int) int {
	return int(C.av_frame_get_buffer((*C.AVFrame)(frame), C.int(align)))
}

/**
 * Check if the frame data is writable.
 *
 * @return A positive value if the frame data is writable (which is true if and
 * only if each of the underlying buffers has only one reference, namely the one
 * stored in this frame). Return 0 otherwise.
 *
 * If 1 is returned the answer is valid until av_buffer_ref() is called on any
 * of the underlying AVBufferRefs (e.g. through av_frame_ref() or directly).
 *
 * @see av_frame_make_writable(), av_buffer_is_writable()
 */
func AvFrameIsWritable(frame *CAVFrame) int {
	return int(C.av_frame_is_writable((*C.AVFrame)(frame)))
}

/**
 * Ensure that the frame data is writable, avoiding data copy if possible.
 *
 * Do nothing if the frame is writable, allocate new buffers and copy the data
 * if it is not. Non-refcounted frames behave as non-writable, i.e. a copy
 * is always made.
 *
 * @return 0 on success, a negative AVERROR on error.
 *
 * @see av_frame_is_writable(), av_buffer_is_writable(),
 * av_buffer_make_writable()
 */
func AvFrameMakeWritable(frame *CAVFrame) int {
	return int(C.av_frame_make_writable((*C.AVFrame)(frame)))
}

/**
 * Copy the frame data from src to dst.
 *
 * This function does not allocate anything, dst must be already initialized and
 * allocated with the same parameters as src.
 *
 * This function only copies the frame data (i.e. the contents of the data /
 * extended data arrays), not any other properties.
 *
 * @return >= 0 on success, a negative AVERROR on error.
 */
func AvFrameCopy(dst *CAVFrame, src *CAVFrame) int {
	return int(C.av_frame_copy((*C.AVFrame)(dst), (*C.AVFrame)(src)))
}

/**
 * Copy only "metadata" fields from src to dst.
 *
 * Metadata for the purpose of this function are those fields that do not affect
 * the data layout in the buffers.  E.g. pts, sample rate (for audio) or sample
 * aspect ratio (for video), but not width/height or channel layout.
 * Side data is also copied.
 */
func AvFrameCopyProps(dst *CAVFrame, src *CAVFrame) int {
	return int(C.av_frame_copy_props((*C.AVFrame)(dst), (*C.AVFrame)(src)))
}

/**
 * Get the buffer reference a given data plane is stored in.
 *
 * @param frame the frame to get the plane's buffer from
 * @param plane index of the data plane of interest in frame->extended_data.
 *
 * @return the buffer reference that contains the plane or NULL if the input
 * frame is not valid.
 */
func AvFrameGetPlaneBuffer(frame *CAVFrame, plane int) *CAVBufferRef {
	return (*CAVBufferRef)(C.av_frame_get_plane_buffer((*C.AVFrame)(frame), C.int(plane)))
}

/**
 * Add a new side data to a frame.
 *
 * @param frame a frame to which the side data should be added
 * @param type type of the added side data
 * @param size size of the side data
 *
 * @return newly added side data on success, NULL on error
 */
func AvFrameNewSideData(frame *CAVFrame, _type CAVFrameSideDataType, size ctypes.SizeT) *CAVFrameSideData {
	return (*CAVFrameSideData)(C.av_frame_new_side_data((*C.AVFrame)(frame), (C.enum_AVFrameSideDataType)(_type), C.size_t(size)))
}

/**
 * Add a new side data to a frame from an existing AVBufferRef
 *
 * @param frame a frame to which the side data should be added
 * @param type  the type of the added side data
 * @param buf   an AVBufferRef to add as side data. The ownership of
 *              the reference is transferred to the frame.
 *
 * @return newly added side data on success, NULL on error. On failure
 *         the frame is unchanged and the AVBufferRef remains owned by
 *         the caller.
 */
func AvFrameNewSideDataFromBuf(frame *CAVFrame, _type CAVFrameSideDataType, buf *CAVBufferRef) *CAVFrameSideData {
	return (*CAVFrameSideData)(C.av_frame_new_side_data_from_buf((*C.AVFrame)(frame), (C.enum_AVFrameSideDataType)(_type), (*C.AVBufferRef)(buf)))
}

/**
 * @return a pointer to the side data of a given type on success, NULL if there
 * is no side data with such type in this frame.
 */
func AvFrameGetSideData(frame *CAVFrame, _type CAVFrameSideDataType) *CAVFrameSideData {
	return (*CAVFrameSideData)(C.av_frame_get_side_data((*C.AVFrame)(frame), (C.enum_AVFrameSideDataType)(_type)))
}

/**
 * Remove and free all side data instances of the given type.
 */
func AvFrameRemoveSideData(frame *CAVFrame, _type CAVFrameSideDataType) {
	C.av_frame_remove_side_data((*C.AVFrame)(frame), (C.enum_AVFrameSideDataType)(_type))
}

/**
 * Flags for frame cropping.
 */
const (
	/**
	 * Apply the maximum possible cropping, even if it requires setting the
	 * AVFrame.data[] entries to unaligned pointers. Passing unaligned data
	 * to FFmpeg API is generally not allowed, and causes undefined behavior
	 * (such as crashes). You can pass unaligned data only to FFmpeg APIs that
	 * are explicitly documented to accept it. Use this flag only if you
	 * absolutely know what you are doing.
	 */
	AV_FRAME_CROP_UNALIGNED = C.AV_FRAME_CROP_UNALIGNED
)

/**
 * Crop the given video AVFrame according to its crop_left/crop_top/crop_right/
 * crop_bottom fields. If cropping is successful, the function will adjust the
 * data pointers and the width/height fields, and set the crop fields to 0.
 *
 * In all cases, the cropping boundaries will be rounded to the inherent
 * alignment of the pixel format. In some cases, such as for opaque hwaccel
 * formats, the left/top cropping is ignored. The crop fields are set to 0 even
 * if the cropping was rounded or ignored.
 *
 * @param frame the frame which should be cropped
 * @param flags Some combination of AV_FRAME_CROP_* flags, or 0.
 *
 * @return >= 0 on success, a negative AVERROR on error. If the cropping fields
 * were invalid, AVERROR(ERANGE) is returned, and nothing is changed.
 */
func AvFrameApplyCropping(frame *CAVFrame, flags int) int {
	return int(C.av_frame_apply_cropping((*C.AVFrame)(frame), C.int(flags)))
}

/**
 * @return a string identifying the side data type
 */
func AvFrameSideDataName(_type CAVFrameSideDataType) string {
	return C.GoString(C.av_frame_side_data_name((C.enum_AVFrameSideDataType)(_type)))
}

/**
 * @}
 */

//  #endif /* AVUTIL_FRAME_H */
