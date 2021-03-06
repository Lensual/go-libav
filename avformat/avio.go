package avformat

/*
#cgo LDFLAGS: -lavformat

#include "libavformat/avformat.h"
*/
import "C"
import "unsafe"

type CFunc unsafe.Pointer

/*
Bytestream IO Context.
New public fields can be added with minor version bumps.
Removal, reordering and changes to existing public fields require
a major version bump.
sizeof(AVIOContext) must not be used outside libav*.

@note None of the function pointers in AVIOContext should be called
      directly, they should only be set by the client application
      when implementing custom I/O. Normally these are set to the
      function pointers specified in avio_alloc_context()
*/
type CAVIOContext C.AVIOContext

func (avioCtx *CAVIOContext) GetCBuffer() unsafe.Pointer {
	return unsafe.Pointer(avioCtx.buffer)
}

/*
Allocate and initialize an AVIOContext for buffered I/O. It must be later
freed with avio_context_free().

@param buffer Memory block for input/output operations via AVIOContext.
       The buffer must be allocated with av_malloc() and friends.
       It may be freed and replaced with a new buffer by libavformat.
       AVIOContext.buffer holds the buffer currently in use,
       which must be later freed with av_free().
@param buffer_size The buffer size is very important for performance.
       For protocols with fixed blocksize it should be set to this blocksize.
       For others a typical size is a cache page, e.g. 4kb.
@param write_flag Set to 1 if the buffer should be writable, 0 otherwise.
@param opaque An opaque pointer to user-specific data.
@param read_packet  A function for refilling the buffer, may be NULL.
                    For stream protocols, must never return 0 but rather
                    a proper AVERROR code.
@param write_packet A function for writing the buffer contents, may be NULL.
       The function may not change the input buffers content.
@param seek A function for seeking to specified byte position, may be NULL.

@return Allocated AVIOContext or NULL on failure.
*/
func AvioAllocContext(buffer unsafe.Pointer, buffer_size int, write_flag int,
	opaque unsafe.Pointer, read_packet CFunc, write_packet CFunc, seek CFunc) *CAVIOContext {

	return (*CAVIOContext)(C.avio_alloc_context((*C.uchar)(buffer), C.int(buffer_size), C.int(write_flag),
		opaque, (*[0]byte)(read_packet), (*[0]byte)(write_packet), (*[0]byte)(seek)))
}

/*
Free the supplied IO context and everything associated with it.

@param s Double pointer to the IO context. This function will write NULL
into s.

!!!!!
NOTE: FOR GO BINDING, s MUST BE C POINTER, SO s MUST SET TO SINGLE POINTER
!!!!!
*/
func AvioContextFree(s *CAVIOContext) {
	C.avio_context_free((**C.AVIOContext)(unsafe.Pointer(&s)))
	s = nil
}
