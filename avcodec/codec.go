package avcodec

/*
#cgo LDFLAGS: -lavcodec

#include "libavcodec/avcodec.h"
*/
import "C"
import "unsafe"

/**
 * AVCodec.
 */
type CAVCodec C.AVCodec

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
