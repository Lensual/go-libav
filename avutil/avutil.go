package avutil

/*
#cgo LDFLAGS: -lavutil

#include "libavutil/avutil.h"
*/
import "C"

const (
	/**
	 * @defgroup lavu_const Constants
	 * @{
	 *
	 * @defgroup lavu_enc Encoding specific
	 *
	 * @note those definition should move to avcodec
	 * @{
	 */

	FF_LAMBDA_SHIFT = C.FF_LAMBDA_SHIFT
	FF_LAMBDA_SCALE = C.FF_LAMBDA_SCALE
	FF_QP2LAMBDA    = C.FF_QP2LAMBDA ///< factor to convert from H.263 QP to lambda
	FF_LAMBDA_MAX   = C.FF_LAMBDA_MAX

	FF_QUALITY_SCALE = C.FF_QUALITY_SCALE //FIXME maybe remove

	/**
	 * @}
	 * @defgroup lavu_time Timestamp specific
	 *
	 * FFmpeg internal timebase and timestamp definitions
	 *
	 * @{
	 */

	/**
	 * @brief Undefined timestamp value
	 *
	 * Usually reported by demuxer that work on containers that do not provide
	 * either pts or dts.
	 */

	AV_NOPTS_VALUE = C.AV_NOPTS_VALUE

	/**
	 * Internal time base represented as integer
	 */

	AV_TIME_BASE = C.AV_TIME_BASE
)
