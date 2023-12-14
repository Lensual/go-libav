package swscale

/*
#cgo pkg-config: libswscale

#include "libswscale/swscale.h"
*/
import "C"
import (
	"unsafe"

	"github.com/Lensual/go-libav/avutil"
)

/*
 * Copyright (C) 2001-2011 Michael Niedermayer <michaelni@gmx.at>
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

//  #ifndef SWSCALE_SWSCALE_H
//  #define SWSCALE_SWSCALE_H

/**
 * @file
 * @ingroup libsws
 * external API header
 */

//  #include <stdint.h>

//  #include "libavutil/avutil.h"
//  #include "libavutil/frame.h"
//  #include "libavutil/log.h"
//  #include "libavutil/pixfmt.h"
//  #include "version_major.h"
//  #ifndef HAVE_AV_CONFIG_H
//  /* When included as part of the ffmpeg build, only include the major version
//   * to avoid unnecessary rebuilds. When included externally, keep including
//   * the full version information. */
//  #include "version.h"
//  #endif

/**
 * @defgroup libsws libswscale
 * Color conversion and scaling library.
 *
 * @{
 *
 * Return the LIBSWSCALE_VERSION_INT constant.
 */
func SwscaleVersion() uint {
	return uint(C.swscale_version())
}

/**
 * Return the libswscale build-time configuration.
 */
func SwscaleConfiguration() string {
	return C.GoString(C.swscale_configuration())
}

/**
 * Return the libswscale license.
 */
func SwscaleLicense() string {
	return C.GoString(C.swscale_license())
}

/* values for the flags, the stuff on the command line is different */
const (
	SWS_FAST_BILINEAR = C.SWS_FAST_BILINEAR
	SWS_BILINEAR      = C.SWS_BILINEAR
	SWS_BICUBIC       = C.SWS_BICUBIC
	SWS_X             = C.SWS_X
	SWS_POINT         = C.SWS_POINT
	SWS_AREA          = C.SWS_AREA
	SWS_BICUBLIN      = C.SWS_BICUBLIN
	SWS_GAUSS         = C.SWS_GAUSS
	SWS_SINC          = C.SWS_SINC
	SWS_LANCZOS       = C.SWS_LANCZOS
	SWS_SPLINE        = C.SWS_SPLINE
)

const (
	SWS_SRC_V_CHR_DROP_MASK  = C.SWS_SRC_V_CHR_DROP_MASK
	SWS_SRC_V_CHR_DROP_SHIFT = C.SWS_SRC_V_CHR_DROP_SHIFT
)

const SWS_PARAM_DEFAULT = C.SWS_PARAM_DEFAULT

const SWS_PRINT_INFO = C.SWS_PRINT_INFO

// the following 3 flags are not completely implemented
// internal chrominance subsampling info
const SWS_FULL_CHR_H_INT = C.SWS_FULL_CHR_H_INT

// input subsampling info
const (
	SWS_FULL_CHR_H_INP  = C.SWS_FULL_CHR_H_INP
	SWS_DIRECT_BGR      = C.SWS_DIRECT_BGR
	SWS_ACCURATE_RND    = C.SWS_ACCURATE_RND
	SWS_BITEXACT        = C.SWS_BITEXACT
	SWS_ERROR_DIFFUSION = C.SWS_ERROR_DIFFUSION
)

const SWS_MAX_REDUCE_CUTOFF = C.SWS_MAX_REDUCE_CUTOFF

const (
	SWS_CS_ITU709    = C.SWS_CS_ITU709
	SWS_CS_FCC       = C.SWS_CS_FCC
	SWS_CS_ITU601    = C.SWS_CS_ITU601
	SWS_CS_ITU624    = C.SWS_CS_ITU624
	SWS_CS_SMPTE170M = C.SWS_CS_SMPTE170M
	SWS_CS_SMPTE240M = C.SWS_CS_SMPTE240M
	SWS_CS_DEFAULT   = C.SWS_CS_DEFAULT
	SWS_CS_BT2020    = C.SWS_CS_BT2020
)

/**
 * Return a pointer to yuv<->rgb coefficients for the given colorspace
 * suitable for sws_setColorspaceDetails().
 *
 * @param colorspace One of the SWS_CS_* macros. If invalid,
 * SWS_CS_DEFAULT is used.
 */
func SwsGetCoefficients(colorspace int) unsafe.Pointer {
	return unsafe.Pointer(C.sws_getCoefficients(C.int(colorspace)))
}

