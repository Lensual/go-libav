package avcodec

/*
#cgo LDFLAGS: -lavcodec

#include "libavcodec/avcodec.h"
*/
import "C"

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
type AVSampleFormat C.enum_AVSampleFormat

const (
	SampleFmt_None   AVSampleFormat = C.AV_SAMPLE_FMT_NONE
	SampleFmt_U8     AVSampleFormat = C.AV_SAMPLE_FMT_U8  ///< unsigned 8 bits
	SampleFmt_S16    AVSampleFormat = C.AV_SAMPLE_FMT_S16 ///< signed 16 bits
	SampleFmt_S32    AVSampleFormat = C.AV_SAMPLE_FMT_S32 ///< signed 32 bits
	SampleFmt_Float  AVSampleFormat = C.AV_SAMPLE_FMT_FLT ///< float
	SampleFmt_Double AVSampleFormat = C.AV_SAMPLE_FMT_DBL ///< double

	SampleFmt_U8P     AVSampleFormat = C.AV_SAMPLE_FMT_U8P  ///< unsigned 8 bits, planar
	SampleFmt_S16P    AVSampleFormat = C.AV_SAMPLE_FMT_S16P ///< signed 16 bits, planar
	SampleFmt_S32P    AVSampleFormat = C.AV_SAMPLE_FMT_S32P ///< signed 32 bits, planar
	SampleFmt_FloatP  AVSampleFormat = C.AV_SAMPLE_FMT_FLTP ///< float, planar
	SampleFmt_DoubleP AVSampleFormat = C.AV_SAMPLE_FMT_DBLP ///< double, planar
	SampleFmt_S64     AVSampleFormat = C.AV_SAMPLE_FMT_S64  ///< signed 64 bits
	SampleFmt_S64P    AVSampleFormat = C.AV_SAMPLE_FMT_S64P ///< signed 64 bits, planar

	SampleFmt_NB AVSampleFormat = C.AV_SAMPLE_FMT_NB ///< Number of sample formats. DO NOT USE if linking dynamically
)
