package avutil

/*
#cgo pkg-config: libavutil

#include "libavutil/pixfmt.h"
*/
import "C"

const AVPALETTE_SIZE = C.AVPALETTE_SIZE
const AVPALETTE_COUNT = C.AVPALETTE_COUNT

/**
 * Pixel format.
 *
 * @note
 * AV_PIX_FMT_RGB32 is handled in an endian-specific manner. An RGBA
 * color is put together as:
 *  (A << 24) | (R << 16) | (G << 8) | B
 * This is stored as BGRA on little-endian CPU architectures and ARGB on
 * big-endian CPUs.
 *
 * @note
 * If the resolution is not a multiple of the chroma subsampling factor
 * then the chroma plane resolution must be rounded up.
 *
 * @par
 * When the pixel format is palettized RGB32 (AV_PIX_FMT_PAL8), the palettized
 * image data is stored in AVFrame.data[0]. The palette is transported in
 * AVFrame.data[1], is 1024 bytes long (256 4-byte entries) and is
 * formatted the same as in AV_PIX_FMT_RGB32 described above (i.e., it is
 * also endian-specific). Note also that the individual RGB32 palette
 * components stored in AVFrame.data[1] should be in the range 0..255.
 * This is important as many custom PAL8 video codecs that were designed
 * to run on the IBM VGA graphics adapter use 6-bit palette components.
 *
 * @par
 * For all the 8 bits per pixel formats, an RGB32 palette is in data[1] like
 * for pal8. This palette is filled in automatically by the function
 * allocating the picture.
 */
type CAVPixelFormat C.enum_AVPixelFormat

