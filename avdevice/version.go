package avdevice

/*
#cgo pkg-config: libavdevice

#include "libavdevice/avdevice.h"
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

//  #ifndef AVDEVICE_VERSION_H
//  #define AVDEVICE_VERSION_H

/**
 * @file
 * @ingroup lavd
 * Libavdevice version macros
 */

//  #include "libavutil/version.h"

//  #include "version_major.h"

const LIBAVDEVICE_VERSION_MINOR = C.LIBAVDEVICE_VERSION_MINOR
const LIBAVDEVICE_VERSION_MICRO = C.LIBAVDEVICE_VERSION_MICRO

const LIBAVDEVICE_VERSION_INT = C.LIBAVDEVICE_VERSION_INT

//const LIBAVDEVICE_VERSION = C.LIBAVDEVICE_VERSION

const LIBAVDEVICE_BUILD = C.LIBAVDEVICE_BUILD

const LIBAVDEVICE_IDENT = C.LIBAVDEVICE_IDENT

//  #endif /* AVDEVICE_VERSION_H */
