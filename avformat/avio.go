package avformat

/*
#cgo pkg-config: libavformat

#include "libavformat/avformat.h"
*/
import "C"
import (
	"unsafe"

	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/ctypes"
)

type CFunc unsafe.Pointer

/*
 * copyright (c) 2001 Fabrice Bellard
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
//  #ifndef AVFORMAT_AVIO_H
//  #define AVFORMAT_AVIO_H

/**
 * @file
 * @ingroup lavf_io
 * Buffered I/O operations
 */

//  #include <stdint.h>
//  #include <stdio.h>

//  #include "libavutil/attributes.h"
//  #include "libavutil/dict.h"
//  #include "libavutil/log.h"

//  #include "libavformat/version_major.h"

/**
 * Seeking works like for a local file.
 */
const AVIO_SEEKABLE_NORMAL = C.AVIO_SEEKABLE_NORMAL

/**
 * Seeking by timestamp with avio_seek_time() is possible.
 */
const AVIO_SEEKABLE_TIME = C.AVIO_SEEKABLE_TIME

/**
 * Callback for checking whether to abort blocking functions.
 * AVERROR_EXIT is returned in this case by the interrupted
 * function. During blocking operations, callback is called with
 * opaque as parameter. If the callback returns 1, the
 * blocking operation will be aborted.
 *
 * No members can be added to this struct without a major bump, if
 * new elements have been added after this struct in AVFormatContext
 * or AVIOContext.
 */
type CAVIOInterruptCB C.AVIOInterruptCB

//#region CAVIOInterruptCB

// int (*callback)(void*)
func (intCb *CAVIOInterruptCB) GetCallback() CFunc {
	return CFunc(unsafe.Pointer(intCb.callback))
}

// int (*callback)(void*)
func (intCb *CAVIOInterruptCB) SetCallback(callback CFunc) {
	intCb.callback = (*[0]byte)(callback)
}

func (intCb *CAVIOInterruptCB) GetOpaque() unsafe.Pointer {
	return unsafe.Pointer(intCb.opaque)
}

func (intCb *CAVIOInterruptCB) SetOpaque(opaque unsafe.Pointer) {
	intCb.opaque = opaque
}

//#endregion CAVIOInterruptCB

/**
 * Directory entry types.
 */
type CAVIODirEntryType C.enum_AVIODirEntryType

const (
	AVIO_ENTRY_UNKNOWN          CAVIODirEntryType = C.AVIO_ENTRY_UNKNOWN
	AVIO_ENTRY_BLOCK_DEVICE     CAVIODirEntryType = C.AVIO_ENTRY_BLOCK_DEVICE
	AVIO_ENTRY_CHARACTER_DEVICE CAVIODirEntryType = C.AVIO_ENTRY_CHARACTER_DEVICE
	AVIO_ENTRY_DIRECTORY        CAVIODirEntryType = C.AVIO_ENTRY_DIRECTORY
	AVIO_ENTRY_NAMED_PIPE       CAVIODirEntryType = C.AVIO_ENTRY_NAMED_PIPE
	AVIO_ENTRY_SYMBOLIC_LINK    CAVIODirEntryType = C.AVIO_ENTRY_SYMBOLIC_LINK
	AVIO_ENTRY_SOCKET           CAVIODirEntryType = C.AVIO_ENTRY_SOCKET
	AVIO_ENTRY_FILE             CAVIODirEntryType = C.AVIO_ENTRY_FILE
	AVIO_ENTRY_SERVER           CAVIODirEntryType = C.AVIO_ENTRY_SERVER
	AVIO_ENTRY_SHARE            CAVIODirEntryType = C.AVIO_ENTRY_SHARE
	AVIO_ENTRY_WORKGROUP        CAVIODirEntryType = C.AVIO_ENTRY_WORKGROUP
)

/**
 * Describes single entry of the directory.
 *
 * Only name and type fields are guaranteed be set.
 * Rest of fields are protocol or/and platform dependent and might be unknown.
 */
type CAVIODirEntry C.AVIODirEntry

// #region CAVIODirEntry

/**< Filename */
func (e *CAVIODirEntry) GetName() string {
	return C.GoString(e.name)
}

/**< Type of the entry */
func (e *CAVIODirEntry) GetType() int {
	return int(e._type)
}

/*
*< Set to 1 when name is encoded with UTF-8, 0 otherwise.
Name can be encoded with UTF-8 even though 0 is set.
*/
func (e *CAVIODirEntry) GetUtf8() int {
	return int(e.utf8)
}

/**< File size in bytes, -1 if unknown. */
func (e *CAVIODirEntry) GetSize() int64 {
	return int64(e.size)
}

/*
*< Time of last modification in microseconds since unix
epoch, -1 if unknown.
*/
func (e *CAVIODirEntry) GetModificationTimestamp() int64 {
	return int64(e.modification_timestamp)
}

/*
*< Time of last access in microseconds since unix epoch,
-1 if unknown.
*/
func (e *CAVIODirEntry) GetAccessTimestamp() int64 {
	return int64(e.access_timestamp)
}

/*
*< Time of last status change in microseconds since unix
epoch, -1 if unknown.
*/
func (e *CAVIODirEntry) GetStatusChangeTimestamp() int64 {
	return int64(e.status_change_timestamp)
}

/**< User ID of owner, -1 if unknown. */
func (e *CAVIODirEntry) GetUserId() int64 {
	return int64(e.user_id)
}