const (
	AV_PIX_FMT_NONE      CAVPixelFormat = C.AV_PIX_FMT_NONE
	AV_PIX_FMT_YUV420P   CAVPixelFormat = C.AV_PIX_FMT_YUV420P   ///< planar YUV 4:2:0, 12bpp, (1 Cr & Cb sample per 2x2 Y samples)
	AV_PIX_FMT_YUYV422   CAVPixelFormat = C.AV_PIX_FMT_YUYV422   ///< packed YUV 4:2:2, 16bpp, Y0 Cb Y1 Cr
	AV_PIX_FMT_RGB24     CAVPixelFormat = C.AV_PIX_FMT_RGB24     ///< packed RGB 8:8:8, 24bpp, RGBRGB...
	AV_PIX_FMT_BGR24     CAVPixelFormat = C.AV_PIX_FMT_BGR24     ///< packed RGB 8:8:8, 24bpp, BGRBGR...
	AV_PIX_FMT_YUV422P   CAVPixelFormat = C.AV_PIX_FMT_YUV422P   ///< planar YUV 4:2:2, 16bpp, (1 Cr & Cb sample per 2x1 Y samples)
	AV_PIX_FMT_YUV444P   CAVPixelFormat = C.AV_PIX_FMT_YUV444P   ///< planar YUV 4:4:4, 24bpp, (1 Cr & Cb sample per 1x1 Y samples)
	AV_PIX_FMT_YUV410P   CAVPixelFormat = C.AV_PIX_FMT_YUV410P   ///< planar YUV 4:1:0,  9bpp, (1 Cr & Cb sample per 4x4 Y samples)
	AV_PIX_FMT_YUV411P   CAVPixelFormat = C.AV_PIX_FMT_YUV411P   ///< planar YUV 4:1:1, 12bpp, (1 Cr & Cb sample per 4x1 Y samples)
	AV_PIX_FMT_GRAY8     CAVPixelFormat = C.AV_PIX_FMT_GRAY8     ///<        Y        ,  8bpp
	AV_PIX_FMT_MONOWHITE CAVPixelFormat = C.AV_PIX_FMT_MONOWHITE ///<        Y        ,  1bpp, 0 is white, 1 is black, in each byte pixels are ordered from the msb to the lsb
	AV_PIX_FMT_MONOBLACK CAVPixelFormat = C.AV_PIX_FMT_MONOBLACK ///<        Y        ,  1bpp, 0 is black, 1 is white, in each byte pixels are ordered from the msb to the lsb
	AV_PIX_FMT_PAL8      CAVPixelFormat = C.AV_PIX_FMT_PAL8      ///< 8 bits with AV_PIX_FMT_RGB32 palette
	AV_PIX_FMT_YUVJ420P  CAVPixelFormat = C.AV_PIX_FMT_YUVJ420P  ///< planar YUV 4:2:0, 12bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV420P and setting color_range
	AV_PIX_FMT_YUVJ422P  CAVPixelFormat = C.AV_PIX_FMT_YUVJ422P  ///< planar YUV 4:2:2, 16bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV422P and setting color_range
	AV_PIX_FMT_YUVJ444P  CAVPixelFormat = C.AV_PIX_FMT_YUVJ444P  ///< planar YUV 4:4:4, 24bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV444P and setting color_range
	AV_PIX_FMT_UYVY422   CAVPixelFormat = C.AV_PIX_FMT_UYVY422   ///< packed YUV 4:2:2, 16bpp, Cb Y0 Cr Y1
	AV_PIX_FMT_UYYVYY411 CAVPixelFormat = C.AV_PIX_FMT_UYYVYY411 ///< packed YUV 4:1:1, 12bpp, Cb Y0 Y1 Cr Y2 Y3
	AV_PIX_FMT_BGR8      CAVPixelFormat = C.AV_PIX_FMT_BGR8      ///< packed RGB 3:3:2,  8bpp, (msb)2B 3G 3R(lsb)
	AV_PIX_FMT_BGR4      CAVPixelFormat = C.AV_PIX_FMT_BGR4      ///< packed RGB 1:2:1 bitstream,  4bpp, (msb)1B 2G 1R(lsb), a byte contains two pixels, the first pixel in the byte is the one composed by the 4 msb bits
	AV_PIX_FMT_BGR4_BYTE CAVPixelFormat = C.AV_PIX_FMT_BGR4_BYTE ///< packed RGB 1:2:1,  8bpp, (msb)1B 2G 1R(lsb)
	AV_PIX_FMT_RGB8      CAVPixelFormat = C.AV_PIX_FMT_RGB8      ///< packed RGB 3:3:2,  8bpp, (msb)2R 3G 3B(lsb)
	AV_PIX_FMT_RGB4      CAVPixelFormat = C.AV_PIX_FMT_RGB4      ///< packed RGB 1:2:1 bitstream,  4bpp, (msb)1R 2G 1B(lsb), a byte contains two pixels, the first pixel in the byte is the one composed by the 4 msb bits
	AV_PIX_FMT_RGB4_BYTE CAVPixelFormat = C.AV_PIX_FMT_RGB4_BYTE ///< packed RGB 1:2:1,  8bpp, (msb)1R 2G 1B(lsb)
	AV_PIX_FMT_NV12      CAVPixelFormat = C.AV_PIX_FMT_NV12      ///< planar YUV 4:2:0, 12bpp, 1 plane for Y and 1 plane for the UV components, which are interleaved (first byte U and the following byte V)
	AV_PIX_FMT_NV21      CAVPixelFormat = C.AV_PIX_FMT_NV21      ///< as above, but U and V bytes are swapped

	AV_PIX_FMT_ARGB CAVPixelFormat = C.AV_PIX_FMT_ARGB ///< packed ARGB 8:8:8:8, 32bpp, ARGBARGB...
	AV_PIX_FMT_RGBA CAVPixelFormat = C.AV_PIX_FMT_RGBA ///< packed RGBA 8:8:8:8, 32bpp, RGBARGBA...
	AV_PIX_FMT_ABGR CAVPixelFormat = C.AV_PIX_FMT_ABGR ///< packed ABGR 8:8:8:8, 32bpp, ABGRABGR...
	AV_PIX_FMT_BGRA CAVPixelFormat = C.AV_PIX_FMT_BGRA ///< packed BGRA 8:8:8:8, 32bpp, BGRABGRA...

	AV_PIX_FMT_GRAY16BE CAVPixelFormat = C.AV_PIX_FMT_GRAY16BE ///<        Y        , 16bpp, big-endian
	AV_PIX_FMT_GRAY16LE CAVPixelFormat = C.AV_PIX_FMT_GRAY16LE ///<        Y        , 16bpp, little-endian
	AV_PIX_FMT_YUV440P  CAVPixelFormat = C.AV_PIX_FMT_YUV440P  ///< planar YUV 4:4:0 (1 Cr & Cb sample per 1x2 Y samples)
	AV_PIX_FMT_YUVJ440P CAVPixelFormat = C.AV_PIX_FMT_YUVJ440P ///< planar YUV 4:4:0 full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV440P and setting color_range
	AV_PIX_FMT_YUVA420P CAVPixelFormat = C.AV_PIX_FMT_YUVA420P ///< planar YUV 4:2:0, 20bpp, (1 Cr & Cb sample per 2x2 Y & A samples)
	AV_PIX_FMT_RGB48BE  CAVPixelFormat = C.AV_PIX_FMT_RGB48BE  ///< packed RGB 16:16:16, 48bpp, 16R, 16G, 16B, the 2-byte value for each R/G/B component is stored as big-endian
	AV_PIX_FMT_RGB48LE  CAVPixelFormat = C.AV_PIX_FMT_RGB48LE  ///< packed RGB 16:16:16, 48bpp, 16R, 16G, 16B, the 2-byte value for each R/G/B component is stored as little-endian

	AV_PIX_FMT_RGB565BE CAVPixelFormat = C.AV_PIX_FMT_RGB565BE ///< packed RGB 5:6:5, 16bpp, (msb)   5R 6G 5B(lsb), big-endian
	AV_PIX_FMT_RGB565LE CAVPixelFormat = C.AV_PIX_FMT_RGB565LE ///< packed RGB 5:6:5, 16bpp, (msb)   5R 6G 5B(lsb), little-endian
	AV_PIX_FMT_RGB555BE CAVPixelFormat = C.AV_PIX_FMT_RGB555BE ///< packed RGB 5:5:5, 16bpp, (msb)1X 5R 5G 5B(lsb), big-endian   , X=unused/undefined
	AV_PIX_FMT_RGB555LE CAVPixelFormat = C.AV_PIX_FMT_RGB555LE ///< packed RGB 5:5:5, 16bpp, (msb)1X 5R 5G 5B(lsb), little-endian, X=unused/undefined

	AV_PIX_FMT_BGR565BE CAVPixelFormat = C.AV_PIX_FMT_BGR565BE ///< packed BGR 5:6:5, 16bpp, (msb)   5B 6G 5R(lsb), big-endian
	AV_PIX_FMT_BGR565LE CAVPixelFormat = C.AV_PIX_FMT_BGR565LE ///< packed BGR 5:6:5, 16bpp, (msb)   5B 6G 5R(lsb), little-endian
	AV_PIX_FMT_BGR555BE CAVPixelFormat = C.AV_PIX_FMT_BGR555BE ///< packed BGR 5:5:5, 16bpp, (msb)1X 5B 5G 5R(lsb), big-endian   , X=unused/undefined
	AV_PIX_FMT_BGR555LE CAVPixelFormat = C.AV_PIX_FMT_BGR555LE ///< packed BGR 5:5:5, 16bpp, (msb)1X 5B 5G 5R(lsb), little-endian, X=unused/undefined

	/**
	 *  Hardware acceleration through VA-API, data[3] contains a
	 *  VASurfaceID.
	 */
	AV_PIX_FMT_VAAPI CAVPixelFormat = C.AV_PIX_FMT_VAAPI

	AV_PIX_FMT_YUV420P16LE CAVPixelFormat = C.AV_PIX_FMT_YUV420P16LE ///< planar YUV 4:2:0, 24bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV420P16BE CAVPixelFormat = C.AV_PIX_FMT_YUV420P16BE ///< planar YUV 4:2:0, 24bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV422P16LE CAVPixelFormat = C.AV_PIX_FMT_YUV422P16LE ///< planar YUV 4:2:2, 32bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_YUV422P16BE CAVPixelFormat = C.AV_PIX_FMT_YUV422P16BE ///< planar YUV 4:2:2, 32bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P16LE CAVPixelFormat = C.AV_PIX_FMT_YUV444P16LE ///< planar YUV 4:4:4, 48bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P16BE CAVPixelFormat = C.AV_PIX_FMT_YUV444P16BE ///< planar YUV 4:4:4, 48bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_DXVA2_VLD   CAVPixelFormat = C.AV_PIX_FMT_DXVA2_VLD   ///< HW decoding through DXVA2, Picture.data[3] contains a LPDIRECT3DSURFACE9 pointer

	AV_PIX_FMT_RGB444LE CAVPixelFormat = C.AV_PIX_FMT_RGB444LE ///< packed RGB 4:4:4, 16bpp, (msb)4X 4R 4G 4B(lsb), little-endian, X=unused/undefined
	AV_PIX_FMT_RGB444BE CAVPixelFormat = C.AV_PIX_FMT_RGB444BE ///< packed RGB 4:4:4, 16bpp, (msb)4X 4R 4G 4B(lsb), big-endian,    X=unused/undefined
	AV_PIX_FMT_BGR444LE CAVPixelFormat = C.AV_PIX_FMT_BGR444LE ///< packed BGR 4:4:4, 16bpp, (msb)4X 4B 4G 4R(lsb), little-endian, X=unused/undefined
	AV_PIX_FMT_BGR444BE CAVPixelFormat = C.AV_PIX_FMT_BGR444BE ///< packed BGR 4:4:4, 16bpp, (msb)4X 4B 4G 4R(lsb), big-endian,    X=unused/undefined
	AV_PIX_FMT_YA8      CAVPixelFormat = C.AV_PIX_FMT_YA8      ///< 8 bits gray, 8 bits alpha

	AV_PIX_FMT_Y400A  CAVPixelFormat = C.AV_PIX_FMT_Y400A  ///< alias for AV_PIX_FMT_YA8
	AV_PIX_FMT_GRAY8A CAVPixelFormat = C.AV_PIX_FMT_GRAY8A ///< alias for AV_PIX_FMT_YA8

	AV_PIX_FMT_BGR48BE CAVPixelFormat = C.AV_PIX_FMT_BGR48BE ///< packed RGB 16:16:16, 48bpp, 16B, 16G, 16R, the 2-byte value for each R/G/B component is stored as big-endian
	AV_PIX_FMT_BGR48LE CAVPixelFormat = C.AV_PIX_FMT_BGR48LE ///< packed RGB 16:16:16, 48bpp, 16B, 16G, 16R, the 2-byte value for each R/G/B component is stored as little-endian

	/**
	 * The following 12 formats have the disadvantage of needing 1 format for each bit depth.
	 * Notice that each 9/10 bits sample is stored in 16 bits with extra padding.
	 * If you want to support multiple bit depths, then using AV_PIX_FMT_YUV420P16* with the bpp stored separately is better.
	 */
	AV_PIX_FMT_YUV420P9BE   CAVPixelFormat = C.AV_PIX_FMT_YUV420P9BE   ///< planar YUV 4:2:0, 13.5bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV420P9LE   CAVPixelFormat = C.AV_PIX_FMT_YUV420P9LE   ///< planar YUV 4:2:0, 13.5bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV420P10BE  CAVPixelFormat = C.AV_PIX_FMT_YUV420P10BE  ///< planar YUV 4:2:0, 15bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV420P10LE  CAVPixelFormat = C.AV_PIX_FMT_YUV420P10LE  ///< planar YUV 4:2:0, 15bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV422P10BE  CAVPixelFormat = C.AV_PIX_FMT_YUV422P10BE  ///< planar YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV422P10LE  CAVPixelFormat = C.AV_PIX_FMT_YUV422P10LE  ///< planar YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P9BE   CAVPixelFormat = C.AV_PIX_FMT_YUV444P9BE   ///< planar YUV 4:4:4, 27bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P9LE   CAVPixelFormat = C.AV_PIX_FMT_YUV444P9LE   ///< planar YUV 4:4:4, 27bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P10BE  CAVPixelFormat = C.AV_PIX_FMT_YUV444P10BE  ///< planar YUV 4:4:4, 30bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P10LE  CAVPixelFormat = C.AV_PIX_FMT_YUV444P10LE  ///< planar YUV 4:4:4, 30bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_YUV422P9BE   CAVPixelFormat = C.AV_PIX_FMT_YUV422P9BE   ///< planar YUV 4:2:2, 18bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV422P9LE   CAVPixelFormat = C.AV_PIX_FMT_YUV422P9LE   ///< planar YUV 4:2:2, 18bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_GBRP         CAVPixelFormat = C.AV_PIX_FMT_GBRP         ///< planar GBR 4:4:4 24bpp
	AV_PIX_FMT_GBR24P       CAVPixelFormat = C.AV_PIX_FMT_GBR24P       // alias for #AV_PIX_FMT_GBRP
	AV_PIX_FMT_GBRP9BE      CAVPixelFormat = C.AV_PIX_FMT_GBRP9BE      ///< planar GBR 4:4:4 27bpp, big-endian
	AV_PIX_FMT_GBRP9LE      CAVPixelFormat = C.AV_PIX_FMT_GBRP9LE      ///< planar GBR 4:4:4 27bpp, little-endian
	AV_PIX_FMT_GBRP10BE     CAVPixelFormat = C.AV_PIX_FMT_GBRP10BE     ///< planar GBR 4:4:4 30bpp, big-endian
	AV_PIX_FMT_GBRP10LE     CAVPixelFormat = C.AV_PIX_FMT_GBRP10LE     ///< planar GBR 4:4:4 30bpp, little-endian
	AV_PIX_FMT_GBRP16BE     CAVPixelFormat = C.AV_PIX_FMT_GBRP16BE     ///< planar GBR 4:4:4 48bpp, big-endian
	AV_PIX_FMT_GBRP16LE     CAVPixelFormat = C.AV_PIX_FMT_GBRP16LE     ///< planar GBR 4:4:4 48bpp, little-endian
	AV_PIX_FMT_YUVA422P     CAVPixelFormat = C.AV_PIX_FMT_YUVA422P     ///< planar YUV 4:2:2 24bpp, (1 Cr & Cb sample per 2x1 Y & A samples)
	AV_PIX_FMT_YUVA444P     CAVPixelFormat = C.AV_PIX_FMT_YUVA444P     ///< planar YUV 4:4:4 32bpp, (1 Cr & Cb sample per 1x1 Y & A samples)
	AV_PIX_FMT_YUVA420P9BE  CAVPixelFormat = C.AV_PIX_FMT_YUVA420P9BE  ///< planar YUV 4:2:0 22.5bpp, (1 Cr & Cb sample per 2x2 Y & A samples), big-endian
	AV_PIX_FMT_YUVA420P9LE  CAVPixelFormat = C.AV_PIX_FMT_YUVA420P9LE  ///< planar YUV 4:2:0 22.5bpp, (1 Cr & Cb sample per 2x2 Y & A samples), little-endian
	AV_PIX_FMT_YUVA422P9BE  CAVPixelFormat = C.AV_PIX_FMT_YUVA422P9BE  ///< planar YUV 4:2:2 27bpp, (1 Cr & Cb sample per 2x1 Y & A samples), big-endian
	AV_PIX_FMT_YUVA422P9LE  CAVPixelFormat = C.AV_PIX_FMT_YUVA422P9LE  ///< planar YUV 4:2:2 27bpp, (1 Cr & Cb sample per 2x1 Y & A samples), little-endian
	AV_PIX_FMT_YUVA444P9BE  CAVPixelFormat = C.AV_PIX_FMT_YUVA444P9BE  ///< planar YUV 4:4:4 36bpp, (1 Cr & Cb sample per 1x1 Y & A samples), big-endian
	AV_PIX_FMT_YUVA444P9LE  CAVPixelFormat = C.AV_PIX_FMT_YUVA444P9LE  ///< planar YUV 4:4:4 36bpp, (1 Cr & Cb sample per 1x1 Y & A samples), little-endian
	AV_PIX_FMT_YUVA420P10BE CAVPixelFormat = C.AV_PIX_FMT_YUVA420P10BE ///< planar YUV 4:2:0 25bpp, (1 Cr & Cb sample per 2x2 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA420P10LE CAVPixelFormat = C.AV_PIX_FMT_YUVA420P10LE ///< planar YUV 4:2:0 25bpp, (1 Cr & Cb sample per 2x2 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA422P10BE CAVPixelFormat = C.AV_PIX_FMT_YUVA422P10BE ///< planar YUV 4:2:2 30bpp, (1 Cr & Cb sample per 2x1 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA422P10LE CAVPixelFormat = C.AV_PIX_FMT_YUVA422P10LE ///< planar YUV 4:2:2 30bpp, (1 Cr & Cb sample per 2x1 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA444P10BE CAVPixelFormat = C.AV_PIX_FMT_YUVA444P10BE ///< planar YUV 4:4:4 40bpp, (1 Cr & Cb sample per 1x1 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA444P10LE CAVPixelFormat = C.AV_PIX_FMT_YUVA444P10LE ///< planar YUV 4:4:4 40bpp, (1 Cr & Cb sample per 1x1 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA420P16BE CAVPixelFormat = C.AV_PIX_FMT_YUVA420P16BE ///< planar YUV 4:2:0 40bpp, (1 Cr & Cb sample per 2x2 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA420P16LE CAVPixelFormat = C.AV_PIX_FMT_YUVA420P16LE ///< planar YUV 4:2:0 40bpp, (1 Cr & Cb sample per 2x2 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA422P16BE CAVPixelFormat = C.AV_PIX_FMT_YUVA422P16BE ///< planar YUV 4:2:2 48bpp, (1 Cr & Cb sample per 2x1 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA422P16LE CAVPixelFormat = C.AV_PIX_FMT_YUVA422P16LE ///< planar YUV 4:2:2 48bpp, (1 Cr & Cb sample per 2x1 Y & A samples, little-endian)
	AV_PIX_FMT_YUVA444P16BE CAVPixelFormat = C.AV_PIX_FMT_YUVA444P16BE ///< planar YUV 4:4:4 64bpp, (1 Cr & Cb sample per 1x1 Y & A samples, big-endian)
	AV_PIX_FMT_YUVA444P16LE CAVPixelFormat = C.AV_PIX_FMT_YUVA444P16LE ///< planar YUV 4:4:4 64bpp, (1 Cr & Cb sample per 1x1 Y & A samples, little-endian)

	AV_PIX_FMT_VDPAU CAVPixelFormat = C.AV_PIX_FMT_VDPAU ///< HW acceleration through VDPAU, Picture.data[3] contains a VdpVideoSurface

	AV_PIX_FMT_XYZ12LE CAVPixelFormat = C.AV_PIX_FMT_XYZ12LE ///< packed XYZ 4:4:4, 36 bpp, (msb) 12X, 12Y, 12Z (lsb), the 2-byte value for each X/Y/Z is stored as little-endian, the 4 lower bits are set to 0
	AV_PIX_FMT_XYZ12BE CAVPixelFormat = C.AV_PIX_FMT_XYZ12BE ///< packed XYZ 4:4:4, 36 bpp, (msb) 12X, 12Y, 12Z (lsb), the 2-byte value for each X/Y/Z is stored as big-endian, the 4 lower bits are set to 0
	AV_PIX_FMT_NV16    CAVPixelFormat = C.AV_PIX_FMT_NV16    ///< interleaved chroma YUV 4:2:2, 16bpp, (1 Cr & Cb sample per 2x1 Y samples)
	AV_PIX_FMT_NV20LE  CAVPixelFormat = C.AV_PIX_FMT_NV20LE  ///< interleaved chroma YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_NV20BE  CAVPixelFormat = C.AV_PIX_FMT_NV20BE  ///< interleaved chroma YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian

	AV_PIX_FMT_RGBA64BE CAVPixelFormat = C.AV_PIX_FMT_RGBA64BE ///< packed RGBA 16:16:16:16, 64bpp, 16R, 16G, 16B, 16A, the 2-byte value for each R/G/B/A component is stored as big-endian
	AV_PIX_FMT_RGBA64LE CAVPixelFormat = C.AV_PIX_FMT_RGBA64LE ///< packed RGBA 16:16:16:16, 64bpp, 16R, 16G, 16B, 16A, the 2-byte value for each R/G/B/A component is stored as little-endian
	AV_PIX_FMT_BGRA64BE CAVPixelFormat = C.AV_PIX_FMT_BGRA64BE ///< packed RGBA 16:16:16:16, 64bpp, 16B, 16G, 16R, 16A, the 2-byte value for each R/G/B/A component is stored as big-endian
	AV_PIX_FMT_BGRA64LE CAVPixelFormat = C.AV_PIX_FMT_BGRA64LE ///< packed RGBA 16:16:16:16, 64bpp, 16B, 16G, 16R, 16A, the 2-byte value for each R/G/B/A component is stored as little-endian

	AV_PIX_FMT_YVYU422 CAVPixelFormat = C.AV_PIX_FMT_YVYU422 ///< packed YUV 4:2:2, 16bpp, Y0 Cr Y1 Cb

	AV_PIX_FMT_YA16BE CAVPixelFormat = C.AV_PIX_FMT_YA16BE ///< 16 bits gray, 16 bits alpha (big-endian)
	AV_PIX_FMT_YA16LE CAVPixelFormat = C.AV_PIX_FMT_YA16LE ///< 16 bits gray, 16 bits alpha (little-endian)

	AV_PIX_FMT_GBRAP     CAVPixelFormat = C.AV_PIX_FMT_GBRAP     ///< planar GBRA 4:4:4:4 32bpp
	AV_PIX_FMT_GBRAP16BE CAVPixelFormat = C.AV_PIX_FMT_GBRAP16BE ///< planar GBRA 4:4:4:4 64bpp, big-endian
	AV_PIX_FMT_GBRAP16LE CAVPixelFormat = C.AV_PIX_FMT_GBRAP16LE ///< planar GBRA 4:4:4:4 64bpp, little-endian
	/**
	 * HW acceleration through QSV, data[3] contains a pointer to the
	 * mfxFrameSurface1 structure.
	 *
	 * Before FFmpeg 5.0:
	 * mfxFrameSurface1.Data.MemId contains a pointer when importing
	 * the following frames as QSV frames:
	 *
	 * VAAPI:
	 * mfxFrameSurface1.Data.MemId contains a pointer to VASurfaceID
	 *
	 * DXVA2:
	 * mfxFrameSurface1.Data.MemId contains a pointer to IDirect3DSurface9
	 *
	 * FFmpeg 5.0 and above:
	 * mfxFrameSurface1.Data.MemId contains a pointer to the mfxHDLPair
	 * structure when importing the following frames as QSV frames:
	 *
	 * VAAPI:
	 * mfxHDLPair.first contains a VASurfaceID pointer.
	 * mfxHDLPair.second is always MFX_INFINITE.
	 *
	 * DXVA2:
	 * mfxHDLPair.first contains IDirect3DSurface9 pointer.
	 * mfxHDLPair.second is always MFX_INFINITE.
	 *
	 * D3D11:
	 * mfxHDLPair.first contains a ID3D11Texture2D pointer.
	 * mfxHDLPair.second contains the texture array index of the frame if the
	 * ID3D11Texture2D is an array texture, or always MFX_INFINITE if it is a
	 * normal texture.
	 */
	AV_PIX_FMT_QSV CAVPixelFormat = C.AV_PIX_FMT_QSV

	/**
	 * HW acceleration though MMAL, data[3] contains a pointer to the
	 * MMAL_BUFFER_HEADER_T structure.
	 */
	AV_PIX_FMT_MMAL CAVPixelFormat = C.AV_PIX_FMT_MMAL

	AV_PIX_FMT_D3D11VA_VLD CAVPixelFormat = C.AV_PIX_FMT_D3D11VA_VLD ///< HW decoding through Direct3D11 via old API, Picture.data[3] contains a ID3D11VideoDecoderOutputView pointer

	/**
	 * HW acceleration through CUDA. data[i] contain CUdeviceptr pointers
	 * exactly as for system memory frames.
	 */
	AV_PIX_FMT_CUDA CAVPixelFormat = C.AV_PIX_FMT_CUDA

	AV_PIX_FMT_0RGB CAVPixelFormat = C.AV_PIX_FMT_0RGB ///< packed RGB 8:8:8, 32bpp, XRGBXRGB...   X=unused/undefined
	AV_PIX_FMT_RGB0 CAVPixelFormat = C.AV_PIX_FMT_RGB0 ///< packed RGB 8:8:8, 32bpp, RGBXRGBX...   X=unused/undefined
	AV_PIX_FMT_0BGR CAVPixelFormat = C.AV_PIX_FMT_0BGR ///< packed BGR 8:8:8, 32bpp, XBGRXBGR...   X=unused/undefined
	AV_PIX_FMT_BGR0 CAVPixelFormat = C.AV_PIX_FMT_BGR0 ///< packed BGR 8:8:8, 32bpp, BGRXBGRX...   X=unused/undefined

	AV_PIX_FMT_YUV420P12BE CAVPixelFormat = C.AV_PIX_FMT_YUV420P12BE ///< planar YUV 4:2:0,18bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV420P12LE CAVPixelFormat = C.AV_PIX_FMT_YUV420P12LE ///< planar YUV 4:2:0,18bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV420P14BE CAVPixelFormat = C.AV_PIX_FMT_YUV420P14BE ///< planar YUV 4:2:0,21bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AV_PIX_FMT_YUV420P14LE CAVPixelFormat = C.AV_PIX_FMT_YUV420P14LE ///< planar YUV 4:2:0,21bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AV_PIX_FMT_YUV422P12BE CAVPixelFormat = C.AV_PIX_FMT_YUV422P12BE ///< planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV422P12LE CAVPixelFormat = C.AV_PIX_FMT_YUV422P12LE ///< planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_YUV422P14BE CAVPixelFormat = C.AV_PIX_FMT_YUV422P14BE ///< planar YUV 4:2:2,28bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AV_PIX_FMT_YUV422P14LE CAVPixelFormat = C.AV_PIX_FMT_YUV422P14LE ///< planar YUV 4:2:2,28bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P12BE CAVPixelFormat = C.AV_PIX_FMT_YUV444P12BE ///< planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P12LE CAVPixelFormat = C.AV_PIX_FMT_YUV444P12LE ///< planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_YUV444P14BE CAVPixelFormat = C.AV_PIX_FMT_YUV444P14BE ///< planar YUV 4:4:4,42bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AV_PIX_FMT_YUV444P14LE CAVPixelFormat = C.AV_PIX_FMT_YUV444P14LE ///< planar YUV 4:4:4,42bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AV_PIX_FMT_GBRP12BE    CAVPixelFormat = C.AV_PIX_FMT_GBRP12BE    ///< planar GBR 4:4:4 36bpp, big-endian
	AV_PIX_FMT_GBRP12LE    CAVPixelFormat = C.AV_PIX_FMT_GBRP12LE    ///< planar GBR 4:4:4 36bpp, little-endian
	AV_PIX_FMT_GBRP14BE    CAVPixelFormat = C.AV_PIX_FMT_GBRP14BE    ///< planar GBR 4:4:4 42bpp, big-endian
	AV_PIX_FMT_GBRP14LE    CAVPixelFormat = C.AV_PIX_FMT_GBRP14LE    ///< planar GBR 4:4:4 42bpp, little-endian
	AV_PIX_FMT_YUVJ411P    CAVPixelFormat = C.AV_PIX_FMT_YUVJ411P    ///< planar YUV 4:1:1, 12bpp, (1 Cr & Cb sample per 4x1 Y samples) full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV411P and setting color_range

	AV_PIX_FMT_BAYER_BGGR8    CAVPixelFormat = C.AV_PIX_FMT_BAYER_BGGR8    ///< bayer, BGBG..(odd line), GRGR..(even line), 8-bit samples
	AV_PIX_FMT_BAYER_RGGB8    CAVPixelFormat = C.AV_PIX_FMT_BAYER_RGGB8    ///< bayer, RGRG..(odd line), GBGB..(even line), 8-bit samples
	AV_PIX_FMT_BAYER_GBRG8    CAVPixelFormat = C.AV_PIX_FMT_BAYER_GBRG8    ///< bayer, GBGB..(odd line), RGRG..(even line), 8-bit samples
	AV_PIX_FMT_BAYER_GRBG8    CAVPixelFormat = C.AV_PIX_FMT_BAYER_GRBG8    ///< bayer, GRGR..(odd line), BGBG..(even line), 8-bit samples
	AV_PIX_FMT_BAYER_BGGR16LE CAVPixelFormat = C.AV_PIX_FMT_BAYER_BGGR16LE ///< bayer, BGBG..(odd line), GRGR..(even line), 16-bit samples, little-endian
	AV_PIX_FMT_BAYER_BGGR16BE CAVPixelFormat = C.AV_PIX_FMT_BAYER_BGGR16BE ///< bayer, BGBG..(odd line), GRGR..(even line), 16-bit samples, big-endian
	AV_PIX_FMT_BAYER_RGGB16LE CAVPixelFormat = C.AV_PIX_FMT_BAYER_RGGB16LE ///< bayer, RGRG..(odd line), GBGB..(even line), 16-bit samples, little-endian
	AV_PIX_FMT_BAYER_RGGB16BE CAVPixelFormat = C.AV_PIX_FMT_BAYER_RGGB16BE ///< bayer, RGRG..(odd line), GBGB..(even line), 16-bit samples, big-endian
	AV_PIX_FMT_BAYER_GBRG16LE CAVPixelFormat = C.AV_PIX_FMT_BAYER_GBRG16LE ///< bayer, GBGB..(odd line), RGRG..(even line), 16-bit samples, little-endian
	AV_PIX_FMT_BAYER_GBRG16BE CAVPixelFormat = C.AV_PIX_FMT_BAYER_GBRG16BE ///< bayer, GBGB..(odd line), RGRG..(even line), 16-bit samples, big-endian
	AV_PIX_FMT_BAYER_GRBG16LE CAVPixelFormat = C.AV_PIX_FMT_BAYER_GRBG16LE ///< bayer, GRGR..(odd line), BGBG..(even line), 16-bit samples, little-endian
	AV_PIX_FMT_BAYER_GRBG16BE CAVPixelFormat = C.AV_PIX_FMT_BAYER_GRBG16BE ///< bayer, GRGR..(odd line), BGBG..(even line), 16-bit samples, big-endian

	// #if FF_API_XVMC
	AV_PIX_FMT_XVMC CAVPixelFormat = C.AV_PIX_FMT_XVMC ///< XVideo Motion Acceleration via common packet passing
	// #endif

	AV_PIX_FMT_YUV440P10LE CAVPixelFormat = C.AV_PIX_FMT_YUV440P10LE ///< planar YUV 4:4:0,20bpp, (1 Cr & Cb sample per 1x2 Y samples), little-endian
	AV_PIX_FMT_YUV440P10BE CAVPixelFormat = C.AV_PIX_FMT_YUV440P10BE ///< planar YUV 4:4:0,20bpp, (1 Cr & Cb sample per 1x2 Y samples), big-endian
	AV_PIX_FMT_YUV440P12LE CAVPixelFormat = C.AV_PIX_FMT_YUV440P12LE ///< planar YUV 4:4:0,24bpp, (1 Cr & Cb sample per 1x2 Y samples), little-endian
	AV_PIX_FMT_YUV440P12BE CAVPixelFormat = C.AV_PIX_FMT_YUV440P12BE ///< planar YUV 4:4:0,24bpp, (1 Cr & Cb sample per 1x2 Y samples), big-endian
	AV_PIX_FMT_AYUV64LE    CAVPixelFormat = C.AV_PIX_FMT_AYUV64LE    ///< packed AYUV 4:4:4,64bpp (1 Cr & Cb sample per 1x1 Y & A samples), little-endian
	AV_PIX_FMT_AYUV64BE    CAVPixelFormat = C.AV_PIX_FMT_AYUV64BE    ///< packed AYUV 4:4:4,64bpp (1 Cr & Cb sample per 1x1 Y & A samples), big-endian

	AV_PIX_FMT_VIDEOTOOLBOX CAVPixelFormat = C.AV_PIX_FMT_VIDEOTOOLBOX ///< hardware decoding through Videotoolbox

	AV_PIX_FMT_P010LE CAVPixelFormat = C.AV_PIX_FMT_P010LE ///< like NV12, with 10bpp per component, data in the high bits, zeros in the low bits, little-endian
	AV_PIX_FMT_P010BE CAVPixelFormat = C.AV_PIX_FMT_P010BE ///< like NV12, with 10bpp per component, data in the high bits, zeros in the low bits, big-endian

	AV_PIX_FMT_GBRAP12BE CAVPixelFormat = C.AV_PIX_FMT_GBRAP12BE ///< planar GBR 4:4:4:4 48bpp, big-endian
	AV_PIX_FMT_GBRAP12LE CAVPixelFormat = C.AV_PIX_FMT_GBRAP12LE ///< planar GBR 4:4:4:4 48bpp, little-endian

	AV_PIX_FMT_GBRAP10BE CAVPixelFormat = C.AV_PIX_FMT_GBRAP10BE ///< planar GBR 4:4:4:4 40bpp, big-endian
	AV_PIX_FMT_GBRAP10LE CAVPixelFormat = C.AV_PIX_FMT_GBRAP10LE ///< planar GBR 4:4:4:4 40bpp, little-endian

	AV_PIX_FMT_MEDIACODEC CAVPixelFormat = C.AV_PIX_FMT_MEDIACODEC ///< hardware decoding through MediaCodec

	AV_PIX_FMT_GRAY12BE CAVPixelFormat = C.AV_PIX_FMT_GRAY12BE ///<        Y        , 12bpp, big-endian
	AV_PIX_FMT_GRAY12LE CAVPixelFormat = C.AV_PIX_FMT_GRAY12LE ///<        Y        , 12bpp, little-endian
	AV_PIX_FMT_GRAY10BE CAVPixelFormat = C.AV_PIX_FMT_GRAY10BE ///<        Y        , 10bpp, big-endian
	AV_PIX_FMT_GRAY10LE CAVPixelFormat = C.AV_PIX_FMT_GRAY10LE ///<        Y        , 10bpp, little-endian

	AV_PIX_FMT_P016LE CAVPixelFormat = C.AV_PIX_FMT_P016LE ///< like NV12, with 16bpp per component, little-endian
	AV_PIX_FMT_P016BE CAVPixelFormat = C.AV_PIX_FMT_P016BE ///< like NV12, with 16bpp per component, big-endian

	/**
	 * Hardware surfaces for Direct3D11.
	 *
	 * This is preferred over the legacy AV_PIX_FMT_D3D11VA_VLD. The new D3D11
	 * hwaccel API and filtering support AV_PIX_FMT_D3D11 only.
	 *
	 * data[0] contains a ID3D11Texture2D pointer, and data[1] contains the
	 * texture array index of the frame as intptr_t if the ID3D11Texture2D is
	 * an array texture (or always 0 if it's a normal texture).
	 */
	AV_PIX_FMT_D3D11 CAVPixelFormat = C.AV_PIX_FMT_D3D11

	AV_PIX_FMT_GRAY9BE CAVPixelFormat = C.AV_PIX_FMT_GRAY9BE ///<        Y        , 9bpp, big-endian
	AV_PIX_FMT_GRAY9LE CAVPixelFormat = C.AV_PIX_FMT_GRAY9LE ///<        Y        , 9bpp, little-endian

	AV_PIX_FMT_GBRPF32BE  CAVPixelFormat = C.AV_PIX_FMT_GBRPF32BE  ///< IEEE-754 single precision planar GBR 4:4:4,     96bpp, big-endian
	AV_PIX_FMT_GBRPF32LE  CAVPixelFormat = C.AV_PIX_FMT_GBRPF32LE  ///< IEEE-754 single precision planar GBR 4:4:4,     96bpp, little-endian
	AV_PIX_FMT_GBRAPF32BE CAVPixelFormat = C.AV_PIX_FMT_GBRAPF32BE ///< IEEE-754 single precision planar GBRA 4:4:4:4, 128bpp, big-endian
	AV_PIX_FMT_GBRAPF32LE CAVPixelFormat = C.AV_PIX_FMT_GBRAPF32LE ///< IEEE-754 single precision planar GBRA 4:4:4:4, 128bpp, little-endian

	/**
	 * DRM-managed buffers exposed through PRIME buffer sharing.
	 *
	 * data[0] points to an AVDRMFrameDescriptor.
	 */
	AV_PIX_FMT_DRM_PRIME CAVPixelFormat = C.AV_PIX_FMT_DRM_PRIME
	/**
	 * Hardware surfaces for OpenCL.
	 *
	 * data[i] contain 2D image objects (typed in C as cl_mem, used
	 * in OpenCL as image2d_t) for each plane of the surface.
	 */
	AV_PIX_FMT_OPENCL CAVPixelFormat = C.AV_PIX_FMT_OPENCL

	AV_PIX_FMT_GRAY14BE CAVPixelFormat = C.AV_PIX_FMT_GRAY14BE ///<        Y        , 14bpp, big-endian
	AV_PIX_FMT_GRAY14LE CAVPixelFormat = C.AV_PIX_FMT_GRAY14LE ///<        Y        , 14bpp, little-endian

	AV_PIX_FMT_GRAYF32BE CAVPixelFormat = C.AV_PIX_FMT_GRAYF32BE ///< IEEE-754 single precision Y, 32bpp, big-endian
	AV_PIX_FMT_GRAYF32LE CAVPixelFormat = C.AV_PIX_FMT_GRAYF32LE ///< IEEE-754 single precision Y, 32bpp, little-endian

	AV_PIX_FMT_YUVA422P12BE CAVPixelFormat = C.AV_PIX_FMT_YUVA422P12BE ///< planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), 12b alpha, big-endian
	AV_PIX_FMT_YUVA422P12LE CAVPixelFormat = C.AV_PIX_FMT_YUVA422P12LE ///< planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), 12b alpha, little-endian
	AV_PIX_FMT_YUVA444P12BE CAVPixelFormat = C.AV_PIX_FMT_YUVA444P12BE ///< planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), 12b alpha, big-endian
	AV_PIX_FMT_YUVA444P12LE CAVPixelFormat = C.AV_PIX_FMT_YUVA444P12LE ///< planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), 12b alpha, little-endian

	AV_PIX_FMT_NV24 CAVPixelFormat = C.AV_PIX_FMT_NV24 ///< planar YUV 4:4:4, 24bpp, 1 plane for Y and 1 plane for the UV components, which are interleaved (first byte U and the following byte V)
	AV_PIX_FMT_NV42 CAVPixelFormat = C.AV_PIX_FMT_NV42 ///< as above, but U and V bytes are swapped

	/**
	 * Vulkan hardware images.
	 *
	 * data[0] points to an AVVkFrame
	 */
	AV_PIX_FMT_VULKAN CAVPixelFormat = C.AV_PIX_FMT_VULKAN

	AV_PIX_FMT_Y210BE CAVPixelFormat = C.AV_PIX_FMT_Y210BE ///< packed YUV 4:2:2 like YUYV422, 20bpp, data in the high bits, big-endian
	AV_PIX_FMT_Y210LE CAVPixelFormat = C.AV_PIX_FMT_Y210LE ///< packed YUV 4:2:2 like YUYV422, 20bpp, data in the high bits, little-endian

	AV_PIX_FMT_X2RGB10LE CAVPixelFormat = C.AV_PIX_FMT_X2RGB10LE ///< packed RGB 10:10:10, 30bpp, (msb)2X 10R 10G 10B(lsb), little-endian, X=unused/undefined
	AV_PIX_FMT_X2RGB10BE CAVPixelFormat = C.AV_PIX_FMT_X2RGB10BE ///< packed RGB 10:10:10, 30bpp, (msb)2X 10R 10G 10B(lsb), big-endian, X=unused/undefined
	AV_PIX_FMT_X2BGR10LE CAVPixelFormat = C.AV_PIX_FMT_X2BGR10LE ///< packed BGR 10:10:10, 30bpp, (msb)2X 10B 10G 10R(lsb), little-endian, X=unused/undefined
	AV_PIX_FMT_X2BGR10BE CAVPixelFormat = C.AV_PIX_FMT_X2BGR10BE ///< packed BGR 10:10:10, 30bpp, (msb)2X 10B 10G 10R(lsb), big-endian, X=unused/undefined

	AV_PIX_FMT_P210BE CAVPixelFormat = C.AV_PIX_FMT_P210BE ///< interleaved chroma YUV 4:2:2, 20bpp, data in the high bits, big-endian
	AV_PIX_FMT_P210LE CAVPixelFormat = C.AV_PIX_FMT_P210LE ///< interleaved chroma YUV 4:2:2, 20bpp, data in the high bits, little-endian

	AV_PIX_FMT_P410BE CAVPixelFormat = C.AV_PIX_FMT_P410BE ///< interleaved chroma YUV 4:4:4, 30bpp, data in the high bits, big-endian
	AV_PIX_FMT_P410LE CAVPixelFormat = C.AV_PIX_FMT_P410LE ///< interleaved chroma YUV 4:4:4, 30bpp, data in the high bits, little-endian

	AV_PIX_FMT_P216BE CAVPixelFormat = C.AV_PIX_FMT_P216BE ///< interleaved chroma YUV 4:2:2, 32bpp, big-endian
	AV_PIX_FMT_P216LE CAVPixelFormat = C.AV_PIX_FMT_P216LE ///< interleaved chroma YUV 4:2:2, 32bpp, little-endian

	AV_PIX_FMT_P416BE CAVPixelFormat = C.AV_PIX_FMT_P416BE ///< interleaved chroma YUV 4:4:4, 48bpp, big-endian
	AV_PIX_FMT_P416LE CAVPixelFormat = C.AV_PIX_FMT_P416LE ///< interleaved chroma YUV 4:4:4, 48bpp, little-endian

	AV_PIX_FMT_VUYA CAVPixelFormat = C.AV_PIX_FMT_VUYA ///< packed VUYA 4:4:4, 32bpp, VUYAVUYA...

	AV_PIX_FMT_RGBAF16BE CAVPixelFormat = C.AV_PIX_FMT_RGBAF16BE ///< IEEE-754 half precision packed RGBA 16:16:16:16, 64bpp, RGBARGBA..., big-endian
	AV_PIX_FMT_RGBAF16LE CAVPixelFormat = C.AV_PIX_FMT_RGBAF16LE ///< IEEE-754 half precision packed RGBA 16:16:16:16, 64bpp, RGBARGBA..., little-endian

	AV_PIX_FMT_VUYX CAVPixelFormat = C.AV_PIX_FMT_VUYX ///< packed VUYX 4:4:4, 32bpp, Variant of VUYA where alpha channel is left undefined

	AV_PIX_FMT_P012LE CAVPixelFormat = C.AV_PIX_FMT_P012LE ///< like NV12, with 12bpp per component, data in the high bits, zeros in the low bits, little-endian
	AV_PIX_FMT_P012BE CAVPixelFormat = C.AV_PIX_FMT_P012BE ///< like NV12, with 12bpp per component, data in the high bits, zeros in the low bits, big-endian

	AV_PIX_FMT_Y212BE CAVPixelFormat = C.AV_PIX_FMT_Y212BE ///< packed YUV 4:2:2 like YUYV422, 24bpp, data in the high bits, zeros in the low bits, big-endian
	AV_PIX_FMT_Y212LE CAVPixelFormat = C.AV_PIX_FMT_Y212LE ///< packed YUV 4:2:2 like YUYV422, 24bpp, data in the high bits, zeros in the low bits, little-endian

	AV_PIX_FMT_XV30BE CAVPixelFormat = C.AV_PIX_FMT_XV30BE ///< packed XVYU 4:4:4, 32bpp, (msb)2X 10V 10Y 10U(lsb), big-endian, variant of Y410 where alpha channel is left undefined
	AV_PIX_FMT_XV30LE CAVPixelFormat = C.AV_PIX_FMT_XV30LE ///< packed XVYU 4:4:4, 32bpp, (msb)2X 10V 10Y 10U(lsb), little-endian, variant of Y410 where alpha channel is left undefined

	AV_PIX_FMT_XV36BE CAVPixelFormat = C.AV_PIX_FMT_XV36BE ///< packed XVYU 4:4:4, 48bpp, data in the high bits, zeros in the low bits, big-endian, variant of Y412 where alpha channel is left undefined
	AV_PIX_FMT_XV36LE CAVPixelFormat = C.AV_PIX_FMT_XV36LE ///< packed XVYU 4:4:4, 48bpp, data in the high bits, zeros in the low bits, little-endian, variant of Y412 where alpha channel is left undefined

	AV_PIX_FMT_RGBF32BE CAVPixelFormat = C.AV_PIX_FMT_RGBF32BE ///< IEEE-754 single precision packed RGB 32:32:32, 96bpp, RGBRGB..., big-endian
	AV_PIX_FMT_RGBF32LE CAVPixelFormat = C.AV_PIX_FMT_RGBF32LE ///< IEEE-754 single precision packed RGB 32:32:32, 96bpp, RGBRGB..., little-endian

	AV_PIX_FMT_RGBAF32BE CAVPixelFormat = C.AV_PIX_FMT_RGBAF32BE ///< IEEE-754 single precision packed RGBA 32:32:32:32, 128bpp, RGBARGBA..., big-endian
	AV_PIX_FMT_RGBAF32LE CAVPixelFormat = C.AV_PIX_FMT_RGBAF32LE ///< IEEE-754 single precision packed RGBA 32:32:32:32, 128bpp, RGBARGBA..., little-endian

	AV_PIX_FMT_P212BE CAVPixelFormat = C.AV_PIX_FMT_P212BE ///< interleaved chroma YUV 4:2:2, 24bpp, data in the high bits, big-endian
	AV_PIX_FMT_P212LE CAVPixelFormat = C.AV_PIX_FMT_P212LE ///< interleaved chroma YUV 4:2:2, 24bpp, data in the high bits, little-endian

	AV_PIX_FMT_P412BE CAVPixelFormat = C.AV_PIX_FMT_P412BE ///< interleaved chroma YUV 4:4:4, 36bpp, data in the high bits, big-endian
	AV_PIX_FMT_P412LE CAVPixelFormat = C.AV_PIX_FMT_P412LE ///< interleaved chroma YUV 4:4:4, 36bpp, data in the high bits, little-endian

	AV_PIX_FMT_GBRAP14BE CAVPixelFormat = C.AV_PIX_FMT_GBRAP14BE ///< planar GBR 4:4:4:4 56bpp, big-endian
	AV_PIX_FMT_GBRAP14LE CAVPixelFormat = C.AV_PIX_FMT_GBRAP14LE ///< planar GBR 4:4:4:4 56bpp, little-endian

	AV_PIX_FMT_NB CAVPixelFormat = C.AV_PIX_FMT_NB ///< number of pixel formats, DO NOT USE THIS if you want to link with shared libav* because the number of formats might differ between versions
)

