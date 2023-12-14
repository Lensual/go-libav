package avcodec

/*
#cgo pkg-config: libavcodec

#include "libavcodec/avcodec.h"
*/
import "C"
import (
	"unsafe"

	"github.com/Lensual/go-libav/avutil"
)

/**
 * @addtogroup lavc_core
 * @{
 */

const (
	/**
	 * Decoder can use draw_horiz_band callback.
	 */
	AV_CODEC_CAP_DRAW_HORIZ_BAND = C.AV_CODEC_CAP_DRAW_HORIZ_BAND
	/**
	 * Codec uses get_buffer() or get_encode_buffer() for allocating buffers and
	 * supports custom allocators.
	 * If not set, it might not use get_buffer() or get_encode_buffer() at all, or
	 * use operations that assume the buffer was allocated by
	 * avcodec_default_get_buffer2 or avcodec_default_get_encode_buffer.
	 */
	AV_CODEC_CAP_DR1 = C.AV_CODEC_CAP_DR1
	/**
	 * Encoder or decoder requires flushing with NULL input at the end in order to
	 * give the complete and correct output.
	 *
	 * NOTE: If this flag is not set, the codec is guaranteed to never be fed with
	 *       with NULL data. The user can still send NULL data to the public encode
	 *       or decode function, but libavcodec will not pass it along to the codec
	 *       unless this flag is set.
	 *
	 * Decoders:
	 * The decoder has a non-zero delay and needs to be fed with avpkt->data=NULL,
	 * avpkt->size=0 at the end to get the delayed data until the decoder no longer
	 * returns frames.
	 *
	 * Encoders:
	 * The encoder needs to be fed with NULL data at the end of encoding until the
	 * encoder no longer returns data.
	 *
	 * NOTE: For encoders implementing the AVCodec.encode2() function, setting this
	 *       flag also means that the encoder must set the pts and duration for
	 *       each output packet. If this flag is not set, the pts and duration will
	 *       be determined by libavcodec from the input frame.
	 */
	AV_CODEC_CAP_DELAY = C.AV_CODEC_CAP_DELAY
	/**
	 * Codec can be fed a final frame with a smaller size.
	 * This can be used to prevent truncation of the last audio samples.
	 */
	AV_CODEC_CAP_SMALL_LAST_FRAME = C.AV_CODEC_CAP_SMALL_LAST_FRAME

	//  #if FF_API_SUBFRAMES
	//  /**
	//   * Codec can output multiple frames per AVPacket
	//   * Normally demuxers return one frame at a time, demuxers which do not do
	//   * are connected to a parser to split what they return into proper frames.
	//   * This flag is reserved to the very rare category of codecs which have a
	//   * bitstream that cannot be split into frames without timeconsuming
	//   * operations like full decoding. Demuxers carrying such bitstreams thus
	//   * may return multiple frames in a packet. This has many disadvantages like
	//   * prohibiting stream copy in many cases thus it should only be considered
	//   * as a last resort.
	//   */
	//  #define AV_CODEC_CAP_SUBFRAMES           (1 <<  8)
	//  #endif

	/**
	 * Codec is experimental and is thus avoided in favor of non experimental
	 * encoders
	 */
	AV_CODEC_CAP_EXPERIMENTAL = C.AV_CODEC_CAP_EXPERIMENTAL
	/**
	 * Codec should fill in channel configuration and samplerate instead of container
	 */
	AV_CODEC_CAP_CHANNEL_CONF = C.AV_CODEC_CAP_CHANNEL_CONF
	/**
	 * Codec supports frame-level multithreading.
	 */
	AV_CODEC_CAP_FRAME_THREADS = C.AV_CODEC_CAP_FRAME_THREADS
	/**
	 * Codec supports slice-based (or partition-based) multithreading.
	 */
	AV_CODEC_CAP_SLICE_THREADS = C.AV_CODEC_CAP_SLICE_THREADS
	/**
	 * Codec supports changed parameters at any point.
	 */
	AV_CODEC_CAP_PARAM_CHANGE = C.AV_CODEC_CAP_PARAM_CHANGE
	/**
	 * Codec supports multithreading through a method other than slice- or
	 * frame-level multithreading. Typically this marks wrappers around
	 * multithreading-capable external libraries.
	 */
	AV_CODEC_CAP_OTHER_THREADS = C.AV_CODEC_CAP_OTHER_THREADS
	/**
	 * Audio encoder supports receiving a different number of samples in each call.
	 */
	AV_CODEC_CAP_VARIABLE_FRAME_SIZE = C.AV_CODEC_CAP_VARIABLE_FRAME_SIZE
	/**
	 * Decoder is not a preferred choice for probing.
	 * This indicates that the decoder is not a good choice for probing.
	 * It could for example be an expensive to spin up hardware decoder,
	 * or it could simply not provide a lot of useful information about
	 * the stream.
	 * A decoder marked with this flag should only be used as last resort
	 * choice for probing.
	 */
	AV_CODEC_CAP_AVOID_PROBING = C.AV_CODEC_CAP_AVOID_PROBING

	/**
	 * Codec is backed by a hardware implementation. Typically used to
	 * identify a non-hwaccel hardware decoder. For information about hwaccels, use
	 * avcodec_get_hw_config() instead.
	 */
	AV_CODEC_CAP_HARDWARE = C.AV_CODEC_CAP_HARDWARE

	/**
	 * Codec is potentially backed by a hardware implementation, but not
	 * necessarily. This is used instead of AV_CODEC_CAP_HARDWARE, if the
	 * implementation provides some sort of internal fallback.
	 */
	AV_CODEC_CAP_HYBRID = C.AV_CODEC_CAP_HYBRID

	/**
	 * This encoder can reorder user opaque values from input AVFrames and return
	 * them with corresponding output packets.
	 * @see AV_CODEC_FLAG_COPY_OPAQUE
	 */
	AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE = C.AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE

	/**
	 * This encoder can be flushed using avcodec_flush_buffers(). If this flag is
	 * not set, the encoder must be closed and reopened to ensure that no frames
	 * remain pending.
	 */
	AV_CODEC_CAP_ENCODER_FLUSH = C.AV_CODEC_CAP_ENCODER_FLUSH

	/**
	 * The encoder is able to output reconstructed frame data, i.e. raw frames that
	 * would be produced by decoding the encoded bitstream.
	 *
	 * Reconstructed frame output is enabled by the AV_CODEC_FLAG_RECON_FRAME flag.
	 */
	AV_CODEC_CAP_ENCODER_RECON_FRAME = C.AV_CODEC_CAP_ENCODER_RECON_FRAME
)

