package swscale

/*
#cgo pkg-config: libswscale

#include "libswscale/version.h"
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

//  #ifndef SWSCALE_VERSION_H
//  #define SWSCALE_VERSION_H

/**
 * @file
 * swscale version macros
 */

//  #include "libavutil/version.h"

//  #include "version_major.h"

const LIBSWSCALE_VERSION_MINOR = C.LIBSWSCALE_VERSION_MINOR
const LIBSWSCALE_VERSION_MICRO = C.LIBSWSCALE_VERSION_MICRO

const LIBSWSCALE_VERSION_INT = C.LIBSWSCALE_VERSION_INT

// const LIBSWSCALE_VERSION = C.LIBSWSCALE_VERSION

const LIBSWSCALE_BUILD = C.LIBSWSCALE_BUILD

const LIBSWSCALE_IDENT = C.LIBSWSCALE_IDENT

//  #endif /* SWSCALE_VERSION_H */