// #if AV_HAVE_BIGENDIAN
// #   define AV_PIX_FMT_NE(be, le) AV_PIX_FMT_##be
// #else
const AV_HAVE_BIGENDIAN = C.AV_HAVE_BIGENDIAN

// #endif

const (
	AV_PIX_FMT_RGB32   CAVPixelFormat = C.AV_PIX_FMT_RGB32
	AV_PIX_FMT_RGB32_1 CAVPixelFormat = C.AV_PIX_FMT_RGB32_1
	AV_PIX_FMT_BGR32   CAVPixelFormat = C.AV_PIX_FMT_BGR32
	AV_PIX_FMT_BGR32_1 CAVPixelFormat = C.AV_PIX_FMT_BGR32_1
	AV_PIX_FMT_0RGB32  CAVPixelFormat = C.AV_PIX_FMT_0RGB32
	AV_PIX_FMT_0BGR32  CAVPixelFormat = C.AV_PIX_FMT_0BGR32

	AV_PIX_FMT_GRAY9  CAVPixelFormat = C.AV_PIX_FMT_GRAY9
	AV_PIX_FMT_GRAY10 CAVPixelFormat = C.AV_PIX_FMT_GRAY10
	AV_PIX_FMT_GRAY12 CAVPixelFormat = C.AV_PIX_FMT_GRAY12
	AV_PIX_FMT_GRAY14 CAVPixelFormat = C.AV_PIX_FMT_GRAY14
	AV_PIX_FMT_GRAY16 CAVPixelFormat = C.AV_PIX_FMT_GRAY16
	AV_PIX_FMT_YA16   CAVPixelFormat = C.AV_PIX_FMT_YA16
	AV_PIX_FMT_RGB48  CAVPixelFormat = C.AV_PIX_FMT_RGB48
	AV_PIX_FMT_RGB565 CAVPixelFormat = C.AV_PIX_FMT_RGB565
	AV_PIX_FMT_RGB555 CAVPixelFormat = C.AV_PIX_FMT_RGB555
	AV_PIX_FMT_RGB444 CAVPixelFormat = C.AV_PIX_FMT_RGB444
	AV_PIX_FMT_RGBA64 CAVPixelFormat = C.AV_PIX_FMT_RGBA64
	AV_PIX_FMT_BGR48  CAVPixelFormat = C.AV_PIX_FMT_BGR48
	AV_PIX_FMT_BGR565 CAVPixelFormat = C.AV_PIX_FMT_BGR565
	AV_PIX_FMT_BGR555 CAVPixelFormat = C.AV_PIX_FMT_BGR555
	AV_PIX_FMT_BGR444 CAVPixelFormat = C.AV_PIX_FMT_BGR444
	AV_PIX_FMT_BGRA64 CAVPixelFormat = C.AV_PIX_FMT_BGRA64

	AV_PIX_FMT_YUV420P9  CAVPixelFormat = C.AV_PIX_FMT_YUV420P9
	AV_PIX_FMT_YUV422P9  CAVPixelFormat = C.AV_PIX_FMT_YUV422P9
	AV_PIX_FMT_YUV444P9  CAVPixelFormat = C.AV_PIX_FMT_YUV444P9
	AV_PIX_FMT_YUV420P10 CAVPixelFormat = C.AV_PIX_FMT_YUV420P10
	AV_PIX_FMT_YUV422P10 CAVPixelFormat = C.AV_PIX_FMT_YUV422P10
	AV_PIX_FMT_YUV440P10 CAVPixelFormat = C.AV_PIX_FMT_YUV440P10
	AV_PIX_FMT_YUV444P10 CAVPixelFormat = C.AV_PIX_FMT_YUV444P10
	AV_PIX_FMT_YUV420P12 CAVPixelFormat = C.AV_PIX_FMT_YUV420P12
	AV_PIX_FMT_YUV422P12 CAVPixelFormat = C.AV_PIX_FMT_YUV422P12
	AV_PIX_FMT_YUV440P12 CAVPixelFormat = C.AV_PIX_FMT_YUV440P12
	AV_PIX_FMT_YUV444P12 CAVPixelFormat = C.AV_PIX_FMT_YUV444P12
	AV_PIX_FMT_YUV420P14 CAVPixelFormat = C.AV_PIX_FMT_YUV420P14
	AV_PIX_FMT_YUV422P14 CAVPixelFormat = C.AV_PIX_FMT_YUV422P14
	AV_PIX_FMT_YUV444P14 CAVPixelFormat = C.AV_PIX_FMT_YUV444P14
	AV_PIX_FMT_YUV420P16 CAVPixelFormat = C.AV_PIX_FMT_YUV420P16
	AV_PIX_FMT_YUV422P16 CAVPixelFormat = C.AV_PIX_FMT_YUV422P16
	AV_PIX_FMT_YUV444P16 CAVPixelFormat = C.AV_PIX_FMT_YUV444P16

	AV_PIX_FMT_GBRP9   CAVPixelFormat = C.AV_PIX_FMT_GBRP9
	AV_PIX_FMT_GBRP10  CAVPixelFormat = C.AV_PIX_FMT_GBRP10
	AV_PIX_FMT_GBRP12  CAVPixelFormat = C.AV_PIX_FMT_GBRP12
	AV_PIX_FMT_GBRP14  CAVPixelFormat = C.AV_PIX_FMT_GBRP14
	AV_PIX_FMT_GBRP16  CAVPixelFormat = C.AV_PIX_FMT_GBRP16
	AV_PIX_FMT_GBRAP10 CAVPixelFormat = C.AV_PIX_FMT_GBRAP10
	AV_PIX_FMT_GBRAP12 CAVPixelFormat = C.AV_PIX_FMT_GBRAP12
	AV_PIX_FMT_GBRAP14 CAVPixelFormat = C.AV_PIX_FMT_GBRAP14
	AV_PIX_FMT_GBRAP16 CAVPixelFormat = C.AV_PIX_FMT_GBRAP16

	AV_PIX_FMT_BAYER_BGGR16 CAVPixelFormat = C.AV_PIX_FMT_BAYER_BGGR16
	AV_PIX_FMT_BAYER_RGGB16 CAVPixelFormat = C.AV_PIX_FMT_BAYER_RGGB16
	AV_PIX_FMT_BAYER_GBRG16 CAVPixelFormat = C.AV_PIX_FMT_BAYER_GBRG16
	AV_PIX_FMT_BAYER_GRBG16 CAVPixelFormat = C.AV_PIX_FMT_BAYER_GRBG16

	AV_PIX_FMT_GBRPF32  CAVPixelFormat = C.AV_PIX_FMT_GBRPF32
	AV_PIX_FMT_GBRAPF32 CAVPixelFormat = C.AV_PIX_FMT_GBRAPF32

	AV_PIX_FMT_GRAYF32 CAVPixelFormat = C.AV_PIX_FMT_GRAYF32

	AV_PIX_FMT_YUVA420P9  CAVPixelFormat = C.AV_PIX_FMT_YUVA420P9
	AV_PIX_FMT_YUVA422P9  CAVPixelFormat = C.AV_PIX_FMT_YUVA422P9
	AV_PIX_FMT_YUVA444P9  CAVPixelFormat = C.AV_PIX_FMT_YUVA444P9
	AV_PIX_FMT_YUVA420P10 CAVPixelFormat = C.AV_PIX_FMT_YUVA420P10
	AV_PIX_FMT_YUVA422P10 CAVPixelFormat = C.AV_PIX_FMT_YUVA422P10
	AV_PIX_FMT_YUVA444P10 CAVPixelFormat = C.AV_PIX_FMT_YUVA444P10
	AV_PIX_FMT_YUVA422P12 CAVPixelFormat = C.AV_PIX_FMT_YUVA422P12
	AV_PIX_FMT_YUVA444P12 CAVPixelFormat = C.AV_PIX_FMT_YUVA444P12
	AV_PIX_FMT_YUVA420P16 CAVPixelFormat = C.AV_PIX_FMT_YUVA420P16
	AV_PIX_FMT_YUVA422P16 CAVPixelFormat = C.AV_PIX_FMT_YUVA422P16
	AV_PIX_FMT_YUVA444P16 CAVPixelFormat = C.AV_PIX_FMT_YUVA444P16

	AV_PIX_FMT_XYZ12  CAVPixelFormat = C.AV_PIX_FMT_XYZ12
	AV_PIX_FMT_NV20   CAVPixelFormat = C.AV_PIX_FMT_NV20
	AV_PIX_FMT_AYUV64 CAVPixelFormat = C.AV_PIX_FMT_AYUV64
	AV_PIX_FMT_P010   CAVPixelFormat = C.AV_PIX_FMT_P010
	AV_PIX_FMT_P012   CAVPixelFormat = C.AV_PIX_FMT_P012
	AV_PIX_FMT_P016   CAVPixelFormat = C.AV_PIX_FMT_P016

	AV_PIX_FMT_Y210    CAVPixelFormat = C.AV_PIX_FMT_Y210
	AV_PIX_FMT_Y212    CAVPixelFormat = C.AV_PIX_FMT_Y212
	AV_PIX_FMT_XV30    CAVPixelFormat = C.AV_PIX_FMT_XV30
	AV_PIX_FMT_XV36    CAVPixelFormat = C.AV_PIX_FMT_XV36
	AV_PIX_FMT_X2RGB10 CAVPixelFormat = C.AV_PIX_FMT_X2RGB10
	AV_PIX_FMT_X2BGR10 CAVPixelFormat = C.AV_PIX_FMT_X2BGR10

	AV_PIX_FMT_P210 CAVPixelFormat = C.AV_PIX_FMT_P210
	AV_PIX_FMT_P410 CAVPixelFormat = C.AV_PIX_FMT_P410
	AV_PIX_FMT_P212 CAVPixelFormat = C.AV_PIX_FMT_P212
	AV_PIX_FMT_P412 CAVPixelFormat = C.AV_PIX_FMT_P412
	AV_PIX_FMT_P216 CAVPixelFormat = C.AV_PIX_FMT_P216
	AV_PIX_FMT_P416 CAVPixelFormat = C.AV_PIX_FMT_P416

	AV_PIX_FMT_RGBAF16 CAVPixelFormat = C.AV_PIX_FMT_RGBAF16

	AV_PIX_FMT_RGBF32  CAVPixelFormat = C.AV_PIX_FMT_RGBF32
	AV_PIX_FMT_RGBAF32 CAVPixelFormat = C.AV_PIX_FMT_RGBAF32
)

