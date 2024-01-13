package avformat

/*
#cgo pkg-config: libavformat

#include "libavformat/avformat.h"
*/
import "C"
import (
	"unsafe"

	"github.com/Lensual/go-libav/avcodec"
	"github.com/Lensual/go-libav/avutil"
	"github.com/Lensual/go-libav/ctypes"
)

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

//  #ifndef AVFORMAT_AVFORMAT_H
//  #define AVFORMAT_AVFORMAT_H

/**
 * @file
 * @ingroup libavf
 * Main libavformat public API header
 */

/**
 * @defgroup libavf libavformat
 * I/O and Muxing/Demuxing Library
 *
 * Libavformat (lavf) is a library for dealing with various media container
 * formats. Its main two purposes are demuxing - i.e. splitting a media file
 * into component streams, and the reverse process of muxing - writing supplied
 * data in a specified container format. It also has an @ref lavf_io
 * "I/O module" which supports a number of protocols for accessing the data (e.g.
 * file, tcp, http and others).
 * Unless you are absolutely sure you won't use libavformat's network
 * capabilities, you should also call avformat_network_init().
 *
 * A supported input format is described by an AVInputFormat struct, conversely
 * an output format is described by AVOutputFormat. You can iterate over all
 * input/output formats using the  av_demuxer_iterate / av_muxer_iterate() functions.
 * The protocols layer is not part of the public API, so you can only get the names
 * of supported protocols with the avio_enum_protocols() function.
 *
 * Main lavf structure used for both muxing and demuxing is AVFormatContext,
 * which exports all information about the file being read or written. As with
 * most Libavformat structures, its size is not part of public ABI, so it cannot be
 * allocated on stack or directly with av_malloc(). To create an
 * AVFormatContext, use avformat_alloc_context() (some functions, like
 * avformat_open_input() might do that for you).
 *
 * Most importantly an AVFormatContext contains:
 * @li the @ref AVFormatContext.iformat "input" or @ref AVFormatContext.oformat
 * "output" format. It is either autodetected or set by user for input;
 * always set by user for output.
 * @li an @ref AVFormatContext.streams "array" of AVStreams, which describe all
 * elementary streams stored in the file. AVStreams are typically referred to
 * using their index in this array.
 * @li an @ref AVFormatContext.pb "I/O context". It is either opened by lavf or
 * set by user for input, always set by user for output (unless you are dealing
 * with an AVFMT_NOFILE format).
 *
 * @section lavf_options Passing options to (de)muxers
 * It is possible to configure lavf muxers and demuxers using the @ref avoptions
 * mechanism. Generic (format-independent) libavformat options are provided by
 * AVFormatContext, they can be examined from a user program by calling
 * av_opt_next() / av_opt_find() on an allocated AVFormatContext (or its AVClass
 * from avformat_get_class()). Private (format-specific) options are provided by
 * AVFormatContext.priv_data if and only if AVInputFormat.priv_class /
 * AVOutputFormat.priv_class of the corresponding format struct is non-NULL.
 * Further options may be provided by the @ref AVFormatContext.pb "I/O context",
 * if its AVClass is non-NULL, and the protocols layer. See the discussion on
 * nesting in @ref avoptions documentation to learn how to access those.
 *
 * @section urls
 * URL strings in libavformat are made of a scheme/protocol, a ':', and a
 * scheme specific string. URLs without a scheme and ':' used for local files
 * are supported but deprecated. "file:" should be used for local files.
 *
 * It is important that the scheme string is not taken from untrusted
 * sources without checks.
 *
 * Note that some schemes/protocols are quite powerful, allowing access to
 * both local and remote files, parts of them, concatenations of them, local
 * audio and video devices and so on.
 *
 * @{
 *
 * @defgroup lavf_decoding Demuxing
 * @{
 * Demuxers read a media file and split it into chunks of data (@em packets). A
 * @ref AVPacket "packet" contains one or more encoded frames which belongs to a
 * single elementary stream. In the lavf API this process is represented by the
 * avformat_open_input() function for opening a file, av_read_frame() for
 * reading a single packet and finally avformat_close_input(), which does the
 * cleanup.
 *
 * @section lavf_decoding_open Opening a media file
 * The minimum information required to open a file is its URL, which
 * is passed to avformat_open_input(), as in the following code:
 * @code
 * const char    *url = "file:in.mp3";
 * AVFormatContext *s = NULL;
 * int ret = avformat_open_input(&s, url, NULL, NULL);
 * if (ret < 0)
 *     abort();
 * @endcode
 * The above code attempts to allocate an AVFormatContext, open the
 * specified file (autodetecting the format) and read the header, exporting the
 * information stored there into s. Some formats do not have a header or do not
 * store enough information there, so it is recommended that you call the
 * avformat_find_stream_info() function which tries to read and decode a few
 * frames to find missing information.
 *
 * In some cases you might want to preallocate an AVFormatContext yourself with
 * avformat_alloc_context() and do some tweaking on it before passing it to
 * avformat_open_input(). One such case is when you want to use custom functions
 * for reading input data instead of lavf internal I/O layer.
 * To do that, create your own AVIOContext with avio_alloc_context(), passing
 * your reading callbacks to it. Then set the @em pb field of your
 * AVFormatContext to newly created AVIOContext.
 *
 * Since the format of the opened file is in general not known until after
 * avformat_open_input() has returned, it is not possible to set demuxer private
 * options on a preallocated context. Instead, the options should be passed to
 * avformat_open_input() wrapped in an AVDictionary:
 * @code
 * AVDictionary *options = NULL;
 * av_dict_set(&options, "video_size", "640x480", 0);
 * av_dict_set(&options, "pixel_format", "rgb24", 0);
 *
 * if (avformat_open_input(&s, url, NULL, &options) < 0)
 *     abort();
 * av_dict_free(&options);
 * @endcode
 * This code passes the private options 'video_size' and 'pixel_format' to the
 * demuxer. They would be necessary for e.g. the rawvideo demuxer, since it
 * cannot know how to interpret raw video data otherwise. If the format turns
 * out to be something different than raw video, those options will not be
 * recognized by the demuxer and therefore will not be applied. Such unrecognized
 * options are then returned in the options dictionary (recognized options are
 * consumed). The calling program can handle such unrecognized options as it
 * wishes, e.g.
 * @code
 * AVDictionaryEntry *e;
 * if (e = av_dict_get(options, "", NULL, AV_DICT_IGNORE_SUFFIX)) {
 *     fprintf(stderr, "Option %s not recognized by the demuxer.\n", e->key);
 *     abort();
 * }
 * @endcode
 *
 * After you have finished reading the file, you must close it with
 * avformat_close_input(). It will free everything associated with the file.
 *
 * @section lavf_decoding_read Reading from an opened file
 * Reading data from an opened AVFormatContext is done by repeatedly calling
 * av_read_frame() on it. Each call, if successful, will return an AVPacket
 * containing encoded data for one AVStream, identified by
 * AVPacket.stream_index. This packet may be passed straight into the libavcodec
 * decoding functions avcodec_send_packet() or avcodec_decode_subtitle2() if the
 * caller wishes to decode the data.
 *
 * AVPacket.pts, AVPacket.dts and AVPacket.duration timing information will be
 * set if known. They may also be unset (i.e. AV_NOPTS_VALUE for
 * pts/dts, 0 for duration) if the stream does not provide them. The timing
 * information will be in AVStream.time_base units, i.e. it has to be
 * multiplied by the timebase to convert them to seconds.
 *
 * A packet returned by av_read_frame() is always reference-counted,
 * i.e. AVPacket.buf is set and the user may keep it indefinitely.
 * The packet must be freed with av_packet_unref() when it is no
 * longer needed.
 *
 * @section lavf_decoding_seek Seeking
 * @}
 *
 * @defgroup lavf_encoding Muxing
 * @{
 * Muxers take encoded data in the form of @ref AVPacket "AVPackets" and write
 * it into files or other output bytestreams in the specified container format.
 *
 * The main API functions for muxing are avformat_write_header() for writing the
 * file header, av_write_frame() / av_interleaved_write_frame() for writing the
 * packets and av_write_trailer() for finalizing the file.
 *
 * At the beginning of the muxing process, the caller must first call
 * avformat_alloc_context() to create a muxing context. The caller then sets up
 * the muxer by filling the various fields in this context:
 *
 * - The @ref AVFormatContext.oformat "oformat" field must be set to select the
 *   muxer that will be used.
 * - Unless the format is of the AVFMT_NOFILE type, the @ref AVFormatContext.pb
 *   "pb" field must be set to an opened IO context, either returned from
 *   avio_open2() or a custom one.
 * - Unless the format is of the AVFMT_NOSTREAMS type, at least one stream must
 *   be created with the avformat_new_stream() function. The caller should fill
 *   the @ref AVStream.codecpar "stream codec parameters" information, such as the
 *   codec @ref AVCodecParameters.codec_type "type", @ref AVCodecParameters.codec_id
 *   "id" and other parameters (e.g. width / height, the pixel or sample format,
 *   etc.) as known. The @ref AVStream.time_base "stream timebase" should
 *   be set to the timebase that the caller desires to use for this stream (note
 *   that the timebase actually used by the muxer can be different, as will be
 *   described later).
 * - It is advised to manually initialize only the relevant fields in
 *   AVCodecParameters, rather than using @ref avcodec_parameters_copy() during
 *   remuxing: there is no guarantee that the codec context values remain valid
 *   for both input and output format contexts.
 * - The caller may fill in additional information, such as @ref
 *   AVFormatContext.metadata "global" or @ref AVStream.metadata "per-stream"
 *   metadata, @ref AVFormatContext.chapters "chapters", @ref
 *   AVFormatContext.programs "programs", etc. as described in the
 *   AVFormatContext documentation. Whether such information will actually be
 *   stored in the output depends on what the container format and the muxer
 *   support.
 *
 * When the muxing context is fully set up, the caller must call
 * avformat_write_header() to initialize the muxer internals and write the file
 * header. Whether anything actually is written to the IO context at this step
 * depends on the muxer, but this function must always be called. Any muxer
 * private options must be passed in the options parameter to this function.
 *
 * The data is then sent to the muxer by repeatedly calling av_write_frame() or
 * av_interleaved_write_frame() (consult those functions' documentation for
 * discussion on the difference between them; only one of them may be used with
 * a single muxing context, they should not be mixed). Do note that the timing
 * information on the packets sent to the muxer must be in the corresponding
 * AVStream's timebase. That timebase is set by the muxer (in the
 * avformat_write_header() step) and may be different from the timebase
 * requested by the caller.
 *
 * Once all the data has been written, the caller must call av_write_trailer()
 * to flush any buffered packets and finalize the output file, then close the IO
 * context (if any) and finally free the muxing context with
 * avformat_free_context().
 * @}
 *
 * @defgroup lavf_io I/O Read/Write
 * @{
 * @section lavf_io_dirlist Directory listing
 * The directory listing API makes it possible to list files on remote servers.
 *
 * Some of possible use cases:
 * - an "open file" dialog to choose files from a remote location,
 * - a recursive media finder providing a player with an ability to play all
 * files from a given directory.
 *
 * @subsection lavf_io_dirlist_open Opening a directory
 * At first, a directory needs to be opened by calling avio_open_dir()
 * supplied with a URL and, optionally, ::AVDictionary containing
 * protocol-specific parameters. The function returns zero or positive
 * integer and allocates AVIODirContext on success.
 *
 * @code
 * AVIODirContext *ctx = NULL;
 * if (avio_open_dir(&ctx, "smb://example.com/some_dir", NULL) < 0) {
 *     fprintf(stderr, "Cannot open directory.\n");
 *     abort();
 * }
 * @endcode
 *
 * This code tries to open a sample directory using smb protocol without
 * any additional parameters.
 *
 * @subsection lavf_io_dirlist_read Reading entries
 * Each directory's entry (i.e. file, another directory, anything else
 * within ::AVIODirEntryType) is represented by AVIODirEntry.
 * Reading consecutive entries from an opened AVIODirContext is done by
 * repeatedly calling avio_read_dir() on it. Each call returns zero or
 * positive integer if successful. Reading can be stopped right after the
 * NULL entry has been read -- it means there are no entries left to be
 * read. The following code reads all entries from a directory associated
 * with ctx and prints their names to standard output.
 * @code
 * AVIODirEntry *entry = NULL;
 * for (;;) {
 *     if (avio_read_dir(ctx, &entry) < 0) {
 *         fprintf(stderr, "Cannot list directory.\n");
 *         abort();
 *     }
 *     if (!entry)
 *         break;
 *     printf("%s\n", entry->name);
 *     avio_free_directory_entry(&entry);
 * }
 * @endcode
 * @}
 *
 * @defgroup lavf_codec Demuxers
 * @{
 * @defgroup lavf_codec_native Native Demuxers
 * @{
 * @}
 * @defgroup lavf_codec_wrappers External library wrappers
 * @{
 * @}
 * @}
 * @defgroup lavf_protos I/O Protocols
 * @{
 * @}
 * @defgroup lavf_internal Internal
 * @{
 * @}
 * @}
 */

//  struct AVFormatContext;
//  struct AVFrame;
//  struct AVDeviceInfoList;

/**
  * @defgroup metadata_api Public Metadata API
  * @{
  * @ingroup libavf
  * The metadata API allows libavformat to export metadata tags to a client
  * application when demuxing. Conversely it allows a client application to
  * set metadata when muxing.
  *
  * Metadata is exported or set as pairs of key/value strings in the 'metadata'
  * fields of the AVFormatContext, AVStream, AVChapter and AVProgram structs
  * using the @ref lavu_dict "AVDictionary" API. Like all strings in FFmpeg,
  * metadata is assumed to be UTF-8 encoded Unicode. Note that metadata
  * exported by demuxers isn't checked to be valid UTF-8 in most cases.
  *
  * Important concepts to keep in mind:
  * -  Keys are unique; there can never be 2 tags with the same key. This is
  *    also meant semantically, i.e., a demuxer should not knowingly produce
  *    several keys that are literally different but semantically identical.
  *    E.g., key=Author5, key=Author6. In this example, all authors must be
  *    placed in the same tag.
  * -  Metadata is flat, not hierarchical; there are no subtags. If you
  *    want to store, e.g., the email address of the child of producer Alice
  *    and actor Bob, that could have key=alice_and_bobs_childs_email_address.
  * -  Several modifiers can be applied to the tag name. This is done by
  *    appending a dash character ('-') and the modifier name in the order
  *    they appear in the list below -- e.g. foo-eng-sort, not foo-sort-eng.
  *    -  language -- a tag whose value is localized for a particular language
  *       is appended with the ISO 639-2/B 3-letter language code.
  *       For example: Author-ger=Michael, Author-eng=Mike
  *       The original/default language is in the unqualified "Author" tag.
  *       A demuxer should set a default if it sets any translated tag.
  *    -  sorting  -- a modified version of a tag that should be used for
  *       sorting will have '-sort' appended. E.g. artist="The Beatles",
  *       artist-sort="Beatles, The".
  * - Some protocols and demuxers support metadata updates. After a successful
  *   call to av_read_frame(), AVFormatContext.event_flags or AVStream.event_flags
  *   will be updated to indicate if metadata changed. In order to detect metadata
  *   changes on a stream, you need to loop through all streams in the AVFormatContext
  *   and check their individual event_flags.
  *
  * -  Demuxers attempt to export metadata in a generic format, however tags
  *    with no generic equivalents are left as they are stored in the container.
  *    Follows a list of generic tag names:
  *
  @verbatim
  album        -- name of the set this work belongs to
  album_artist -- main creator of the set/album, if different from artist.
				  e.g. "Various Artists" for compilation albums.
  artist       -- main creator of the work
  comment      -- any additional description of the file.
  composer     -- who composed the work, if different from artist.
  copyright    -- name of copyright holder.
  creation_time-- date when the file was created, preferably in ISO 8601.
  date         -- date when the work was created, preferably in ISO 8601.
  disc         -- number of a subset, e.g. disc in a multi-disc collection.
  encoder      -- name/settings of the software/hardware that produced the file.
  encoded_by   -- person/group who created the file.
  filename     -- original name of the file.
  genre        -- <self-evident>.
  language     -- main language in which the work is performed, preferably
				  in ISO 639-2 format. Multiple languages can be specified by
				  separating them with commas.
  performer    -- artist who performed the work, if different from artist.
				  E.g for "Also sprach Zarathustra", artist would be "Richard
				  Strauss" and performer "London Philharmonic Orchestra".
  publisher    -- name of the label/publisher.
  service_name     -- name of the service in broadcasting (channel name).
  service_provider -- name of the service provider in broadcasting.
  title        -- name of the work.
  track        -- number of this work in the set, can be in form current/total.
  variant_bitrate -- the total bitrate of the bitrate variant that the current stream is part of
  @endverbatim
  *
  * Look in the examples section for an application example how to use the Metadata API.
  *
  * @}
*/

/* packet functions */

/**
 * Allocate and read the payload of a packet and initialize its
 * fields with default values.
 *
 * @param s    associated IO context
 * @param pkt packet
 * @param size desired payload size
 * @return >0 (read size) if OK, AVERROR_xxx otherwise
 */
func AvGetPacket(s *CAVIOContext, pkt *avcodec.CAVPacket, size int) int {
	return int(C.av_get_packet((*C.AVIOContext)(s), (*C.AVPacket)(unsafe.Pointer(pkt)), C.int(size)))
}

/**
 * Read data and append it to the current content of the AVPacket.
 * If pkt->size is 0 this is identical to av_get_packet.
 * Note that this uses av_grow_packet and thus involves a realloc
 * which is inefficient. Thus this function should only be used
 * when there is no reasonable way to know (an upper bound of)
 * the final size.
 *
 * @param s    associated IO context
 * @param pkt packet
 * @param size amount of data to read
 * @return >0 (read size) if OK, AVERROR_xxx otherwise, previous data
 *         will not be lost even if an error occurs.
 */
func AvAppendPacket(s *CAVIOContext, pkt *avcodec.CAVPacket, size int) int {
	return int(C.av_append_packet((*C.AVIOContext)(s), (*C.AVPacket)(unsafe.Pointer(pkt)), C.int(size)))
}

/*************************************************/
/* input/output formats */

type CAVCodecTag C.struct_AVCodecTag

/**
 * This structure contains the data a format has to probe a file.
 */
type CAVProbeData C.AVProbeData

//#region CAVProbeData

func (pd *CAVProbeData) GetFilename() string {
	return C.GoString(pd.filename)
}

/**< Buffer must have AVPROBE_PADDING_SIZE of extra allocated bytes filled with zero. */
func (pd *CAVProbeData) GetBuf() unsafe.Pointer {
	return unsafe.Pointer(pd.buf)
}

/**< Buffer must have AVPROBE_PADDING_SIZE of extra allocated bytes filled with zero. */
func (pd *CAVProbeData) SetBuf(buf unsafe.Pointer) {
	pd.buf = (*C.uchar)(buf)
}

/**< Size of buf except extra allocated bytes */
func (pd *CAVProbeData) GetBufSize() int {
	return int(pd.buf_size)
}

/**< Size of buf except extra allocated bytes */
func (pd *CAVProbeData) SetBufSize(bufSize int) {
	pd.buf_size = C.int(bufSize)
}

/**< mime_type, when known. */
func (pd *CAVProbeData) GetMimeType() string {
	return C.GoString(pd.mime_type)
}

//#endregion CAVProbeData

const (
	AVPROBE_SCORE_RETRY        = C.AVPROBE_SCORE_RETRY
	AVPROBE_SCORE_STREAM_RETRY = C.AVPROBE_SCORE_STREAM_RETRY

	AVPROBE_SCORE_EXTENSION = C.AVPROBE_SCORE_EXTENSION ///< score for file extension
	AVPROBE_SCORE_MIME      = C.AVPROBE_SCORE_MIME      ///< score for file mime type
	AVPROBE_SCORE_MAX       = C.AVPROBE_SCORE_MAX       ///< maximum score
)

const AVPROBE_PADDING_SIZE = C.AVPROBE_PADDING_SIZE ///< extra allocated bytes at the end of the probe buffer

// / Demuxer will use avio_open, no opened file should be provided by the caller.
const (
	AVFMT_NOFILE     = C.AVFMT_NOFILE
	AVFMT_NEEDNUMBER = C.AVFMT_NEEDNUMBER /**< Needs '%d' in filename. */
)

/**
 * The muxer/demuxer is experimental and should be used with caution.
 *
 * - demuxers: will not be selected automatically by probing, must be specified
 *             explicitly.
 */
const (
	AVFMT_EXPERIMENTAL  = C.AVFMT_EXPERIMENTAL
	AVFMT_SHOW_IDS      = C.AVFMT_SHOW_IDS      /**< Show format stream IDs numbers. */
	AVFMT_GLOBALHEADER  = C.AVFMT_GLOBALHEADER  /**< Format wants global header. */
	AVFMT_NOTIMESTAMPS  = C.AVFMT_NOTIMESTAMPS  /**< Format does not need / have any timestamps. */
	AVFMT_GENERIC_INDEX = C.AVFMT_GENERIC_INDEX /**< Use generic index building code. */
	AVFMT_TS_DISCONT    = C.AVFMT_TS_DISCONT    /**< Format allows timestamp discontinuities. Note, muxers always require valid (monotone) timestamps */
	AVFMT_VARIABLE_FPS  = C.AVFMT_VARIABLE_FPS  /**< Format allows variable fps. */
	AVFMT_NODIMENSIONS  = C.AVFMT_NODIMENSIONS  /**< Format does not need width/height */
	AVFMT_NOSTREAMS     = C.AVFMT_NOSTREAMS     /**< Format does not require any streams */
	AVFMT_NOBINSEARCH   = C.AVFMT_NOBINSEARCH   /**< Format does not allow to fall back on binary search via read_timestamp */
	AVFMT_NOGENSEARCH   = C.AVFMT_NOGENSEARCH   /**< Format does not allow to fall back on generic search */
	AVFMT_NO_BYTE_SEEK  = C.AVFMT_NO_BYTE_SEEK  /**< Format does not allow seeking by bytes */
	//  #if FF_API_ALLOW_FLUSH
	//  AVFMT_ALLOW_FLUSH  = C.AVFMT_ALLOW_FLUSH     /**< @deprecated: Just send a NULL packet if you want to flush a muxer. */
	//  #endif

	/**< Format does not require strictly
	increasing timestamps, but they must
	still be monotonic */
	AVFMT_TS_NONSTRICT = C.AVFMT_TS_NONSTRICT
	/**< Format allows muxing negative
	timestamps. If not set the timestamp
	will be shifted in av_write_frame and
	av_interleaved_write_frame so they
	start from 0.
	The user or muxer can override this through
	AVFormatContext.avoid_negative_ts
	*/
	AVFMT_TS_NEGATIVE = C.AVFMT_TS_NEGATIVE
	AVFMT_SEEK_TO_PTS = C.AVFMT_SEEK_TO_PTS /**< Seeking is based on PTS */
)