/**< Group ID of owner, -1 if unknown. */
func (e *CAVIODirEntry) GetGroupId() int64 {
	return int64(e.group_id)
}

/**< Unix file mode, -1 if unknown. */
func (e *CAVIODirEntry) GetFilemode() int64 {
	return int64(e.filemode)
}

//#endregion CAVIODirEntry

// #if FF_API_AVIODIRCONTEXT
type CAVIODirContext C.AVIODirContext

// #region CAVIODirContext

func (dc *CAVIODirContext) GetUrlContext() *C.struct_URLContext {
	return dc.url_context
}

//#endregion CAVIODirContext

//  #else
//  typedef struct AVIODirContext AVIODirContext;
//  #endif

/**
 * Different data types that can be returned via the AVIO
 * write_data_type callback.
 */
type CAVIODataMarkerType = C.enum_AVIODataMarkerType

const (
	/**
	 * Header data; this needs to be present for the stream to be decodeable.
	 */
	AVIO_DATA_MARKER_HEADER CAVIODataMarkerType = C.AVIO_DATA_MARKER_HEADER
	/**
	 * A point in the output bytestream where a decoder can start decoding
	 * (i.e. a keyframe). A demuxer/decoder given the data flagged with
	 * AVIO_DATA_MARKER_HEADER, followed by any AVIO_DATA_MARKER_SYNC_POINT,
	 * should give decodeable results.
	 */
	AVIO_DATA_MARKER_SYNC_POINT CAVIODataMarkerType = C.AVIO_DATA_MARKER_SYNC_POINT
	/**
	 * A point in the output bytestream where a demuxer can start parsing
	 * (for non self synchronizing bytestream formats). That is, any
	 * non-keyframe packet start point.
	 */
	AVIO_DATA_MARKER_BOUNDARY_POINT CAVIODataMarkerType = C.AVIO_DATA_MARKER_BOUNDARY_POINT
	/**
	 * This is any, unlabelled data. It can either be a muxer not marking
	 * any positions at all, it can be an actual boundary/sync point
	 * that the muxer chooses not to mark, or a later part of a packet/fragment
	 * that is cut into multiple write callbacks due to limited IO buffer size.
	 */
	AVIO_DATA_MARKER_UNKNOWN CAVIODataMarkerType = C.AVIO_DATA_MARKER_UNKNOWN
	/**
	 * Trailer data, which doesn't contain actual content, but only for
	 * finalizing the output file.
	 */
	AVIO_DATA_MARKER_TRAILER CAVIODataMarkerType = C.AVIO_DATA_MARKER_TRAILER
	/**
	 * A point in the output bytestream where the underlying AVIOContext might
	 * flush the buffer depending on latency or buffering requirements. Typically
	 * means the end of a packet.
	 */
	AVIO_DATA_MARKER_FLUSH_POINT CAVIODataMarkerType = C.AVIO_DATA_MARKER_FLUSH_POINT
)

/**
 * Bytestream IO Context.
 * New public fields can be added with minor version bumps.
 * Removal, reordering and changes to existing public fields require
 * a major version bump.
 * sizeof(AVIOContext) must not be used outside libav*.
 *
 * @note None of the function pointers in AVIOContext should be called
 *       directly, they should only be set by the client application
 *       when implementing custom I/O. Normally these are set to the
 *       function pointers specified in avio_alloc_context()
 */
type CAVIOContext C.AVIOContext

//#region CAVIOContext

/**
 * A class for private options.
 *
 * If this AVIOContext is created by avio_open2(), av_class is set and
 * passes the options down to protocols.
 *
 * If this AVIOContext is manually allocated, then av_class may be set by
 * the caller.
 *
 * warning -- this field can be NULL, be sure to not pass this AVIOContext
 * to any av_opt_* functions in that case.
 */
func (avioCtx *CAVIOContext) GetAvClass() *avutil.CAVClass {
	return (*avutil.CAVClass)(unsafe.Pointer(avioCtx.av_class))
}

/*
 * The following shows the relationship between buffer, buf_ptr,
 * buf_ptr_max, buf_end, buf_size, and pos, when reading and when writing
 * (since AVIOContext is used for both):
 *
 **********************************************************************************
 *                                   READING
 **********************************************************************************
 *
 *                            |              buffer_size              |
 *                            |---------------------------------------|
 *                            |                                       |
 *
 *                         buffer          buf_ptr       buf_end
 *                            +---------------+-----------------------+
 *                            |/ / / / / / / /|/ / / / / / /|         |
 *  read buffer:              |/ / consumed / | to be read /|         |
 *                            |/ / / / / / / /|/ / / / / / /|         |
 *                            +---------------+-----------------------+
 *
 *                                                         pos
 *              +-------------------------------------------+-----------------+
 *  input file: |                                           |                 |
 *              +-------------------------------------------+-----------------+
 *
 *
 **********************************************************************************
 *                                   WRITING
 **********************************************************************************
 *
 *                             |          buffer_size                 |
 *                             |--------------------------------------|
 *                             |                                      |
 *
 *                                                buf_ptr_max
 *                          buffer                 (buf_ptr)       buf_end
 *                             +-----------------------+--------------+
 *                             |/ / / / / / / / / / / /|              |
 *  write buffer:              | / / to be flushed / / |              |
 *                             |/ / / / / / / / / / / /|              |
 *                             +-----------------------+--------------+
 *                               buf_ptr can be in this
 *                               due to a backward seek
 *
 *                            pos
 *               +-------------+----------------------------------------------+
 *  output file: |             |                                              |
 *               +-------------+----------------------------------------------+
 *
 */