/**
 * Chromaticity coordinates of the source primaries.
 * These values match the ones defined by ISO/IEC 23091-2_2019 subclause 8.1 and ITU-T H.273.
 */
type CAVColorPrimaries C.enum_AVColorPrimaries

const (
	AVCOL_PRI_RESERVED0   CAVColorPrimaries = C.AVCOL_PRI_RESERVED0
	AVCOL_PRI_BT709       CAVColorPrimaries = C.AVCOL_PRI_BT709 ///< also ITU-R BT1361 / IEC 61966-2-4 / SMPTE RP 177 Annex B
	AVCOL_PRI_UNSPECIFIED CAVColorPrimaries = C.AVCOL_PRI_UNSPECIFIED
	AVCOL_PRI_RESERVED    CAVColorPrimaries = C.AVCOL_PRI_RESERVED
	AVCOL_PRI_BT470M      CAVColorPrimaries = C.AVCOL_PRI_BT470M ///< also FCC Title 47 Code of Federal Regulations 73.682 (a)(20)

	AVCOL_PRI_BT470BG      CAVColorPrimaries = C.AVCOL_PRI_BT470BG   ///< also ITU-R BT601-6 625 / ITU-R BT1358 625 / ITU-R BT1700 625 PAL & SECAM
	AVCOL_PRI_SMPTE170M    CAVColorPrimaries = C.AVCOL_PRI_SMPTE170M ///< also ITU-R BT601-6 525 / ITU-R BT1358 525 / ITU-R BT1700 NTSC
	AVCOL_PRI_SMPTE240M    CAVColorPrimaries = C.AVCOL_PRI_SMPTE240M ///< identical to above, also called "SMPTE C" even though it uses D65
	AVCOL_PRI_FILM         CAVColorPrimaries = C.AVCOL_PRI_FILM      ///< colour filters using Illuminant C
	AVCOL_PRI_BT2020       CAVColorPrimaries = C.AVCOL_PRI_BT2020    ///< ITU-R BT2020
	AVCOL_PRI_SMPTE428     CAVColorPrimaries = C.AVCOL_PRI_SMPTE428  ///< SMPTE ST 428-1 (CIE 1931 XYZ)
	AVCOL_PRI_SMPTEST428_1 CAVColorPrimaries = C.AVCOL_PRI_SMPTEST428_1
	AVCOL_PRI_SMPTE431     CAVColorPrimaries = C.AVCOL_PRI_SMPTE431 ///< SMPTE ST 431-2 (2011) / DCI P3
	AVCOL_PRI_SMPTE432     CAVColorPrimaries = C.AVCOL_PRI_SMPTE432 ///< SMPTE ST 432-1 (2010) / P3 D65 / Display P3
	AVCOL_PRI_EBU3213      CAVColorPrimaries = C.AVCOL_PRI_EBU3213  ///< EBU Tech. 3213-E (nothing there) / one of JEDEC P22 group phosphors
	AVCOL_PRI_JEDEC_P22    CAVColorPrimaries = C.AVCOL_PRI_JEDEC_P22
	AVCOL_PRI_NB           CAVColorPrimaries = C.AVCOL_PRI_NB ///< Not part of ABI
)