/**
 * @addtogroup lavf_encoding
 * @{
 */
type CAVOutputFormat C.AVOutputFormat

//#region CAVOutputFormat

func (oFormat *CAVOutputFormat) GetName() string {
	return C.GoString(oFormat.name)
}

/**
 * Descriptive name for the format, meant to be more human-readable
 * than name. You should use the NULL_IF_CONFIG_SMALL() macro
 * to define it.
 */
func (oFormat *CAVOutputFormat) GetLongName() string {
	return C.GoString(oFormat.long_name)
}

func (oFormat *CAVOutputFormat) GetMimeType() string {
	return C.GoString(oFormat.mime_type)
}

/**< comma-separated filename extensions */
func (oFormat *CAVOutputFormat) GetExtensions() string {
	return C.GoString(oFormat.extensions)
}

/* output support */

/**< default audio codec */
func (oFormat *CAVOutputFormat) GetAudioCodec() avcodec.CAVCodecID {
	return avcodec.CAVCodecID(oFormat.audio_codec)
}

/**< default audio codec */
func (oFormat *CAVOutputFormat) SetAudioCodec(audioCodec avcodec.CAVCodecID) {
	oFormat.audio_codec = C.enum_AVCodecID(audioCodec)
}

/**< default video codec */
func (oFormat *CAVOutputFormat) GetVideoCodec() avcodec.CAVCodecID {
	return avcodec.CAVCodecID(oFormat.video_codec)
}

/**< default video codec */
func (oFormat *CAVOutputFormat) SetVideoCodec(videoCodec avcodec.CAVCodecID) {
	oFormat.video_codec = C.enum_AVCodecID(videoCodec)
}

/**< default subtitle codec */
func (oFormat *CAVOutputFormat) GetSubtitleCodec() avcodec.CAVCodecID {
	return avcodec.CAVCodecID(oFormat.subtitle_codec)
}

/**< default subtitle codec */
func (oFormat *CAVOutputFormat) SetSubtitleCodec(subtitleCodec avcodec.CAVCodecID) {
	oFormat.subtitle_codec = C.enum_AVCodecID(subtitleCodec)
}

/**
 * can use flags: AVFMT_NOFILE, AVFMT_NEEDNUMBER,
 * AVFMT_GLOBALHEADER, AVFMT_NOTIMESTAMPS, AVFMT_VARIABLE_FPS,
 * AVFMT_NODIMENSIONS, AVFMT_NOSTREAMS,
 * AVFMT_TS_NONSTRICT, AVFMT_TS_NEGATIVE
 */
func (oFormat *CAVOutputFormat) GetFlags() int {
	return int(oFormat.flags)
}

/**
 * can use flags: AVFMT_NOFILE, AVFMT_NEEDNUMBER,
 * AVFMT_GLOBALHEADER, AVFMT_NOTIMESTAMPS, AVFMT_VARIABLE_FPS,
 * AVFMT_NODIMENSIONS, AVFMT_NOSTREAMS,
 * AVFMT_TS_NONSTRICT, AVFMT_TS_NEGATIVE
 */
func (oFormat *CAVOutputFormat) SetFlags(flags int) {
	oFormat.flags = C.int(flags)
}

/**
 * List of supported codec_id-codec_tag pairs, ordered by "better
 * choice first". The arrays are all terminated by AV_CODEC_ID_NONE.
 */
func (oFormat *CAVOutputFormat) GetCodecTag() **CAVCodecTag {
	return (**CAVCodecTag)(unsafe.Pointer(oFormat.codec_tag))
}

// /< AVClass for the private context
func (oFormat *CAVOutputFormat) GetPrivClass() *avutil.CAVClass {
	return (*avutil.CAVClass)(unsafe.Pointer(oFormat.priv_class))
}

//#endregion CAVOutputFormat

/**
 * @}
 */

/**
 * @addtogroup lavf_decoding
 * @{
 */
type CAVInputFormat C.AVInputFormat

//#region CAVInputFormat
/**
 * A comma separated list of short names for the format. New names
 * may be appended with a minor bump.
 */
func (iFormat *CAVInputFormat) GetName() string {
	return C.GoString(iFormat.name)
}

/**
 * Descriptive name for the format, meant to be more human-readable
 * than name. You should use the NULL_IF_CONFIG_SMALL() macro
 * to define it.
 */
func (iFormat *CAVInputFormat) GetLongName() string {
	return C.GoString(iFormat.long_name)
}

/**
 * Can use flags: AVFMT_NOFILE, AVFMT_NEEDNUMBER, AVFMT_SHOW_IDS,
 * AVFMT_NOTIMESTAMPS, AVFMT_GENERIC_INDEX, AVFMT_TS_DISCONT, AVFMT_NOBINSEARCH,
 * AVFMT_NOGENSEARCH, AVFMT_NO_BYTE_SEEK, AVFMT_SEEK_TO_PTS.
 */
func (iFormat *CAVInputFormat) GetFlags() int {
	return int(iFormat.flags)
}

/**
 * Can use flags: AVFMT_NOFILE, AVFMT_NEEDNUMBER, AVFMT_SHOW_IDS,
 * AVFMT_NOTIMESTAMPS, AVFMT_GENERIC_INDEX, AVFMT_TS_DISCONT, AVFMT_NOBINSEARCH,
 * AVFMT_NOGENSEARCH, AVFMT_NO_BYTE_SEEK, AVFMT_SEEK_TO_PTS.
 */
func (iFormat *CAVInputFormat) SetFlags(flags int) {
	iFormat.flags = C.int(flags)
}

/**
 * If extensions are defined, then no probe is done. You should
 * usually not use extension format guessing because it is not
 * reliable enough
 */
func (iFormat *CAVInputFormat) GetExtensions() string {
	return C.GoString(iFormat.extensions)
}

func (iFormat *CAVInputFormat) GetCodecTag() **CAVCodecTag {
	return (**CAVCodecTag)(unsafe.Pointer(iFormat.codec_tag))
}

// /< AVClass for the private context
func (iFormat *CAVInputFormat) GetPrivClass() *avutil.CAVClass {
	return (*avutil.CAVClass)(unsafe.Pointer(iFormat.priv_class))
}

/**
 * Comma-separated list of mime types.
 * It is used check for matching mime types while probing.
 * @see av_probe_input_format2
 */
func (iFormat *CAVInputFormat) GetMimeType() string {
	return C.GoString(iFormat.mime_type)
}

/*****************************************************************
 * No fields below this line are part of the public API. They
 * may not be used outside of libavformat and can be changed and
 * removed at will.
 * New public fields should be added right above.
 *****************************************************************
 */

/**
 * Raw demuxers store their codec ID here.
 */
func (iFormat *CAVInputFormat) GetRawCodecId() int {
	return int(iFormat.raw_codec_id)
}

/**
 * Raw demuxers store their codec ID here.
 */
func (iFormat *CAVInputFormat) SetRawCodecId(rawCodecId int) {
	iFormat.raw_codec_id = C.int(rawCodecId)
}

/**
 * Size of private data so that it can be allocated in the wrapper.
 */
func (iFormat *CAVInputFormat) GetPrivDataSize() int {
	return int(iFormat.priv_data_size)
}

/**
 * Size of private data so that it can be allocated in the wrapper.
 */
func (iFormat *CAVInputFormat) SetPrivDataSize(privDataSize int) {
	iFormat.priv_data_size = C.int(privDataSize)
}

/**
 * Internal flags. See FF_FMT_FLAG_* in internal.h.
 */
func (iFormat *CAVInputFormat) GetFlagsInternal() int {
	return int(iFormat.flags_internal)
}

/**
 * Internal flags. See FF_FMT_FLAG_* in internal.h.
 */
func (iFormat *CAVInputFormat) SetFlagsInternal(flagsInternal int) {
	iFormat.flags_internal = C.int(flagsInternal)
}

/**
 * Tell if a given file has a chance of being parsed as this format.
 * The buffer provided is guaranteed to be AVPROBE_PADDING_SIZE bytes
 * big so you do not have to check for that unless you need more.
 */
// int (*read_probe)(const AVProbeData *);
func (iFormat *CAVInputFormat) GetReadProbe() ctypes.CFunc {
	return ctypes.CFunc(iFormat.read_probe)
}

/**
 * Tell if a given file has a chance of being parsed as this format.
 * The buffer provided is guaranteed to be AVPROBE_PADDING_SIZE bytes
 * big so you do not have to check for that unless you need more.
 */
// int (*read_probe)(const AVProbeData *);
func (iFormat *CAVInputFormat) SetReadProbe(readProbe ctypes.CFunc) {
	iFormat.read_probe = (*[0]byte)(readProbe)
}

/**
 * Read the format header and initialize the AVFormatContext
 * structure. Return 0 if OK. 'avformat_new_stream' should be
 * called to create new streams.
 */
// int (*read_header)(struct AVFormatContext *);
func (iFormat *CAVInputFormat) GetReadHeader() ctypes.CFunc {
	return ctypes.CFunc(iFormat.read_header)
}

/**
 * Read the format header and initialize the AVFormatContext
 * structure. Return 0 if OK. 'avformat_new_stream' should be
 * called to create new streams.
 */
// int (*read_header)(struct AVFormatContext *);
func (iFormat *CAVInputFormat) SetReadHeader(readHeader ctypes.CFunc) {
	iFormat.read_header = (*[0]byte)(readHeader)
}

/**
 * Read one packet and put it in 'pkt'. pts and flags are also
 * set. 'avformat_new_stream' can be called only if the flag
 * AVFMTCTX_NOHEADER is used and only in the calling thread (not in a
 * background thread).
 * @return 0 on success, < 0 on error.
 *         Upon returning an error, pkt must be unreferenced by the caller.
 */
// int (*read_packet)(struct AVFormatContext *, AVPacket *pkt);
func (iFormat *CAVInputFormat) GetReadPacket() ctypes.CFunc {
	return ctypes.CFunc(iFormat.read_packet)
}

/**
 * Read one packet and put it in 'pkt'. pts and flags are also
 * set. 'avformat_new_stream' can be called only if the flag
 * AVFMTCTX_NOHEADER is used and only in the calling thread (not in a
 * background thread).
 * @return 0 on success, < 0 on error.
 *         Upon returning an error, pkt must be unreferenced by the caller.
 */
// int (*read_packet)(struct AVFormatContext *, AVPacket *pkt);
func (iFormat *CAVInputFormat) SetReadPacket(readPacket ctypes.CFunc) {
	iFormat.read_packet = (*[0]byte)(readPacket)
}

/**
 * Close the stream. The AVFormatContext and AVStreams are not
 * freed by this function
 */
// int (*read_close)(struct AVFormatContext *);
func (iFormat *CAVInputFormat) GetReadClose() ctypes.CFunc {
	return ctypes.CFunc(iFormat.read_close)
}

/**
 * Close the stream. The AVFormatContext and AVStreams are not
 * freed by this function
 */
// int (*read_close)(struct AVFormatContext *);
func (iFormat *CAVInputFormat) SetReadClose(readClose ctypes.CFunc) {
	iFormat.read_close = (*[0]byte)(readClose)
}

/**
 * Seek to a given timestamp relative to the frames in
 * stream component stream_index.
 * @param stream_index Must not be -1.
 * @param flags Selects which direction should be preferred if no exact
 *              match is available.
 * @return >= 0 on success (but not necessarily the new offset)
 */
// int (*read_seek)(struct AVFormatContext *,
//                  int stream_index, int64_t timestamp, int flags);
func (iFormat *CAVInputFormat) GetReadSeek() ctypes.CFunc {
	return ctypes.CFunc(iFormat.read_seek)
}

/**
 * Seek to a given timestamp relative to the frames in
 * stream component stream_index.
 * @param stream_index Must not be -1.
 * @param flags Selects which direction should be preferred if no exact
 *              match is available.
 * @return >= 0 on success (but not necessarily the new offset)
 */
// int (*read_seek)(struct AVFormatContext *,
//                  int stream_index, int64_t timestamp, int flags);
func (iFormat *CAVInputFormat) SetReadSeek(readSeek ctypes.CFunc) {
	iFormat.read_seek = (*[0]byte)(readSeek)
}

/**
 * Get the next timestamp in stream[stream_index].time_base units.
 * @return the timestamp or AV_NOPTS_VALUE if an error occurred
 */
// int64_t (*read_timestamp)(struct AVFormatContext *s, int stream_index,
//                           int64_t *pos, int64_t pos_limit);
func (iFormat *CAVInputFormat) GetReadTimestamp() ctypes.CFunc {
	return ctypes.CFunc(iFormat.read_timestamp)
}

/**
 * Get the next timestamp in stream[stream_index].time_base units.
 * @return the timestamp or AV_NOPTS_VALUE if an error occurred
 */
// int64_t (*read_timestamp)(struct AVFormatContext *s, int stream_index,
//                           int64_t *pos, int64_t pos_limit);
func (iFormat *CAVInputFormat) SetReadTimestamp(readTimestamp ctypes.CFunc) {
	iFormat.read_timestamp = (*[0]byte)(readTimestamp)
}

/**
 * Start/resume playing - only meaningful if using a network-based format
 * (RTSP).
 */
// int (*read_play)(struct AVFormatContext *);
func (iFormat *CAVInputFormat) GetReadPlay() ctypes.CFunc {
	return ctypes.CFunc(iFormat.read_play)
}

/**
 * Start/resume playing - only meaningful if using a network-based format
 * (RTSP).
 */
// int (*read_play)(struct AVFormatContext *);
func (iFormat *CAVInputFormat) SetReadPlay(readPlay ctypes.CFunc) {
	iFormat.read_play = (*[0]byte)(readPlay)
}

/**
 * Pause playing - only meaningful if using a network-based format
 * (RTSP).
 */
// int (*read_pause)(struct AVFormatContext *);
func (iFormat *CAVInputFormat) GetReadPause() ctypes.CFunc {
	return ctypes.CFunc(iFormat.read_pause)
}

/**
 * Pause playing - only meaningful if using a network-based format
 * (RTSP).
 */
// int (*read_pause)(struct AVFormatContext *);
func (iFormat *CAVInputFormat) SetReadPause(readPause ctypes.CFunc) {
	iFormat.read_pause = (*[0]byte)(readPause)
}

/**
 * Seek to timestamp ts.
 * Seeking will be done so that the point from which all active streams
 * can be presented successfully will be closest to ts and within min/max_ts.
 * Active streams are all streams that have AVStream.discard < AVDISCARD_ALL.
 */
// int (*read_seek2)(struct AVFormatContext *s, int stream_index, int64_t min_ts, int64_t ts, int64_t max_ts, int flags);
func (iFormat *CAVInputFormat) GetReadSeek2() ctypes.CFunc {
	return ctypes.CFunc(iFormat.read_seek2)
}

/**
 * Seek to timestamp ts.
 * Seeking will be done so that the point from which all active streams
 * can be presented successfully will be closest to ts and within min/max_ts.
 * Active streams are all streams that have AVStream.discard < AVDISCARD_ALL.
 */
// int (*read_seek2)(struct AVFormatContext *s, int stream_index, int64_t min_ts, int64_t ts, int64_t max_ts, int flags);
func (iFormat *CAVInputFormat) SetReadSeek2(readSeek2 ctypes.CFunc) {
	iFormat.read_seek2 = (*[0]byte)(readSeek2)
}

/**
 * Returns device list with it properties.
 * @see avdevice_list_devices() for more details.
 */
// int (*get_device_list)(struct AVFormatContext *s, struct AVDeviceInfoList *device_list);
func (iFormat *CAVInputFormat) GetGetDeviceList() ctypes.CFunc {
	return ctypes.CFunc(iFormat.get_device_list)
}

/**
 * Returns device list with it properties.
 * @see avdevice_list_devices() for more details.
 */
// int (*get_device_list)(struct AVFormatContext *s, struct AVDeviceInfoList *device_list);
func (iFormat *CAVInputFormat) SetGetDeviceList(getDeviceList ctypes.CFunc) {
	iFormat.get_device_list = (*[0]byte)(getDeviceList)
}

//#endregion CAVInputFormat

/**
 * @}
 */

type CAVStreamParseType C.enum_AVStreamParseType

const (
	AVSTREAM_PARSE_NONE       CAVStreamParseType = C.AVSTREAM_PARSE_NONE
	AVSTREAM_PARSE_FULL       CAVStreamParseType = C.AVSTREAM_PARSE_FULL       /**< full parsing and repack */
	AVSTREAM_PARSE_HEADERS    CAVStreamParseType = C.AVSTREAM_PARSE_HEADERS    /**< Only parse headers, do not repack. */
	AVSTREAM_PARSE_TIMESTAMPS CAVStreamParseType = C.AVSTREAM_PARSE_TIMESTAMPS /**< full parsing and interpolation of timestamps for frames not starting on a packet boundary */
	AVSTREAM_PARSE_FULL_ONCE  CAVStreamParseType = C.AVSTREAM_PARSE_FULL_ONCE  /**< full parsing and repack of the first frame only, only implemented for H.264 currently */
	AVSTREAM_PARSE_FULL_RAW   CAVStreamParseType = C.AVSTREAM_PARSE_FULL_RAW   /**< full parsing and repack with timestamp and position generation by parser for raw
	  this assumes that each packet in the file contains no demuxer level headers and
	  just codec level data, otherwise position generation would fail */
)

type CAVIndexEntry C.AVIndexEntry

//#region CAVIndexEntry

func (e *CAVIndexEntry) GetPos() int64 {
	return int64(e.pos)
}

func (e *CAVIndexEntry) SetPos(pos int64) {
	e.pos = C.int64_t(pos)
}

/**<
 * Timestamp in AVStream.time_base units, preferably the time from which on correctly decoded frames are available
 * when seeking to this entry. That means preferable PTS on keyframe based formats.
 * But demuxers can choose to store a different timestamp, if it is more convenient for the implementation or nothing better
 * is known
 */
func (e *CAVIndexEntry) GetTimestamp() int64 {
	return int64(e.timestamp)
}

/**<
 * Timestamp in AVStream.time_base units, preferably the time from which on correctly decoded frames are available
 * when seeking to this entry. That means preferable PTS on keyframe based formats.
 * But demuxers can choose to store a different timestamp, if it is more convenient for the implementation or nothing better
 * is known
 */
func (e *CAVIndexEntry) SetTimestamp(timestamp int64) {
	e.timestamp = C.int64_t(timestamp)
}

const AVINDEX_KEYFRAME = C.AVINDEX_KEYFRAME

/**
 * Flag is used to indicate which frame should be discarded after decoding.
 */
const AVINDEX_DISCARD_FRAME = C.AVINDEX_DISCARD_FRAME

//TODO GO不可访问C的位域

// func (e *CAVIndexEntry) GetFlags() int {
// 	return int(e.flags)
// }

// // Yeah, trying to keep the size of this small to reduce memory requirements (it is 24 vs. 32 bytes due to possible 8-byte alignment).
// func (e *CAVIndexEntry) GetSize() int {
// 	return int(e.size)
// }

/**< Minimum distance between this and the previous keyframe, used to avoid unneeded searching. */
func (e *CAVIndexEntry) GetMinDistance() int {
	return int(e.min_distance)
}

/**< Minimum distance between this and the previous keyframe, used to avoid unneeded searching. */
func (e *CAVIndexEntry) SetMinDistance(minDistance int) {
	e.min_distance = C.int(minDistance)
}

//#endregion CAVIndexEntry

