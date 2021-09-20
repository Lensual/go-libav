package avformat

/*
#cgo LDFLAGS: -lavformat

#include "libavformat/avformat.h"
*/
import "C"
import "unsafe"

/**
 * Format I/O context.
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVFormatContext) must not be used outside libav*, use
 * avformat_alloc_context() to create an AVFormatContext.
 *
 * Fields can be accessed through AVOptions (av_opt*),
 * the name string used matches the associated command line parameter name and
 * can be found in libavformat/options_table.h.
 * The AVOption/command line parameter names differ in some cases from the C
 * structure field names for historic reasons or brevity.
 */
type CAVFormatContext C.AVFormatContext

/*
Allocate an AVFormatContext.

avformat_free_context() can be used to free the context and everything
allocated by the framework within it.
*/
func AvformatAllocContext() *CAVFormatContext {
	return (*CAVFormatContext)(C.avformat_alloc_context())
}

/*
Free an AVFormatContext and all its streams.

@param s context to free
*/
func AvformatFreeContext(s *CAVFormatContext) {
	C.avformat_free_context((*C.AVFormatContext)(s))
}

/*
Open an input stream and read the header. The codecs are not opened.
The stream must be closed with avformat_close_input().

@param ps Pointer to user-supplied AVFormatContext (allocated by avformat_alloc_context).
          May be a pointer to NULL, in which case an AVFormatContext is allocated by this
          function and written into ps.
          Note that a user-supplied AVFormatContext will be freed on failure.
@param url URL of the stream to open.
@param fmt If non-NULL, this parameter forces a specific input format.
           Otherwise the format is autodetected.
@param options  A dictionary filled with AVFormatContext and demuxer-private options.
                On return this parameter will be destroyed and replaced with a dict containing
                options that were not found. May be NULL.

@return 0 on success, a negative AVERROR on failure.

@note If you want to use custom IO, preallocate the format context and set its pb field.
*/
func AvformatOpenInput(ps **CAVFormatContext, url string, fmt *C.AVInputFormat, options **C.AVDictionary) int {
	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
		defer C.free(unsafe.Pointer(cUrl))
	}

	//TODO AVInputFormat AVDictionary
	return int(C.avformat_open_input((**C.AVFormatContext)(unsafe.Pointer(ps)), cUrl, fmt, options))
}

/*
Read packets of a media file to get stream information. This
is useful for file formats with no headers such as MPEG. This
function also computes the real framerate in case of MPEG-2 repeat
frame mode.
The logical file position is not changed by this function;
examined packets may be buffered for later processing.

@param ic media file handle
@param options  If non-NULL, an ic.nb_streams long array of pointers to
                dictionaries, where i-th member contains options for
                codec corresponding to i-th stream.
                On return each dictionary will be filled with options that were not found.
@return >=0 if OK, AVERROR_xxx on error

@note this function isn't guaranteed to open all the codecs, so
      options being non-empty at return is a perfectly normal behavior.

@todo Let the user decide somehow what information is needed so that
      we do not waste time getting stuff the user does not need.
*/
func AvformatFindStreamInfo(ic *CAVFormatContext, options **C.AVDictionary) int {
	//TODO AVDictionary
	return int(C.avformat_find_stream_info((*C.AVFormatContext)(ic), options))
}

/**
 * Close an opened input AVFormatContext. Free it and all its contents
 * and set *s to NULL.
 */
func AvformatCloseInput(s **CAVFormatContext) {
	C.avformat_close_input((**C.AVFormatContext)(unsafe.Pointer(s)))
}

/*
Print detailed information about the input or output format, such as
duration, bitrate, streams, container, programs, metadata, side data,
codec and time base.

 @param ic        the context to analyze
 @param index     index of the stream to dump information about
 @param url       the URL to print, such as source or destination file
 @param is_output Select whether the specified context is an input(0) or output(1)
*/
func AvDumpFormat(ic *CAVFormatContext, index int, url string, is_output int) {
	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
		defer C.free(unsafe.Pointer(cUrl))
	}

	C.av_dump_format((*C.AVFormatContext)(ic), C.int(index), cUrl, C.int(is_output))
}