// when used for filters they must have an odd number of elements
// coeffs cannot be shared between vectors
type CSwsVector C.SwsVector

//#region CSwsVector

// /< pointer to the list of coefficients
func (sv CSwsVector) GetCoeff() *float64 {
	return (*float64)(sv.coeff)
}

// /< number of coefficients in the vector
func (sv CSwsVector) GetLength() int {
	return int(sv.length)
}

//#endregion CSwsVector

// vectors can be shared
type CSwsFilter C.SwsFilter

//region CSwsFilter

func (sf CSwsFilter) GetLumH() *CSwsVector {
	return (*CSwsVector)(sf.lumH)
}
func (sf CSwsFilter) GetLumV() *CSwsVector {
	return (*CSwsVector)(sf.lumV)
}
func (sf CSwsFilter) GetChrH() *CSwsVector {
	return (*CSwsVector)(sf.chrH)
}
func (sf CSwsFilter) GetChrV() *CSwsVector {
	return (*CSwsVector)(sf.chrV)
}

//endregion CSwsFilter

type CSwsContext C.struct_SwsContext

/**
 * Return a positive value if pix_fmt is a supported input format, 0
 * otherwise.
 */
func SwsIsSupportedInput(pixFmt avutil.CAVPixelFormat) int {
	return int(C.sws_isSupportedInput(C.enum_AVPixelFormat(pixFmt)))
}

/**
 * Return a positive value if pix_fmt is a supported output format, 0
 * otherwise.
 */
func SwsIsSupportedOutput(pixFmt avutil.CAVPixelFormat) int {
	return int(C.sws_isSupportedOutput(C.enum_AVPixelFormat(pixFmt)))
}

/**
 * @param[in]  pix_fmt the pixel format
 * @return a positive value if an endianness conversion for pix_fmt is
 * supported, 0 otherwise.
 */
func SwsIsSupportedEndiannessConversion(pixFmt avutil.CAVPixelFormat) int {
	return int(C.sws_isSupportedEndiannessConversion(C.enum_AVPixelFormat(pixFmt)))
}

/**
 * Allocate an empty SwsContext. This must be filled and passed to
 * sws_init_context(). For filling see AVOptions, options.c and
 * sws_setColorspaceDetails().
 */
func SwsAllocContext() *CSwsContext {
	return (*CSwsContext)(C.sws_alloc_context())
}

/**
 * Initialize the swscaler context sws_context.
 *
 * @return zero or positive value on success, a negative value on
 * error
 */
func SwsInitContext(sws_context *CSwsContext, srcFilter *CSwsFilter, dstFilter *CSwsFilter) int {
	return int(C.sws_init_context((*C.struct_SwsContext)(sws_context), (*C.SwsFilter)(srcFilter), (*C.SwsFilter)(dstFilter)))
}

/**
 * Free the swscaler context swsContext.
 * If swsContext is NULL, then does nothing.
 */
func SwsFreeContext(swsContext *CSwsContext) {
	C.sws_freeContext((*C.struct_SwsContext)(swsContext))
}

/**
 * Allocate and return an SwsContext. You need it to perform
 * scaling/conversion operations using sws_scale().
 *
 * @param srcW the width of the source image
 * @param srcH the height of the source image
 * @param srcFormat the source image format
 * @param dstW the width of the destination image
 * @param dstH the height of the destination image
 * @param dstFormat the destination image format
 * @param flags specify which algorithm and options to use for rescaling
 * @param param extra parameters to tune the used scaler
 *              For SWS_BICUBIC param[0] and [1] tune the shape of the basis
 *              function, param[0] tunes f(1) and param[1] f´(1)
 *              For SWS_GAUSS param[0] tunes the exponent and thus cutoff
 *              frequency
 *              For SWS_LANCZOS param[0] tunes the width of the window function
 * @return a pointer to an allocated context, or NULL in case of error
 * @note this function is to be removed after a saner alternative is
 *       written
 */
func SwsGetContext(srcW int, srcH int, srcFormat avutil.CAVPixelFormat,
	dstW int, dstH int, dstFormat avutil.CAVPixelFormat,
	flags int, srcFilter *CSwsFilter,
	dstFilter *CSwsFilter, param *float64) *CSwsContext {
	return (*CSwsContext)(C.sws_getContext(C.int(srcW), C.int(srcH), (C.enum_AVPixelFormat)(srcFormat), C.int(dstW), C.int(dstH), (C.enum_AVPixelFormat)(dstFormat), C.int(flags), (*C.SwsFilter)(srcFilter), (*C.SwsFilter)(dstFilter), (*C.double)(param)))
}