const (
	/**
	 * The stream should be chosen by default among other streams of the same type,
	 * unless the user has explicitly specified otherwise.
	 */
	AV_DISPOSITION_DEFAULT = C.AV_DISPOSITION_DEFAULT
	/**
	 * The stream is not in original language.
	 *
	 * @note AV_DISPOSITION_ORIGINAL is the inverse of this disposition. At most
	 *       one of them should be set in properly tagged streams.
	 * @note This disposition may apply to any stream type, not just audio.
	 */
	AV_DISPOSITION_DUB = C.AV_DISPOSITION_DUB
	/**
	 * The stream is in original language.
	 *
	 * @see the notes for AV_DISPOSITION_DUB
	 */
	AV_DISPOSITION_ORIGINAL = C.AV_DISPOSITION_ORIGINAL
	/**
	 * The stream is a commentary track.
	 */
	AV_DISPOSITION_COMMENT = C.AV_DISPOSITION_COMMENT
	/**
	 * The stream contains song lyrics.
	 */
	AV_DISPOSITION_LYRICS = C.AV_DISPOSITION_LYRICS
	/**
	 * The stream contains karaoke audio.
	 */
	AV_DISPOSITION_KARAOKE = C.AV_DISPOSITION_KARAOKE

	/**
	 * Track should be used during playback by default.
	 * Useful for subtitle track that should be displayed
	 * even when user did not explicitly ask for subtitles.
	 */
	AV_DISPOSITION_FORCED = C.AV_DISPOSITION_FORCED
	/**
	 * The stream is intended for hearing impaired audiences.
	 */
	AV_DISPOSITION_HEARING_IMPAIRED = C.AV_DISPOSITION_HEARING_IMPAIRED
	/**
	 * The stream is intended for visually impaired audiences.
	 */
	AV_DISPOSITION_VISUAL_IMPAIRED = C.AV_DISPOSITION_VISUAL_IMPAIRED
	/**
	 * The audio stream contains music and sound effects without voice.
	 */
	AV_DISPOSITION_CLEAN_EFFECTS = C.AV_DISPOSITION_CLEAN_EFFECTS
	/**
	 * The stream is stored in the file as an attached picture/"cover art" (e.g.
	 * APIC frame in ID3v2). The first (usually only) packet associated with it
	 * will be returned among the first few packets read from the file unless
	 * seeking takes place. It can also be accessed at any time in
	 * AVStream.attached_pic.
	 */
	AV_DISPOSITION_ATTACHED_PIC = C.AV_DISPOSITION_ATTACHED_PIC
	/**
	 * The stream is sparse, and contains thumbnail images, often corresponding
	 * to chapter markers. Only ever used with AV_DISPOSITION_ATTACHED_PIC.
	 */
	AV_DISPOSITION_TIMED_THUMBNAILS = C.AV_DISPOSITION_TIMED_THUMBNAILS

	/**
	 * The stream is intended to be mixed with a spatial audio track. For example,
	 * it could be used for narration or stereo music, and may remain unchanged by
	 * listener head rotation.
	 */
	AV_DISPOSITION_NON_DIEGETIC = C.AV_DISPOSITION_NON_DIEGETIC

	/**
	 * The subtitle stream contains captions, providing a transcription and possibly
	 * a translation of audio. Typically intended for hearing-impaired audiences.
	 */
	AV_DISPOSITION_CAPTIONS = C.AV_DISPOSITION_CAPTIONS
	/**
	 * The subtitle stream contains a textual description of the video content.
	 * Typically intended for visually-impaired audiences or for the cases where the
	 * video cannot be seen.
	 */
	AV_DISPOSITION_DESCRIPTIONS = C.AV_DISPOSITION_DESCRIPTIONS
	/**
	 * The subtitle stream contains time-aligned metadata that is not intended to be
	 * directly presented to the user.
	 */
	AV_DISPOSITION_METADATA = C.AV_DISPOSITION_METADATA
	/**
	 * The audio stream is intended to be mixed with another stream before
	 * presentation.
	 * Corresponds to mix_type=0 in mpegts.
	 */
	AV_DISPOSITION_DEPENDENT = C.AV_DISPOSITION_DEPENDENT
	/**
	 * The video stream contains still images.
	 */
	AV_DISPOSITION_STILL_IMAGE = C.AV_DISPOSITION_STILL_IMAGE
)

/**
 * @return The AV_DISPOSITION_* flag corresponding to disp or a negative error
 *         code if disp does not correspond to a known stream disposition.
 */
func AvDispositionFromString(disp string) int {
	var cDisp *C.char = nil
	if len(disp) > 0 {
		cDisp = C.CString(disp)
		defer C.free(unsafe.Pointer(cDisp))
	}

	return int(C.av_disposition_from_string(cDisp))
}

/**
 * @param disposition a combination of AV_DISPOSITION_* values
 * @return The string description corresponding to the lowest set bit in
 *         disposition. NULL when the lowest set bit does not correspond
 *         to a known disposition or when disposition is 0.
 */
func AvDispositionToString(disposition int) string {
	return C.GoString(C.av_disposition_to_string(C.int(disposition)))
}

/**
 * Options for behavior on timestamp wrap detection.
 */
const (
	AV_PTS_WRAP_IGNORE     = C.AV_PTS_WRAP_IGNORE     ///< ignore the wrap
	AV_PTS_WRAP_ADD_OFFSET = C.AV_PTS_WRAP_ADD_OFFSET ///< add the format specific offset on wrap detection
	AV_PTS_WRAP_SUB_OFFSET = C.AV_PTS_WRAP_SUB_OFFSET ///< subtract the format specific offset on wrap detection
)

/**
 * Stream structure.
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVStream) must not be used outside libav*.
 */
type CAVStream C.AVStream

//#region CAVStream

/**
 * A class for @ref avoptions. Set on stream creation.
 */
func (st *CAVStream) GetAvClass() *avutil.CAVClass {
	return (*avutil.CAVClass)(unsafe.Pointer(st.av_class))
}

/**< stream index in AVFormatContext */
func (st *CAVStream) GetIndex() int {
	return int(st.index)
}

/**< stream index in AVFormatContext */
func (st *CAVStream) SetIndex(index int) {
	st.index = C.int(index)
}

/**
 * Format-specific stream ID.
 * decoding: set by libavformat
 * encoding: set by the user, replaced by libavformat if left unset
 */
func (st *CAVStream) GetId() int {
	return int(st.id)
}

/**
 * Format-specific stream ID.
 * decoding: set by libavformat
 * encoding: set by the user, replaced by libavformat if left unset
 */
func (st *CAVStream) SetId(id int) {
	st.id = C.int(id)
}

/**
 * Codec parameters associated with this stream. Allocated and freed by
 * libavformat in avformat_new_stream() and avformat_free_context()
 * respectively.
 *
 * - demuxing: filled by libavformat on stream creation or in
 *             avformat_find_stream_info()
 * - muxing: filled by the caller before avformat_write_header()
 */
func (st *CAVStream) GetCodecPar() *avcodec.CAVCodecParameters {
	return (*avcodec.CAVCodecParameters)(unsafe.Pointer(st.codecpar))
}

/**
 * Codec parameters associated with this stream. Allocated and freed by
 * libavformat in avformat_new_stream() and avformat_free_context()
 * respectively.
 *
 * - demuxing: filled by libavformat on stream creation or in
 *             avformat_find_stream_info()
 * - muxing: filled by the caller before avformat_write_header()
 */
func (st *CAVStream) SetCodecPar(codecPar *avcodec.CAVCodecParameters) {
	st.codecpar = (*C.AVCodecParameters)(unsafe.Pointer(codecPar))
}

func (st *CAVStream) GetPrivData() unsafe.Pointer {
	return st.priv_data
}

func (st *CAVStream) SetPrivData(privData unsafe.Pointer) {
	st.priv_data = privData
}

/**
 * This is the fundamental unit of time (in seconds) in terms
 * of which frame timestamps are represented.
 *
 * decoding: set by libavformat
 * encoding: May be set by the caller before avformat_write_header() to
 *           provide a hint to the muxer about the desired timebase. In
 *           avformat_write_header(), the muxer will overwrite this field
 *           with the timebase that will actually be used for the timestamps
 *           written into the file (which may or may not be related to the
 *           user-provided one, depending on the format).
 */
func (st *CAVStream) GetTimeBase() avutil.CAVRational {
	return *(*avutil.CAVRational)(unsafe.Pointer(&st.time_base))
}

/**
 * This is the fundamental unit of time (in seconds) in terms
 * of which frame timestamps are represented.
 *
 * decoding: set by libavformat
 * encoding: May be set by the caller before avformat_write_header() to
 *           provide a hint to the muxer about the desired timebase. In
 *           avformat_write_header(), the muxer will overwrite this field
 *           with the timebase that will actually be used for the timestamps
 *           written into the file (which may or may not be related to the
 *           user-provided one, depending on the format).
 */
func (st *CAVStream) GetTimeBasePtr() *avutil.CAVRational {
	return (*avutil.CAVRational)(unsafe.Pointer(&st.time_base))
}

/**
 * This is the fundamental unit of time (in seconds) in terms
 * of which frame timestamps are represented.
 *
 * decoding: set by libavformat
 * encoding: May be set by the caller before avformat_write_header() to
 *           provide a hint to the muxer about the desired timebase. In
 *           avformat_write_header(), the muxer will overwrite this field
 *           with the timebase that will actually be used for the timestamps
 *           written into the file (which may or may not be related to the
 *           user-provided one, depending on the format).
 */
func (st *CAVStream) SetTimeBase(timeBase avutil.CAVRational) {
	st.time_base = *(*C.AVRational)(unsafe.Pointer(&timeBase))
}

/**
 * Decoding: pts of the first frame of the stream in presentation order, in stream time base.
 * Only set this if you are absolutely 100% sure that the value you set
 * it to really is the pts of the first frame.
 * This may be undefined (AV_NOPTS_VALUE).
 * @note The ASF header does NOT contain a correct start_time the ASF
 * demuxer must NOT set this.
 */
func (st *CAVStream) GetStartTime() int64 {
	return int64(st.start_time)
}

/**
 * Decoding: pts of the first frame of the stream in presentation order, in stream time base.
 * Only set this if you are absolutely 100% sure that the value you set
 * it to really is the pts of the first frame.
 * This may be undefined (AV_NOPTS_VALUE).
 * @note The ASF header does NOT contain a correct start_time the ASF
 * demuxer must NOT set this.
 */
func (st *CAVStream) SetStartTime(startTime int64) {
	st.start_time = C.int64_t(startTime)
}

/**
 * Decoding: duration of the stream, in stream time base.
 * If a source file does not specify a duration, but does specify
 * a bitrate, this value will be estimated from bitrate and file size.
 *
 * Encoding: May be set by the caller before avformat_write_header() to
 * provide a hint to the muxer about the estimated duration.
 */
func (st *CAVStream) GetDuration() int64 {
	return int64(st.duration)
}

/**
 * Decoding: duration of the stream, in stream time base.
 * If a source file does not specify a duration, but does specify
 * a bitrate, this value will be estimated from bitrate and file size.
 *
 * Encoding: May be set by the caller before avformat_write_header() to
 * provide a hint to the muxer about the estimated duration.
 */
func (st *CAVStream) SetDuration(duration int64) {
	st.duration = C.int64_t(duration)
}

// /< number of frames in this stream if known or 0
func (st *CAVStream) GetNbFrames() int64 {
	return int64(st.nb_frames)
}

// /< number of frames in this stream if known or 0
func (st *CAVStream) SetNbFrames(nbFrames int64) {
	st.nb_frames = C.int64_t(nbFrames)
}

/**
 * Stream disposition - a combination of AV_DISPOSITION_* flags.
 * - demuxing: set by libavformat when creating the stream or in
 *             avformat_find_stream_info().
 * - muxing: may be set by the caller before avformat_write_header().
 */
func (st *CAVStream) GetDisposition() int {
	return int(st.disposition)
}

/**
 * Stream disposition - a combination of AV_DISPOSITION_* flags.
 * - demuxing: set by libavformat when creating the stream or in
 *             avformat_find_stream_info().
 * - muxing: may be set by the caller before avformat_write_header().
 */
func (st *CAVStream) SetDisposition(disposition int) {
	st.disposition = C.int(disposition)
}

// /< Selects which packets can be discarded at will and do not need to be demuxed.
func (st *CAVStream) GetDiscard() avcodec.CAVDiscard {
	return avcodec.CAVDiscard(st.discard)
}

// /< Selects which packets can be discarded at will and do not need to be demuxed.
func (st *CAVStream) SetDiscard(discard avcodec.CAVDiscard) {
	st.discard = C.enum_AVDiscard(discard)
}

/**
 * sample aspect ratio (0 if unknown)
 * - encoding: Set by user.
 * - decoding: Set by libavformat.
 */
func (st *CAVStream) GetSampleAspectRatio() avutil.CAVRational {
	return *(*avutil.CAVRational)(unsafe.Pointer(&st.sample_aspect_ratio))
}

/**
 * sample aspect ratio (0 if unknown)
 * - encoding: Set by user.
 * - decoding: Set by libavformat.
 */
func (st *CAVStream) SetSampleAspectRatio(sampleAspectRatio avutil.CAVRational) {
	st.sample_aspect_ratio = *(*C.AVRational)(unsafe.Pointer(&sampleAspectRatio))
}

func (st *CAVStream) GetMetadata() *avutil.CAVDictionary {
	return (*avutil.CAVDictionary)(st.metadata)
}

func (st *CAVStream) SetMetadata(metadata *avutil.CAVDictionary) {
	st.metadata = (*C.AVDictionary)(metadata)
}

/**
 * Average framerate
 *
 * - demuxing: May be set by libavformat when creating the stream or in
 *             avformat_find_stream_info().
 * - muxing: May be set by the caller before avformat_write_header().
 */
func (st *CAVStream) GetAvgFrameRate() avutil.CAVRational {
	return *(*avutil.CAVRational)(unsafe.Pointer(&st.avg_frame_rate))
}

/**
 * Average framerate
 *
 * - demuxing: May be set by libavformat when creating the stream or in
 *             avformat_find_stream_info().
 * - muxing: May be set by the caller before avformat_write_header().
 */
func (st *CAVStream) SetAvgFrameRate(avgFrameRate avutil.CAVRational) {
	st.avg_frame_rate = *(*C.AVRational)(unsafe.Pointer(&avgFrameRate))
}

/**
 * For streams with AV_DISPOSITION_ATTACHED_PIC disposition, this packet
 * will contain the attached picture.
 *
 * decoding: set by libavformat, must not be modified by the caller.
 * encoding: unused
 */
func (st *CAVStream) GetAttachedPic() avcodec.CAVPacket {
	return *(*avcodec.CAVPacket)(unsafe.Pointer(&st.attached_pic))
}

/**
 * For streams with AV_DISPOSITION_ATTACHED_PIC disposition, this packet
 * will contain the attached picture.
 *
 * decoding: set by libavformat, must not be modified by the caller.
 * encoding: unused
 */
func (st *CAVStream) GetAttachedPicPtr() *avcodec.CAVPacket {
	return (*avcodec.CAVPacket)(unsafe.Pointer(&st.attached_pic))
}

//  #if FF_API_AVSTREAM_SIDE_DATA
// 	 /**
// 	  * An array of side data that applies to the whole stream (i.e. the
// 	  * container does not allow it to change between packets).
// 	  *
// 	  * There may be no overlap between the side data in this array and side data
// 	  * in the packets. I.e. a given side data is either exported by the muxer
// 	  * (demuxing) / set by the caller (muxing) in this array, then it never
// 	  * appears in the packets, or the side data is exported / sent through
// 	  * the packets (always in the first packet where the value becomes known or
// 	  * changes), then it does not appear in this array.
// 	  *
// 	  * - demuxing: Set by libavformat when the stream is created.
// 	  * - muxing: May be set by the caller before avformat_write_header().
// 	  *
// 	  * Freed by libavformat in avformat_free_context().
// 	  *
// 	  * @deprecated use AVStream's @ref AVCodecParameters.coded_side_data
// 	  *             "codecpar side data".
// 	  */
// 	 attribute_deprecated
// 	 AVPacketSideData *side_data;
// 	 /**
// 	  * The number of elements in the AVStream.side_data array.
// 	  *
// 	  * @deprecated use AVStream's @ref AVCodecParameters.nb_coded_side_data
// 	  *             "codecpar side data".
// 	  */
// 	 attribute_deprecated
// 	 int            nb_side_data;
//  #endif

/**
 * Flags indicating events happening on the stream, a combination of
 * AVSTREAM_EVENT_FLAG_*.
 *
 * - demuxing: may be set by the demuxer in avformat_open_input(),
 *   avformat_find_stream_info() and av_read_frame(). Flags must be cleared
 *   by the user once the event has been handled.
 * - muxing: may be set by the user after avformat_write_header(). to
 *   indicate a user-triggered event.  The muxer will clear the flags for
 *   events it has handled in av_[interleaved]_write_frame().
 */
func (st *CAVStream) GetEventFlags() int {
	return int(st.event_flags)
}

/**
 * Flags indicating events happening on the stream, a combination of
 * AVSTREAM_EVENT_FLAG_*.
 *
 * - demuxing: may be set by the demuxer in avformat_open_input(),
 *   avformat_find_stream_info() and av_read_frame(). Flags must be cleared
 *   by the user once the event has been handled.
 * - muxing: may be set by the user after avformat_write_header(). to
 *   indicate a user-triggered event.  The muxer will clear the flags for
 *   events it has handled in av_[interleaved]_write_frame().
 */
func (st *CAVStream) SetEventFlags(eventFlags int) {
	st.event_flags = C.int(eventFlags)
}

/**
 * - demuxing: the demuxer read new metadata from the file and updated
 *     AVStream.metadata accordingly
 * - muxing: the user updated AVStream.metadata and wishes the muxer to write
 *     it into the file
 */
const AVSTREAM_EVENT_FLAG_METADATA_UPDATED = C.AVSTREAM_EVENT_FLAG_METADATA_UPDATED

/**
 * - demuxing: new packets for this stream were read from the file. This
 *   event is informational only and does not guarantee that new packets
 *   for this stream will necessarily be returned from av_read_frame().
 */
const AVSTREAM_EVENT_FLAG_NEW_PACKETS = C.AVSTREAM_EVENT_FLAG_NEW_PACKETS

/**
 * Real base framerate of the stream.
 * This is the lowest framerate with which all timestamps can be
 * represented accurately (it is the least common multiple of all
 * framerates in the stream). Note, this value is just a guess!
 * For example, if the time base is 1/90000 and all frames have either
 * approximately 3600 or 1800 timer ticks, then r_frame_rate will be 50/1.
 */
func (st *CAVStream) GetRFrameRate() avutil.CAVRational {
	return *(*avutil.CAVRational)(unsafe.Pointer(&st.r_frame_rate))
}

/**
 * Real base framerate of the stream.
 * This is the lowest framerate with which all timestamps can be
 * represented accurately (it is the least common multiple of all
 * framerates in the stream). Note, this value is just a guess!
 * For example, if the time base is 1/90000 and all frames have either
 * approximately 3600 or 1800 timer ticks, then r_frame_rate will be 50/1.
 */
func (st *CAVStream) SetRFrameRate(rFrameRate avutil.CAVRational) {
	st.r_frame_rate = *(*C.AVRational)(unsafe.Pointer(&rFrameRate))
}

/**
 * Number of bits in timestamps. Used for wrapping control.
 *
 * - demuxing: set by libavformat
 * - muxing: set by libavformat
 *
 */
func (st *CAVStream) GetPtsWrapBits() int {
	return int(st.pts_wrap_bits)
}

/**
 * Number of bits in timestamps. Used for wrapping control.
 *
 * - demuxing: set by libavformat
 * - muxing: set by libavformat
 *
 */
func (st *CAVStream) SetPtsWrapBits(ptsWrapBit int) {
	st.pts_wrap_bits = C.int(ptsWrapBit)
}

//#endregion CAVStream

func AvStreamGetParser(s *CAVStream) *avcodec.CAVCodecParserContext {
	return (*avcodec.CAVCodecParserContext)(unsafe.Pointer(C.av_stream_get_parser((*C.AVStream)(s))))
}

// #if FF_API_GET_END_PTS
// /**
//  * Returns the pts of the last muxed packet + its duration
//  *
//  * the retuned value is undefined when used with a demuxer.
//  */
// attribute_deprecated
// int64_t    av_stream_get_end_pts(const AVStream *st);
// #endif

const AV_PROGRAM_RUNNING = C.AV_PROGRAM_RUNNING

/**
 * New fields can be added to the end with minor version bumps.
 * Removal, reordering and changes to existing fields require a major
 * version bump.
 * sizeof(AVProgram) must not be used outside libav*.
 */
type CAVProgram C.AVProgram

// #region CAVProgram

func (p *CAVProgram) GetId() int {
	return int(p.id)
}
func (p *CAVProgram) SetId(id int) {
	p.id = C.int(id)
}

func (p *CAVProgram) GetFlags() int {
	return int(p.flags)
}
func (p *CAVProgram) SetFlags(flags int) {
	p.flags = C.int(flags)
}

// /< selects which program to discard and which to feed to the caller
func (p *CAVProgram) GetDiscard() avcodec.CAVDiscard {
	return avcodec.CAVDiscard(p.discard)
}

// /< selects which program to discard and which to feed to the caller
func (p *CAVProgram) SetDiscard(discard avcodec.CAVDiscard) {
	p.discard = C.enum_AVDiscard(discard)
}

func (p *CAVProgram) GetStreamIndex() *ctypes.UInt {
	return (*ctypes.UInt)(p.stream_index)
}
func (p *CAVProgram) SetStreamIndex(streamIndex *ctypes.UInt) {
	p.stream_index = (*C.uint)(streamIndex)
}

func (p *CAVProgram) GetNbStreamIndexes() uint {
	return uint(p.nb_stream_indexes)
}
func (p *CAVProgram) SetNbStreamIndexes(nbStreamIndexs uint) {
	p.nb_stream_indexes = C.uint(nbStreamIndexs)
}

func (p *CAVProgram) GetMetadata() *avutil.CAVDictionary {
	return (*avutil.CAVDictionary)(p.metadata)
}
func (p *CAVProgram) SetMetadata(metadata *avutil.CAVDictionary) {
	p.metadata = (*C.AVDictionary)(metadata)
}

func (p *CAVProgram) GetProgramNum() int {
	return int(p.program_num)
}
func (p *CAVProgram) SetProgramNum(programNum int) {
	p.program_num = C.int(programNum)
}

func (p *CAVProgram) GetPmtPid() int {
	return int(p.pmt_pid)
}
func (p *CAVProgram) SetPmtPid(pmtPid int) {
	p.pmt_pid = C.int(pmtPid)
}

func (p *CAVProgram) GetPcrPid() int {
	return int(p.pcr_pid)
}
func (p *CAVProgram) SetPcrPid(pcrPid int) {
	p.pcr_pid = C.int(pcrPid)
}

func (p *CAVProgram) GetPmtVersion() int {
	return int(p.pmt_version)
}
func (p *CAVProgram) SetPmtVersion(pmtVersion int) {
	p.pmt_version = C.int(pmtVersion)
}

/*****************************************************************
 * All fields below this line are not part of the public API. They
 * may not be used outside of libavformat and can be changed and
 * removed at will.
 * New public fields should be added right above.
 *****************************************************************
 */

func (p *CAVProgram) GetStartTime() int64 {
	return int64(p.start_time)
}
func (p *CAVProgram) SetStartTime(startTime int64) {
	p.start_time = C.int64_t(startTime)
}

func (p *CAVProgram) GetEndTime() int64 {
	return int64(p.end_time)
}
func (p *CAVProgram) SetEndTime(endTime int64) {
	p.end_time = C.int64_t(endTime)
}

// /< reference dts for wrap detection
func (p *CAVProgram) GetPtsWrapRefrence() int64 {
	return int64(p.pts_wrap_reference)
}

// /< reference dts for wrap detection
func (p *CAVProgram) SetPtsWrapRefrence(ptsWrapReference int64) {
	p.pts_wrap_reference = C.int64_t(ptsWrapReference)
}