/**
 * AVProfile.
 */
type CAVProfile C.AVProfile

// #region CAVCodec
func (p *CAVProfile) GetProfile() int {
	return int(p.profile)
}

// /< short name for the profile
func (p *CAVProfile) GetName() string {
	return C.GoString(p.name)
}

//#endregion CAVCodec

/**
 * AVCodec.
 */
type CAVCodec C.AVCodec

//#region CAVCodec

/**
 * Name of the codec implementation.
 * The name is globally unique among encoders and among decoders (but an
 * encoder and a decoder can share the same name).
 * This is the primary way to find a codec from the user perspective.
 */
func (c *CAVCodec) GetName() string {
	return C.GoString(c.name)
}

/**
 * Descriptive name for the codec, meant to be more human readable than name.
 * You should use the NULL_IF_CONFIG_SMALL() macro to define it.
 */
func (c *CAVCodec) GetLongName() string {
	return C.GoString(c.long_name)
}

func (c *CAVCodec) GetType() avutil.CAVMediaType {
	return avutil.CAVMediaType(c._type)
}

func (c *CAVCodec) GetId() CAVCodecID {
	return CAVCodecID(c.id)
}

/**
 * Codec capabilities.
 * see AV_CODEC_CAP_*
 */
func (c *CAVCodec) GetCapabilities() int {
	return int(c.capabilities)
}

// /< maximum value for lowres supported by the decoder
func (c *CAVCodec) GetMaxLowres() uint8 {
	return uint8(c.max_lowres)
}

// /< array of supported framerates, or NULL if any, array is terminated by {0,0}
func (c *CAVCodec) GetSupportedFramerates() *avutil.CAVRational {
	return (*avutil.CAVRational)(unsafe.Pointer(c.supported_framerates))
}

// /< array of supported pixel formats, or NULL if unknown, array is terminated by -1
func (c *CAVCodec) GetPixFmts() *avutil.CAVPixelFormat {
	return (*avutil.CAVPixelFormat)(c.pix_fmts)
}

// /< array of supported audio samplerates, or NULL if unknown, array is terminated by 0
func (c *CAVCodec) GetSupportedSamplerates() *C.int {
	return (*C.int)(unsafe.Pointer(c.supported_samplerates))
}

// /< array of supported sample formats, or NULL if unknown, array is terminated by -1
func (c *CAVCodec) GetSampleFmts() *avutil.CAVSampleFormat {
	return (*avutil.CAVSampleFormat)(c.sample_fmts)
}

//  #if FF_API_OLD_CHANNEL_LAYOUT
// 	 /**
// 	  * @deprecated use ch_layouts instead
// 	  */
// 	 attribute_deprecated
// 	 const uint64_t *channel_layouts;         ///< array of support channel layouts, or NULL if unknown. array is terminated by 0
//  #endif

// /< AVClass for the private context
func (c *CAVCodec) GetPrivClass() *avutil.CAVClass {
	return (*avutil.CAVClass)(unsafe.Pointer(c.priv_class))
}

// /< array of recognized profiles, or NULL if unknown, array is terminated by {AV_PROFILE_UNKNOWN}
func (c *CAVCodec) GetProfiles() *C.AVProfile {
	return c.profiles
}

