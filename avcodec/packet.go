package avcodec

/*
#cgo LDFLAGS: -lavcodec

#include "libavcodec/avcodec.h"
*/
import "C"
import "unsafe"

type CAVPacket C.AVPacket

func (pkt *CAVPacket) GetData() unsafe.Pointer {
	return unsafe.Pointer(pkt.data)
}

func (pkt *CAVPacket) GetSize() int {
	return int(pkt.size)
}

/*
Allocate an AVPacket and set its fields to default values.  The resulting
struct must be freed using av_packet_free().

@return An AVPacket filled with default values or NULL on failure.

@note this only allocates the AVPacket itself, not the data buffers. Those
must be allocated through other means such as av_new_packet.

@see av_new_packet
*/
func AvPacketAlloc() *CAVPacket {
	return (*CAVPacket)(C.av_packet_alloc())
}

/*
 Create a new packet that references the same data as src.

 This is a shortcut for av_packet_alloc()+av_packet_ref().

 @return newly created AVPacket on success, NULL on error.

 @see av_packet_alloc
 @see av_packet_ref
*/
func AvPacketClone(src *CAVPacket) *CAVPacket {
	return (*CAVPacket)(C.av_packet_clone((*C.AVPacket)(src)))
}

/*
Free the packet, if the packet is reference counted, it will be
 unreferenced first.

 @param pkt packet to be freed. The pointer will be set to NULL.
 @note passing NULL is a no-op.
*/
func AvPacketFree(pkt **CAVPacket) {
	C.av_packet_free((**C.AVPacket)(unsafe.Pointer(pkt)))
}

/*
Allocate the payload of a packet and initialize its fields with
 default values.

 @param pkt packet
 @param size wanted payload size
 @return 0 if OK, AVERROR_xxx otherwise
*/
func AvNewPacket(pkt *CAVPacket, size int) int {
	return int(C.av_new_packet((*C.AVPacket)(pkt), C.int(size)))
}