// /< behavior on wrap detection
func (p *CAVProgram) GetPtsWrapBehavior() int {
	return int(p.pts_wrap_behavior)
}

// /< behavior on wrap detection
func (p *CAVProgram) SetPtsWrapBehavior(ptsWrapBehavior int) {
	p.pts_wrap_behavior = C.int(ptsWrapBehavior)
}

//#endregion CAVProgram

const AVFMTCTX_NOHEADER = C.AVFMTCTX_NOHEADER /**< signal that no header is present
  (streams are added dynamically) */
const AVFMTCTX_UNSEEKABLE = C.AVFMTCTX_UNSEEKABLE /**< signal that the stream is definitely
  not seekable, and attempts to call the
  seek function will fail. For some
  network protocols (e.g. HLS), this can
  change dynamically at runtime. */

type CAVChapter C.AVChapter

//#region CAVChapter

// /< unique ID to identify the chapter
func (c *CAVChapter) GetId() int64 {
	return int64(c.id)
}

// /< unique ID to identify the chapter
func (c *CAVChapter) SetId(id int64) {
	c.id = C.int64_t(id)
}

// /< time base in which the start/end timestamps are specified
func (c *CAVChapter) GetTimeBase() avutil.CAVRational {
	return *(*avutil.CAVRational)(unsafe.Pointer(&c.time_base))
}

// /< time base in which the start/end timestamps are specified
func (c *CAVChapter) SetTimeBase(timeBase avutil.CAVRational) {
	c.time_base = *(*C.AVRational)(unsafe.Pointer(&timeBase))
}

// /< chapter start/end time in time_base units
func (c *CAVChapter) GetStart() int64 {
	return int64(c.start)
}

// /< chapter start/end time in time_base units
func (c *CAVChapter) SetStart(start int64) {
	c.start = C.int64_t(start)
}

// /< chapter start/end time in time_base units
func (c *CAVChapter) GetEnd() int64 {
	return int64(c.end)
}

// /< chapter start/end time in time_base units
func (c *CAVChapter) SetEnd(end int64) {
	c.end = C.int64_t(end)
}

func (c *CAVChapter) GetMetadata() *avutil.CAVDictionary {
	return (*avutil.CAVDictionary)(c.metadata)
}
func (c *CAVChapter) SetMetadata(metadata *avutil.CAVDictionary) {
	c.metadata = (*C.AVDictionary)(metadata)
}

//#endregion CAVChapter

/**
 * Callback used by devices to communicate with application.
 */

//	 typedef int (*av_format_control_message)(struct AVFormatContext *s, int type,
//		void *data, size_t data_size);
type CAvFormatControlMessage C.av_format_control_message

//	 typedef int (*AVOpenCallback)(struct AVFormatContext *s, AVIOContext **pb, const char *url, int flags,
//		const AVIOInterruptCB *int_cb, AVDictionary **options);
type CAVOpenCallback ctypes.CFunc

/**
* The duration of a video can be estimated through various ways, and this enum can be used
* to know how the duration was estimated.
 */
type CAVDurationEstimationMethod C.enum_AVDurationEstimationMethod

const (
	AVFMT_DURATION_FROM_PTS     CAVDurationEstimationMethod = C.AVFMT_DURATION_FROM_PTS     ///< Duration accurately estimated from PTSes
	AVFMT_DURATION_FROM_STREAM  CAVDurationEstimationMethod = C.AVFMT_DURATION_FROM_STREAM  ///< Duration estimated from a stream with a known duration
	AVFMT_DURATION_FROM_BITRATE CAVDurationEstimationMethod = C.AVFMT_DURATION_FROM_BITRATE ///< Duration estimated from bitrate (less accurate)
)

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

//#region CAVFormatContext

/**
* A class for logging and @ref avoptions. Set by avformat_alloc_context().
* Exports (de)muxer private options if they exist.
 */
func (fmtCtx *CAVFormatContext) GetAvClass() *avutil.CAVClass {
	return (*avutil.CAVClass)(unsafe.Pointer(fmtCtx.av_class))
}

/**
 * The input container format.
 *
 * Demuxing only, set by avformat_open_input().
 */
func (fmtCtx *CAVFormatContext) GetIformat() *CAVInputFormat {
	return (*CAVInputFormat)(fmtCtx.iformat)
}

/**
 * The output container format.
 *
 * Muxing only, must be set by the caller before avformat_write_header().
 */
func (fmtCtx *CAVFormatContext) GetOformat() *CAVOutputFormat {
	return (*CAVOutputFormat)(fmtCtx.oformat)
}

/**
 * The output container format.
 *
 * Muxing only, must be set by the caller before avformat_write_header().
 */
func (fmtCtx *CAVFormatContext) SetOformat(oformat *CAVOutputFormat) {
	fmtCtx.oformat = (*C.AVOutputFormat)(oformat)
}

/**
 * Format private data. This is an AVOptions-enabled struct
 * if and only if iformat/oformat.priv_class is not NULL.
 *
 * - muxing: set by avformat_write_header()
 * - demuxing: set by avformat_open_input()
 */
func (fmtCtx *CAVFormatContext) GetPrivData() unsafe.Pointer {
	return fmtCtx.priv_data
}

/**
 * Format private data. This is an AVOptions-enabled struct
 * if and only if iformat/oformat.priv_class is not NULL.
 *
 * - muxing: set by avformat_write_header()
 * - demuxing: set by avformat_open_input()
 */
func (fmtCtx *CAVFormatContext) SetPrivData(privData unsafe.Pointer) {
	fmtCtx.priv_data = privData
}

/**
 * I/O context.
 *
 * - demuxing: either set by the user before avformat_open_input() (then
 *             the user must close it manually) or set by avformat_open_input().
 * - muxing: set by the user before avformat_write_header(). The caller must
 *           take care of closing / freeing the IO context.
 *
 * Do NOT set this field if AVFMT_NOFILE flag is set in
 * iformat/oformat.flags. In such a case, the (de)muxer will handle
 * I/O in some other way and this field will be NULL.
 */
func (fmtCtx *CAVFormatContext) GetPb() *CAVIOContext {
	return (*CAVIOContext)(fmtCtx.pb)
}

/**
 * I/O context.
 *
 * - demuxing: either set by the user before avformat_open_input() (then
 *             the user must close it manually) or set by avformat_open_input().
 * - muxing: set by the user before avformat_write_header(). The caller must
 *           take care of closing / freeing the IO context.
 *
 * Do NOT set this field if AVFMT_NOFILE flag is set in
 * iformat/oformat.flags. In such a case, the (de)muxer will handle
 * I/O in some other way and this field will be NULL.
 */
func (fmtCtx *CAVFormatContext) GetPbPtr() **CAVIOContext {
	return (**CAVIOContext)(unsafe.Pointer(&fmtCtx.pb))
}

/**
 * I/O context.
 *
 * - demuxing: either set by the user before avformat_open_input() (then
 *             the user must close it manually) or set by avformat_open_input().
 * - muxing: set by the user before avformat_write_header(). The caller must
 *           take care of closing / freeing the IO context.
 *
 * Do NOT set this field if AVFMT_NOFILE flag is set in
 * iformat/oformat.flags. In such a case, the (de)muxer will handle
 * I/O in some other way and this field will be NULL.
 */
func (fmtCtx *CAVFormatContext) SetPb(avioCtx *CAVIOContext) {
	fmtCtx.pb = (*C.AVIOContext)(avioCtx)
}

/* stream info */

/**
 * Flags signalling stream properties. A combination of AVFMTCTX_*.
 * Set by libavformat.
 */
func (fmtCtx *CAVFormatContext) GetCtxFlags() int {
	return int(fmtCtx.ctx_flags)
}

/**
 * Flags signalling stream properties. A combination of AVFMTCTX_*.
 * Set by libavformat.
 */
func (fmtCtx *CAVFormatContext) SetCtxFlags(ctxFlags int) {
	fmtCtx.ctx_flags = C.int(ctxFlags)
}

/**
 * Number of elements in AVFormatContext.streams.
 *
 * Set by avformat_new_stream(), must not be modified by any other code.
 */
func (fmtCtx *CAVFormatContext) GetNbStreams() uint {
	return uint(fmtCtx.nb_streams)
}

/**
 * Number of elements in AVFormatContext.streams.
 *
 * Set by avformat_new_stream(), must not be modified by any other code.
 */
func (fmtCtx *CAVFormatContext) SetNbStreams(nbStreams uint) {
	fmtCtx.nb_streams = C.uint(nbStreams)
}

/**
 * A list of all streams in the file. New streams are created with
 * avformat_new_stream().
 *
 * - demuxing: streams are created by libavformat in avformat_open_input().
 *             If AVFMTCTX_NOHEADER is set in ctx_flags, then new streams may also
 *             appear in av_read_frame().
 * - muxing: streams are created by the user before avformat_write_header().
 *
 * Freed by libavformat in avformat_free_context().
 */
func (fmtCtx *CAVFormatContext) GetStreams() **CAVStream {
	return (**CAVStream)(unsafe.Pointer(fmtCtx.streams))
}

/**
 * A list of all streams in the file. New streams are created with
 * avformat_new_stream().
 *
 * - demuxing: streams are created by libavformat in avformat_open_input().
 *             If AVFMTCTX_NOHEADER is set in ctx_flags, then new streams may also
 *             appear in av_read_frame().
 * - muxing: streams are created by the user before avformat_write_header().
 *
 * Freed by libavformat in avformat_free_context().
 */
func (fmtCtx *CAVFormatContext) SetStreams(streams **CAVStream) {
	fmtCtx.streams = (**C.AVStream)(unsafe.Pointer(streams))
}

/**
 * input or output URL. Unlike the old filename field, this field has no
 * length restriction.
 *
 * - demuxing: set by avformat_open_input(), initialized to an empty
 *             string if url parameter was NULL in avformat_open_input().
 * - muxing: may be set by the caller before calling avformat_write_header()
 *           (or avformat_init_output() if that is called first) to a string
 *           which is freeable by av_free(). Set to an empty string if it
 *           was NULL in avformat_init_output().
 *
 * Freed by libavformat in avformat_free_context().
 */
func (fmtCtx *CAVFormatContext) GetUrl() string {
	return C.GoString(fmtCtx.url)
}

/**
 * input or output URL. Unlike the old filename field, this field has no
 * length restriction.
 *
 * - demuxing: set by avformat_open_input(), initialized to an empty
 *             string if url parameter was NULL in avformat_open_input().
 * - muxing: may be set by the caller before calling avformat_write_header()
 *           (or avformat_init_output() if that is called first) to a string
 *           which is freeable by av_free(). Set to an empty string if it
 *           was NULL in avformat_init_output().
 *
 * Freed by libavformat in avformat_free_context().
 */
func (fmtCtx *CAVFormatContext) SetUrl(url string) {
	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
	}
	fmtCtx.url = cUrl
}

/**
 * Position of the first frame of the component, in
 * AV_TIME_BASE fractional seconds. NEVER set this value directly:
 * It is deduced from the AVStream values.
 *
 * Demuxing only, set by libavformat.
 */
func (fmtCtx *CAVFormatContext) GetStartTime() int64 {
	return int64(fmtCtx.start_time)
}

/**
 * Position of the first frame of the component, in
 * AV_TIME_BASE fractional seconds. NEVER set this value directly:
 * It is deduced from the AVStream values.
 *
 * Demuxing only, set by libavformat.
 */
func (fmtCtx *CAVFormatContext) SetStartTime(startTime int64) {
	fmtCtx.start_time = C.int64_t(startTime)
}

/**
 * Duration of the stream, in AV_TIME_BASE fractional
 * seconds. Only set this value if you know none of the individual stream
 * durations and also do not set any of them. This is deduced from the
 * AVStream values if not set.
 *
 * Demuxing only, set by libavformat.
 */
func (fmtCtx *CAVFormatContext) GetDuration() int64 {
	return int64(fmtCtx.duration)
}

/**
 * Duration of the stream, in AV_TIME_BASE fractional
 * seconds. Only set this value if you know none of the individual stream
 * durations and also do not set any of them. This is deduced from the
 * AVStream values if not set.
 *
 * Demuxing only, set by libavformat.
 */
func (fmtCtx *CAVFormatContext) SetDuration(duration int64) {
	fmtCtx.duration = C.int64_t(duration)
}

/**
 * Total stream bitrate in bit/s, 0 if not
 * available. Never set it directly if the file_size and the
 * duration are known as FFmpeg can compute it automatically.
 */
func (fmtCtx *CAVFormatContext) GetBitRate() int64 {
	return int64(fmtCtx.bit_rate)
}

/**
 * Total stream bitrate in bit/s, 0 if not
 * available. Never set it directly if the file_size and the
 * duration are known as FFmpeg can compute it automatically.
 */
func (fmtCtx *CAVFormatContext) SetBitRate(bitRate int64) {
	fmtCtx.bit_rate = C.int64_t(bitRate)
}

func (fmtCtx *CAVFormatContext) GetPacketSize() uint {
	return uint(fmtCtx.packet_size)
}
func (fmtCtx *CAVFormatContext) SetPacketSize(packetSize uint) {
	fmtCtx.packet_size = C.uint(packetSize)
}

func (fmtCtx *CAVFormatContext) GetMaxDelay() int {
	return int(fmtCtx.max_delay)
}
func (fmtCtx *CAVFormatContext) SetMaxDelay(maxDelay int) {
	fmtCtx.max_delay = C.int(maxDelay)
}

/**
 * Flags modifying the (de)muxer behaviour. A combination of AVFMT_FLAG_*.
 * Set by the user before avformat_open_input() / avformat_write_header().
 */
func (fmtCtx *CAVFormatContext) GetFlags() int {
	return int(fmtCtx.flags)
}

/**
 * Flags modifying the (de)muxer behaviour. A combination of AVFMT_FLAG_*.
 * Set by the user before avformat_open_input() / avformat_write_header().
 */
func (fmtCtx *CAVFormatContext) SetFlags(flags int) {
	fmtCtx.flags = C.int(flags)
}

const (
	AVFMT_FLAG_GENPTS          = C.AVFMT_FLAG_GENPTS          ///< Generate missing pts even if it requires parsing future frames.
	AVFMT_FLAG_IGNIDX          = C.AVFMT_FLAG_IGNIDX          ///< Ignore index.
	AVFMT_FLAG_NONBLOCK        = C.AVFMT_FLAG_NONBLOCK        ///< Do not block when reading packets from input.
	AVFMT_FLAG_IGNDTS          = C.AVFMT_FLAG_IGNDTS          ///< Ignore DTS on frames that contain both DTS & PTS
	AVFMT_FLAG_NOFILLIN        = C.AVFMT_FLAG_NOFILLIN        ///< Do not infer any values from other values, just return what is stored in the container
	AVFMT_FLAG_NOPARSE         = C.AVFMT_FLAG_NOPARSE         ///< Do not use AVParsers, you also must set AVFMT_FLAG_NOFILLIN as the fillin code works on frames and no parsing -> no frames. Also seeking to frames can not work if parsing to find frame boundaries has been disabled
	AVFMT_FLAG_NOBUFFER        = C.AVFMT_FLAG_NOBUFFER        ///< Do not buffer frames when possible
	AVFMT_FLAG_CUSTOM_IO       = C.AVFMT_FLAG_CUSTOM_IO       ///< The caller has supplied a custom AVIOContext, don't avio_close() it.
	AVFMT_FLAG_DISCARD_CORRUPT = C.AVFMT_FLAG_DISCARD_CORRUPT ///< Discard frames marked corrupted
	AVFMT_FLAG_FLUSH_PACKETS   = C.AVFMT_FLAG_FLUSH_PACKETS   ///< Flush the AVIOContext every packet.
	/**
	 * When muxing, try to avoid writing any random/volatile data to the output.
	 * This includes any random IDs, real-time timestamps/dates, muxer version, etc.
	 *
	 * This flag is mainly intended for testing.
	 */
	AVFMT_FLAG_BITEXACT  = C.AVFMT_FLAG_BITEXACT
	AVFMT_FLAG_SORT_DTS  = C.AVFMT_FLAG_SORT_DTS  ///< try to interleave outputted packets by dts (using this flag can slow demuxing down)
	AVFMT_FLAG_FAST_SEEK = C.AVFMT_FLAG_FAST_SEEK ///< Enable fast, but inaccurate seeks for some formats
	//	#if FF_API_LAVF_SHORTEST
	AVFMT_FLAG_SHORTEST = C.AVFMT_FLAG_SHORTEST ///< Stop muxing when the shortest stream stops.
	//  #endif
	AVFMT_FLAG_AUTO_BSF = C.AVFMT_FLAG_AUTO_BSF ///< Add bitstream filters as requested by the muxer
)

/**
 * Maximum number of bytes read from input in order to determine stream
 * properties. Used when reading the global header and in
 * avformat_find_stream_info().
 *
 * Demuxing only, set by the caller before avformat_open_input().
 *
 * @note this is \e not  used for determining the \ref AVInputFormat
 *       "input format"
 * @sa format_probesize
 */
func (fmtCtx *CAVFormatContext) GetProbesize() int64 {
	return int64(fmtCtx.probesize)
}

/**
 * Maximum number of bytes read from input in order to determine stream
 * properties. Used when reading the global header and in
 * avformat_find_stream_info().
 *
 * Demuxing only, set by the caller before avformat_open_input().
 *
 * @note this is \e not  used for determining the \ref AVInputFormat
 *       "input format"
 * @sa format_probesize
 */
func (fmtCtx *CAVFormatContext) SetProbesize(probesize int64) {
	fmtCtx.probesize = C.int64_t(probesize)
}

/**
 * Maximum duration (in AV_TIME_BASE units) of the data read
 * from input in avformat_find_stream_info().
 * Demuxing only, set by the caller before avformat_find_stream_info().
 * Can be set to 0 to let avformat choose using a heuristic.
 */
func (fmtCtx *CAVFormatContext) GetMaxAnalyzeDuration() int64 {
	return int64(fmtCtx.max_analyze_duration)
}

/**
 * Maximum duration (in AV_TIME_BASE units) of the data read
 * from input in avformat_find_stream_info().
 * Demuxing only, set by the caller before avformat_find_stream_info().
 * Can be set to 0 to let avformat choose using a heuristic.
 */
func (fmtCtx *CAVFormatContext) SetMaxAnalyzeDuration(maxAnalyzeDuration int64) {
	fmtCtx.max_analyze_duration = C.int64_t(maxAnalyzeDuration)
}

func (fmtCtx *CAVFormatContext) GetKey() unsafe.Pointer {
	return unsafe.Pointer(fmtCtx.key)
}
func (fmtCtx *CAVFormatContext) SetKey(key unsafe.Pointer) {
	fmtCtx.key = (*C.uint8_t)(key)
}

func (fmtCtx *CAVFormatContext) GetKeylen() int {
	return int(fmtCtx.keylen)
}
func (fmtCtx *CAVFormatContext) SetKeylen(keylen int) {
	fmtCtx.keylen = C.int(keylen)
}

func (fmtCtx *CAVFormatContext) GetNbPrograms() uint {
	return uint(fmtCtx.nb_programs)
}
func (fmtCtx *CAVFormatContext) SetNbPrograms(nbPrograms uint) {
	fmtCtx.nb_programs = C.uint(nbPrograms)
}

func (fmtCtx *CAVFormatContext) GetPrograms() **CAVProgram {
	return (**CAVProgram)(unsafe.Pointer(fmtCtx.programs))
}
func (fmtCtx *CAVFormatContext) SetPrograms(programs **CAVProgram) {
	fmtCtx.programs = (**C.AVProgram)(unsafe.Pointer(programs))
}

/**
 * Forced video codec_id.
 * Demuxing: Set by user.
 */
func (fmtCtx *CAVFormatContext) GetVideoCodecId() avcodec.CAVCodecID {
	return (avcodec.CAVCodecID)(fmtCtx.video_codec_id)
}

/**
 * Forced video codec_id.
 * Demuxing: Set by user.
 */
func (fmtCtx *CAVFormatContext) SetVideoCodecId(videoCodecId avcodec.CAVCodecID) {
	fmtCtx.video_codec_id = C.enum_AVCodecID(videoCodecId)
}

/**
 * Forced audio codec_id.
 * Demuxing: Set by user.
 */
func (fmtCtx *CAVFormatContext) GetAudioCodecId() avcodec.CAVCodecID {
	return (avcodec.CAVCodecID)(fmtCtx.audio_codec_id)
}

/**
 * Forced audio codec_id.
 * Demuxing: Set by user.
 */
func (fmtCtx *CAVFormatContext) SetAudioCodecId(audioCodecId avcodec.CAVCodecID) {
	fmtCtx.audio_codec_id = C.enum_AVCodecID(audioCodecId)
}

/**
 * Forced subtitle codec_id.
 * Demuxing: Set by user.
 */
func (fmtCtx *CAVFormatContext) GetSubtitleCodecId() avcodec.CAVCodecID {
	return (avcodec.CAVCodecID)(fmtCtx.subtitle_codec_id)
}

/**
 * Forced subtitle codec_id.
 * Demuxing: Set by user.
 */
func (fmtCtx *CAVFormatContext) SetSubtitleCodecId(subtitleCodecId avcodec.CAVCodecID) {
	fmtCtx.subtitle_codec_id = C.enum_AVCodecID(subtitleCodecId)
}

/**
 * Maximum amount of memory in bytes to use for the index of each stream.
 * If the index exceeds this size, entries will be discarded as
 * needed to maintain a smaller size. This can lead to slower or less
 * accurate seeking (depends on demuxer).
 * Demuxers for which a full in-memory index is mandatory will ignore
 * this.
 * - muxing: unused
 * - demuxing: set by user
 */
func (fmtCtx *CAVFormatContext) GetMaxIndexSize() uint {
	return uint(fmtCtx.max_index_size)
}

/**
 * Maximum amount of memory in bytes to use for the index of each stream.
 * If the index exceeds this size, entries will be discarded as
 * needed to maintain a smaller size. This can lead to slower or less
 * accurate seeking (depends on demuxer).
 * Demuxers for which a full in-memory index is mandatory will ignore
 * this.
 * - muxing: unused
 * - demuxing: set by user
 */
