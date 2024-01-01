package avcodec

/*
#cgo pkg-config: libavcodec

#include "libavcodec/codec_desc.h"
*/
import "C"
import (
	"unsafe"

	"github.com/Lensual/go-libav/avutil"
)

/*
 * Codec descriptors public API
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

//  #ifndef AVCODEC_CODEC_DESC_H
//  #define AVCODEC_CODEC_DESC_H

//  #include "libavutil/avutil.h"

//  #include "codec_id.h"

/**
 * @addtogroup lavc_core
 * @{
 */

/**
 * This struct describes the properties of a single codec described by an
 * AVCodecID.
 * @see avcodec_descriptor_get()
 */
type CAVCodecDescriptor C.AVCodecDescriptor

// #region CAVCodecDescriptor

func (cd *CAVCodecDescriptor) GetId() CAVCodecID {
	return CAVCodecID(cd.id)
}
func (cd *CAVCodecDescriptor) SetId(id CAVCodecID) {
	cd.id = C.enum_AVCodecID(id)
}

func (cd *CAVCodecDescriptor) GetType() avutil.CAVMediaType {
	return avutil.CAVMediaType(cd._type)
}
func (cd *CAVCodecDescriptor) SetType(_type avutil.CAVMediaType) {
	cd._type = C.enum_AVMediaType(_type)
}

/**
 * Name of the codec described by this descriptor. It is non-empty and
 * unique for each codec descriptor. It should contain alphanumeric
 * characters and '_' only.
 */
func (cd *CAVCodecDescriptor) GetName() string {
	return C.GoString(cd.name)
}

/**
 * A more descriptive name for this codec. May be NULL.
 */
func (cd *CAVCodecDescriptor) GetLongName() string {
	return C.GoString(cd.long_name)
}

/**
 * Codec properties, a combination of AV_CODEC_PROP_* flags.
 */
func (cd *CAVCodecDescriptor) GetProps() int {
	return int(cd.props)
}

/**
 * Codec properties, a combination of AV_CODEC_PROP_* flags.
 */
func (cd *CAVCodecDescriptor) SetProps(props int) {
	cd.props = C.int(props)
}

/**
 * MIME type(s) associated with the codec.
 * May be NULL; if not, a NULL-terminated array of MIME types.
 * The first item is always non-NULL and is the preferred MIME type.
 */
func (cd *CAVCodecDescriptor) GetMimeTypes() string {
	return C.GoString(*cd.mime_types)
}

/**
 * If non-NULL, an array of profiles recognized for this codec.
 * Terminated with AV_PROFILE_UNKNOWN.
 */
func (cd *CAVCodecDescriptor) GetProfiles() *CAVProfile {
	return (*CAVProfile)(cd.profiles)
}

//#endregion CAVCodecDescriptor

const (
	/**
	 * Codec uses only intra compression.
	 * Video and audio codecs only.
	 */
	AV_CODEC_PROP_INTRA_ONLY = C.AV_CODEC_PROP_INTRA_ONLY
	/**
	 * Codec supports lossy compression. Audio and video codecs only.
	 * @note a codec may support both lossy and lossless
	 * compression modes
	 */
	AV_CODEC_PROP_LOSSY = C.AV_CODEC_PROP_LOSSY
	/**
	 * Codec supports lossless compression. Audio and video codecs only.
	 */
	AV_CODEC_PROP_LOSSLESS = C.AV_CODEC_PROP_LOSSLESS
	/**
	 * Codec supports frame reordering. That is, the coded order (the order in which
	 * the encoded packets are output by the encoders / stored / input to the
	 * decoders) may be different from the presentation order of the corresponding
	 * frames.
	 *
	 * For codecs that do not have this property set, PTS and DTS should always be
	 * equal.
	 */
	AV_CODEC_PROP_REORDER = C.AV_CODEC_PROP_REORDER

	/**
	 * Video codec supports separate coding of fields in interlaced frames.
	 */
	AV_CODEC_PROP_FIELDS = C.AV_CODEC_PROP_FIELDS

	/**
	 * Subtitle codec is bitmap based
	 * Decoded AVSubtitle data can be read from the AVSubtitleRect->pict field.
	 */
	AV_CODEC_PROP_BITMAP_SUB = C.AV_CODEC_PROP_BITMAP_SUB
	/**
	 * Subtitle codec is text based.
	 * Decoded AVSubtitle data can be read from the AVSubtitleRect->ass field.
	 */
	AV_CODEC_PROP_TEXT_SUB = C.AV_CODEC_PROP_TEXT_SUB
)

/**
 * @return descriptor for given codec ID or NULL if no descriptor exists.
 */
func AvcodecDescriptorGet(id CAVCodecID) *CAVCodecDescriptor {
	return (*CAVCodecDescriptor)(C.avcodec_descriptor_get(C.enum_AVCodecID(id)))
}

/**
 * Iterate over all codec descriptors known to libavcodec.
 *
 * @param prev previous descriptor. NULL to get the first descriptor.
 *
 * @return next descriptor or NULL after the last descriptor
 */
func AvcodecDescriptorNext(prev *CAVCodecDescriptor) *CAVCodecDescriptor {
	return (*CAVCodecDescriptor)(C.avcodec_descriptor_next((*C.AVCodecDescriptor)(prev)))
}

/**
 * @return codec descriptor with the given name or NULL if no such descriptor
 *         exists.
 */
func AvcodecDescriptorGetByName(name string) *CAVCodecDescriptor {
	var cName *C.char = nil
	if len(name) > 0 {
		cName = C.CString(name)
		defer C.free(unsafe.Pointer(cName))
	}

	return (*CAVCodecDescriptor)(C.avcodec_descriptor_get_by_name(cName))
}

/**
 * @}
 */

//  #endif // AVCODEC_CODEC_DESC_H
