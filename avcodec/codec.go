package avcodec

/*
#cgo LDFLAGS: -lavcodec

#include "libavcodec/avcodec.h"
*/
import "C"
import (
	"unsafe"
)

/**
 * AVCodec.
 */
type CAVCodec C.AVCodec

func (cavcodec *CAVCodec) ID() uint32 {
	return cavcodec.id
}

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
	 * Subtitle codec is bitmap based
	 * Decoded AVSubtitle data can be read from the AVSubtitleRect->pict field.
	 */
	AV_CODEC_PROP_BITMAP_SUB = C.AV_CODEC_PROP_BITMAP_SUB

	/**
	 * Subtitle codec is text based.
	 * Decoded AVSubtitle data can be read from the AVSubtitleRect->ass field.
	 */
	AV_CODEC_PROP_TEXT_SUB = C.AV_CODEC_PROP_TEXT_SUB

	/**
	 * @ingroup lavc_decoding
	 * Required number of additionally allocated bytes at the end of the input bitstream for decoding.
	 * This is mainly needed because some optimized bitstream readers read
	 * 32 or 64 bit at once and could read over the end.<br>
	 * Note: If the first 23 bits of the additional bytes are not 0, then damaged
	 * MPEG bitstreams could cause overread and segfault.
	 */
	AV_INPUT_BUFFER_PADDING_SIZE = C.AV_INPUT_BUFFER_PADDING_SIZE

	/**
	 * @ingroup lavc_encoding
	 * minimum encoding buffer size
	 * Used to avoid some checks during header writing.
	 */
	AV_INPUT_BUFFER_MIN_SIZE = C.AV_INPUT_BUFFER_MIN_SIZE
)

/*
Iterate over all registered codecs.

@param opaque a pointer where libavcodec will store the iteration state. Must
               point to NULL to start the iteration.

@return the next registered codec or NULL when the iteration is
         finished
*/
func AvCodecIterate(opaque *unsafe.Pointer) *CAVCodec {
	return (*CAVCodec)(C.av_codec_iterate(opaque))
}

/*
Find a registered decoder with a matching codec ID.

 @param id AVCodecID of the requested decoder
 @return A decoder if one was found, NULL otherwise.
*/
func AvcodecFindDecoder(id AVCodecID) *CAVCodec {
	return (*CAVCodec)(C.avcodec_find_decoder(uint32(id)))
}

/*
Find a registered decoder with the specified name.

 @param name name of the requested decoder
 @return A decoder if one was found, NULL otherwise.
*/
func AvcodecFindDecoderByName(name string) *CAVCodec {
	var cname *C.char = nil
	if len(name) > 0 {
		cname = C.CString(name)
		defer C.free(unsafe.Pointer(cname))
	}

	return (*CAVCodec)(C.avcodec_find_decoder_by_name(cname))
}

/*
Find a registered encoder with a matching codec ID.

 @param id AVCodecID of the requested encoder
 @return An encoder if one was found, NULL otherwise.
*/
func AvcodecFindEncoder(id AVCodecID) *CAVCodec {
	return (*CAVCodec)(C.avcodec_find_encoder(uint32(id)))
}

/*
Find a registered encoder with the specified name.

 @param name name of the requested encoder
 @return An encoder if one was found, NULL otherwise.
*/
func AvcodecFindEncoderByName(name string) *CAVCodec {
	var cname *C.char = nil
	if len(name) > 0 {
		cname = C.CString(name)
		defer C.free(unsafe.Pointer(cname))
	}

	return (*CAVCodec)(C.avcodec_find_encoder_by_name(cname))
}