func (fmtCtx *CAVFormatContext) SetMaxIndexSize(maxIndexSize uint) {
	fmtCtx.max_index_size = C.uint(maxIndexSize)
}

/**
 * Maximum amount of memory in bytes to use for buffering frames
 * obtained from realtime capture devices.
 */
func (fmtCtx *CAVFormatContext) GetMaxPictureBuffer() uint {
	return uint(fmtCtx.max_picture_buffer)
}

/**
 * Maximum amount of memory in bytes to use for buffering frames
 * obtained from realtime capture devices.
 */
func (fmtCtx *CAVFormatContext) SetMaxPictureBuffer(maxPictureBuffer uint) {
	fmtCtx.max_picture_buffer = C.uint(maxPictureBuffer)
}

/**
 * Number of chapters in AVChapter array.
 * When muxing, chapters are normally written in the file header,
 * so nb_chapters should normally be initialized before write_header
 * is called. Some muxers (e.g. mov and mkv) can also write chapters
 * in the trailer.  To write chapters in the trailer, nb_chapters
 * must be zero when write_header is called and non-zero when
 * write_trailer is called.
 * - muxing: set by user
 * - demuxing: set by libavformat
 */
func (fmtCtx *CAVFormatContext) GetNbChapters() uint {
	return uint(fmtCtx.nb_chapters)
}

/**
 * Number of chapters in AVChapter array.
 * When muxing, chapters are normally written in the file header,
 * so nb_chapters should normally be initialized before write_header
 * is called. Some muxers (e.g. mov and mkv) can also write chapters
 * in the trailer.  To write chapters in the trailer, nb_chapters
 * must be zero when write_header is called and non-zero when
 * write_trailer is called.
 * - muxing: set by user
 * - demuxing: set by libavformat
 */
func (fmtCtx *CAVFormatContext) SetNbChapters(nbChapter uint) {
	fmtCtx.nb_chapters = C.uint(nbChapter)
}

func (fmtCtx *CAVFormatContext) GetChapters() **CAVChapter {
	return (**CAVChapter)(unsafe.Pointer(fmtCtx.chapters))
}
func (fmtCtx *CAVFormatContext) SetChapters(chapters **CAVChapter) {
	fmtCtx.chapters = (**C.AVChapter)(unsafe.Pointer(chapters))
}

/**
 * Metadata that applies to the whole file.
 *
 * - demuxing: set by libavformat in avformat_open_input()
 * - muxing: may be set by the caller before avformat_write_header()
 *
 * Freed by libavformat in avformat_free_context().
 */
func (fmtCtx *CAVFormatContext) GetMetadata() *avutil.CAVDictionary {
	return (*avutil.CAVDictionary)(fmtCtx.metadata)
}

/**
 * Metadata that applies to the whole file.
 *
 * - demuxing: set by libavformat in avformat_open_input()
 * - muxing: may be set by the caller before avformat_write_header()
 *
 * Freed by libavformat in avformat_free_context().
 */
func (fmtCtx *CAVFormatContext) SetMetadata(metadata *avutil.CAVDictionary) {
	fmtCtx.metadata = (*C.AVDictionary)(metadata)
}

/**
 * Start time of the stream in real world time, in microseconds
 * since the Unix epoch (00:00 1st January 1970). That is, pts=0 in the
 * stream was captured at this real world time.
 * - muxing: Set by the caller before avformat_write_header(). If set to
 *           either 0 or AV_NOPTS_VALUE, then the current wall-time will
 *           be used.
 * - demuxing: Set by libavformat. AV_NOPTS_VALUE if unknown. Note that
 *             the value may become known after some number of frames
 *             have been received.
 */
func (fmtCtx *CAVFormatContext) GetStartTimeRealtime() int64 {
	return int64(fmtCtx.start_time_realtime)
}

/**
 * Start time of the stream in real world time, in microseconds
 * since the Unix epoch (00:00 1st January 1970). That is, pts=0 in the
 * stream was captured at this real world time.
 * - muxing: Set by the caller before avformat_write_header(). If set to
 *           either 0 or AV_NOPTS_VALUE, then the current wall-time will
 *           be used.
 * - demuxing: Set by libavformat. AV_NOPTS_VALUE if unknown. Note that
 *             the value may become known after some number of frames
 *             have been received.
 */
func (fmtCtx *CAVFormatContext) SetStartTimeRealtime(startTimeRealtime int64) {
	fmtCtx.start_time_realtime = C.int64_t(startTimeRealtime)
}

/**
 * The number of frames used for determining the framerate in
 * avformat_find_stream_info().
 * Demuxing only, set by the caller before avformat_find_stream_info().
 */
func (fmtCtx *CAVFormatContext) GetFpsProbeSize() int {
	return int(fmtCtx.fps_probe_size)
}

/**
 * The number of frames used for determining the framerate in
 * avformat_find_stream_info().
 * Demuxing only, set by the caller before avformat_find_stream_info().
 */
func (fmtCtx *CAVFormatContext) SetFpsProbeSize(fpsProbeSize int) {
	fmtCtx.fps_probe_size = C.int(fpsProbeSize)
}

/**
 * Error recognition; higher values will detect more errors but may
 * misdetect some more or less valid parts as errors.
 * Demuxing only, set by the caller before avformat_open_input().
 */
func (fmtCtx *CAVFormatContext) GetErrorRecognition() int {
	return int(fmtCtx.error_recognition)
}

/**
 * Error recognition; higher values will detect more errors but may
 * misdetect some more or less valid parts as errors.
 * Demuxing only, set by the caller before avformat_open_input().
 */
func (fmtCtx *CAVFormatContext) SetErrorRecognition(errorRecognition int) {
	fmtCtx.error_recognition = C.int(errorRecognition)
}

/**
 * Custom interrupt callbacks for the I/O layer.
 *
 * demuxing: set by the user before avformat_open_input().
 * muxing: set by the user before avformat_write_header()
 * (mainly useful for AVFMT_NOFILE formats). The callback
 * should also be passed to avio_open2() if it's used to
 * open the file.
 */
func (fmtCtx *CAVFormatContext) GetInterruptCallback() CAVIOInterruptCB {
	return CAVIOInterruptCB(fmtCtx.interrupt_callback)
}

/**
 * Custom interrupt callbacks for the I/O layer.
 *
 * demuxing: set by the user before avformat_open_input().
 * muxing: set by the user before avformat_write_header()
 * (mainly useful for AVFMT_NOFILE formats). The callback
 * should also be passed to avio_open2() if it's used to
 * open the file.
 */
func (fmtCtx *CAVFormatContext) SetInterruptCallback(interruptCall CAVIOInterruptCB) {
	fmtCtx.interrupt_callback = C.AVIOInterruptCB(interruptCall)
}

/**
 * Flags to enable debugging.
 */
func (fmtCtx *CAVFormatContext) GetDebug() int {
	return int(fmtCtx.debug)
}

/**
 * Flags to enable debugging.
 */
func (fmtCtx *CAVFormatContext) SetDebug(debug int) {
	fmtCtx.debug = C.int(debug)
}

const FF_FDEBUG_TS = C.FF_FDEBUG_TS

/**
 * Maximum buffering duration for interleaving.
 *
 * To ensure all the streams are interleaved correctly,
 * av_interleaved_write_frame() will wait until it has at least one packet
 * for each stream before actually writing any packets to the output file.
 * When some streams are "sparse" (i.e. there are large gaps between
 * successive packets), this can result in excessive buffering.
 *
 * This field specifies the maximum difference between the timestamps of the
 * first and the last packet in the muxing queue, above which libavformat
 * will output a packet regardless of whether it has queued a packet for all
 * the streams.
 *
 * Muxing only, set by the caller before avformat_write_header().
 */
func (fmtCtx *CAVFormatContext) GetMaxInterleaveDelta() int64 {
	return int64(fmtCtx.max_interleave_delta)
}

/**
 * Maximum buffering duration for interleaving.
 *
 * To ensure all the streams are interleaved correctly,
 * av_interleaved_write_frame() will wait until it has at least one packet
 * for each stream before actually writing any packets to the output file.
 * When some streams are "sparse" (i.e. there are large gaps between
 * successive packets), this can result in excessive buffering.
 *
 * This field specifies the maximum difference between the timestamps of the
 * first and the last packet in the muxing queue, above which libavformat
 * will output a packet regardless of whether it has queued a packet for all
 * the streams.
 *
 * Muxing only, set by the caller before avformat_write_header().
 */
func (fmtCtx *CAVFormatContext) SetMaxInterleaveDelta(maxInterleaveDelta int64) {
	fmtCtx.max_interleave_delta = C.int64_t(maxInterleaveDelta)
}

/**
 * Allow non-standard and experimental extension
 * @see AVCodecContext.strict_std_compliance
 */
func (fmtCtx *CAVFormatContext) GetStrictStdCompliance() int {
	return int(fmtCtx.strict_std_compliance)
}

/**
 * Allow non-standard and experimental extension
 * @see AVCodecContext.strict_std_compliance
 */
func (fmtCtx *CAVFormatContext) SetStrictStdCompliance(strictStdCompliance int) {
	fmtCtx.strict_std_compliance = C.int(strictStdCompliance)
}

/**
 * Flags indicating events happening on the file, a combination of
 * AVFMT_EVENT_FLAG_*.
 *
 * - demuxing: may be set by the demuxer in avformat_open_input(),
 *   avformat_find_stream_info() and av_read_frame(). Flags must be cleared
 *   by the user once the event has been handled.
 * - muxing: may be set by the user after avformat_write_header() to
 *   indicate a user-triggered event.  The muxer will clear the flags for
 *   events it has handled in av_[interleaved]_write_frame().
 */
func (fmtCtx *CAVFormatContext) GetEventFlags() int {
	return int(fmtCtx.event_flags)
}

/**
 * Flags indicating events happening on the file, a combination of
 * AVFMT_EVENT_FLAG_*.
 *
 * - demuxing: may be set by the demuxer in avformat_open_input(),
 *   avformat_find_stream_info() and av_read_frame(). Flags must be cleared
 *   by the user once the event has been handled.
 * - muxing: may be set by the user after avformat_write_header() to
 *   indicate a user-triggered event.  The muxer will clear the flags for
 *   events it has handled in av_[interleaved]_write_frame().
 */
func (fmtCtx *CAVFormatContext) SetEventFlags(eventFlags int) {
	fmtCtx.event_flags = C.int(eventFlags)
}

/**
 * - demuxing: the demuxer read new metadata from the file and updated
 *   AVFormatContext.metadata accordingly
 * - muxing: the user updated AVFormatContext.metadata and wishes the muxer to
 *   write it into the file
 */
const AVFMT_EVENT_FLAG_METADATA_UPDATED = C.AVFMT_EVENT_FLAG_METADATA_UPDATED

/**
 * Maximum number of packets to read while waiting for the first timestamp.
 * Decoding only.
 */
func (fmtCtx *CAVFormatContext) GetMaxTsProbe() int {
	return int(fmtCtx.max_ts_probe)
}

/**
 * Maximum number of packets to read while waiting for the first timestamp.
 * Decoding only.
 */
func (fmtCtx *CAVFormatContext) SetMaxTsProbe(maxTsProbe int) {
	fmtCtx.max_ts_probe = C.int(maxTsProbe)
}

/**
 * Avoid negative timestamps during muxing.
 * Any value of the AVFMT_AVOID_NEG_TS_* constants.
 * Note, this works better when using av_interleaved_write_frame().
 * - muxing: Set by user
 * - demuxing: unused
 */
func (fmtCtx *CAVFormatContext) GetAvoidNegativeTs() int {
	return int(fmtCtx.avoid_negative_ts)
}

/**
 * Avoid negative timestamps during muxing.
 * Any value of the AVFMT_AVOID_NEG_TS_* constants.
 * Note, this works better when using av_interleaved_write_frame().
 * - muxing: Set by user
 * - demuxing: unused
 */
func (fmtCtx *CAVFormatContext) SetAvoidNegativeTs(avoidNegativeTs int) {
	fmtCtx.avoid_negative_ts = C.int(avoidNegativeTs)
}

const (
	AVFMT_AVOID_NEG_TS_AUTO              = C.AVFMT_AVOID_NEG_TS_AUTO              ///< Enabled when required by target format
	AVFMT_AVOID_NEG_TS_DISABLED          = C.AVFMT_AVOID_NEG_TS_DISABLED          ///< Do not shift timestamps even when they are negative.
	AVFMT_AVOID_NEG_TS_MAKE_NON_NEGATIVE = C.AVFMT_AVOID_NEG_TS_MAKE_NON_NEGATIVE ///< Shift timestamps so they are non negative
	AVFMT_AVOID_NEG_TS_MAKE_ZERO         = C.AVFMT_AVOID_NEG_TS_MAKE_ZERO         ///< Shift timestamps so that they start at 0
)

/**
 * Transport stream id.
 * This will be moved into demuxer private options. Thus no API/ABI compatibility
 */
func (fmtCtx *CAVFormatContext) GetTsId() int {
	return int(fmtCtx.ts_id)
}

/**
 * Transport stream id.
 * This will be moved into demuxer private options. Thus no API/ABI compatibility
 */
func (fmtCtx *CAVFormatContext) SetTsId(tsId int) {
	fmtCtx.ts_id = C.int(tsId)
}

/**
 * Audio preload in microseconds.
 * Note, not all formats support this and unpredictable things may happen if it is used when not supported.
 * - encoding: Set by user
 * - decoding: unused
 */
func (fmtCtx *CAVFormatContext) GetAudioPreload() int {
	return int(fmtCtx.audio_preload)
}

/**
 * Audio preload in microseconds.
 * Note, not all formats support this and unpredictable things may happen if it is used when not supported.
 * - encoding: Set by user
 * - decoding: unused
 */
func (fmtCtx *CAVFormatContext) SetAudioPreload(audioPreload int) {
	fmtCtx.audio_preload = C.int(audioPreload)
}

/**
 * Max chunk time in microseconds.
 * Note, not all formats support this and unpredictable things may happen if it is used when not supported.
 * - encoding: Set by user
 * - decoding: unused
 */
func (fmtCtx *CAVFormatContext) GetMaxChunkDuration() int {
	return int(fmtCtx.max_chunk_duration)
}

/**
 * Max chunk time in microseconds.
 * Note, not all formats support this and unpredictable things may happen if it is used when not supported.
 * - encoding: Set by user
 * - decoding: unused
 */
func (fmtCtx *CAVFormatContext) SetMaxChunkDuration(maxChunkDuration int) {
	fmtCtx.max_chunk_duration = C.int(maxChunkDuration)
}

/**
 * Max chunk size in bytes
 * Note, not all formats support this and unpredictable things may happen if it is used when not supported.
 * - encoding: Set by user
 * - decoding: unused
 */
func (fmtCtx *CAVFormatContext) GetMaxChunkSize() int {
	return int(fmtCtx.max_chunk_size)
}

/**
 * Max chunk size in bytes
 * Note, not all formats support this and unpredictable things may happen if it is used when not supported.
 * - encoding: Set by user
 * - decoding: unused
 */
func (fmtCtx *CAVFormatContext) SetMaxChunkSize(maxChunkSize int) {
	fmtCtx.max_chunk_size = C.int(maxChunkSize)
}

/**
 * forces the use of wallclock timestamps as pts/dts of packets
 * This has undefined results in the presence of B frames.
 * - encoding: unused
 * - decoding: Set by user
 */
func (fmtCtx *CAVFormatContext) GetUseWallclockAsTimestamps() int {
	return int(fmtCtx.use_wallclock_as_timestamps)
}

/**
 * forces the use of wallclock timestamps as pts/dts of packets
 * This has undefined results in the presence of B frames.
 * - encoding: unused
 * - decoding: Set by user
 */
func (fmtCtx *CAVFormatContext) SetUseWallclockAsTimestamps(useWallclockAsTimestamps int) {
	fmtCtx.use_wallclock_as_timestamps = C.int(useWallclockAsTimestamps)
}

/**
 * avio flags, used to force AVIO_FLAG_DIRECT.
 * - encoding: unused
 * - decoding: Set by user
 */
func (fmtCtx *CAVFormatContext) GetAvioFlags() int {
	return int(fmtCtx.avio_flags)
}

/**
 * avio flags, used to force AVIO_FLAG_DIRECT.
 * - encoding: unused
 * - decoding: Set by user
 */
func (fmtCtx *CAVFormatContext) SetAvioFlags(avioFlags int) {
	fmtCtx.avio_flags = C.int(avioFlags)
}

/**
 * The duration field can be estimated through various ways, and this field can be used
 * to know how the duration was estimated.
 * - encoding: unused
 * - decoding: Read by user
 */
func (fmtCtx *CAVFormatContext) GetDurationEstimationMethod() CAVDurationEstimationMethod {
	return CAVDurationEstimationMethod(fmtCtx.duration_estimation_method)
}

/**
 * The duration field can be estimated through various ways, and this field can be used
 * to know how the duration was estimated.
 * - encoding: unused
 * - decoding: Read by user
 */
func (fmtCtx *CAVFormatContext) SetDurationEstimationMethod(durationEstimationMethod CAVDurationEstimationMethod) {
	fmtCtx.duration_estimation_method = C.enum_AVDurationEstimationMethod(durationEstimationMethod)
}

/**
 * Skip initial bytes when opening stream
 * - encoding: unused
 * - decoding: Set by user
 */
func (fmtCtx *CAVFormatContext) GetSkipInitialBytes() int64 {
	return int64(fmtCtx.skip_initial_bytes)
}

/**
 * Skip initial bytes when opening stream
 * - encoding: unused
 * - decoding: Set by user
 */
func (fmtCtx *CAVFormatContext) SetSkipInitialBytes(skipInitialBytes int64) {
	fmtCtx.skip_initial_bytes = C.int64_t(skipInitialBytes)
}

/**
 * Correct single timestamp overflows
 * - encoding: unused
 * - decoding: Set by user
 */
func (fmtCtx *CAVFormatContext) GetCorrectTsOverflow() uint {
	return uint(fmtCtx.correct_ts_overflow)
}

/**
 * Correct single timestamp overflows
 * - encoding: unused
 * - decoding: Set by user
 */
func (fmtCtx *CAVFormatContext) SetCorrectTsOverflow(correctTsOverflow uint) {
	fmtCtx.correct_ts_overflow = C.uint(correctTsOverflow)
}

/**
 * Force seeking to any (also non key) frames.
 * - encoding: unused
 * - decoding: Set by user
 */
func (fmtCtx *CAVFormatContext) GetSeek2Any() int {
	return int(fmtCtx.seek2any)
}

/**
 * Force seeking to any (also non key) frames.
 * - encoding: unused
 * - decoding: Set by user
 */
func (fmtCtx *CAVFormatContext) SetSeek2Any(seek2Any int) {
	fmtCtx.seek2any = C.int(seek2Any)
}

/**
 * Flush the I/O context after each packet.
 * - encoding: Set by user
 * - decoding: unused
 */
func (fmtCtx *CAVFormatContext) GetFlushPackets() int {
	return int(fmtCtx.flush_packets)
}

/**
 * Flush the I/O context after each packet.
 * - encoding: Set by user
 * - decoding: unused
 */
func (fmtCtx *CAVFormatContext) SetFlushPackets(flushPackets int) {
	fmtCtx.flush_packets = C.int(flushPackets)
}

/**
 * format probing score.
 * The maximal score is AVPROBE_SCORE_MAX, its set when the demuxer probes
 * the format.
 * - encoding: unused
 * - decoding: set by avformat, read by user
 */
func (fmtCtx *CAVFormatContext) GetProbeScore() int {
	return int(fmtCtx.probe_score)
}

/**
 * format probing score.
 * The maximal score is AVPROBE_SCORE_MAX, its set when the demuxer probes
 * the format.
 * - encoding: unused
 * - decoding: set by avformat, read by user
 */
func (fmtCtx *CAVFormatContext) SetProbeScore(probeScore int) {
	fmtCtx.probe_score = C.int(probeScore)
}

/**
 * Maximum number of bytes read from input in order to identify the
 * \ref AVInputFormat "input format". Only used when the format is not set
 * explicitly by the caller.
 *
 * Demuxing only, set by the caller before avformat_open_input().
 *
 * @sa probesize
 */
func (fmtCtx *CAVFormatContext) GetFormatProbesize() int {
	return int(fmtCtx.format_probesize)
}

/**
 * Maximum number of bytes read from input in order to identify the
 * \ref AVInputFormat "input format". Only used when the format is not set
 * explicitly by the caller.
 *
 * Demuxing only, set by the caller before avformat_open_input().
 *
 * @sa probesize
 */
func (fmtCtx *CAVFormatContext) SetFormatProbesize(formatProbeSize int) {
	fmtCtx.format_probesize = C.int(formatProbeSize)
}

/**
 * ',' separated list of allowed decoders.
 * If NULL then all are allowed
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) GetCodecWhitelist() string {
	return C.GoString(fmtCtx.codec_whitelist)
}

/**
 * ',' separated list of allowed decoders.
 * If NULL then all are allowed
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) SetCodecWhitelist(codecWhitelist string) {
	var cCodecWhitelist *C.char = nil
	if len(codecWhitelist) > 0 {
		cCodecWhitelist = C.CString(codecWhitelist)
	}

	fmtCtx.codec_whitelist = cCodecWhitelist
}

/**
 * ',' separated list of allowed demuxers.
 * If NULL then all are allowed
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) GetFormatWhitelist() string {
	return C.GoString(fmtCtx.format_whitelist)
}

/**
 * ',' separated list of allowed demuxers.
 * If NULL then all are allowed
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) SetFormatWhitelist(formatWhitelist string) {
	var cFormatWhitelist *C.char = nil
	if len(formatWhitelist) > 0 {
		cFormatWhitelist = C.CString(formatWhitelist)
	}

	fmtCtx.format_whitelist = cFormatWhitelist
}

/**
 * IO repositioned flag.
 * This is set by avformat when the underlaying IO context read pointer
 * is repositioned, for example when doing byte based seeking.
 * Demuxers can use the flag to detect such changes.
 */
func (fmtCtx *CAVFormatContext) GetIoRespositioned() int {
	return int(fmtCtx.io_repositioned)
}

