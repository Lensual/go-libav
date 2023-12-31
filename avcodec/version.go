package avcodec

/*
#cgo pkg-config: libavcodec

#include "libavcodec/version.h"
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

//  #ifndef AVCODEC_VERSION_H
//  #define AVCODEC_VERSION_H

/**
 * @file
 * @ingroup libavc
 * Libavcodec version macros.
 */

//  #include "libavutil/version.h"

//  #include "version_major.h"

const LIBAVCODEC_VERSION_MINOR = C.LIBAVCODEC_VERSION_MINOR
const LIBAVCODEC_VERSION_MICRO = C.LIBAVCODEC_VERSION_MICRO

const LIBAVCODEC_VERSION_INT = C.LIBAVCODEC_VERSION_INT

//const LIBAVCODEC_VERSION = C.LIBAVCODEC_VERSION

const LIBAVCODEC_BUILD = C.LIBAVCODEC_BUILD

const LIBAVCODEC_IDENT = C.LIBAVCODEC_IDENT

//  #endif /* AVCODEC_VERSION_H */
