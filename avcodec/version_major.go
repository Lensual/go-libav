package avcodec

/*
#cgo pkg-config: libavcodec

#include "libavcodec/version_major.h"
*/
import "C"

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

//  #ifndef AVCODEC_VERSION_MAJOR_H
//  AVCODEC_VERSION_MAJOR_H

/**
 * @file
 * @ingroup libavc
 * Libavcodec version macros.
 */

const LIBAVCODEC_VERSION_MAJOR = C.LIBAVCODEC_VERSION_MAJOR

/**
 * FF_API_* defines may be placed below to indicate public API that will be
 * dropped at a future version bump. The defines themselves are not part of
 * the public API and may change, break or disappear at any time.
 *
 * @note, when bumping the major version it is recommended to manually
 * disable each FF_API_* in its own commit instead of disabling them all
 * at once through the bump. This improves the git bisect-ability of the change.
 */

const (
	FF_API_INIT_PACKET        = C.FF_API_INIT_PACKET
	FF_API_IDCT_NONE          = C.FF_API_IDCT_NONE
	FF_API_SVTAV1_OPTS        = C.FF_API_SVTAV1_OPTS
	FF_API_AYUV_CODECID       = C.FF_API_AYUV_CODECID
	FF_API_VT_OUTPUT_CALLBACK = C.FF_API_VT_OUTPUT_CALLBACK
	FF_API_AVCODEC_CHROMA_POS = C.FF_API_AVCODEC_CHROMA_POS
	FF_API_VT_HWACCEL_CONTEXT = C.FF_API_VT_HWACCEL_CONTEXT
	FF_API_AVCTX_FRAME_NUMBER = C.FF_API_AVCTX_FRAME_NUMBER
	FF_API_SLICE_OFFSET       = C.FF_API_SLICE_OFFSET
	FF_API_SUBFRAMES          = C.FF_API_SUBFRAMES
	FF_API_TICKS_PER_FRAME    = C.FF_API_TICKS_PER_FRAME
	FF_API_DROPCHANGED        = C.FF_API_DROPCHANGED

	FF_API_AVFFT            = C.FF_API_AVFFT
	FF_API_FF_PROFILE_LEVEL = C.FF_API_FF_PROFILE_LEVEL

	// reminder to remove CrystalHD decoders on next major bump
	FF_CODEC_CRYSTAL_HD = C.FF_CODEC_CRYSTAL_HD
)

//  #endif /* AVCODEC_VERSION_MAJOR_H */