/**
 * IO repositioned flag.
 * This is set by avformat when the underlaying IO context read pointer
 * is repositioned, for example when doing byte based seeking.
 * Demuxers can use the flag to detect such changes.
 */
func (fmtCtx *CAVFormatContext) SetIoRespositioned(ioRespositioned int) {
	fmtCtx.io_repositioned = C.int(ioRespositioned)
}

/**
 * Forced video codec.
 * This allows forcing a specific decoder, even when there are multiple with
 * the same codec_id.
 * Demuxing: Set by user
 */
func (fmtCtx *CAVFormatContext) GetVideoCodec() *avcodec.CAVCodec {
	return (*avcodec.CAVCodec)(unsafe.Pointer(fmtCtx.video_codec))
}

/**
 * Forced video codec.
 * This allows forcing a specific decoder, even when there are multiple with
 * the same codec_id.
 * Demuxing: Set by user
 */
func (fmtCtx *CAVFormatContext) SetVideoCodec(videoCodec *avcodec.CAVCodec) {
	fmtCtx.video_codec = (*C.AVCodec)(unsafe.Pointer(videoCodec))
}

/**
 * Forced audio codec.
 * This allows forcing a specific decoder, even when there are multiple with
 * the same codec_id.
 * Demuxing: Set by user
 */
func (fmtCtx *CAVFormatContext) GetAudioCodec() *avcodec.CAVCodec {
	return (*avcodec.CAVCodec)(unsafe.Pointer(fmtCtx.audio_codec))
}

/**
 * Forced audio codec.
 * This allows forcing a specific decoder, even when there are multiple with
 * the same codec_id.
 * Demuxing: Set by user
 */
func (fmtCtx *CAVFormatContext) SetAudioCodec(audioCodec *avcodec.CAVCodec) {
	fmtCtx.audio_codec = (*C.AVCodec)(unsafe.Pointer(audioCodec))
}

/**
 * Forced subtitle codec.
 * This allows forcing a specific decoder, even when there are multiple with
 * the same codec_id.
 * Demuxing: Set by user
 */
func (fmtCtx *CAVFormatContext) GetSubtitleCodec() *avcodec.CAVCodec {
	return (*avcodec.CAVCodec)(unsafe.Pointer(fmtCtx.subtitle_codec))
}

/**
 * Forced subtitle codec.
 * This allows forcing a specific decoder, even when there are multiple with
 * the same codec_id.
 * Demuxing: Set by user
 */
func (fmtCtx *CAVFormatContext) SetSubtitleCodec(subtitleCodec *avcodec.CAVCodec) {
	fmtCtx.subtitle_codec = (*C.AVCodec)(unsafe.Pointer(subtitleCodec))
}

/**
 * Forced data codec.
 * This allows forcing a specific decoder, even when there are multiple with
 * the same codec_id.
 * Demuxing: Set by user
 */
func (fmtCtx *CAVFormatContext) GetDataCodec() *avcodec.CAVCodec {
	return (*avcodec.CAVCodec)(unsafe.Pointer(fmtCtx.data_codec))
}

/**
 * Forced data codec.
 * This allows forcing a specific decoder, even when there are multiple with
 * the same codec_id.
 * Demuxing: Set by user
 */
func (fmtCtx *CAVFormatContext) SetDataCodec(dataCodec *avcodec.CAVCodec) {
	fmtCtx.data_codec = (*C.AVCodec)(unsafe.Pointer(dataCodec))
}

/**
 * Number of bytes to be written as padding in a metadata header.
 * Demuxing: Unused.
 * Muxing: Set by user via av_format_set_metadata_header_padding.
 */
func (fmtCtx *CAVFormatContext) GetMetadataHeaderPadding() int {
	return int(fmtCtx.metadata_header_padding)
}

/**
 * Number of bytes to be written as padding in a metadata header.
 * Demuxing: Unused.
 * Muxing: Set by user via av_format_set_metadata_header_padding.
 */
func (fmtCtx *CAVFormatContext) SetMetadataHeaderPadding(metadataHeaderPadding int) {
	fmtCtx.metadata_header_padding = C.int(metadataHeaderPadding)
}

/**
 * User data.
 * This is a place for some private data of the user.
 */
func (fmtCtx *CAVFormatContext) GetOpaque() unsafe.Pointer {
	return fmtCtx.opaque
}

/**
 * User data.
 * This is a place for some private data of the user.
 */
func (fmtCtx *CAVFormatContext) SetOpaque(opaque unsafe.Pointer) {
	fmtCtx.opaque = opaque
}

/**
 * Callback used by devices to communicate with application.
 */
func (fmtCtx *CAVFormatContext) GetControlMessageCb() CAvFormatControlMessage {
	return CAvFormatControlMessage(fmtCtx.control_message_cb)
}

/**
 * Callback used by devices to communicate with application.
 */
func (fmtCtx *CAVFormatContext) SetControlMessageCb(controlMessageCb CAvFormatControlMessage) {
	fmtCtx.control_message_cb = C.av_format_control_message(controlMessageCb)
}

/**
 * Output timestamp offset, in microseconds.
 * Muxing: set by user
 */
func (fmtCtx *CAVFormatContext) GetOutputTsOffset() int64 {
	return int64(fmtCtx.output_ts_offset)
}

/**
 * Output timestamp offset, in microseconds.
 * Muxing: set by user
 */
func (fmtCtx *CAVFormatContext) SetOutputTsOffset(outputTsOffset int64) {
	fmtCtx.output_ts_offset = C.int64_t(outputTsOffset)
}

/**
 * dump format separator.
 * can be ", " or "\n      " or anything else
 * - muxing: Set by user.
 * - demuxing: Set by user.
 */
func (fmtCtx *CAVFormatContext) GetDumpSeparator() string {
	return C.GoString((*C.char)(unsafe.Pointer(fmtCtx.dump_separator)))
}

/**
 * dump format separator.
 * can be ", " or "\n      " or anything else
 * - muxing: Set by user.
 * - demuxing: Set by user.
 */
func (fmtCtx *CAVFormatContext) SetDumpSeparator(dumpSeparator string) {
	var cDumpSeparator *C.char = nil
	if len(dumpSeparator) > 0 {
		cDumpSeparator = C.CString(dumpSeparator)
	}

	fmtCtx.dump_separator = (*C.uchar)(unsafe.Pointer(cDumpSeparator))
}

/**
 * Forced Data codec_id.
 * Demuxing: Set by user.
 */
func (fmtCtx *CAVFormatContext) GetDataCodecId() avcodec.CAVCodecID {
	return avcodec.CAVCodecID(fmtCtx.data_codec_id)
}

/**
 * Forced Data codec_id.
 * Demuxing: Set by user.
 */
func (fmtCtx *CAVFormatContext) SetDataCodecId(dataCodecid avcodec.CAVCodecID) {
	fmtCtx.data_codec_id = C.enum_AVCodecID(dataCodecid)
}

/**
 * ',' separated list of allowed protocols.
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) GetProtocolWhitelist() string {
	return C.GoString(fmtCtx.protocol_whitelist)
}

/**
 * ',' separated list of allowed protocols.
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) SetProtocolWhitelist(protocolWhitelist string) {
	var cProtocolWhitelist *C.char = nil
	if len(protocolWhitelist) > 0 {
		cProtocolWhitelist = C.CString(protocolWhitelist)
	}

	fmtCtx.protocol_whitelist = cProtocolWhitelist
}

/**
 * A callback for opening new IO streams.
 *
 * Whenever a muxer or a demuxer needs to open an IO stream (typically from
 * avformat_open_input() for demuxers, but for certain formats can happen at
 * other times as well), it will call this callback to obtain an IO context.
 *
 * @param s the format context
 * @param pb on success, the newly opened IO context should be returned here
 * @param url the url to open
 * @param flags a combination of AVIO_FLAG_*
 * @param options a dictionary of additional options, with the same
 *                semantics as in avio_open2()
 * @return 0 on success, a negative AVERROR code on failure
 *
 * @note Certain muxers and demuxers do nesting, i.e. they open one or more
 * additional internal format contexts. Thus the AVFormatContext pointer
 * passed to this callback may be different from the one facing the caller.
 * It will, however, have the same 'opaque' field.
 */
// 	 int (*io_open)(struct AVFormatContext *s, AVIOContext **pb, const char *url,
// 					int flags, AVDictionary **options);
func (fmtCtx *CAVFormatContext) GetIoOpen() ctypes.CFunc {
	return ctypes.CFunc(fmtCtx.io_open)
}

/**
 * A callback for opening new IO streams.
 *
 * Whenever a muxer or a demuxer needs to open an IO stream (typically from
 * avformat_open_input() for demuxers, but for certain formats can happen at
 * other times as well), it will call this callback to obtain an IO context.
 *
 * @param s the format context
 * @param pb on success, the newly opened IO context should be returned here
 * @param url the url to open
 * @param flags a combination of AVIO_FLAG_*
 * @param options a dictionary of additional options, with the same
 *                semantics as in avio_open2()
 * @return 0 on success, a negative AVERROR code on failure
 *
 * @note Certain muxers and demuxers do nesting, i.e. they open one or more
 * additional internal format contexts. Thus the AVFormatContext pointer
 * passed to this callback may be different from the one facing the caller.
 * It will, however, have the same 'opaque' field.
 */
// 	 int (*io_open)(struct AVFormatContext *s, AVIOContext **pb, const char *url,
// 					int flags, AVDictionary **options);
func (fmtCtx *CAVFormatContext) SetIoOpen(ioOpen ctypes.CFunc) {
	fmtCtx.io_open = (*[0]byte)(ioOpen)
}

//  #if FF_API_AVFORMAT_IO_CLOSE
// 	 /**
// 	  * A callback for closing the streams opened with AVFormatContext.io_open().
// 	  *
// 	  * @deprecated use io_close2
// 	  */
// 	 attribute_deprecated
// 	 void (*io_close)(struct AVFormatContext *s, AVIOContext *pb);
//  #endif

/**
 * ',' separated list of disallowed protocols.
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) GetProtocolBlacklist() string {
	return C.GoString(fmtCtx.protocol_blacklist)
}

/**
 * ',' separated list of disallowed protocols.
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) SetProtocolBlacklist(protocolBlackList string) {
	var cProtocolBlackList *C.char = nil
	if len(protocolBlackList) > 0 {
		cProtocolBlackList = C.CString(protocolBlackList)
	}
	fmtCtx.protocol_blacklist = cProtocolBlackList
}

/**
 * The maximum number of streams.
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) GetMaxStreams() int {
	return int(fmtCtx.max_streams)
}

/**
 * The maximum number of streams.
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) SetMaxStreams(maxStreams int) {
	fmtCtx.max_streams = C.int(maxStreams)
}

/**
 * Skip duration calcuation in estimate_timings_from_pts.
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) GetSkipEstimateDurationFromPts() int {
	return int(fmtCtx.skip_estimate_duration_from_pts)
}

/**
 * Skip duration calcuation in estimate_timings_from_pts.
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) SetSkipEstimateDurationFromPts(skipEstimateDurationFromPts int) {
	fmtCtx.skip_estimate_duration_from_pts = C.int(skipEstimateDurationFromPts)
}

/**
 * Maximum number of packets that can be probed
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) GetMaxProbePackets() int {
	return int(fmtCtx.max_probe_packets)
}

/**
 * Maximum number of packets that can be probed
 * - encoding: unused
 * - decoding: set by user
 */
func (fmtCtx *CAVFormatContext) SetMaxProbePackets(maxProbePackets int) {
	fmtCtx.max_probe_packets = C.int(maxProbePackets)
}

/**
 * A callback for closing the streams opened with AVFormatContext.io_open().
 *
 * Using this is preferred over io_close, because this can return an error.
 * Therefore this callback is used instead of io_close by the generic
 * libavformat code if io_close is NULL or the default.
 *
 * @param s the format context
 * @param pb IO context to be closed and freed
 * @return 0 on success, a negative AVERROR code on failure
 */
// 	 int (*io_close2)(struct AVFormatContext *s, AVIOContext *pb);
func (fmtCtx *CAVFormatContext) GetIoClose2() ctypes.CFunc {
	return ctypes.CFunc(fmtCtx.io_close2)
}

/**
 * A callback for closing the streams opened with AVFormatContext.io_open().
 *
 * Using this is preferred over io_close, because this can return an error.
 * Therefore this callback is used instead of io_close by the generic
 * libavformat code if io_close is NULL or the default.
 *
 * @param s the format context
 * @param pb IO context to be closed and freed
 * @return 0 on success, a negative AVERROR code on failure
 */
// 	 int (*io_close2)(struct AVFormatContext *s, AVIOContext *pb);
func (fmtCtx *CAVFormatContext) SetIoClose2(ioClose2 ctypes.CFunc) {
	fmtCtx.io_close2 = (*[0]byte)(ioClose2)
}

//#endregion CAVFormatContext

/**
 * This function will cause global side data to be injected in the next packet
 * of each stream as well as after any subsequent seek.
 *
 * @note global side data is always available in every AVStream's
 *       @ref AVCodecParameters.coded_side_data "codecpar side data" array, and
 *       in a @ref AVCodecContext.coded_side_data "decoder's side data" array if
 *       initialized with said stream's codecpar.
 * @see av_packet_side_data_get()
 */
func AvFormatInjectGlobalSideData(s *CAVFormatContext) {
	C.av_format_inject_global_side_data((*C.AVFormatContext)(s))
}

/**
 * Returns the method used to set ctx->duration.
 *
 * @return AVFMT_DURATION_FROM_PTS, AVFMT_DURATION_FROM_STREAM, or AVFMT_DURATION_FROM_BITRATE.
 */
func AvFmtCtxGetDurationEstimationMethod(ctx *CAVFormatContext) CAVDurationEstimationMethod {
	return CAVDurationEstimationMethod(C.av_fmt_ctx_get_duration_estimation_method((*C.AVFormatContext)(ctx)))
}

/**
 * @defgroup lavf_core Core functions
 * @ingroup libavf
 *
 * Functions for querying libavformat capabilities, allocating core structures,
 * etc.
 * @{
 */

/**
 * Return the LIBAVFORMAT_VERSION_INT constant.
 */
func AvformatVersion() uint {
	return uint(C.avformat_version())
}

/**
 * Return the libavformat build-time configuration.
 */
func AvformatConfiguration() string {
	return C.GoString(C.avformat_configuration())
}

/**
 * Return the libavformat license.
 */
func AvformatLicense() string {
	return C.GoString(C.avformat_license())
}

/**
 * Do global initialization of network libraries. This is optional,
 * and not recommended anymore.
 *
 * This functions only exists to work around thread-safety issues
 * with older GnuTLS or OpenSSL libraries. If libavformat is linked
 * to newer versions of those libraries, or if you do not use them,
 * calling this function is unnecessary. Otherwise, you need to call
 * this function before any other threads using them are started.
 *
 * This function will be deprecated once support for older GnuTLS and
 * OpenSSL libraries is removed, and this function has no purpose
 * anymore.
 */
func AvformatNetworkInit() int {
	return int(C.avformat_network_init())
}

/**
 * Undo the initialization done by avformat_network_init. Call it only
 * once for each time you called avformat_network_init.
 */
func AvformatNetworkDeinit() int {
	return int(C.avformat_network_deinit())
}

/**
 * Iterate over all registered muxers.
 *
 * @param opaque a pointer where libavformat will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered muxer or NULL when the iteration is
 *         finished
 */
func AvMuxerIterate(opaque *unsafe.Pointer) *CAVOutputFormat {
	return (*CAVOutputFormat)(C.av_muxer_iterate(opaque))
}

/**
 * Iterate over all registered demuxers.
 *
 * @param opaque a pointer where libavformat will store the iteration state.
 *               Must point to NULL to start the iteration.
 *
 * @return the next registered demuxer or NULL when the iteration is
 *         finished
 */
func AvDeuxerIterate(opaque *unsafe.Pointer) *CAVInputFormat {
	return (*CAVInputFormat)(C.av_demuxer_iterate(opaque))
}

/**
 * Allocate an AVFormatContext.
 * avformat_free_context() can be used to free the context and everything
 * allocated by the framework within it.
 */
func AvformatAllocContext() *CAVFormatContext {
	return (*CAVFormatContext)(C.avformat_alloc_context())
}

/**
 * Free an AVFormatContext and all its streams.
 * @param s context to free
 */
func AvformatFreeContext(s *CAVFormatContext) {
	C.avformat_free_context((*C.AVFormatContext)(s))
}

/**
 * Get the AVClass for AVFormatContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
func AvformatGetClass() *avutil.CAVClass {
	return (*avutil.CAVClass)(unsafe.Pointer((C.avformat_get_class())))
}

/**
 * Get the AVClass for AVStream. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
func AvStreamGetClass() *avutil.CAVClass {
	return (*avutil.CAVClass)(unsafe.Pointer((C.av_stream_get_class())))
}

/**
 * Add a new stream to a media file.
 *
 * When demuxing, it is called by the demuxer in read_header(). If the
 * flag AVFMTCTX_NOHEADER is set in s.ctx_flags, then it may also
 * be called in read_packet().
 *
 * When muxing, should be called by the user before avformat_write_header().
 *
 * User is required to call avformat_free_context() to clean up the allocation
 * by avformat_new_stream().
 *
 * @param s media file handle
 * @param c unused, does nothing
 *
 * @return newly created stream or NULL on error.
 */
func AvformatNewStream(s *CAVFormatContext, c *avcodec.CAVCodec) *CAVStream {
	return (*CAVStream)(C.avformat_new_stream((*C.AVFormatContext)(s), (*C.AVCodec)(unsafe.Pointer(c))))
}

// #if FF_API_AVSTREAM_SIDE_DATA
// /**
//  * Wrap an existing array as stream side data.
//  *
//  * @param st   stream
//  * @param type side information type
//  * @param data the side data array. It must be allocated with the av_malloc()
//  *             family of functions. The ownership of the data is transferred to
//  *             st.
//  * @param size side information size
//  *
//  * @return zero on success, a negative AVERROR code on failure. On failure,
//  *         the stream is unchanged and the data remains owned by the caller.
//  * @deprecated use av_packet_side_data_add() with the stream's
//  *             @ref AVCodecParameters.coded_side_data "codecpar side data"
//  */
// attribute_deprecated
// int av_stream_add_side_data(AVStream *st, enum AVPacketSideDataType type,
//                             uint8_t *data, size_t size);

// /**
//  * Allocate new information from stream.
//  *
//  * @param stream stream
//  * @param type   desired side information type
//  * @param size   side information size
//  *
//  * @return pointer to fresh allocated data or NULL otherwise
//  * @deprecated use av_packet_side_data_new() with the stream's
//  *             @ref AVCodecParameters.coded_side_data "codecpar side data"
//  */
// attribute_deprecated
// uint8_t *av_stream_new_side_data(AVStream *stream,
//                                  enum AVPacketSideDataType type, size_t size);
// /**
//  * Get side information from stream.
//  *
//  * @param stream stream
//  * @param type   desired side information type
//  * @param size   If supplied, *size will be set to the size of the side data
//  *               or to zero if the desired side data is not present.
//  *
//  * @return pointer to data if present or NULL otherwise
//  * @deprecated use av_packet_side_data_get() with the stream's
//  *             @ref AVCodecParameters.coded_side_data "codecpar side data"
//  */
// attribute_deprecated
// uint8_t *av_stream_get_side_data(const AVStream *stream,
//                                  enum AVPacketSideDataType type, size_t *size);
// #endif

func AvNewProgram(s *CAVFormatContext, id int) *CAVProgram {
	return (*CAVProgram)(C.av_new_program((*C.AVFormatContext)(s), C.int(id)))
}

/**
 * @}
 */

/**
 * Allocate an AVFormatContext for an output format.
 * avformat_free_context() can be used to free the context and
 * everything allocated by the framework within it.
 *
 * @param ctx           pointee is set to the created format context,
 *                      or to NULL in case of failure
 * @param oformat       format to use for allocating the context, if NULL
 *                      format_name and filename are used instead
 * @param format_name   the name of output format to use for allocating the
 *                      context, if NULL filename is used instead
 * @param filename      the name of the filename to use for allocating the
 *                      context, may be NULL
 *
 * @return  >= 0 in case of success, a negative AVERROR code in case of
 *          failure
 */
func AvformatAllocOutputContext2(ctx **CAVFormatContext, oformat *CAVOutputFormat, format_name string, filename string) int {
	var cFormatName *C.char = nil
	if len(format_name) > 0 {
		cFormatName = C.CString(format_name)
		defer C.free(unsafe.Pointer(cFormatName))
	}

	var cFilename *C.char = nil
	if len(filename) > 0 {
		cFilename = C.CString(filename)
		defer C.free(unsafe.Pointer(cFilename))
	}

	return int(C.avformat_alloc_output_context2((**C.AVFormatContext)(unsafe.Pointer(ctx)), (*C.AVOutputFormat)(unsafe.Pointer(oformat)), cFormatName, cFilename))
}

/**
 * @addtogroup lavf_decoding
 * @{
 */

/**
 * Find AVInputFormat based on the short name of the input format.
 */
func AvFindInputFormat(shortName string) *CAVInputFormat {
	var cShortName *C.char = nil
	if len(shortName) > 0 {
		cShortName = C.CString(shortName)
		defer C.free(unsafe.Pointer(cShortName))
	}

	return (*CAVInputFormat)(C.av_find_input_format(cShortName))
}

/**
 * Guess the file format.
 *
 * @param pd        data to be probed
 * @param is_opened Whether the file is already opened; determines whether
 *                  demuxers with or without AVFMT_NOFILE are probed.
 */
func AvProbeInputFormat(pd *CAVProbeData, isOpened int) *CAVInputFormat {
	return (*CAVInputFormat)(C.av_probe_input_format((*C.AVProbeData)(pd), C.int(isOpened)))
}

/**
 * Guess the file format.
 *
 * @param pd        data to be probed
 * @param is_opened Whether the file is already opened; determines whether
 *                  demuxers with or without AVFMT_NOFILE are probed.
 * @param score_max A probe score larger that this is required to accept a
 *                  detection, the variable is set to the actual detection
 *                  score afterwards.
 *                  If the score is <= AVPROBE_SCORE_MAX / 4 it is recommended
 *                  to retry with a larger probe buffer.
 */