/**
 * Color Transfer Characteristic.
 * These values match the ones defined by ISO/IEC 23091-2_2019 subclause 8.2.
 */
type CAVColorTransferCharacteristic C.enum_AVColorTransferCharacteristic

const (
	AVCOL_TRC_RESERVED0    CAVColorTransferCharacteristic = C.AVCOL_TRC_RESERVED0
	AVCOL_TRC_BT709        CAVColorTransferCharacteristic = C.AVCOL_TRC_BT709 ///< also ITU-R BT1361
	AVCOL_TRC_UNSPECIFIED  CAVColorTransferCharacteristic = C.AVCOL_TRC_UNSPECIFIED
	AVCOL_TRC_RESERVED     CAVColorTransferCharacteristic = C.AVCOL_TRC_RESERVED
	AVCOL_TRC_GAMMA22      CAVColorTransferCharacteristic = C.AVCOL_TRC_GAMMA22   ///< also ITU-R BT470M / ITU-R BT1700 625 PAL & SECAM
	AVCOL_TRC_GAMMA28      CAVColorTransferCharacteristic = C.AVCOL_TRC_GAMMA28   ///< also ITU-R BT470BG
	AVCOL_TRC_SMPTE170M    CAVColorTransferCharacteristic = C.AVCOL_TRC_SMPTE170M ///< also ITU-R BT601-6 525 or 625 / ITU-R BT1358 525 or 625 / ITU-R BT1700 NTSC
	AVCOL_TRC_SMPTE240M    CAVColorTransferCharacteristic = C.AVCOL_TRC_SMPTE240M
	AVCOL_TRC_LINEAR       CAVColorTransferCharacteristic = C.AVCOL_TRC_LINEAR       ///< "Linear transfer characteristics"
	AVCOL_TRC_LOG          CAVColorTransferCharacteristic = C.AVCOL_TRC_LOG          ///< "Logarithmic transfer characteristic (100:1 range)"
	AVCOL_TRC_LOG_SQRT     CAVColorTransferCharacteristic = C.AVCOL_TRC_LOG_SQRT     ///< "Logarithmic transfer characteristic (100 * Sqrt(10) : 1 range)"
	AVCOL_TRC_IEC61966_2_4 CAVColorTransferCharacteristic = C.AVCOL_TRC_IEC61966_2_4 ///< IEC 61966-2-4
	AVCOL_TRC_BT1361_ECG   CAVColorTransferCharacteristic = C.AVCOL_TRC_BT1361_ECG   ///< ITU-R BT1361 Extended Colour Gamut
	AVCOL_TRC_IEC61966_2_1 CAVColorTransferCharacteristic = C.AVCOL_TRC_IEC61966_2_1 ///< IEC 61966-2-1 (sRGB or sYCC)
	AVCOL_TRC_BT2020_10    CAVColorTransferCharacteristic = C.AVCOL_TRC_BT2020_10    ///< ITU-R BT2020 for 10-bit system
	AVCOL_TRC_BT2020_12    CAVColorTransferCharacteristic = C.AVCOL_TRC_BT2020_12    ///< ITU-R BT2020 for 12-bit system
	AVCOL_TRC_SMPTE2084    CAVColorTransferCharacteristic = C.AVCOL_TRC_SMPTE2084    ///< SMPTE ST 2084 for 10-, 12-, 14- and 16-bit systems
	AVCOL_TRC_SMPTEST2084  CAVColorTransferCharacteristic = C.AVCOL_TRC_SMPTEST2084
	AVCOL_TRC_SMPTE428     CAVColorTransferCharacteristic = C.AVCOL_TRC_SMPTE428 ///< SMPTE ST 428-1
	AVCOL_TRC_SMPTEST428_1 CAVColorTransferCharacteristic = C.AVCOL_TRC_SMPTEST428_1
	AVCOL_TRC_ARIB_STD_B67 CAVColorTransferCharacteristic = C.AVCOL_TRC_ARIB_STD_B67 ///< ARIB STD-B67, known as "Hybrid log-gamma"
	AVCOL_TRC_NB           CAVColorTransferCharacteristic = C.AVCOL_TRC_NB           ///< Not part of ABI
)