/**
 * Scale the image slice in srcSlice and put the resulting scaled
 * slice in the image in dst. A slice is a sequence of consecutive
 * rows in an image.
 *
 * Slices have to be provided in sequential order, either in
 * top-bottom or bottom-top order. If slices are provided in
 * non-sequential order the behavior of the function is undefined.
 *
 * @param c         the scaling context previously created with
 *                  sws_getContext()
 * @param srcSlice  the array containing the pointers to the planes of
 *                  the source slice
 * @param srcStride the array containing the strides for each plane of
 *                  the source image
 * @param srcSliceY the position in the source image of the slice to
 *                  process, that is the number (counted starting from
 *                  zero) in the image of the first row of the slice
 * @param srcSliceH the height of the source slice, that is the number
 *                  of rows in the slice
 * @param dst       the array containing the pointers to the planes of
 *                  the destination image
 * @param dstStride the array containing the strides for each plane of
 *                  the destination image
 * @return          the height of the output slice
 */
func SwsScale(c *CSwsContext, srcSlice []*C.uint8_t, srcStride *C.int, srcSliceY int, srcSliceH int, dst **C.uint8_t, dstStride *C.int) int {
	return int(C.sws_scale((*C.struct_SwsContext)(c), &srcSlice[0], srcStride, C.int(srcSliceY), C.int(srcSliceH), dst, dstStride))
}

/**
 * Scale source data from src and write the output to dst.
 *
 * This is merely a convenience wrapper around
 * - sws_frame_start()
 * - sws_send_slice(0, src->height)
 * - sws_receive_slice(0, dst->height)
 * - sws_frame_end()
 *
 * @param c   The scaling context
 * @param dst The destination frame. See documentation for sws_frame_start() for
 *            more details.
 * @param src The source frame.
 *
 * @return 0 on success, a negative AVERROR code on failure
 */
func SwsScaleFrame(c *CSwsContext, dst *avutil.CAVFrame, src *avutil.CAVFrame) int {
	return int(C.sws_scale_frame((*C.struct_SwsContext)(c), (*C.AVFrame)(unsafe.Pointer(dst)), (*C.AVFrame)(unsafe.Pointer(src))))
}

/**
 * Initialize the scaling process for a given pair of source/destination frames.
 * Must be called before any calls to sws_send_slice() and sws_receive_slice().
 *
 * This function will retain references to src and dst, so they must both use
 * refcounted buffers (if allocated by the caller, in case of dst).
 *
 * @param c   The scaling context
 * @param dst The destination frame.
 *
 *            The data buffers may either be already allocated by the caller or
 *            left clear, in which case they will be allocated by the scaler.
 *            The latter may have performance advantages - e.g. in certain cases
 *            some output planes may be references to input planes, rather than
 *            copies.
 *
 *            Output data will be written into this frame in successful
 *            sws_receive_slice() calls.
 * @param src The source frame. The data buffers must be allocated, but the
 *            frame data does not have to be ready at this point. Data
 *            availability is then signalled by sws_send_slice().
 * @return 0 on success, a negative AVERROR code on failure
 *
 * @see sws_frame_end()
 */
func SwsFrameStart(c *CSwsContext, dst *avutil.CAVFrame, src *avutil.CAVFrame) int {
	return int(C.sws_frame_start((*C.struct_SwsContext)(c), (*C.AVFrame)(unsafe.Pointer(dst)), (*C.AVFrame)(unsafe.Pointer(src))))
}

/**
 * Finish the scaling process for a pair of source/destination frames previously
 * submitted with sws_frame_start(). Must be called after all sws_send_slice()
 * and sws_receive_slice() calls are done, before any new sws_frame_start()
 * calls.
 *
 * @param c   The scaling context
 */
func SwsFrameEnd(c *CSwsContext) {
	C.sws_frame_end((*C.struct_SwsContext)(c))
}

/**
 * Indicate that a horizontal slice of input data is available in the source
 * frame previously provided to sws_frame_start(). The slices may be provided in
 * any order, but may not overlap. For vertically subsampled pixel formats, the
 * slices must be aligned according to subsampling.
 *
 * @param c   The scaling context
 * @param slice_start first row of the slice
 * @param slice_height number of rows in the slice
 *
 * @return a non-negative number on success, a negative AVERROR code on failure.
 */