/**< Start of the buffer. */
func (avioCtx *CAVIOContext) GetBuffer() unsafe.Pointer {
	return unsafe.Pointer(avioCtx.buffer)
}

/**< Maximum buffer size */
func (avioCtx *CAVIOContext) GetBufferSize() int {
	return int(avioCtx.buffer_size)
}

/**< Current position in the buffer */
func (avioCtx *CAVIOContext) GetBufPtr() unsafe.Pointer {
	return unsafe.Pointer(avioCtx.buf_ptr)
}

/*
*< End of the data, may be less than

	buffer+buffer_size if the read function returned
	less data than requested, e.g. for streams where
	no more data has been received yet.
*/
func (avioCtx *CAVIOContext) GetBufEnd() unsafe.Pointer {
	return unsafe.Pointer(avioCtx.buf_end)
}

/*
*< A private pointer, passed to the read/write/seek/...

	functions.
*/
func (avioCtx *CAVIOContext) GetOpaque() unsafe.Pointer {
	return unsafe.Pointer(avioCtx.opaque)
}

// int (*read_packet)(void *opaque, uint8_t *buf, int buf_size);
func (avioCtx *CAVIOContext) GetReadPacket() CFunc {
	return CFunc(avioCtx.read_packet)
}

//  #if FF_API_AVIO_WRITE_NONCONST

// int (*write_packet)(void *opaque, uint8_t *buf, int buf_size);
func (avioCtx *CAVIOContext) GetWritePacket() CFunc {
	return CFunc(avioCtx.write_packet)
}

//  #else
// 	 int (*write_packet)(void *opaque, const uint8_t *buf, int buf_size);
//  #endif

// int64_t (*seek)(void *opaque, int64_t offset, int whence);
func (avioCtx *CAVIOContext) GetSeek() CFunc {
	return CFunc(avioCtx.seek)
}

/**< position in the file of the current buffer */
func (avioCtx *CAVIOContext) GetPos() int64 {
	return int64(avioCtx.pos)
}

/**< true if was unable to read due to error or eof */
func (avioCtx *CAVIOContext) GetEofReached() int {
	return int(avioCtx.eof_reached)
}

/**< contains the error code or 0 if no error happened */
func (avioCtx *CAVIOContext) GetError() int {
	return int(avioCtx.error)
}

/**< true if open for writing */
func (avioCtx *CAVIOContext) GetWriteFlag() int {
	return int(avioCtx.write_flag)
}

func (avioCtx *CAVIOContext) GetMaxPacketSize() int {
	return int(avioCtx.max_packet_size)
}

/*
*< Try to buffer at least this amount of data
before flushing it.
*/
func (avioCtx *CAVIOContext) GetMinPacketSize() int {
	return int(avioCtx.min_packet_size)
}

func (avioCtx *CAVIOContext) GetCheckSum() uint32 {
	return uint32(avioCtx.checksum)
}

func (avioCtx *CAVIOContext) GetCheckSumPtr() unsafe.Pointer {
	return unsafe.Pointer(avioCtx.checksum_ptr)
}

// unsigned long (*update_checksum)(unsigned long checksum, const uint8_t *buf, unsigned int size);
func (avioCtx *CAVIOContext) GetUpdateCheckSum() CFunc {
	return CFunc(avioCtx.update_checksum)
}

/**
 * Pause or resume playback for network streaming protocols - e.g. MMS.
 */
// int (*read_pause)(void *opaque, int pause);
func (avioCtx *CAVIOContext) GetReadPause() CFunc {
	return CFunc(avioCtx.read_pause)
}

/**
 * Seek to a given timestamp in stream with the specified stream_index.
 * Needed for some network streaming protocols which don't support seeking
 * to byte position.
 */
// int64_t (*read_seek)(void *opaque, int stream_index, int64_t timestamp, int flags);
func (avioCtx *CAVIOContext) GetReadSeek() CFunc {
	return CFunc(avioCtx.read_seek)
}

/**
 * A combination of AVIO_SEEKABLE_ flags or 0 when the stream is not seekable.
 */
func (avioCtx *CAVIOContext) GetSeekable() int {
	return int(avioCtx.seekable)
}

/**
 * avio_read and avio_write should if possible be satisfied directly
 * instead of going through a buffer, and avio_seek will always
 * call the underlying seek function directly.
 */
func (avioCtx *CAVIOContext) GetDirect() int {
	return int(avioCtx.direct)
}

/**
 * ',' separated list of allowed protocols.
 */
func (avioCtx *CAVIOContext) GetProtocolWhitelist() string {
	return C.GoString(avioCtx.protocol_whitelist)
}

/**
 * ',' separated list of disallowed protocols.
 */
func (avioCtx *CAVIOContext) GetProtocolBlacklist() string {
	return C.GoString(avioCtx.protocol_blacklist)
}

/**
 * A callback that is used instead of write_packet.
 */
//  #if FF_API_AVIO_WRITE_NONCONST

//	 int (*write_data_type)(void *opaque, uint8_t *buf, int buf_size,
//							enum AVIODataMarkerType type, int64_t time);
func (avioCtx *CAVIOContext) GetWriteDataType() CFunc {
	return CFunc(avioCtx.write_data_type)
}