func AvProbeInputFormat2(pd *CAVProbeData, isOpened int, scoreMax *ctypes.Int) *CAVInputFormat {
	return (*CAVInputFormat)(C.av_probe_input_format2((*C.AVProbeData)(pd), C.int(isOpened), (*C.int)(scoreMax)))
}

/**
 * Guess the file format.
 *
 * @param is_opened Whether the file is already opened; determines whether
 *                  demuxers with or without AVFMT_NOFILE are probed.
 * @param score_ret The score of the best detection.
 */
func AvProbeInputFormat3(pd *CAVProbeData, isOpened int, scoreRet *ctypes.Int) *CAVInputFormat {
	return (*CAVInputFormat)(C.av_probe_input_format3((*C.AVProbeData)(pd), C.int(isOpened), (*C.int)(scoreRet)))
}

/**
 * Probe a bytestream to determine the input format. Each time a probe returns
 * with a score that is too low, the probe buffer size is increased and another
 * attempt is made. When the maximum probe size is reached, the input format
 * with the highest score is returned.
 *
 * @param pb             the bytestream to probe
 * @param fmt            the input format is put here
 * @param url            the url of the stream
 * @param logctx         the log context
 * @param offset         the offset within the bytestream to probe from
 * @param max_probe_size the maximum probe buffer size (zero for default)
 *
 * @return the score in case of success, a negative value corresponding to an
 *         the maximal score is AVPROBE_SCORE_MAX
 *         AVERROR code otherwise
 */
func AvProbeInputBuffer2(pd *CAVIOContext, fmt **CAVInputFormat, url string, logctx unsafe.Pointer, offset uint, maxProbeSize uint) int {
	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
		defer C.free(unsafe.Pointer(cUrl))
	}

	return int(C.av_probe_input_buffer2(
		(*C.AVIOContext)(pd),
		(**C.AVInputFormat)(unsafe.Pointer(fmt)),
		cUrl,
		logctx,
		C.uint(offset),
		C.uint(maxProbeSize),
	))
}

/**
 * Like av_probe_input_buffer2() but returns 0 on success
 */
func AvProbeInputBuffer(pd *CAVIOContext, fmt **CAVInputFormat, url string, logctx unsafe.Pointer, offset uint, maxProbeSize uint) int {
	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
		defer C.free(unsafe.Pointer(cUrl))
	}

	return int(C.av_probe_input_buffer(
		(*C.AVIOContext)(pd),
		(**C.AVInputFormat)(unsafe.Pointer(fmt)),
		cUrl,
		logctx,
		C.uint(offset),
		C.uint(maxProbeSize),
	))
}

/**
 * Open an input stream and read the header. The codecs are not opened.
 * The stream must be closed with avformat_close_input().
 *
 * @param ps       Pointer to user-supplied AVFormatContext (allocated by
 *                 avformat_alloc_context). May be a pointer to NULL, in
 *                 which case an AVFormatContext is allocated by this
 *                 function and written into ps.
 *                 Note that a user-supplied AVFormatContext will be freed
 *                 on failure.
 * @param url      URL of the stream to open.
 * @param fmt      If non-NULL, this parameter forces a specific input format.
 *                 Otherwise the format is autodetected.
 * @param options  A dictionary filled with AVFormatContext and demuxer-private
 *                 options.
 *                 On return this parameter will be destroyed and replaced with
 *                 a dict containing options that were not found. May be NULL.
 *
 * @return 0 on success, a negative AVERROR on failure.
 *
 * @note If you want to use custom IO, preallocate the format context and set its pb field.
 */
