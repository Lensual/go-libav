package avformat

/*
#cgo LDFLAGS: -lavformat

#include "libavformat/avformat.h"
*/
import "C"
import "unsafe"

type AVIOContext struct {
	CAVIOContext *C.AVIOContext
}

/*
buffer		Memory block for input/output operations via AVIOContext. The buffer must be allocated with av_malloc() and friends. It may be freed and replaced with a new buffer by libavformat. AVIOContext.buffer holds the buffer currently in use, which must be later freed with av_free().
buffer_size	The buffer size is very important for performance. For protocols with fixed blocksize it should be set to this blocksize. For others a typical size is a cache page, e.g. 4kb.
write_flag	Set to 1 if the buffer should be writable, 0 otherwise.
opaque		An opaque pointer to user-specific data.
read_packet	A function for refilling the buffer, may be NULL. For stream protocols, must never return 0 but rather a proper AVERROR code.
write_packet	A function for writing the buffer contents, may be NULL. The function may not change the input buffers content.
seek		A function for seeking to specified byte position, may be NULL.
*/
func AvioAllocContext(buffer unsafe.Pointer, buffer_size int, write_flag int,
	opaque unsafe.Pointer, read_packet unsafe.Pointer, write_packet unsafe.Pointer, seek unsafe.Pointer) *AVIOContext {

	return &AVIOContext{
		CAVIOContext: C.avio_alloc_context((*C.uchar)(buffer), C.int(buffer_size), C.int(write_flag),
			opaque, (*[0]byte)(read_packet), (*[0]byte)(write_packet), (*[0]byte)(seek)),
	}

}