/**
 * YUV colorspace type.
 * These values match the ones defined by ISO/IEC 23091-2_2019 subclause 8.3.
 */
type CAVColorSpace C.enum_AVColorSpace

const (
	AVCOL_SPC_RGB                CAVColorSpace = C.AVCOL_SPC_RGB   ///< order of coefficients is actually GBR, also IEC 61966-2-1 (sRGB), YZX and ST 428-1
	AVCOL_SPC_BT709              CAVColorSpace = C.AVCOL_SPC_BT709 ///< also ITU-R BT1361 / IEC 61966-2-4 xvYCC709 / derived in SMPTE RP 177 Annex B
	AVCOL_SPC_UNSPECIFIED        CAVColorSpace = C.AVCOL_SPC_UNSPECIFIED
	AVCOL_SPC_RESERVED           CAVColorSpace = C.AVCOL_SPC_RESERVED  ///< reserved for future use by ITU-T and ISO/IEC just like 15-255 are
	AVCOL_SPC_FCC                CAVColorSpace = C.AVCOL_SPC_FCC       ///< FCC Title 47 Code of Federal Regulations 73.682 (a)(20)
	AVCOL_SPC_BT470BG            CAVColorSpace = C.AVCOL_SPC_BT470BG   ///< also ITU-R BT601-6 625 / ITU-R BT1358 625 / ITU-R BT1700 625 PAL & SECAM / IEC 61966-2-4 xvYCC601
	AVCOL_SPC_SMPTE170M          CAVColorSpace = C.AVCOL_SPC_SMPTE170M ///< also ITU-R BT601-6 525 / ITU-R BT1358 525 / ITU-R BT1700 NTSC / functionally identical to above
	AVCOL_SPC_SMPTE240M          CAVColorSpace = C.AVCOL_SPC_SMPTE240M ///< derived from 170M primaries and D65 white point, 170M is derived from BT470 System M's primaries
	AVCOL_SPC_YCGCO              CAVColorSpace = C.AVCOL_SPC_YCGCO     ///< used by Dirac / VC-2 and H.264 FRext, see ITU-T SG16
	AVCOL_SPC_YCOCG              CAVColorSpace = C.AVCOL_SPC_YCOCG
	AVCOL_SPC_BT2020_NCL         CAVColorSpace = C.AVCOL_SPC_BT2020_NCL         ///< ITU-R BT2020 non-constant luminance system
	AVCOL_SPC_BT2020_CL          CAVColorSpace = C.AVCOL_SPC_BT2020_CL          ///< ITU-R BT2020 constant luminance system
	AVCOL_SPC_SMPTE2085          CAVColorSpace = C.AVCOL_SPC_SMPTE2085          ///< SMPTE 2085, Y'D'zD'x
	AVCOL_SPC_CHROMA_DERIVED_NCL CAVColorSpace = C.AVCOL_SPC_CHROMA_DERIVED_NCL ///< Chromaticity-derived non-constant luminance system
	AVCOL_SPC_CHROMA_DERIVED_CL  CAVColorSpace = C.AVCOL_SPC_CHROMA_DERIVED_CL  ///< Chromaticity-derived constant luminance system
	AVCOL_SPC_ICTCP              CAVColorSpace = C.AVCOL_SPC_ICTCP              ///< ITU-R BT.2100-0, ICtCp
	AVCOL_SPC_NB                 CAVColorSpace = C.AVCOL_SPC_NB                 ///< Not part of ABI
)

