package avformat

/*
#cgo pkg-config: libavformat

#include "libavformat/avformat.h"
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

//  #ifndef AVFORMAT_VERSION_H
//  #define AVFORMAT_VERSION_H

/**
 * @file
 * @ingroup libavf
 * Libavformat version macros
 */

//  #include "libavutil/version.h"

//  #include "version_major.h"

const LIBAVFORMAT_VERSION_MINOR = C.LIBAVFORMAT_VERSION_MINOR
const LIBAVFORMAT_VERSION_MICRO = C.LIBAVFORMAT_VERSION_MICRO

const LIBAVFORMAT_VERSION_INT = C.LIBAVFORMAT_VERSION_INT

// const LIBAVFORMAT_VERSION = C.LIBAVFORMAT_VERSION

const LIBAVFORMAT_BUILD = C.LIBAVFORMAT_BUILD

const LIBAVFORMAT_IDENT = C.LIBAVFORMAT_IDENT

//  #endif /* AVFORMAT_VERSION_H */