func AvformatOpenInput(ps **CAVFormatContext, url string, fmt *CAVInputFormat, options **avutil.CAVDictionary) int {
	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
		defer C.free(unsafe.Pointer(cUrl))
	}

	return int(C.avformat_open_input((**C.AVFormatContext)(unsafe.Pointer(ps)), cUrl, (*C.AVInputFormat)(fmt), (**C.AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Read packets of a media file to get stream information. This
 * is useful for file formats with no headers such as MPEG. This
 * function also computes the real framerate in case of MPEG-2 repeat
 * frame mode.
 * The logical file position is not changed by this function;
 * examined packets may be buffered for later processing.
 *
 * @param ic media file handle
 * @param options  If non-NULL, an ic.nb_streams long array of pointers to
 *                 dictionaries, where i-th member contains options for
 *                 codec corresponding to i-th stream.
 *                 On return each dictionary will be filled with options that were not found.
 * @return >=0 if OK, AVERROR_xxx on error
 *
 * @note this function isn't guaranteed to open all the codecs, so
 *       options being non-empty at return is a perfectly normal behavior.
 *
 * @todo Let the user decide somehow what information is needed so that
 *       we do not waste time getting stuff the user does not need.
 */
func AvformatFindStreamInfo(ic *CAVFormatContext, options **avutil.CAVDictionary) int {
	return int(C.avformat_find_stream_info((*C.AVFormatContext)(ic), (**C.AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Find the programs which belong to a given stream.
 *
 * @param ic    media file handle
 * @param last  the last found program, the search will start after this
 *              program, or from the beginning if it is NULL
 * @param s     stream index
 *
 * @return the next program which belongs to s, NULL if no program is found or
 *         the last program is not among the programs of ic.
 */
func AvFindProgramFromStream(ic *CAVFormatContext, last *CAVProgram, s int) *CAVProgram {
	return (*CAVProgram)(C.av_find_program_from_stream((*C.AVFormatContext)(ic), (*C.AVProgram)(last), C.int(s)))
}

func AvProgramAddStreamIndex(ac *CAVFormatContext, progid int, idx uint) {
	C.av_program_add_stream_index((*C.AVFormatContext)(ac), C.int(progid), C.uint(idx))
}

/**
 * Find the "best" stream in the file.
 * The best stream is determined according to various heuristics as the most
 * likely to be what the user expects.
 * If the decoder parameter is non-NULL, av_find_best_stream will find the
 * default decoder for the stream's codec; streams for which no decoder can
 * be found are ignored.
 *
 * @param ic                media file handle
 * @param type              stream type: video, audio, subtitles, etc.
 * @param wanted_stream_nb  user-requested stream number,
 *                          or -1 for automatic selection
 * @param related_stream    try to find a stream related (eg. in the same
 *                          program) to this one, or -1 if none
 * @param decoder_ret       if non-NULL, returns the decoder for the
 *                          selected stream
 * @param flags             flags; none are currently defined
 *
 * @return  the non-negative stream number in case of success,
 *          AVERROR_STREAM_NOT_FOUND if no stream with the requested type
 *          could be found,
 *          AVERROR_DECODER_NOT_FOUND if streams were found but no decoder
 *
 * @note  If av_find_best_stream returns successfully and decoder_ret is not
 *        NULL, then *decoder_ret is guaranteed to be set to a valid AVCodec.
 */
func AvFindBestStream(ic *CAVFormatContext, _type avutil.CAVMediaType, wanted_stream_nb int,
	related_stream int, decoder_ret **avcodec.CAVCodec, flags int) int {
	return int(C.av_find_best_stream((*C.AVFormatContext)(ic), (C.enum_AVMediaType)(_type), C.int(wanted_stream_nb),
		C.int(related_stream), (**C.AVCodec)(unsafe.Pointer(decoder_ret)), C.int(flags)))
}

/**
 * Return the next frame of a stream.
 * This function returns what is stored in the file, and does not validate
 * that what is there are valid frames for the decoder. It will split what is
 * stored in the file into frames and return one for each call. It will not
 * omit invalid data between valid frames so as to give the decoder the maximum
 * information possible for decoding.
 *
 * On success, the returned packet is reference-counted (pkt->buf is set) and
 * valid indefinitely. The packet must be freed with av_packet_unref() when
 * it is no longer needed. For video, the packet contains exactly one frame.
 * For audio, it contains an integer number of frames if each frame has
 * a known fixed size (e.g. PCM or ADPCM data). If the audio frames have
 * a variable size (e.g. MPEG audio), then it contains one frame.
 *
 * pkt->pts, pkt->dts and pkt->duration are always set to correct
 * values in AVStream.time_base units (and guessed if the format cannot
 * provide them). pkt->pts can be AV_NOPTS_VALUE if the video format
 * has B-frames, so it is better to rely on pkt->dts if you do not
 * decompress the payload.
 *
 * @return 0 if OK, < 0 on error or end of file. On error, pkt will be blank
 *         (as if it came from av_packet_alloc()).
 *
 * @note pkt will be initialized, so it may be uninitialized, but it must not
 *       contain data that needs to be freed.
 */
func AvReadFrame(s *CAVFormatContext, pkt *avcodec.CAVPacket) int {
	return int(C.av_read_frame((*C.AVFormatContext)(s), (*C.AVPacket)(unsafe.Pointer(pkt))))
}

/**
 * Seek to the keyframe at timestamp.
 * 'timestamp' in 'stream_index'.
 *
 * @param s            media file handle
 * @param stream_index If stream_index is (-1), a default stream is selected,
 *                     and timestamp is automatically converted from
 *                     AV_TIME_BASE units to the stream specific time_base.
 * @param timestamp    Timestamp in AVStream.time_base units or, if no stream
 *                     is specified, in AV_TIME_BASE units.
 * @param flags        flags which select direction and seeking mode
 *
 * @return >= 0 on success
 */
func AvSeekFrame(s *CAVFormatContext, streamIndex int, timestamp int64, flags int) int {
	return int(C.av_seek_frame((*C.AVFormatContext)(s), C.int(streamIndex), C.int64_t(timestamp), C.int(flags)))
}

/**
* Seek to timestamp ts.
* Seeking will be done so that the point from which all active streams
* can be presented successfully will be closest to ts and within min/max_ts.
* Active streams are all streams that have AVStream.discard < AVDISCARD_ALL.
*
* If flags contain AVSEEK_FLAG_BYTE, then all timestamps are in bytes and
* are the file position (this may not be supported by all demuxers).
* If flags contain AVSEEK_FLAG_FRAME, then all timestamps are in frames
* in the stream with stream_index (this may not be supported by all demuxers).
* Otherwise all timestamps are in units of the stream selected by stream_index
* or if stream_index is -1, in AV_TIME_BASE units.
* If flags contain AVSEEK_FLAG_ANY, then non-keyframes are treated as
* keyframes (this may not be supported by all demuxers).
* If flags contain AVSEEK_FLAG_BACKWARD, it is ignored.
*
* @param s            media file handle
* @param stream_index index of the stream which is used as time base reference
* @param min_ts       smallest acceptable timestamp
* @param ts           target timestamp
* @param max_ts       largest acceptable timestamp
* @param flags        flags
* @return >=0 on success, error code otherwise
*
* @note This is part of the new seek API which is still under construction.
 */
func AvformatSeekFile(s *CAVFormatContext, streamIndex int, minTs int64, ts int64, maxTs int64, flags int) int {
	return int(C.avformat_seek_file((*C.AVFormatContext)(s), C.int(streamIndex), C.int64_t(minTs), C.int64_t(ts), C.int64_t(maxTs), C.int(flags)))
}

/**
* Discard all internally buffered data. This can be useful when dealing with
* discontinuities in the byte stream. Generally works only with formats that
* can resync. This includes headerless formats like MPEG-TS/TS but should also
* work with NUT, Ogg and in a limited way AVI for example.
*
* The set of streams, the detected duration, stream parameters and codecs do
* not change when calling this function. If you want a complete reset, it's
* better to open a new AVFormatContext.
*
* This does not flush the AVIOContext (s->pb). If necessary, call
* avio_flush(s->pb) before calling this function.
*
* @param s media file handle
* @return >=0 on success, error code otherwise
 */
func AvformatFlush(s *CAVFormatContext) int {
	return int(C.avformat_flush((*C.AVFormatContext)(s)))
}

/**
* Start playing a network-based stream (e.g. RTSP stream) at the
* current position.
 */
func AvReadPlay(s *CAVFormatContext) int {
	return int(C.av_read_play((*C.AVFormatContext)(s)))
}

/**
* Pause a network-based stream (e.g. RTSP stream).
*
* Use av_read_play() to resume it.
 */
func AvReadPause(s *CAVFormatContext) int {
	return int(C.av_read_pause((*C.AVFormatContext)(s)))
}

/**
 * Close an opened input AVFormatContext. Free it and all its contents
 * and set *s to NULL.
 */
func AvformatCloseInput(s **CAVFormatContext) {
	C.avformat_close_input((**C.AVFormatContext)(unsafe.Pointer(s)))
}

/**
 * @}
 */

const (
	AVSEEK_FLAG_BACKWARD = C.AVSEEK_FLAG_BACKWARD ///< seek backward
	AVSEEK_FLAG_BYTE     = C.AVSEEK_FLAG_BYTE     ///< seeking based on position in bytes
	AVSEEK_FLAG_ANY      = C.AVSEEK_FLAG_ANY      ///< seek to any frame, even non-keyframes
	AVSEEK_FLAG_FRAME    = C.AVSEEK_FLAG_FRAME    ///< seeking based on frame number
)

/**
 * @addtogroup lavf_encoding
 * @{
 */

const (
	AVSTREAM_INIT_IN_WRITE_HEADER = C.AVSTREAM_INIT_IN_WRITE_HEADER ///< stream parameters initialized in avformat_write_header
	AVSTREAM_INIT_IN_INIT_OUTPUT  = C.AVSTREAM_INIT_IN_INIT_OUTPUT  ///< stream parameters initialized in avformat_init_output
)

/**
 * Allocate the stream private data and write the stream header to
 * an output media file.
 *
 * @param s        Media file handle, must be allocated with
 *                 avformat_alloc_context().
 *                 Its \ref AVFormatContext.oformat "oformat" field must be set
 *                 to the desired output format;
 *                 Its \ref AVFormatContext.pb "pb" field must be set to an
 *                 already opened ::AVIOContext.
 * @param options  An ::AVDictionary filled with AVFormatContext and
 *                 muxer-private options.
 *                 On return this parameter will be destroyed and replaced with
 *                 a dict containing options that were not found. May be NULL.
 *
 * @retval AVSTREAM_INIT_IN_WRITE_HEADER On success, if the codec had not already been
 *                                       fully initialized in avformat_init_output().
 * @retval AVSTREAM_INIT_IN_INIT_OUTPUT  On success, if the codec had already been fully
 *                                       initialized in avformat_init_output().
 * @retval AVERROR                       A negative AVERROR on failure.
 *
 * @see av_opt_find, av_dict_set, avio_open, av_oformat_next, avformat_init_output.
 */
func AvformatWriteHeader(s *CAVFormatContext, options **avutil.CAVDictionary) int {
	return int(C.avformat_write_header((*C.AVFormatContext)(s), (**C.AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Allocate the stream private data and initialize the codec, but do not write the header.
 * May optionally be used before avformat_write_header() to initialize stream parameters
 * before actually writing the header.
 * If using this function, do not pass the same options to avformat_write_header().
 *
 * @param s        Media file handle, must be allocated with
 *                 avformat_alloc_context().
 *                 Its \ref AVFormatContext.oformat "oformat" field must be set
 *                 to the desired output format;
 *                 Its \ref AVFormatContext.pb "pb" field must be set to an
 *                 already opened ::AVIOContext.
 * @param options  An ::AVDictionary filled with AVFormatContext and
 *                 muxer-private options.
 *                 On return this parameter will be destroyed and replaced with
 *                 a dict containing options that were not found. May be NULL.
 *
 * @retval AVSTREAM_INIT_IN_WRITE_HEADER On success, if the codec requires
 *                                       avformat_write_header to fully initialize.
 * @retval AVSTREAM_INIT_IN_INIT_OUTPUT  On success, if the codec has been fully
 *                                       initialized.
 * @retval AVERROR                       Anegative AVERROR on failure.
 *
 * @see av_opt_find, av_dict_set, avio_open, av_oformat_next, avformat_write_header.
 */
func AvformatInitOutput(s *CAVFormatContext, options **avutil.CAVDictionary) int {
	return int(C.avformat_init_output((*C.AVFormatContext)(s), (**C.AVDictionary)(unsafe.Pointer(options))))
}

/**
 * Write a packet to an output media file.
 *
 * This function passes the packet directly to the muxer, without any buffering
 * or reordering. The caller is responsible for correctly interleaving the
 * packets if the format requires it. Callers that want libavformat to handle
 * the interleaving should call av_interleaved_write_frame() instead of this
 * function.
 *
 * @param s media file handle
 * @param pkt The packet containing the data to be written. Note that unlike
 *            av_interleaved_write_frame(), this function does not take
 *            ownership of the packet passed to it (though some muxers may make
 *            an internal reference to the input packet).
 *            <br>
 *            This parameter can be NULL (at any time, not just at the end), in
 *            order to immediately flush data buffered within the muxer, for
 *            muxers that buffer up data internally before writing it to the
 *            output.
 *            <br>
 *            Packet's @ref AVPacket.stream_index "stream_index" field must be
 *            set to the index of the corresponding stream in @ref
 *            AVFormatContext.streams "s->streams".
 *            <br>
 *            The timestamps (@ref AVPacket.pts "pts", @ref AVPacket.dts "dts")
 *            must be set to correct values in the stream's timebase (unless the
 *            output format is flagged with the AVFMT_NOTIMESTAMPS flag, then
 *            they can be set to AV_NOPTS_VALUE).
 *            The dts for subsequent packets passed to this function must be strictly
 *            increasing when compared in their respective timebases (unless the
 *            output format is flagged with the AVFMT_TS_NONSTRICT, then they
 *            merely have to be nondecreasing).  @ref AVPacket.duration
 *            "duration") should also be set if known.
 * @return < 0 on error, = 0 if OK, 1 if flushed and there is no more data to flush
 *
 * @see av_interleaved_write_frame()
 */
func AvWriteFrame(s *CAVFormatContext, pkt *avcodec.CAVPacket) int {
	return int(C.av_write_frame((*C.AVFormatContext)(s), (*C.AVPacket)(unsafe.Pointer(pkt))))
}

/**
 * Write a packet to an output media file ensuring correct interleaving.
 *
 * This function will buffer the packets internally as needed to make sure the
 * packets in the output file are properly interleaved, usually ordered by
 * increasing dts. Callers doing their own interleaving should call
 * av_write_frame() instead of this function.
 *
 * Using this function instead of av_write_frame() can give muxers advance
 * knowledge of future packets, improving e.g. the behaviour of the mp4
 * muxer for VFR content in fragmenting mode.
 *
 * @param s media file handle
 * @param pkt The packet containing the data to be written.
 *            <br>
 *            If the packet is reference-counted, this function will take
 *            ownership of this reference and unreference it later when it sees
 *            fit. If the packet is not reference-counted, libavformat will
 *            make a copy.
 *            The returned packet will be blank (as if returned from
 *            av_packet_alloc()), even on error.
 *            <br>
 *            This parameter can be NULL (at any time, not just at the end), to
 *            flush the interleaving queues.
 *            <br>
 *            Packet's @ref AVPacket.stream_index "stream_index" field must be
 *            set to the index of the corresponding stream in @ref
 *            AVFormatContext.streams "s->streams".
 *            <br>
 *            The timestamps (@ref AVPacket.pts "pts", @ref AVPacket.dts "dts")
 *            must be set to correct values in the stream's timebase (unless the
 *            output format is flagged with the AVFMT_NOTIMESTAMPS flag, then
 *            they can be set to AV_NOPTS_VALUE).
 *            The dts for subsequent packets in one stream must be strictly
 *            increasing (unless the output format is flagged with the
 *            AVFMT_TS_NONSTRICT, then they merely have to be nondecreasing).
 *            @ref AVPacket.duration "duration" should also be set if known.
 *
 * @return 0 on success, a negative AVERROR on error.
 *
 * @see av_write_frame(), AVFormatContext.max_interleave_delta
 */
func AvInterleavedWriteFrame(s *CAVFormatContext, pkt *avcodec.CAVPacket) int {
	return int(C.av_interleaved_write_frame((*C.AVFormatContext)(s), (*C.AVPacket)(unsafe.Pointer(pkt))))
}

/**
 * Write an uncoded frame to an output media file.
 *
 * The frame must be correctly interleaved according to the container
 * specification; if not, av_interleaved_write_uncoded_frame() must be used.
 *
 * See av_interleaved_write_uncoded_frame() for details.
 */
func AvWriteUncodedFrame(s *CAVFormatContext, streamIndex int, frame *avutil.CAVFrame) int {
	return int(C.av_write_uncoded_frame((*C.AVFormatContext)(s), C.int(streamIndex), (*C.AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Write an uncoded frame to an output media file.
 *
 * If the muxer supports it, this function makes it possible to write an AVFrame
 * structure directly, without encoding it into a packet.
 * It is mostly useful for devices and similar special muxers that use raw
 * video or PCM data and will not serialize it into a byte stream.
 *
 * To test whether it is possible to use it with a given muxer and stream,
 * use av_write_uncoded_frame_query().
 *
 * The caller gives up ownership of the frame and must not access it
 * afterwards.
 *
 * @return  >=0 for success, a negative code on error
 */
func AvInterleavedWriteUncodedFrame(s *CAVFormatContext, streamIndex int, frame *avutil.CAVFrame) int {
	return int(C.av_interleaved_write_uncoded_frame((*C.AVFormatContext)(s), C.int(streamIndex), (*C.AVFrame)(unsafe.Pointer(frame))))
}

/**
 * Test whether a muxer supports uncoded frame.
 *
 * @return  >=0 if an uncoded frame can be written to that muxer and stream,
 *          <0 if not
 */
func AvWriteUncodedFrameQuery(s *CAVFormatContext, streamIndex int) int {
	return int(C.av_write_uncoded_frame_query((*C.AVFormatContext)(s), C.int(streamIndex)))
}

/**
 * Write the stream trailer to an output media file and free the
 * file private data.
 *
 * May only be called after a successful call to avformat_write_header.
 *
 * @param s media file handle
 * @return 0 if OK, AVERROR_xxx on error
 */
func AvWriteTrailer(s *CAVFormatContext) int {
	return int(C.av_write_trailer((*C.AVFormatContext)(s)))
}

/**
 * Return the output format in the list of registered output formats
 * which best matches the provided parameters, or return NULL if
 * there is no match.
 *
 * @param short_name if non-NULL checks if short_name matches with the
 *                   names of the registered formats
 * @param filename   if non-NULL checks if filename terminates with the
 *                   extensions of the registered formats
 * @param mime_type  if non-NULL checks if mime_type matches with the
 *                   MIME type of the registered formats
 */
func AvGuessFormat(shortName string, filename string, mimeType string) *CAVOutputFormat {
	var cShortName *C.char = nil
	if len(shortName) > 0 {
		cShortName = C.CString(shortName)
		defer C.free(unsafe.Pointer(cShortName))
	}

	var cFilename *C.char = nil
	if len(filename) > 0 {
		cFilename = C.CString(filename)
		defer C.free(unsafe.Pointer(cFilename))
	}

	var cMimeType *C.char = nil
	if len(mimeType) > 0 {
		cMimeType = C.CString(mimeType)
		defer C.free(unsafe.Pointer(cMimeType))
	}

	return (*CAVOutputFormat)(C.av_guess_format(cShortName, cFilename, cMimeType))
}

/**
 * Guess the codec ID based upon muxer and filename.
 */
func AvGuessCodec(fmt *CAVOutputFormat, shortName string, filename string, mimeType string, _type avutil.CAVMediaType) avcodec.CAVCodecID {
	var cShortName *C.char = nil
	if len(shortName) > 0 {
		cShortName = C.CString(shortName)
		defer C.free(unsafe.Pointer(cShortName))
	}

	var cFilename *C.char = nil
	if len(filename) > 0 {
		cFilename = C.CString(filename)
		defer C.free(unsafe.Pointer(cFilename))
	}

	var cMimeType *C.char = nil
	if len(mimeType) > 0 {
		cMimeType = C.CString(mimeType)
		defer C.free(unsafe.Pointer(cMimeType))
	}

	return avcodec.CAVCodecID(C.av_guess_codec((*C.AVOutputFormat)(fmt), cShortName, cFilename, cMimeType, C.enum_AVMediaType(_type)))
}

/**
 * Get timing information for the data currently output.
 * The exact meaning of "currently output" depends on the format.
 * It is mostly relevant for devices that have an internal buffer and/or
 * work in real time.
 * @param s          media file handle
 * @param stream     stream in the media file
 * @param[out] dts   DTS of the last packet output for the stream, in stream
 *                   time_base units
 * @param[out] wall  absolute time when that packet whas output,
 *                   in microsecond
 * @retval  0               Success
 * @retval  AVERROR(ENOSYS) The format does not support it
 *
 * @note Some formats or devices may not allow to measure dts and wall
 *       atomically.
 */
func AvGetOutputTimestamp(s *CAVFormatContext, stream int, dts *ctypes.Int64, wall *ctypes.Int64) int {
	return int(C.av_get_output_timestamp((*C.AVFormatContext)(s), C.int(stream), (*C.int64_t)(dts), (*C.int64_t)(wall)))
}

/**
 * @}
 */

/**
 * @defgroup lavf_misc Utility functions
 * @ingroup libavf
 * @{
 *
 * Miscellaneous utility functions related to both muxing and demuxing
 * (or neither).
 */

/**
 * Send a nice hexadecimal dump of a buffer to the specified file stream.
 *
 * @param f The file stream pointer where the dump should be sent to.
 * @param buf buffer
 * @param size buffer size
 *
 * @see av_hex_dump_log, av_pkt_dump2, av_pkt_dump_log2
 */
func AvHexDump(f *C.FILE, buf unsafe.Pointer, size int) {
	C.av_hex_dump(f, (*C.uint8_t)(buf), C.int(size))
}

/**
 * Send a nice hexadecimal dump of a buffer to the log.
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 * pointer to an AVClass struct.
 * @param level The importance level of the message, lower values signifying
 * higher importance.
 * @param buf buffer
 * @param size buffer size
 *
 * @see av_hex_dump, av_pkt_dump2, av_pkt_dump_log2
 */
func AvHexDumpLog(avcl unsafe.Pointer, level int, buf unsafe.Pointer, size int) {
	C.av_hex_dump_log(avcl, C.int(level), (*C.uint8_t)(buf), C.int(size))
}

/**
 * Send a nice dump of a packet to the specified file stream.
 *
 * @param f The file stream pointer where the dump should be sent to.
 * @param pkt packet to dump
 * @param dump_payload True if the payload must be displayed, too.
 * @param st AVStream that the packet belongs to
 */
func AvPktDump2(f *C.FILE, pkt *avcodec.CAVPacket, dumpPayload int, st *CAVStream) {
	C.av_pkt_dump2(f, (*C.AVPacket)(unsafe.Pointer(pkt)), C.int(dumpPayload), (*C.AVStream)(st))
}

/**
 * Send a nice dump of a packet to the log.
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 * pointer to an AVClass struct.
 * @param level The importance level of the message, lower values signifying
 * higher importance.
 * @param pkt packet to dump
 * @param dump_payload True if the payload must be displayed, too.
 * @param st AVStream that the packet belongs to
 */
func AvPktDumpLog2(avcl unsafe.Pointer, level int, pkt *avcodec.CAVPacket, dumpPayload int, st *CAVStream) {
	C.av_pkt_dump_log2(avcl, C.int(level), (*C.AVPacket)(unsafe.Pointer(pkt)), C.int(dumpPayload), (*C.AVStream)(st))
}

/**
 * Get the AVCodecID for the given codec tag tag.
 * If no codec id is found returns AV_CODEC_ID_NONE.
 *
 * @param tags list of supported codec_id-codec_tag pairs, as stored
 * in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
 * @param tag  codec tag to match to a codec ID
 */
func AvCodecGetId(tags **CAVCodecTag, tag uint) avcodec.CAVCodecID {
	return avcodec.CAVCodecID(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(tags)), C.uint(tag)))
}

/**
 * Get the codec tag for the given codec id id.
 * If no codec tag is found returns 0.
 *
 * @param tags list of supported codec_id-codec_tag pairs, as stored
 * in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
 * @param id   codec ID to match to a codec tag
 */
func AvCodecGetTag(tags **CAVCodecTag, id avcodec.CAVCodecID) uint {
	return uint(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(tags)), (C.enum_AVCodecID)(id)))
}

/**
 * Get the codec tag for the given codec id.
 *
 * @param tags list of supported codec_id - codec_tag pairs, as stored
 * in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
 * @param id codec id that should be searched for in the list
 * @param tag A pointer to the found tag
 * @return 0 if id was not found in tags, > 0 if it was found
 */
func AvCodecGetTag2(tags **CAVCodecTag, id avcodec.CAVCodecID, tag *ctypes.UInt) int {
	return int(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(tags)), (C.enum_AVCodecID)(id), (*C.uint)(unsafe.Pointer(tag))))
}

func AvFindDefaultStreamIndex(s *CAVFormatContext) int {
	return int(C.av_find_default_stream_index((*C.AVFormatContext)(s)))
}

/**
 * Get the index for a specific timestamp.
 *
 * @param st        stream that the timestamp belongs to
 * @param timestamp timestamp to retrieve the index for
 * @param flags if AVSEEK_FLAG_BACKWARD then the returned index will correspond
 *                 to the timestamp which is <= the requested one, if backward
 *                 is 0, then it will be >=
 *              if AVSEEK_FLAG_ANY seek to any frame, only keyframes otherwise
 * @return < 0 if no such timestamp could be found
 */
func AvIndexSearchTimestamp(st *CAVStream, timestamp int64, flags int) int {
	return int(C.av_index_search_timestamp((*C.AVStream)(st), C.int64_t(timestamp), C.int(flags)))
}

/**
 * Get the index entry count for the given AVStream.
 *
 * @param st stream
 * @return the number of index entries in the stream
 */
func AvformatIndexGetEntriesCount(st *CAVStream) int {
	return int(C.avformat_index_get_entries_count((*C.AVStream)(st)))
}

/**
 * Get the AVIndexEntry corresponding to the given index.
 *
 * @param st          Stream containing the requested AVIndexEntry.
 * @param idx         The desired index.
 * @return A pointer to the requested AVIndexEntry if it exists, NULL otherwise.
 *
 * @note The pointer returned by this function is only guaranteed to be valid
 *       until any function that takes the stream or the parent AVFormatContext
 *       as input argument is called.
 */
func AvformatIndexGetEntry(st *CAVStream, idx int) *CAVIndexEntry {
	return (*CAVIndexEntry)(C.avformat_index_get_entry((*C.AVStream)(st), C.int(idx)))
}

/**
 * Get the AVIndexEntry corresponding to the given timestamp.
 *
 * @param st          Stream containing the requested AVIndexEntry.
 * @param wanted_timestamp   Timestamp to retrieve the index entry for.
 * @param flags       If AVSEEK_FLAG_BACKWARD then the returned entry will correspond
 *                    to the timestamp which is <= the requested one, if backward
 *                    is 0, then it will be >=
 *                    if AVSEEK_FLAG_ANY seek to any frame, only keyframes otherwise.
 * @return A pointer to the requested AVIndexEntry if it exists, NULL otherwise.
 *
 * @note The pointer returned by this function is only guaranteed to be valid
 *       until any function that takes the stream or the parent AVFormatContext
 *       as input argument is called.
 */
func AvformatIndexGetEntryFromTimestamp(st *CAVStream, wantedTimestamp int64, flags int) *CAVIndexEntry {
	return (*CAVIndexEntry)(C.avformat_index_get_entry_from_timestamp((*C.AVStream)(st), C.int64_t(wantedTimestamp), C.int(flags)))
}

/**
 * Add an index entry into a sorted list. Update the entry if the list
 * already contains it.
 *
 * @param timestamp timestamp in the time base of the given stream
 */
func AvAddIndexEntry(st *CAVStream, pos int64, timestamp int64,
	size int, distance int, flags int) int {
	return int(C.av_add_index_entry((*C.AVStream)(st), C.int64_t(pos), C.int64_t(timestamp),
		C.int(size), C.int(distance), C.int(flags)))
}

/**
 * Split a URL string into components.
 *
 * The pointers to buffers for storing individual components may be null,
 * in order to ignore that component. Buffers for components not found are
 * set to empty strings. If the port is not found, it is set to a negative
 * value.
 *
 * @param proto the buffer for the protocol
 * @param proto_size the size of the proto buffer
 * @param authorization the buffer for the authorization
 * @param authorization_size the size of the authorization buffer
 * @param hostname the buffer for the host name
 * @param hostname_size the size of the hostname buffer
 * @param port_ptr a pointer to store the port number in
 * @param path the buffer for the path
 * @param path_size the size of the path buffer
 * @param url the URL to split
 */
func AvUrlSplit(proto unsafe.Pointer, protoSize int,
	authorization unsafe.Pointer, authorizationSize int,
	hostname unsafe.Pointer, hostnameSize int,
	portPtr *ctypes.Int,
	path unsafe.Pointer, pathSize int,
	url string) {

	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
		defer C.free(unsafe.Pointer(cUrl))
	}

	C.av_url_split((*C.char)(proto), C.int(protoSize),
		(*C.char)(authorization), C.int(authorizationSize),
		(*C.char)(hostname), C.int(hostnameSize),
		(*C.int)(portPtr),
		(*C.char)(path), C.int(pathSize),
		cUrl)
}

/**
 * Print detailed information about the input or output format, such as
 * duration, bitrate, streams, container, programs, metadata, side data,
 * codec and time base.
 *
 * @param ic        the context to analyze
 * @param index     index of the stream to dump information about
 * @param url       the URL to print, such as source or destination file
 * @param is_output Select whether the specified context is an input(0) or output(1)
 */
func AvDumpFormat(ic *CAVFormatContext, index int, url string, is_output int) {
	var cUrl *C.char = nil
	if len(url) > 0 {
		cUrl = C.CString(url)
		defer C.free(unsafe.Pointer(cUrl))
	}

	C.av_dump_format((*C.AVFormatContext)(ic), C.int(index), cUrl, C.int(is_output))
}

const AV_FRAME_FILENAME_FLAGS_MULTIPLE = C.AV_FRAME_FILENAME_FLAGS_MULTIPLE ///< Allow multiple %d

/**
 * Return in 'buf' the path with '%d' replaced by a number.
 *
 * Also handles the '%0nd' format where 'n' is the total number
 * of digits and '%%'.
 *
 * @param buf destination buffer
 * @param buf_size destination buffer size
 * @param path numbered sequence string
 * @param number frame number
 * @param flags AV_FRAME_FILENAME_FLAGS_*
 * @return 0 if OK, -1 on format error
 */
func AvGetFrameFilename2(buf unsafe.Pointer, bufSize int,
	path string, number int, flags int) int {
	var cPath *C.char = nil
	if len(path) > 0 {
		cPath = C.CString(path)
		defer C.free(unsafe.Pointer(cPath))
	}

	return int(C.av_get_frame_filename2((*C.char)(buf), C.int(bufSize),
		cPath, C.int(number), C.int(flags)))
}

func AvGetFrameFilename(buf unsafe.Pointer, bufSize int,
	path string, number int) int {
	var cPath *C.char = nil
	if len(path) > 0 {
		cPath = C.CString(path)
		defer C.free(unsafe.Pointer(cPath))
	}

	return int(C.av_get_frame_filename((*C.char)(buf), C.int(bufSize),
		cPath, C.int(number)))
}

/**
 * Check whether filename actually is a numbered sequence generator.
 *
 * @param filename possible numbered sequence string
 * @return 1 if a valid numbered sequence string, 0 otherwise
 */
func AvFilenameNumberTest(filename string) int {
	var cFilename *C.char = nil
	if len(filename) > 0 {
		cFilename = C.CString(filename)
		defer C.free(unsafe.Pointer(cFilename))
	}

	return int(C.av_filename_number_test(cFilename))
}

/**
 * Generate an SDP for an RTP session.
 *
 * Note, this overwrites the id values of AVStreams in the muxer contexts
 * for getting unique dynamic payload types.
 *
 * @param ac array of AVFormatContexts describing the RTP streams. If the
 *           array is composed by only one context, such context can contain
 *           multiple AVStreams (one AVStream per RTP stream). Otherwise,
 *           all the contexts in the array (an AVCodecContext per RTP stream)
 *           must contain only one AVStream.
 * @param n_files number of AVCodecContexts contained in ac
 * @param buf buffer where the SDP will be stored (must be allocated by
 *            the caller)
 * @param size the size of the buffer
 * @return 0 if OK, AVERROR_xxx on error
 */
func AvSdpCreate(ac []*CAVFormatContext, nFiles int, buf unsafe.Pointer, size int) int {
	cAc := unsafe.SliceData(ac)
	return int(C.av_sdp_create((**C.AVFormatContext)(unsafe.Pointer(cAc)), C.int(nFiles), (*C.char)(buf), C.int(size)))
}

/**
 * Return a positive value if the given filename has one of the given
 * extensions, 0 otherwise.
 *
 * @param filename   file name to check against the given extensions
 * @param extensions a comma-separated list of filename extensions
 */
func AvMatchExt(filename string, extensions string) int {
	var cFilename *C.char = nil
	if len(filename) > 0 {
		cFilename = C.CString(filename)
		defer C.free(unsafe.Pointer(cFilename))
	}

	var cExtensions *C.char = nil
	if len(extensions) > 0 {
		cExtensions = C.CString(extensions)
		defer C.free(unsafe.Pointer(cExtensions))
	}

	return int(C.av_match_ext(cFilename, cExtensions))
}

/**
 * Test if the given container can store a codec.
 *
 * @param ofmt           container to check for compatibility
 * @param codec_id       codec to potentially store in container
 * @param std_compliance standards compliance level, one of FF_COMPLIANCE_*
 *
 * @return 1 if codec with ID codec_id can be stored in ofmt, 0 if it cannot.
 *         A negative number if this information is not available.
 */
func AvformatQueryCodec(ofmt *CAVOutputFormat, codecId avcodec.CAVCodecID,
	stdCompliance int) int {
	return int(C.avformat_query_codec((*C.AVOutputFormat)(ofmt), C.enum_AVCodecID(codecId),
		C.int(stdCompliance)))
}

/**
 * @defgroup riff_fourcc RIFF FourCCs
 * @{
 * Get the tables mapping RIFF FourCCs to libavcodec AVCodecIDs. The tables are
 * meant to be passed to av_codec_get_id()/av_codec_get_tag() as in the
 * following code:
 * @code
 * uint32_t tag = MKTAG('H', '2', '6', '4');
 * const struct AVCodecTag *table[] = { avformat_get_riff_video_tags(), 0 };
 * enum AVCodecID id = av_codec_get_id(table, tag);
 * @endcode
 */

/**
 * @return the table mapping RIFF FourCCs for video to libavcodec AVCodecID.
 */
func AvformatGetRiffVideoTags() *CAVCodecTag {
	return (*CAVCodecTag)(C.avformat_get_riff_video_tags())
}

/**
 * @return the table mapping RIFF FourCCs for audio to AVCodecID.
 */
func AvformatGetRiffAudioTags() *CAVCodecTag {
	return (*CAVCodecTag)(C.avformat_get_riff_audio_tags())
}

/**
 * @return the table mapping MOV FourCCs for video to libavcodec AVCodecID.
 */
func AvformatGetMovVideoTags() *CAVCodecTag {
	return (*CAVCodecTag)(C.avformat_get_mov_video_tags())
}

/**
 * @return the table mapping MOV FourCCs for audio to AVCodecID.
 */
func AvformatGetMovAudioTags() *CAVCodecTag {
	return (*CAVCodecTag)(C.avformat_get_mov_audio_tags())
}

/**
 * @}
 */

/**
 * Guess the sample aspect ratio of a frame, based on both the stream and the
 * frame aspect ratio.
 *
 * Since the frame aspect ratio is set by the codec but the stream aspect ratio
 * is set by the demuxer, these two may not be equal. This function tries to
 * return the value that you should use if you would like to display the frame.
 *
 * Basic logic is to use the stream aspect ratio if it is set to something sane
 * otherwise use the frame aspect ratio. This way a container setting, which is
 * usually easy to modify can override the coded value in the frames.
 *
 * @param format the format context which the stream is part of
 * @param stream the stream which the frame is part of
 * @param frame the frame with the aspect ratio to be determined
 * @return the guessed (valid) sample_aspect_ratio, 0/1 if no idea
 */
func AvGuessSampleAspectRatio(format *CAVFormatContext, stream *CAVStream,
	frame *avutil.CAVFrame) avutil.CAVRational {
	ret := C.av_guess_sample_aspect_ratio((*C.AVFormatContext)(format), (*C.AVStream)(stream),
		(*C.AVFrame)(unsafe.Pointer(frame)))
	return *(*avutil.CAVRational)(unsafe.Pointer(&ret))
}

/**
 * Guess the frame rate, based on both the container and codec information.
 *
 * @param ctx the format context which the stream is part of
 * @param stream the stream which the frame is part of
 * @param frame the frame for which the frame rate should be determined, may be NULL
 * @return the guessed (valid) frame rate, 0/1 if no idea
 */
func AvGuessFrameRate(ctx *CAVFormatContext, stream *CAVStream,
	frame *avutil.CAVFrame) avutil.CAVRational {
	ret := C.av_guess_frame_rate((*C.AVFormatContext)(ctx), (*C.AVStream)(stream),
		(*C.AVFrame)(unsafe.Pointer(frame)))
	return *(*avutil.CAVRational)(unsafe.Pointer(&ret))

}

/**
 * Check if the stream st contained in s is matched by the stream specifier
 * spec.
 *
 * See the "stream specifiers" chapter in the documentation for the syntax
 * of spec.
 *
 * @return  >0 if st is matched by spec;
 *          0  if st is not matched by spec;
 *          AVERROR code if spec is invalid
 *
 * @note  A stream specifier can match several streams in the format.
 */
func AvformatMatchStreamSpecifier(s *CAVFormatContext, st *CAVStream,
	spec string) int {
	var cSpec *C.char = nil
	if len(spec) > 0 {
		cSpec = C.CString(spec)
		defer C.free(unsafe.Pointer(cSpec))
	}

	return int(C.avformat_match_stream_specifier((*C.AVFormatContext)(s), (*C.AVStream)(st),
		cSpec))
}

func AvformatQueueAttachedPictures(s *CAVFormatContext) int {
	return int(C.avformat_queue_attached_pictures((*C.AVFormatContext)(s)))
}

type CAVTimebaseSource = C.enum_AVTimebaseSource

const (
	AVFMT_TBCF_AUTO    CAVTimebaseSource = C.AVFMT_TBCF_AUTO
	AVFMT_TBCF_DECODER CAVTimebaseSource = C.AVFMT_TBCF_DECODER
	AVFMT_TBCF_DEMUXER CAVTimebaseSource = C.AVFMT_TBCF_DEMUXER

	// #if FF_API_R_FRAME_RATE

	AVFMT_TBCF_R_FRAMERATE CAVTimebaseSource = C.AVFMT_TBCF_R_FRAMERATE

	// #endif
)

/**
 * Transfer internal timing information from one stream to another.
 *
 * This function is useful when doing stream copy.
 *
 * @param ofmt     target output format for ost
 * @param ost      output stream which needs timings copy and adjustments
 * @param ist      reference input stream to copy timings from
 * @param copy_tb  define from where the stream codec timebase needs to be imported
 */
func AvformatTransferInternalStreamTimingInfo(ofmt *CAVOutputFormat,
	ost *CAVStream, ist *CAVStream,
	copyTb CAVTimebaseSource) int {
	return int(C.avformat_transfer_internal_stream_timing_info((*C.AVOutputFormat)(ofmt),
		(*C.AVStream)(ost), (*C.AVStream)(ist),
		copyTb))
}

/**
 * Get the internal codec timebase from a stream.
 *
 * @param st  input stream to extract the timebase from
 */
func AvStreamGetCodecTimebase(st *CAVStream) avutil.CAVRational {
	ret := C.av_stream_get_codec_timebase((*C.AVStream)(st))
	return *(*avutil.CAVRational)(unsafe.Pointer(&ret))
}

/**
 * @}
 */

// #endif /* AVFORMAT_AVFORMAT_H */