/**
 * Visual content value range.
 *
 * These values are based on definitions that can be found in multiple
 * specifications, such as ITU-T BT.709 (3.4 - Quantization of RGB, luminance
 * and colour-difference signals), ITU-T BT.2020 (Table 5 - Digital
 * Representation) as well as ITU-T BT.2100 (Table 9 - Digital 10- and 12-bit
 * integer representation). At the time of writing, the BT.2100 one is
 * recommended, as it also defines the full range representation.
 *
 * Common definitions:
 *   - For RGB and luma planes such as Y in YCbCr and I in ICtCp,
 *     'E' is the original value in range of 0.0 to 1.0.
 *   - For chroma planes such as Cb,Cr and Ct,Cp, 'E' is the original
 *     value in range of -0.5 to 0.5.
 *   - 'n' is the output bit depth.
 *   - For additional definitions such as rounding and clipping to valid n
 *     bit unsigned integer range, please refer to BT.2100 (Table 9).
 */
type CAVColorRange C.enum_AVColorRange

const (
	AVCOL_RANGE_UNSPECIFIED CAVColorRange = C.AVCOL_RANGE_UNSPECIFIED

	/**
	 * Narrow or limited range content.
	 *
	 * - For luma planes:
	 *
	 *       (219 * E + 16) * 2^(n-8)
	 *
	 *   F.ex. the range of 16-235 for 8 bits
	 *
	 * - For chroma planes:
	 *
	 *       (224 * E + 128) * 2^(n-8)
	 *
	 *   F.ex. the range of 16-240 for 8 bits
	 */
	AVCOL_RANGE_MPEG CAVColorRange = C.AVCOL_RANGE_MPEG

	/**
	 * Full range content.
	 *
	 * - For RGB and luma planes:
	 *
	 *       (2^n - 1) * E
	 *
	 *   F.ex. the range of 0-255 for 8 bits
	 *
	 * - For chroma planes:
	 *
	 *       (2^n - 1) * E + 2^(n - 1)
	 *
	 *   F.ex. the range of 1-255 for 8 bits
	 */
	AVCOL_RANGE_JPEG CAVColorRange = C.AVCOL_RANGE_JPEG
	AVCOL_RANGE_NB   CAVColorRange = C.AVCOL_RANGE_NB ///< Not part of ABI
)

