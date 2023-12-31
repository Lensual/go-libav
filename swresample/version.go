package swresample

/*
#cgo pkg-config: libswresample

#include "libswresample/version.h"
*/
import "C"

/*
 * Version macros.
 *
 * This file is part of libswresample
 *
 * libswresample is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * libswresample is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with libswresample; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

//  #ifndef SWRESAMPLE_VERSION_H
//  #define SWRESAMPLE_VERSION_H

/**
 * @file
 * Libswresample version macros
 */

//  #include "libavutil/version.h"

//  #include "version_major.h"

const LIBSWRESAMPLE_VERSION_MINOR = C.LIBSWRESAMPLE_VERSION_MINOR
const LIBSWRESAMPLE_VERSION_MICRO = C.LIBSWRESAMPLE_VERSION_MICRO

const LIBSWRESAMPLE_VERSION_INT = C.LIBSWRESAMPLE_VERSION_INT

// const LIBSWRESAMPLE_VERSION = C.LIBSWRESAMPLE_VERSION

const LIBSWRESAMPLE_BUILD = C.LIBSWRESAMPLE_BUILD

const LIBSWRESAMPLE_IDENT = C.LIBSWRESAMPLE_IDENT

//  #endif /* SWRESAMPLE_VERSION_H */