func SwsSendSlice(c *CSwsContext, sliceStart uint, sliceHeight uint) int {
	return int(C.sws_send_slice((*C.struct_SwsContext)(c), C.uint(sliceStart), C.uint(sliceHeight)))
}

/**
 * Request a horizontal slice of the output data to be written into the frame
 * previously provided to sws_frame_start().
 *
 * @param c   The scaling context
 * @param slice_start first row of the slice; must be a multiple of
 *                    sws_receive_slice_alignment()
 * @param slice_height number of rows in the slice; must be a multiple of
 *                     sws_receive_slice_alignment(), except for the last slice
 *                     (i.e. when slice_start+slice_height is equal to output
 *                     frame height)
 *
 * @return a non-negative number if the data was successfully written into the output
 *         AVERROR(EAGAIN) if more input data needs to be provided before the
 *                         output can be produced
 *         another negative AVERROR code on other kinds of scaling failure
 */
func SwsReceiveSlice(c *CSwsContext, sliceStart uint, sliceHeight uint) int {
	return int(C.sws_receive_slice((*C.struct_SwsContext)(c), C.uint(sliceStart), C.uint(sliceHeight)))
}

/**
 * Get the alignment required for slices
 *
 * @param c   The scaling context
 * @return alignment required for output slices requested with sws_receive_slice().
 *         Slice offsets and sizes passed to sws_receive_slice() must be
 *         multiples of the value returned from this function.
 */
func SwsReceiveSliceAlignment(c *CSwsContext) uint {
	return uint(C.sws_receive_slice_alignment((*C.struct_SwsContext)(c)))
}

/**
 * @param c the scaling context
 * @param dstRange flag indicating the while-black range of the output (1=jpeg / 0=mpeg)
 * @param srcRange flag indicating the while-black range of the input (1=jpeg / 0=mpeg)
 * @param table the yuv2rgb coefficients describing the output yuv space, normally ff_yuv2rgb_coeffs[x]
 * @param inv_table the yuv2rgb coefficients describing the input yuv space, normally ff_yuv2rgb_coeffs[x]
 * @param brightness 16.16 fixed point brightness correction
 * @param contrast 16.16 fixed point contrast correction
 * @param saturation 16.16 fixed point saturation correction
 *
 * @return A negative error code on error, non negative otherwise.
 *         If `LIBSWSCALE_VERSION_MAJOR < 7`, returns -1 if not supported.
 */
func SwsSetColorspaceDetails(c *CSwsContext, invTable [4]int, srcRange int, table [4]int, dstRange int, brightness int, contrast int, saturation int) int {
	return int(C.sws_setColorspaceDetails(
		(*C.struct_SwsContext)(c),
		(*C.int)(unsafe.Pointer((unsafe.SliceData(invTable[:])))),
		C.int(srcRange),
		(*C.int)(unsafe.Pointer((unsafe.SliceData(table[:])))),
		C.int(dstRange),
		C.int(brightness),
		C.int(contrast),
		C.int(saturation),
	))
}

/**
 * @return A negative error code on error, non negative otherwise.
 *         If `LIBSWSCALE_VERSION_MAJOR < 7`, returns -1 if not supported.
 */
func SwsGetColorspaceDetails(c *CSwsContext, invTable **int,
	srcRange *int, table **int, dstRange *int,
	brightness *int, contrast *int, saturation *int) int {
	return int(C.sws_getColorspaceDetails(
		(*C.struct_SwsContext)(c),
		(**C.int)(unsafe.Pointer(invTable)),
		(*C.int)(unsafe.Pointer(srcRange)),
		(**C.int)(unsafe.Pointer(table)),
		(*C.int)(unsafe.Pointer(dstRange)),
		(*C.int)(unsafe.Pointer(brightness)),
		(*C.int)(unsafe.Pointer(contrast)),
		(*C.int)(unsafe.Pointer(saturation)),
	))
}

/**
 * Allocate and return an uninitialized vector with length coefficients.
 */
func SwsAllocVec(length int) *CSwsVector {
	return (*CSwsVector)(C.sws_allocVec(C.int(length)))
}

/**
 * Return a normalized Gaussian curve used to filter stuff
 * quality = 3 is high quality, lower is lower quality.
 */
