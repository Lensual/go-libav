package avutil

/*
#cgo pkg-config: libavutil

#include "libavutil/error.h"
#include "libavutil/mem.h"
#include <stdlib.h>

// call marco method
int marco_AVERROR(int e) {
	return AVERROR(e);
}

// call marco method
char* marco_av_err2str(int errnum) {
	return av_strdup(av_err2str(errnum));
}
*/
import "C"
import "unsafe"

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
 * error code definitions
 */

//  #ifndef AVUTIL_ERROR_H
//  #define AVUTIL_ERROR_H

//  #include <errno.h>
//  #include <stddef.h>

//  #include "macros.h"

//  /**
//   * @addtogroup lavu_error
//   *
//   * @{
//   */

//	/* error handling */
//
// /< Returns a negative error code from a POSIX error code, to return from library functions.
func AVERROR(e int) int {
	return int(C.marco_AVERROR(C.int(e)))
}

//  #if EDOM > 0
//  #define AVERROR(e) (-(e))   ///< Returns a negative error code from a POSIX error code, to return from library functions.
//  #define AVUNERROR(e) (-(e)) ///< Returns a POSIX error code from a library function error return value.
//  #else
//  /* Some platforms have E* and errno already negated. */
//  #define AVERROR(e) (e)
//  #define AVUNERROR(e) (e)
//  #endif

//  #define FFERRTAG(a, b, c, d) (-(int)MKTAG(a, b, c, d))

const (
	AVERROR_BSF_NOT_FOUND      = C.AVERROR_BSF_NOT_FOUND
	AVERROR_BUG                = C.AVERROR_BUG
	AVERROR_BUFFER_TOO_SMALL   = C.AVERROR_BUFFER_TOO_SMALL
	AVERROR_DECODER_NOT_FOUND  = C.AVERROR_DECODER_NOT_FOUND
	AVERROR_DEMUXER_NOT_FOUND  = C.AVERROR_DEMUXER_NOT_FOUND
	AVERROR_ENCODER_NOT_FOUND  = C.AVERROR_ENCODER_NOT_FOUND
	AVERROR_EOF                = C.AVERROR_EOF
	AVERROR_EXIT               = C.AVERROR_EXIT
	AVERROR_EXTERNAL           = C.AVERROR_EXTERNAL
	AVERROR_FILTER_NOT_FOUND   = C.AVERROR_FILTER_NOT_FOUND
	AVERROR_INVALIDDATA        = C.AVERROR_INVALIDDATA
	AVERROR_MUXER_NOT_FOUND    = C.AVERROR_MUXER_NOT_FOUND
	AVERROR_OPTION_NOT_FOUND   = C.AVERROR_OPTION_NOT_FOUND
	AVERROR_PATCHWELCOME       = C.AVERROR_PATCHWELCOME
	AVERROR_PROTOCOL_NOT_FOUND = C.AVERROR_PROTOCOL_NOT_FOUND

	AVERROR_STREAM_NOT_FOUND = C.AVERROR_STREAM_NOT_FOUND
	/**
	 * This is semantically identical to AVERROR_BUG
	 * it has been introduced in Libav after our AVERROR_BUG and with a modified value.
	 */
	AVERROR_BUG2           = C.AVERROR_BUG2
	AVERROR_UNKNOWN        = C.AVERROR_UNKNOWN
	AVERROR_EXPERIMENTAL   = C.AVERROR_EXPERIMENTAL
	AVERROR_INPUT_CHANGED  = C.AVERROR_INPUT_CHANGED
	AVERROR_OUTPUT_CHANGED = C.AVERROR_OUTPUT_CHANGED
	/* HTTP & RTSP errors */
	AVERROR_HTTP_BAD_REQUEST  = C.AVERROR_HTTP_BAD_REQUEST
	AVERROR_HTTP_UNAUTHORIZED = C.AVERROR_HTTP_UNAUTHORIZED
	AVERROR_HTTP_FORBIDDEN    = C.AVERROR_HTTP_FORBIDDEN
	AVERROR_HTTP_NOT_FOUND    = C.AVERROR_HTTP_NOT_FOUND
	AVERROR_HTTP_OTHER_4XX    = C.AVERROR_HTTP_OTHER_4XX
	AVERROR_HTTP_SERVER_ERROR = C.AVERROR_HTTP_SERVER_ERROR
)

const AV_ERROR_MAX_STRING_SIZE = C.AV_ERROR_MAX_STRING_SIZE

//  /**
//   * Put a description of the AVERROR code errnum in errbuf.
//   * In case of failure the global variable errno is set to indicate the
//   * error. Even in case of failure av_strerror() will print a generic
//   * error message indicating the errnum provided to errbuf.
//   *
//   * @param errnum      error code to describe
//   * @param errbuf      buffer to which description is written
//   * @param errbuf_size the size in bytes of errbuf
//   * @return 0 on success, a negative value if a description for errnum
//   * cannot be found
//   */
//  int av_strerror(int errnum, char *errbuf, size_t errbuf_size);

//  /**
//   * Fill the provided buffer with a string containing an error string
//   * corresponding to the AVERROR code errnum.
//   *
//   * @param errbuf         a buffer
//   * @param errbuf_size    size in bytes of errbuf
//   * @param errnum         error code to describe
//   * @return the buffer in input, filled with the error description
//   * @see av_strerror()
//   */
//  static inline char *av_make_error_string(char *errbuf, size_t errbuf_size, int errnum)
//  {
// 	 av_strerror(errnum, errbuf, errbuf_size);
// 	 return errbuf;
//  }

/**
 * Convenience macro, the return value should be used only directly in
 * function arguments but never stand-alone.
 */
func AvErr2str(code int) string {
	cStr := C.marco_av_err2str(C.int(code))
	defer C.free(unsafe.Pointer(cStr))
	return C.GoString(cStr)
}

/**
 * @}
 */

//  #endif /* AVUTIL_ERROR_H */