/**
 * Group name of the codec implementation.
 * This is a short symbolic name of the wrapper backing this codec. A
 * wrapper uses some kind of external implementation for the codec, such
 * as an external library, or a codec implementation provided by the OS or
 * the hardware.
 * If this field is NULL, this is a builtin, libavcodec native codec.
 * If non-NULL, this will be the suffix in AVCodec.name in most cases
 * (usually AVCodec.name will be of the form "<codec_name>_<wrapper_name>").
 */
func (c *CAVCodec) GetWrapperName() string {
	return C.GoString(c.wrapper_name)
}

/**
 * Array of supported channel layouts, terminated with a zeroed layout.
 */
func (c *CAVCodec) GetChLayouts() *avutil.CAVChannelLayout {
	return (*avutil.CAVChannelLayout)(unsafe.Pointer(c.ch_layouts))
}

//#endregion CAVCodec

//TODO recheck header
// const (
// 	/**
// 	 * Codec uses only intra compression.
// 	 * Video and audio codecs only.
// 	 */
// 	AV_CODEC_PROP_INTRA_ONLY = C.AV_CODEC_PROP_INTRA_ONLY

// 	/**
// 	 * Codec supports lossy compression. Audio and video codecs only.
// 	 * @note a codec may support both lossy and lossless
// 	 * compression modes
// 	 */
// 	AV_CODEC_PROP_LOSSY = C.AV_CODEC_PROP_LOSSY

// 	/**
// 	 * Codec supports lossless compression. Audio and video codecs only.
// 	 */
// 	AV_CODEC_PROP_LOSSLESS = C.AV_CODEC_PROP_LOSSLESS

// 	/**
// 	 * Codec supports frame reordering. That is, the coded order (the order in which
// 	 * the encoded packets are output by the encoders / stored / input to the
// 	 * decoders) may be different from the presentation order of the corresponding
// 	 * frames.
// 	 *
// 	 * For codecs that do not have this property set, PTS and DTS should always be
// 	 * equal.
// 	 */
// 	AV_CODEC_PROP_REORDER = C.AV_CODEC_PROP_REORDER

// 	/**
// 	 * Subtitle codec is bitmap based
// 	 * Decoded AVSubtitle data can be read from the AVSubtitleRect->pict field.
// 	 */
// 	AV_CODEC_PROP_BITMAP_SUB = C.AV_CODEC_PROP_BITMAP_SUB

// 	/**
// 	 * Subtitle codec is text based.
// 	 * Decoded AVSubtitle data can be read from the AVSubtitleRect->ass field.
// 	 */
// 	AV_CODEC_PROP_TEXT_SUB = C.AV_CODEC_PROP_TEXT_SUB

// 	/**
// 	 * @ingroup lavc_decoding
// 	 * Required number of additionally allocated bytes at the end of the input bitstream for decoding.
// 	 * This is mainly needed because some optimized bitstream readers read
// 	 * 32 or 64 bit at once and could read over the end.<br>
// 	 * Note: If the first 23 bits of the additional bytes are not 0, then damaged
// 	 * MPEG bitstreams could cause overread and segfault.
// 	 */
// 	AV_INPUT_BUFFER_PADDING_SIZE = C.AV_INPUT_BUFFER_PADDING_SIZE

// 	/**
// 	 * @ingroup lavc_encoding
// 	 * minimum encoding buffer size
// 	 * Used to avoid some checks during header writing.
// 	 */
// 	AV_INPUT_BUFFER_MIN_SIZE = C.AV_INPUT_BUFFER_MIN_SIZE
// )

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
func AvcodecFindDecoder(id CAVCodecID) *CAVCodec {
	return (*CAVCodec)(C.avcodec_find_decoder(uint32(id)))
}

/*
Find a registered decoder with the specified name.

	@param name name of the requested decoder
	@return A decoder if one was found, NULL otherwise.
*/
func AvcodecFindDecoderByName(name string) *CAVCodec {
	var cName *C.char = nil
	if len(name) > 0 {
		cName = C.CString(name)
		defer C.free(unsafe.Pointer(cName))
	}

	return (*CAVCodec)(C.avcodec_find_decoder_by_name(cName))
}

/*
Find a registered encoder with a matching codec ID.

	@param id AVCodecID of the requested encoder
	@return An encoder if one was found, NULL otherwise.
*/
func AvcodecFindEncoder(id CAVCodecID) *CAVCodec {
	return (*CAVCodec)(C.avcodec_find_encoder(uint32(id)))
}

/*
Find a registered encoder with the specified name.

	@param name name of the requested encoder
	@return An encoder if one was found, NULL otherwise.
*/
func AvcodecFindEncoderByName(name string) *CAVCodec {
	var cName *C.char = nil
	if len(name) > 0 {
		cName = C.CString(name)
		defer C.free(unsafe.Pointer(cName))
	}

	return (*CAVCodec)(C.avcodec_find_encoder_by_name(cName))
}