func SwsGetGaussianVec(variance float64, quality float64) *CSwsVector {
	return (*CSwsVector)(C.sws_getGaussianVec(C.double(variance), C.double(quality)))
}

/**
 * Scale all the coefficients of a by the scalar value.
 */
func SwsScaleVec(a *CSwsVector, scalar float64) {
	C.sws_scaleVec((*C.SwsVector)(a), C.double(scalar))
}

/**
 * Scale all the coefficients of a so that their sum equals height.
 */
func SwsNormalizeVec(a *CSwsVector, height float64) {
	C.sws_normalizeVec((*C.SwsVector)(a), C.double(height))
}

func SwsFreeVec(a *CSwsVector) {
	C.sws_freeVec((*C.SwsVector)(a))
}

func SwsGetDefaultFilter(lumaGBlur float32, chromaGBlur float32,
	lumaSharpen float32, chromaSharpen float32,
	chromaHShift float32, chromaVShift float32,
	verbose int) *CSwsFilter {
	return (*CSwsFilter)(C.sws_getDefaultFilter(
		C.float(lumaGBlur),
		C.float(chromaGBlur),
		C.float(lumaSharpen),
		C.float(chromaSharpen),
		C.float(chromaHShift),
		C.float(chromaVShift),
		C.int(verbose),
	))
}

func SwsFreeFilter(filter *CSwsFilter) {
	C.sws_freeFilter((*C.SwsFilter)(filter))
}

/**
 * Check if context can be reused, otherwise reallocate a new one.
 *
 * If context is NULL, just calls sws_getContext() to get a new
 * context. Otherwise, checks if the parameters are the ones already
 * saved in context. If that is the case, returns the current
 * context. Otherwise, frees context and gets a new context with
 * the new parameters.
 *
 * Be warned that srcFilter and dstFilter are not checked, they
 * are assumed to remain the same.
 */
func SwsGetCachedContext(context *CSwsContext,
	srcW int, srcH int, srcFormat avutil.CAVPixelFormat,
	dstW int, dstH int, dstFormat avutil.CAVPixelFormat,
	flags int, srcFilter *CSwsFilter,
	dstFilter *CSwsFilter, param *float64) *CSwsContext {
	return (*CSwsContext)(C.sws_getCachedContext(
		(*C.struct_SwsContext)(context),
		C.int(srcW),
		C.int(srcH),
		(C.enum_AVPixelFormat)(srcFormat),
		C.int(dstW),
		C.int(dstH),
		(C.enum_AVPixelFormat)(dstFormat),
		C.int(flags),
		(*C.SwsFilter)(srcFilter),
		(*C.SwsFilter)(dstFilter),
		(*C.double)(param),
	))
}

/**
 * Convert an 8-bit paletted frame into a frame with a color depth of 32 bits.
 *
 * The output frame will have the same packed format as the palette.
 *
 * @param src        source frame buffer
 * @param dst        destination frame buffer
 * @param num_pixels number of pixels to convert
 * @param palette    array with [256] entries, which must match color arrangement (RGB or BGR) of src
 */
func SwsConvertPalette8ToPacked32(src unsafe.Pointer, dst unsafe.Pointer, numPixels int, palette *uint8) {
	C.sws_convertPalette8ToPacked32((*C.uint8_t)(src), (*C.uint8_t)(dst), C.int(numPixels), (*C.uint8_t)(palette))
}

/**
 * Convert an 8-bit paletted frame into a frame with a color depth of 24 bits.
 *
 * With the palette format "ABCD", the destination frame ends up with the format "ABC".
 *
 * @param src        source frame buffer
 * @param dst        destination frame buffer
 * @param num_pixels number of pixels to convert
 * @param palette    array with [256] entries, which must match color arrangement (RGB or BGR) of src
 */
func SwsConvertPalette8ToPacked24(src unsafe.Pointer, dst unsafe.Pointer, numPixels int, palette *uint8) {
	C.sws_convertPalette8ToPacked24((*C.uint8_t)(src), (*C.uint8_t)(dst), C.int(numPixels), (*C.uint8_t)(palette))
}

/**
 * Get the AVClass for swsContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
func SwsGetClass() *avutil.CAVClass {
	return (*avutil.CAVClass)(unsafe.Pointer(C.sws_get_class()))
}

/**
 * @}
 */

//  #endif /* SWSCALE_SWSCALE_H */