//  #else
// 	 int (*write_data_type)(void *opaque, const uint8_t *buf, int buf_size,
// 							enum AVIODataMarkerType type, int64_t time);
//  #endif

/**
 * If set, don't call write_data_type separately for AVIO_DATA_MARKER_BOUNDARY_POINT,
 * but ignore them and treat them as AVIO_DATA_MARKER_UNKNOWN (to avoid needlessly
 * small chunks of data returned from the callback).
 */
func (avioCtx *CAVIOContext) GetIgnoreBoundaryPoint() int {
	return int(avioCtx.ignore_boundary_point)
}

/**
 * Maximum reached position before a backward seek in the write buffer,
 * used keeping track of already written data for a later flush.
 */
func (avioCtx *CAVIOContext) GetBufPtrMax() unsafe.Pointer {
	return unsafe.Pointer(avioCtx.buf_ptr_max)
}

/**
 * Read-only statistic of bytes read for this AVIOContext.
 */
func (avioCtx *CAVIOContext) GetBytesRead() int64 {
	return int64(avioCtx.bytes_read)
}

/**
 * Read-only statistic of bytes written for this AVIOContext.
 */
func (avioCtx *CAVIOContext) GetBytesWritten() int64 {
	return int64(avioCtx.bytes_written)
}

//#endregion CAVIOContext

/**
 * Return the name of the protocol that will handle the passed URL.
 *
 * NULL is returned if no protocol could be found for the given URL.
 *
 * @return Name of the protocol or NULL.
 */
func AvioFindProtocolName(url string) string {
	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
		defer C.free(unsafe.Pointer(cUrl))
	}

	return C.GoString(C.avio_find_protocol_name(cUrl))
}

/**
 * Return AVIO_FLAG_* access flags corresponding to the access permissions
 * of the resource in url, or a negative value corresponding to an
 * AVERROR code in case of failure. The returned access flags are
 * masked by the value in flags.
 *
 * @note This function is intrinsically unsafe, in the sense that the
 * checked resource may change its existence or permission status from
 * one call to another. Thus you should not trust the returned value,
 * unless you are sure that no other processes are accessing the
 * checked resource.
 */
func AvioCheck(url string, flags int) int {
	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
		defer C.free(unsafe.Pointer(cUrl))
	}

	return int(C.avio_check(cUrl, C.int(flags)))
}

/**
 * Open directory for reading.
 *
 * @param s       directory read context. Pointer to a NULL pointer must be passed.
 * @param url     directory to be listed.
 * @param options A dictionary filled with protocol-private options. On return
 *                this parameter will be destroyed and replaced with a dictionary
 *                containing options that were not found. May be NULL.
 * @return >=0 on success or negative on error.
 */
