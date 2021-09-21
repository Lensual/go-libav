package avcodec

/*
#cgo LDFLAGS: -lavcodec

#include "libavcodec/avcodec.h"
*/
import "C"
import (
	"unsafe"

	"github.com/Lensual/go-libav/avutil"
)

/**
 * AVCodec.
 */
type CAVCodec C.AVCodec

func (cavcodec *CAVCodec) ID() uint32 {
	return cavcodec.id
}

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

/**
 * Supply raw packet data as input to a decoder.
 *
 * Internally, this call will copy relevant AVCodecContext fields, which can
 * influence decoding per-packet, and apply them when the packet is actually
 * decoded. (For example AVCodecContext.skip_frame, which might direct the
 * decoder to drop the frame contained by the packet sent with this function.)
 *
 * @warning The input buffer, avpkt->data must be AV_INPUT_BUFFER_PADDING_SIZE
 *          larger than the actual read bytes because some optimized bitstream
 *          readers read 32 or 64 bits at once and could read over the end.
 *
 * @note The AVCodecContext MUST have been opened with @ref avcodec_open2()
 *       before packets may be fed to the decoder.
 *
 * @param avctx codec context
 * @param[in] avpkt The input AVPacket. Usually, this will be a single video
 *                  frame, or several complete audio frames.
 *                  Ownership of the packet remains with the caller, and the
 *                  decoder will not write to the packet. The decoder may create
 *                  a reference to the packet data (or copy it if the packet is
 *                  not reference-counted).
 *                  Unlike with older APIs, the packet is always fully consumed,
 *                  and if it contains multiple frames (e.g. some audio codecs),
 *                  will require you to call avcodec_receive_frame() multiple
 *                  times afterwards before you can send a new packet.
 *                  It can be NULL (or an AVPacket with data set to NULL and
 *                  size set to 0); in this case, it is considered a flush
 *                  packet, which signals the end of the stream. Sending the
 *                  first flush packet will return success. Subsequent ones are
 *                  unnecessary and will return AVERROR_EOF. If the decoder
 *                  still has frames buffered, it will return them after sending
 *                  a flush packet.
 *
 * @return 0 on success, otherwise negative error code:
 *      AVERROR(EAGAIN):   input is not accepted in the current state - user
 *                         must read output with avcodec_receive_frame() (once
 *                         all output is read, the packet should be resent, and
 *                         the call will not fail with EAGAIN).
 *      AVERROR_EOF:       the decoder has been flushed, and no new packets can
 *                         be sent to it (also returned if more than 1 flush
 *                         packet is sent)
 *      AVERROR(EINVAL):   codec not opened, it is an encoder, or requires flush
 *      AVERROR(ENOMEM):   failed to add packet to internal queue, or similar
 *      other errors: legitimate decoding errors
 */
func AvcodecSendPacket(avctx *CAVCodecContext, avpkt *CAVPacket) int {
	return int(C.avcodec_send_packet((*C.AVCodecContext)(avctx), (*C.AVPacket)(avpkt)))
}

/**
 * Return decoded output data from a decoder.
 *
 * @param avctx codec context
 * @param frame This will be set to a reference-counted video or audio
 *              frame (depending on the decoder type) allocated by the
 *              decoder. Note that the function will always call
 *              av_frame_unref(frame) before doing anything else.
 *
 * @return
 *      0:                 success, a frame was returned
 *      AVERROR(EAGAIN):   output is not available in this state - user must try
 *                         to send new input
 *      AVERROR_EOF:       the decoder has been fully flushed, and there will be
 *                         no more output frames
 *      AVERROR(EINVAL):   codec not opened, or it is an encoder
 *      AVERROR_INPUT_CHANGED:   current decoded frame has changed parameters
 *                               with respect to first decoded frame. Applicable
 *                               when flag AV_CODEC_FLAG_DROPCHANGED is set.
 *      other negative values: legitimate decoding errors
 */
func AvcodecReceiveFrame(avctx *CAVCodecContext, frame *avutil.CAVFrame) int {
	return int(C.avcodec_receive_frame((*C.AVCodecContext)(avctx), (*C.AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Supply a raw video or audio frame to the encoder. Use avcodec_receive_packet()
 * to retrieve buffered output packets.
 *
 * @param avctx     codec context
 * @param[in] frame AVFrame containing the raw audio or video frame to be encoded.
 *                  Ownership of the frame remains with the caller, and the
 *                  encoder will not write to the frame. The encoder may create
 *                  a reference to the frame data (or copy it if the frame is
 *                  not reference-counted).
 *                  It can be NULL, in which case it is considered a flush
 *                  packet.  This signals the end of the stream. If the encoder
 *                  still has packets buffered, it will return them after this
 *                  call. Once flushing mode has been entered, additional flush
 *                  packets are ignored, and sending frames will return
 *                  AVERROR_EOF.
 *
 *                  For audio:
 *                  If AV_CODEC_CAP_VARIABLE_FRAME_SIZE is set, then each frame
 *                  can have any number of samples.
 *                  If it is not set, frame->nb_samples must be equal to
 *                  avctx->frame_size for all frames except the last.
 *                  The final frame may be smaller than avctx->frame_size.
 * @return 0 on success, otherwise negative error code:
 *      AVERROR(EAGAIN):   input is not accepted in the current state - user
 *                         must read output with avcodec_receive_packet() (once
 *                         all output is read, the packet should be resent, and
 *                         the call will not fail with EAGAIN).
 *      AVERROR_EOF:       the encoder has been flushed, and no new frames can
 *                         be sent to it
 *      AVERROR(EINVAL):   codec not opened, it is a decoder, or requires flush
 *      AVERROR(ENOMEM):   failed to add packet to internal queue, or similar
 *      other errors: legitimate encoding errors
 */
func avcodec_send_frame(avctx *CAVCodecContext, frame *avutil.CAVFrame) int {
	return int(C.avcodec_send_frame((*C.AVCodecContext)(avctx), (*C.AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Read encoded data from the encoder.
 *
 * @param avctx codec context
 * @param avpkt This will be set to a reference-counted packet allocated by the
 *              encoder. Note that the function will always call
 *              av_packet_unref(avpkt) before doing anything else.
 * @return 0 on success, otherwise negative error code:
 *      AVERROR(EAGAIN):   output is not available in the current state - user
 *                         must try to send input
 *      AVERROR_EOF:       the encoder has been fully flushed, and there will be
 *                         no more output packets
 *      AVERROR(EINVAL):   codec not opened, or it is a decoder
 *      other errors: legitimate encoding errors
 */
func avcodec_receive_packet(avctx *CAVCodecContext, avpkt *CAVPacket) int {
	return int(C.avcodec_receive_packet((*C.AVCodecContext)(avctx), (*C.AVPacket)(avpkt)))
}