/**
 * Location of chroma samples.
 *
 * Illustration showing the location of the first (top left) chroma sample of the
 * image, the left shows only luma, the right
 * shows the location of the chroma sample, the 2 could be imagined to overlay
 * each other but are drawn separately due to limitations of ASCII
 *
 *                1st 2nd       1st 2nd horizontal luma sample positions
 *                 v   v         v   v
 *                 ______        ______
 *1st luma line > |X   X ...    |3 4 X ...     X are luma samples,
 *                |             |1 2           1-6 are possible chroma positions
 *2nd luma line > |X   X ...    |5 6 X ...     0 is undefined/unknown position
 */
type CAVChromaLocation C.enum_AVChromaLocation

const (
	AVCHROMA_LOC_UNSPECIFIED CAVChromaLocation = C.AVCHROMA_LOC_UNSPECIFIED
	AVCHROMA_LOC_LEFT        CAVChromaLocation = C.AVCHROMA_LOC_LEFT    ///< MPEG-2/4 4:2:0, H.264 default for 4:2:0
	AVCHROMA_LOC_CENTER      CAVChromaLocation = C.AVCHROMA_LOC_CENTER  ///< MPEG-1 4:2:0, JPEG 4:2:0, H.263 4:2:0
	AVCHROMA_LOC_TOPLEFT     CAVChromaLocation = C.AVCHROMA_LOC_TOPLEFT ///< ITU-R 601, SMPTE 274M 296M S314M(DV 4:1:1), mpeg2 4:2:2
	AVCHROMA_LOC_TOP         CAVChromaLocation = C.AVCHROMA_LOC_TOP
	AVCHROMA_LOC_BOTTOMLEFT  CAVChromaLocation = C.AVCHROMA_LOC_BOTTOMLEFT
	AVCHROMA_LOC_BOTTOM      CAVChromaLocation = C.AVCHROMA_LOC_BOTTOM
	AVCHROMA_LOC_NB          CAVChromaLocation = C.AVCHROMA_LOC_NB ///< Not part of ABI
)