func AvioOpenDir(s **CAVIODirContext, url string, options **avutil.CAVDictionary) int {
	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
		defer C.free(unsafe.Pointer(cUrl))
	}

	return int(C.avio_open_dir((**C.AVIODirContext)(unsafe.Pointer(s)), cUrl, (**C.AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Get next directory entry.
 *
 * Returned entry must be freed with avio_free_directory_entry(). In particular
 * it may outlive AVIODirContext.
 *
 * @param s         directory read context.
 * @param[out] next next entry or NULL when no more entries.
 * @return >=0 on success or negative on error. End of list is not considered an
 *             error.
 */
func AvioReadDir(s *CAVIODirContext, next **CAVIODirEntry) int {
	return int(C.avio_read_dir((*C.AVIODirContext)(s), (**C.AVIODirEntry)(unsafe.Pointer(next))))
}

/**
 * Close directory.
 *
 * @note Entries created using avio_read_dir() are not deleted and must be
 * freeded with avio_free_directory_entry().
 *
 * @param s         directory read context.
 * @return >=0 on success or negative on error.
 */
func AvioCloseDir(s **CAVIODirContext) int {
	return int(C.avio_close_dir((**C.AVIODirContext)(unsafe.Pointer(s))))
}

/**
 * Free entry allocated by avio_read_dir().
 *
 * @param entry entry to be freed.
 */
func AvioFreeDirectoryEntry(s **CAVIODirEntry) {
	C.avio_free_directory_entry((**C.AVIODirEntry)(unsafe.Pointer(s)))
}

/**
 * Allocate and initialize an AVIOContext for buffered I/O. It must be later
 * freed with avio_context_free().
 *
 * @param buffer Memory block for input/output operations via AVIOContext.
 *        The buffer must be allocated with av_malloc() and friends.
 *        It may be freed and replaced with a new buffer by libavformat.
 *        AVIOContext.buffer holds the buffer currently in use,
 *        which must be later freed with av_free().
 * @param buffer_size The buffer size is very important for performance.
 *        For protocols with fixed blocksize it should be set to this blocksize.
 *        For others a typical size is a cache page, e.g. 4kb.
 * @param write_flag Set to 1 if the buffer should be writable, 0 otherwise.
 * @param opaque An opaque pointer to user-specific data.
 * @param read_packet  A function for refilling the buffer, may be NULL.
 *                     For stream protocols, must never return 0 but rather
 *                     a proper AVERROR code.
 * @param write_packet A function for writing the buffer contents, may be NULL.
 *        The function may not change the input buffers content.
 * @param seek A function for seeking to specified byte position, may be NULL.
 *
 * @return Allocated AVIOContext or NULL on failure.
 */
func AvioAllocContext(buffer unsafe.Pointer, bufferSize int, writeFlag int,
	opaque unsafe.Pointer, readPacket CFunc, writePacket CFunc, seek CFunc) *CAVIOContext {

	return (*CAVIOContext)(C.avio_alloc_context((*C.uchar)(buffer), C.int(bufferSize), C.int(writeFlag),
		opaque, (*[0]byte)(readPacket), (*[0]byte)(writePacket), (*[0]byte)(seek)))
}

/**
 * Free the supplied IO context and everything associated with it.
 *
 * @param s Double pointer to the IO context. This function will write NULL
 * into s.
 */
func AvioContextFree(s **CAVIOContext) {
	C.avio_context_free((**C.AVIOContext)(unsafe.Pointer(s)))
}

func AvioW8(s *CAVIOContext, b int) {
	C.avio_w8((*C.AVIOContext)(s), C.int(b))
}
func AvioWrite(s *CAVIOContext, buf unsafe.Pointer, size int) {
	C.avio_write((*C.AVIOContext)(s), (*C.uchar)(buf), C.int(size))
}
func AvioWl64(s *CAVIOContext, val uint64) {
	C.avio_wl64((*C.AVIOContext)(s), C.uint64_t(val))
}
func AvioWb64(s *CAVIOContext, val uint64) {
	C.avio_wb64((*C.AVIOContext)(s), C.uint64_t(val))
}
func AvioWl32(s *CAVIOContext, val uint) {
	C.avio_wl32((*C.AVIOContext)(s), C.uint(val))
}
func AvioWb32(s *CAVIOContext, val uint) {
	C.avio_wb32((*C.AVIOContext)(s), C.uint(val))
}
func AvioWl24(s *CAVIOContext, val uint) {
	C.avio_wl24((*C.AVIOContext)(s), C.uint(val))
}
func AvioWb24(s *CAVIOContext, val uint) {
	C.avio_wb24((*C.AVIOContext)(s), C.uint(val))
}
func AvioWl16(s *CAVIOContext, val uint) {
	C.avio_wb16((*C.AVIOContext)(s), C.uint(val))
}
func AvioWb16(s *CAVIOContext, val uint) {
	C.avio_wb24((*C.AVIOContext)(s), C.uint(val))
}

/**
 * Write a NULL-terminated string.
 * @return number of bytes written.
 */
func AvioPutStr(s *CAVIOContext, str string) int {
	var cStr *C.char = nil
	if len(str) > 0 {
		cStr = C.CString(str)
		defer C.free(unsafe.Pointer(cStr))
	}

	return int(C.avio_put_str((*C.AVIOContext)(s), cStr))
}

/**
 * Convert an UTF-8 string to UTF-16LE and write it.
 * @param s the AVIOContext
 * @param str NULL-terminated UTF-8 string
 *
 * @return number of bytes written.
 */
func AvioPutStr16le(s *CAVIOContext, str string) int {
	var cStr *C.char = nil
	if len(str) > 0 {
		cStr = C.CString(str)
		defer C.free(unsafe.Pointer(cStr))
	}

	return int(C.avio_put_str16le((*C.AVIOContext)(s), cStr))
}

/**
 * Convert an UTF-8 string to UTF-16BE and write it.
 * @param s the AVIOContext
 * @param str NULL-terminated UTF-8 string
 *
 * @return number of bytes written.
 */
func AvioPutStr16be(s *CAVIOContext, str string) int {
	var cStr *C.char = nil
	if len(str) > 0 {
		cStr = C.CString(str)
		defer C.free(unsafe.Pointer(cStr))
	}

	return int(C.avio_put_str16be((*C.AVIOContext)(s), cStr))
}

/**
 * Mark the written bytestream as a specific type.
 *
 * Zero-length ranges are omitted from the output.
 *
 * @param s    the AVIOContext
 * @param time the stream time the current bytestream pos corresponds to
 *             (in AV_TIME_BASE units), or AV_NOPTS_VALUE if unknown or not
 *             applicable
 * @param type the kind of data written starting at the current pos
 */
func AvioWriteMarker(s *CAVIOContext, time int64, _type CAVIODataMarkerType) {
	C.avio_write_marker((*C.AVIOContext)(s), C.int64_t(time), _type)
}

/**
 * ORing this as the "whence" parameter to a seek function causes it to
 * return the filesize without seeking anywhere. Supporting this is optional.
 * If it is not supported then the seek function will return <0.
 */
const AVSEEK_SIZE = C.AVSEEK_SIZE

/**
 * Passing this flag as the "whence" parameter to a seek function causes it to
 * seek by any means (like reopening and linear reading) or other normally unreasonable
 * means that can be extremely slow.
 * This may be ignored by the seek code.
 */
const AVSEEK_FORCE = C.AVSEEK_FORCE

/**
 * fseek() equivalent for AVIOContext.
 * @return new position or AVERROR.
 */
func AvioSeek(s *CAVIOContext, offset int64, whence int) int64 {
	return int64(C.avio_seek((*C.AVIOContext)(s), C.int64_t(offset), C.int(whence)))
}

/**
 * Skip given number of bytes forward
 * @return new position or AVERROR.
 */
func AvioSkip(s *CAVIOContext, offset int64) int64 {
	return int64(C.avio_skip((*C.AVIOContext)(s), C.int64_t(offset)))
}

/**
 * ftell() equivalent for AVIOContext.
 * @return position or AVERROR.
 */
func AvioTell(s *CAVIOContext) int64 {
	return int64(C.avio_tell((*C.AVIOContext)(s)))
}

/**
 * Get the filesize.
 * @return filesize or AVERROR
 */
func AvioSize(s *CAVIOContext) int64 {
	return int64(C.avio_size((*C.AVIOContext)(s)))
}

/**
 * Similar to feof() but also returns nonzero on read errors.
 * @return non zero if and only if at end of file or a read error happened when reading.
 */
func AvioFeof(s *CAVIOContext) int {
	return int(C.avio_feof((*C.AVIOContext)(s)))
}

/**
 * Writes a formatted string to the context taking a va_list.
 * @return number of bytes written, < 0 on error.
 */
func AvioVprintf(s *CAVIOContext, fmt string, ap *C.struct___va_list_tag) int {
	var cFmt *C.char = nil
	if len(fmt) > 0 {
		cFmt = C.CString(fmt)
		defer C.free(unsafe.Pointer(cFmt))
	}

	return int(C.avio_vprintf((*C.AVIOContext)(s), cFmt, ap))
}

//  /**
//   * Writes a formatted string to the context.
//   * @return number of bytes written, < 0 on error.
//   */
//  int avio_printf(AVIOContext *s, const char *fmt, ...) av_printf_format(2, 3);

//  /**
//   * Write a NULL terminated array of strings to the context.
//   * Usually you don't need to use this function directly but its macro wrapper,
//   * avio_print.
//   */
//  void avio_print_string_array(AVIOContext *s, const char *strings[]);

//  /**
//   * Write strings (const char *) to the context.
//   * This is a convenience macro around avio_print_string_array and it
//   * automatically creates the string array from the variable argument list.
//   * For simple string concatenations this function is more performant than using
//   * avio_printf since it does not need a temporary buffer.
//   */
//  #define avio_print(s, ...) \
// 	 avio_print_string_array(s, (const char*[]){__VA_ARGS__, NULL})

/**
 * Force flushing of buffered data.
 *
 * For write streams, force the buffered data to be immediately written to the output,
 * without to wait to fill the internal buffer.
 *
 * For read streams, discard all currently buffered data, and advance the
 * reported file position to that of the underlying stream. This does not
 * read new data, and does not perform any seeks.
 */
func AvioFlush(s *CAVIOContext) {
	C.avio_flush((*C.AVIOContext)(s))
}

/**
 * Read size bytes from AVIOContext into buf.
 * @return number of bytes read or AVERROR
 */
func AvioRead(s *CAVIOContext, buf unsafe.Pointer, size int) int {
	return int(C.avio_read((*C.AVIOContext)(s), (*C.uchar)(buf), C.int(size)))
}

/**
 * Read size bytes from AVIOContext into buf. Unlike avio_read(), this is allowed
 * to read fewer bytes than requested. The missing bytes can be read in the next
 * call. This always tries to read at least 1 byte.
 * Useful to reduce latency in certain cases.
 * @return number of bytes read or AVERROR
 */
func AvioReadPartial(s *CAVIOContext, buf unsafe.Pointer, size int) int {
	return int(C.avio_read_partial((*C.AVIOContext)(s), (*C.uchar)(buf), C.int(size)))
}

/**
 * @name Functions for reading from AVIOContext
 * @{
 *
 * @note return 0 if EOF, so you cannot use it if EOF handling is
 *       necessary
 */

func AvioR8(s *CAVIOContext) int {
	return int(C.avio_r8((*C.AVIOContext)(s)))
}
func AvioRl16(s *CAVIOContext) uint {
	return uint(C.avio_rl16((*C.AVIOContext)(s)))
}
func AvioRl24(s *CAVIOContext) uint {
	return uint(C.avio_rl24((*C.AVIOContext)(s)))
}
func AvioRl32(s *CAVIOContext) uint {
	return uint(C.avio_rl32((*C.AVIOContext)(s)))
}
func AvioRl64(s *CAVIOContext) uint64 {
	return uint64(C.avio_rl64((*C.AVIOContext)(s)))
}
func AvioRb16(s *CAVIOContext) uint {
	return uint(C.avio_rb16((*C.AVIOContext)(s)))
}
func AvioRb24(s *CAVIOContext) uint {
	return uint(C.avio_rb24((*C.AVIOContext)(s)))
}
func AvioRb32(s *CAVIOContext) uint {
	return uint(C.avio_rb32((*C.AVIOContext)(s)))
}
func AvioRb64(s *CAVIOContext) uint64 {
	return uint64(C.avio_rb64((*C.AVIOContext)(s)))
}

/**
 * @}
 */

/**
 * Read a string from pb into buf. The reading will terminate when either
 * a NULL character was encountered, maxlen bytes have been read, or nothing
 * more can be read from pb. The result is guaranteed to be NULL-terminated, it
 * will be truncated if buf is too small.
 * Note that the string is not interpreted or validated in any way, it
 * might get truncated in the middle of a sequence for multi-byte encodings.
 *
 * @return number of bytes read (is always <= maxlen).
 * If reading ends on EOF or error, the return value will be one more than
 * bytes actually read.
 */
func AvioGetStr(pb *CAVIOContext, maxlen int, buf unsafe.Pointer, buflen int) int {
	return int(C.avio_get_str((*C.AVIOContext)(pb), C.int(maxlen), (*C.char)(buf), C.int(buflen)))
}

/**
 * Read a UTF-16 string from pb and convert it to UTF-8.
 * The reading will terminate when either a null or invalid character was
 * encountered or maxlen bytes have been read.
 * @return number of bytes read (is always <= maxlen)
 */
func AvioGetStr16le(pb *CAVIOContext, maxlen int, buf unsafe.Pointer, buflen int) int {
	return int(C.avio_get_str16le((*C.AVIOContext)(pb), C.int(maxlen), (*C.char)(buf), C.int(buflen)))
}
func AvioGetStr16be(pb *CAVIOContext, maxlen int, buf unsafe.Pointer, buflen int) int {
	return int(C.avio_get_str16be((*C.AVIOContext)(pb), C.int(maxlen), (*C.char)(buf), C.int(buflen)))
}

/**
 * @name URL open modes
 * The flags argument to avio_open must be one of the following
 * constants, optionally ORed with other flags.
 * @{
 */
const (
	AVIO_FLAG_READ       = C.AVIO_FLAG_READ       /**< read-only */
	AVIO_FLAG_WRITE      = C.AVIO_FLAG_WRITE      /**< write-only */
	AVIO_FLAG_READ_WRITE = C.AVIO_FLAG_READ_WRITE /**< read-write pseudo flag */
)

/**
 * @}
 */

/**
 * Use non-blocking mode.
 * If this flag is set, operations on the context will return
 * AVERROR(EAGAIN) if they can not be performed immediately.
 * If this flag is not set, operations on the context will never return
 * AVERROR(EAGAIN).
 * Note that this flag does not affect the opening/connecting of the
 * context. Connecting a protocol will always block if necessary (e.g. on
 * network protocols) but never hang (e.g. on busy devices).
 * Warning: non-blocking protocols is work-in-progress; this flag may be
 * silently ignored.
 */
const AVIO_FLAG_NONBLOCK = C.AVIO_FLAG_NONBLOCK

/**
 * Use direct mode.
 * avio_read and avio_write should if possible be satisfied directly
 * instead of going through a buffer, and avio_seek will always
 * call the underlying seek function directly.
 */
const AVIO_FLAG_DIRECT = C.AVIO_FLAG_DIRECT

/**
 * Create and initialize a AVIOContext for accessing the
 * resource indicated by url.
 * @note When the resource indicated by url has been opened in
 * read+write mode, the AVIOContext can be used only for writing.
 *
 * @param s Used to return the pointer to the created AVIOContext.
 * In case of failure the pointed to value is set to NULL.
 * @param url resource to access
 * @param flags flags which control how the resource indicated by url
 * is to be opened
 * @return >= 0 in case of success, a negative value corresponding to an
 * AVERROR code in case of failure
 */
func AvioOpen(s **CAVIOContext, url string, flags int) int {
	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
		defer C.free(unsafe.Pointer(cUrl))
	}

	return int(C.avio_open((**C.AVIOContext)(unsafe.Pointer(s)), cUrl, C.int(flags)))
}

/**
 * Create and initialize a AVIOContext for accessing the
 * resource indicated by url.
 * @note When the resource indicated by url has been opened in
 * read+write mode, the AVIOContext can be used only for writing.
 *
 * @param s Used to return the pointer to the created AVIOContext.
 * In case of failure the pointed to value is set to NULL.
 * @param url resource to access
 * @param flags flags which control how the resource indicated by url
 * is to be opened
 * @param int_cb an interrupt callback to be used at the protocols level
 * @param options  A dictionary filled with protocol-private options. On return
 * this parameter will be destroyed and replaced with a dict containing options
 * that were not found. May be NULL.
 * @return >= 0 in case of success, a negative value corresponding to an
 * AVERROR code in case of failure
 */
func AvioOpen2(s **CAVIOContext, url string, flags int, intCb *CAVIOInterruptCB, options **avutil.CAVDictionary) int {
	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
		defer C.free(unsafe.Pointer(cUrl))
	}

	return int(C.avio_open2((**C.AVIOContext)(unsafe.Pointer(s)), cUrl, C.int(flags), (*C.AVIOInterruptCB)(intCb), (**C.AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Close the resource accessed by the AVIOContext s and free it.
 * This function can only be used if s was opened by avio_open().
 *
 * The internal buffer is automatically flushed before closing the
 * resource.
 *
 * @return 0 on success, an AVERROR < 0 on error.
 * @see avio_closep
 */
func AvioClose(s *CAVIOContext) int {
	return int(C.avio_close((*C.AVIOContext)(s)))
}

/**
 * Close the resource accessed by the AVIOContext *s, free it
 * and set the pointer pointing to it to NULL.
 * This function can only be used if s was opened by avio_open().
 *
 * The internal buffer is automatically flushed before closing the
 * resource.
 *
 * @return 0 on success, an AVERROR < 0 on error.
 * @see avio_close
 */
func AvioClosep(s **CAVIOContext) int {
	return int(C.avio_closep((**C.AVIOContext)(unsafe.Pointer(s))))
}

/**
 * Open a write only memory stream.
 *
 * @param s new IO context
 * @return zero if no error.
 */
func AvioOpenDynBuf(s **CAVIOContext) int {
	return int(C.avio_open_dyn_buf((**C.AVIOContext)(unsafe.Pointer(s))))
}

/**
 * Return the written size and a pointer to the buffer.
 * The AVIOContext stream is left intact.
 * The buffer must NOT be freed.
 * No padding is added to the buffer.
 *
 * @param s IO context
 * @param pbuffer pointer to a byte buffer
 * @return the length of the byte buffer
 */
func AvioGetDynBuf(s *CAVIOContext, pbuffer unsafe.Pointer) int {
	return int(C.avio_get_dyn_buf((*C.AVIOContext)(s), (**C.uchar)(pbuffer)))
}

/**
 * Return the written size and a pointer to the buffer. The buffer
 * must be freed with av_free().
 * Padding of AV_INPUT_BUFFER_PADDING_SIZE is added to the buffer.
 *
 * @param s IO context
 * @param pbuffer pointer to a byte buffer
 * @return the length of the byte buffer
 */
func AvioCloseDynBuf(s *CAVIOContext, pbuffer unsafe.Pointer) int {
	return int(C.avio_close_dyn_buf((*C.AVIOContext)(s), (**C.uchar)(pbuffer)))
}

/**
 * Iterate through names of available protocols.
 *
 * @param opaque A private pointer representing current protocol.
 *        It must be a pointer to NULL on first iteration and will
 *        be updated by successive calls to avio_enum_protocols.
 * @param output If set to 1, iterate over output protocols,
 *               otherwise over input protocols.
 *
 * @return A static string containing the name of current protocol or NULL
 */
func AvioEnumProtocols(opaque *unsafe.Pointer, output int) string {
	return C.GoString(C.avio_enum_protocols(opaque, C.int(output)))
}

/**
 * Get AVClass by names of available protocols.
 *
 * @return A AVClass of input protocol name or NULL
 */
func AvioProtocolGetClass(name string) *avutil.CAVClass {
	var cName *C.char = nil
	if len(name) > 0 {
		cName = C.CString(name)
		defer C.free(unsafe.Pointer(cName))
	}

	return (*avutil.CAVClass)(unsafe.Pointer(C.avio_protocol_get_class(cName)))
}

/**
 * Pause and resume playing - only meaningful if using a network streaming
 * protocol (e.g. MMS).
 *
 * @param h     IO context from which to call the read_pause function pointer
 * @param pause 1 for pause, 0 for resume
 */
func AvioPause(h *CAVIOContext, pause int) int {
	return int(C.avio_pause((*C.AVIOContext)(h), C.int(pause)))
}

/**
 * Seek to a given timestamp relative to some component stream.
 * Only meaningful if using a network streaming protocol (e.g. MMS.).
 *
 * @param h IO context from which to call the seek function pointers
 * @param stream_index The stream index that the timestamp is relative to.
 *        If stream_index is (-1) the timestamp should be in AV_TIME_BASE
 *        units from the beginning of the presentation.
 *        If a stream_index >= 0 is used and the protocol does not support
 *        seeking based on component streams, the call will fail.
 * @param timestamp timestamp in AVStream.time_base units
 *        or if there is no stream specified then in AV_TIME_BASE units.
 * @param flags Optional combination of AVSEEK_FLAG_BACKWARD, AVSEEK_FLAG_BYTE
 *        and AVSEEK_FLAG_ANY. The protocol may silently ignore
 *        AVSEEK_FLAG_BACKWARD and AVSEEK_FLAG_ANY, but AVSEEK_FLAG_BYTE will
 *        fail if used and not supported.
 * @return >= 0 on success
 * @see AVInputFormat::read_seek
 */
func AvioSeekTime(h *CAVIOContext, streamIndex int, timestamp int64, flags int) int64 {
	return int64(C.avio_seek_time((*C.AVIOContext)(h), C.int(streamIndex), C.int64_t(timestamp), C.int(flags)))
}

/* Avoid a warning. The header can not be included because it breaks c++. */
type CAVBPrint C.struct_AVBPrint

/**
 * Read contents of h into print buffer, up to max_size bytes, or up to EOF.
 *
 * @return 0 for success (max_size bytes read or EOF reached), negative error
 * code otherwise
 */
func AvioReadToBprint(h *CAVIOContext, pb *CAVBPrint, maxSize ctypes.SizeT) int {
	return int(C.avio_read_to_bprint((*C.AVIOContext)(h), (*C.struct_AVBPrint)(pb), C.size_t(maxSize)))
}

/**
 * Accept and allocate a client context on a server context.
 * @param  s the server context
 * @param  c the client context, must be unallocated
 * @return   >= 0 on success or a negative value corresponding
 *           to an AVERROR on failure
 */
func AvioAccept(s *CAVIOContext, c **CAVIOContext) int {
	return int(C.avio_accept((*C.AVIOContext)(s), (**C.AVIOContext)(unsafe.Pointer(c))))
}

/**
 * Perform one step of the protocol handshake to accept a new client.
 * This function must be called on a client returned by avio_accept() before
 * using it as a read/write context.
 * It is separate from avio_accept() because it may block.
 * A step of the handshake is defined by places where the application may
 * decide to change the proceedings.
 * For example, on a protocol with a request header and a reply header, each
 * one can constitute a step because the application may use the parameters
 * from the request to change parameters in the reply; or each individual
 * chunk of the request can constitute a step.
 * If the handshake is already finished, avio_handshake() does nothing and
 * returns 0 immediately.
 *
 * @param  c the client context to perform the handshake on
 * @return   0   on a complete and successful handshake
 *           > 0 if the handshake progressed, but is not complete
 *           < 0 for an AVERROR code
 */
func AvioHandshake(c *CAVIOContext) int {
	return int(C.avio_handshake((*C.AVIOContext)(c)))
}

//  #endif /* AVFORMAT_AVIO_H */
