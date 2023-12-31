package avformat

/*
#cgo pkg-config: libavformat

#include "libavformat/version_major.h"
*/
import "C"

/*
 * Version macros.
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

//  #ifndef AVFORMAT_VERSION_MAJOR_H
//  #define AVFORMAT_VERSION_MAJOR_H

/**
 * @file
 * @ingroup libavf
 * Libavformat version macros
 */

// Major bumping may affect Ticket5467, 5421, 5451(compatibility with Chromium)
// Also please add any ticket numbers that you believe might be affected here
const LIBAVFORMAT_VERSION_MAJOR = C.LIBAVFORMAT_VERSION_MAJOR

/**
 * FF_API_* defines may be placed below to indicate public API that will be
 * dropped at a future version bump. The defines themselves are not part of
 * the public API and may change, break or disappear at any time.
 *
 * @note, when bumping the major version it is recommended to manually
 * disable each FF_API_* in its own commit instead of disabling them all
 * at once through the bump. This improves the git bisect-ability of the change.
 *
 */
const (
	FF_API_COMPUTE_PKT_FIELDS2 = C.FF_API_COMPUTE_PKT_FIELDS2
	FF_API_GET_END_PTS         = C.FF_API_GET_END_PTS
	FF_API_AVIODIRCONTEXT      = C.FF_API_AVIODIRCONTEXT
	FF_API_AVFORMAT_IO_CLOSE   = C.FF_API_AVFORMAT_IO_CLOSE
	FF_API_AVIO_WRITE_NONCONST = C.FF_API_AVIO_WRITE_NONCONST
	FF_API_LAVF_SHORTEST       = C.FF_API_LAVF_SHORTEST
	FF_API_ALLOW_FLUSH         = C.FF_API_ALLOW_FLUSH
	FF_API_AVSTREAM_SIDE_DATA  = C.FF_API_AVSTREAM_SIDE_DATA
)

const FF_API_R_FRAME_RATE = C.FF_API_R_FRAME_RATE

//  #endif /* AVFORMAT_VERSION_MAJOR_H */
