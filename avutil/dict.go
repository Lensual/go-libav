package avutil

/*
#cgo pkg-config: libavutil

#include "libavutil/dict.h"
*/
import "C"
import "unsafe"

/**
 * @file
 * Public dictionary API.
 * @deprecated
 *  AVDictionary is provided for compatibility with libav. It is both in
 *  implementation as well as API inefficient. It does not scale and is
 *  extremely slow with large dictionaries.
 *  It is recommended that new code uses our tree container from tree.c/h
 *  where applicable, which uses AVL trees to achieve O(log n) performance.
 */

//  #ifndef AVUTIL_DICT_H
//  #define AVUTIL_DICT_H

//  #include <stdint.h>

/**
  * @addtogroup lavu_dict AVDictionary
  * @ingroup lavu_data
  *
  * @brief Simple key:value store
  *
  * @{
  * Dictionaries are used for storing key-value pairs.
  *
  * - To **create an AVDictionary**, simply pass an address of a NULL
  *   pointer to av_dict_set(). NULL can be used as an empty dictionary
  *   wherever a pointer to an AVDictionary is required.
  * - To **insert an entry**, use av_dict_set().
  * - Use av_dict_get() to **retrieve an entry**.
  * - To **iterate over all entries**, use av_dict_iterate().
  * - In order to **free the dictionary and all its contents**, use av_dict_free().
  *
  @code
	AVDictionary *d = NULL;           // "create" an empty dictionary
	AVDictionaryEntry *t = NULL;

	av_dict_set(&d, "foo", "bar", 0); // add an entry

	char *k = av_strdup("key");       // if your strings are already allocated,
	char *v = av_strdup("value");     // you can avoid copying them like this
	av_dict_set(&d, k, v, AV_DICT_DONT_STRDUP_KEY | AV_DICT_DONT_STRDUP_VAL);

	while ((t = av_dict_iterate(d, t))) {
		<....>                        // iterate over all entries in d
	}
	av_dict_free(&d);
  @endcode
*/

/**
 * @name AVDictionary Flags
 * Flags that influence behavior of the matching of keys or insertion to the dictionary.
 * @{
 */
const (
	AV_DICT_MATCH_CASE    = C.AV_DICT_MATCH_CASE    /**< Only get an entry with exact-case key match. Only relevant in av_dict_get(). */
	AV_DICT_IGNORE_SUFFIX = C.AV_DICT_IGNORE_SUFFIX /**< Return first entry in a dictionary whose first part corresponds to the search key,
	  ignoring the suffix of the found key string. Only relevant in av_dict_get(). */
	AV_DICT_DONT_STRDUP_KEY = C.AV_DICT_DONT_STRDUP_KEY /**< Take ownership of a key that's been
	  allocated with av_malloc() or another memory allocation function. */
	AV_DICT_DONT_STRDUP_VAL = C.AV_DICT_DONT_STRDUP_VAL /**< Take ownership of a value that's been
	  allocated with av_malloc() or another memory allocation function. */
	AV_DICT_DONT_OVERWRITE = C.AV_DICT_DONT_OVERWRITE /**< Don't overwrite existing entries. */
	AV_DICT_APPEND         = C.AV_DICT_APPEND         /**< If the entry already exists, append to it.  Note that no
	  delimiter is added, the strings are simply concatenated. */
	AV_DICT_MULTIKEY = C.AV_DICT_MULTIKEY /**< Allow to store several equal keys in the dictionary */
)

/**
 * @}
 */

type CAVDictionaryEntry C.AVDictionaryEntry

// #region CAVDictionaryEntry
func (e *CAVDictionaryEntry) GetKey() string {
	return C.GoString(e.key)
}
func (e *CAVDictionaryEntry) SetKey(key string) {
	e.key = C.CString(key)
}
func (e *CAVDictionaryEntry) GetValue() string {
	return C.GoString(e.value)
}
func (e *CAVDictionaryEntry) SetValue(value string) {
	e.value = C.CString(value)
}

//#endregion CAVDictionaryEntry

type CAVDictionary C.AVDictionary

/**
 * Get a dictionary entry with matching key.
 *
 * The returned entry key or value must not be changed, or it will
 * cause undefined behavior.
 *
 * @param prev  Set to the previous matching element to find the next.
 *              If set to NULL the first matching element is returned.
 * @param key   Matching key
 * @param flags A collection of AV_DICT_* flags controlling how the
 *              entry is retrieved
 *
 * @return      Found entry or NULL in case no matching entry was found in the dictionary
 */
func AvDictGet(m *CAVDictionary, key string, prev *CAVDictionaryEntry, flags int) *CAVDictionaryEntry {
	return (*CAVDictionaryEntry)(C.av_dict_get((*C.AVDictionary)(m), C.CString(key), (*C.AVDictionaryEntry)(prev), C.int(flags)))
}

/**
 * Iterate over a dictionary
 *
 * Iterates through all entries in the dictionary.
 *
 * @warning The returned AVDictionaryEntry key/value must not be changed.
 *
 * @warning As av_dict_set() invalidates all previous entries returned
 * by this function, it must not be called while iterating over the dict.
 *
 * Typical usage:
 * @code
 * const AVDictionaryEntry *e = NULL;
 * while ((e = av_dict_iterate(m, e))) {
 *     // ...
 * }
 * @endcode
 *
 * @param m     The dictionary to iterate over
 * @param prev  Pointer to the previous AVDictionaryEntry, NULL initially
 *
 * @retval AVDictionaryEntry* The next element in the dictionary
 * @retval NULL               No more elements in the dictionary
 */
func AvDictIterate(m *CAVDictionary, prev *CAVDictionaryEntry) *CAVDictionaryEntry {
	return (*CAVDictionaryEntry)(C.av_dict_iterate((*C.AVDictionary)(m), (*C.AVDictionaryEntry)(prev)))
}

/**
 * Get number of entries in dictionary.
 *
 * @param m dictionary
 * @return  number of entries in dictionary
 */
func AvDictCount(m *CAVDictionary) int {
	return int(C.av_dict_count((*C.AVDictionary)(m)))
}

/**
 * Set the given entry in *pm, overwriting an existing entry.
 *
 * Note: If AV_DICT_DONT_STRDUP_KEY or AV_DICT_DONT_STRDUP_VAL is set,
 * these arguments will be freed on error.
 *
 * @warning Adding a new entry to a dictionary invalidates all existing entries
 * previously returned with av_dict_get() or av_dict_iterate().
 *
 * @param pm        Pointer to a pointer to a dictionary struct. If *pm is NULL
 *                  a dictionary struct is allocated and put in *pm.
 * @param key       Entry key to add to *pm (will either be av_strduped or added as a new key depending on flags)
 * @param value     Entry value to add to *pm (will be av_strduped or added as a new key depending on flags).
 *                  Passing a NULL value will cause an existing entry to be deleted.
 *
 * @return          >= 0 on success otherwise an error code <0
 */
func AvDictSet(pm **CAVDictionary, key string, value string, flags int) int {
	return int(C.av_dict_set((**C.AVDictionary)(unsafe.Pointer(pm)), C.CString(key), C.CString(value), C.int(flags)))
}

/**
 * Convenience wrapper for av_dict_set() that converts the value to a string
 * and stores it.
 *
 * Note: If ::AV_DICT_DONT_STRDUP_KEY is set, key will be freed on error.
 */
func AvDictSetInt(pm **CAVDictionary, key string, value int64, flags int) int {
	return int(C.av_dict_set_int((**C.AVDictionary)(unsafe.Pointer(pm)), C.CString(key), C.int64_t(value), C.int(flags)))
}

/**
 * Parse the key/value pairs list and add the parsed entries to a dictionary.
 *
 * In case of failure, all the successfully set entries are stored in
 * *pm. You may need to manually free the created dictionary.
 *
 * @param key_val_sep  A 0-terminated list of characters used to separate
 *                     key from value
 * @param pairs_sep    A 0-terminated list of characters used to separate
 *                     two pairs from each other
 * @param flags        Flags to use when adding to the dictionary.
 *                     ::AV_DICT_DONT_STRDUP_KEY and ::AV_DICT_DONT_STRDUP_VAL
 *                     are ignored since the key/value tokens will always
 *                     be duplicated.
 *
 * @return             0 on success, negative AVERROR code on failure
 */
func AvDictParseString(pm **CAVDictionary, str string, keyValSep string, pairsSep string, flags int) int {
	return int(C.av_dict_parse_string((**C.AVDictionary)(unsafe.Pointer(pm)), C.CString(str), C.CString(keyValSep), C.CString(pairsSep), C.int(flags)))
}

/**
 * Copy entries from one AVDictionary struct into another.
 *
 * @note Metadata is read using the ::AV_DICT_IGNORE_SUFFIX flag
 *
 * @param dst   Pointer to a pointer to a AVDictionary struct to copy into. If *dst is NULL,
 *              this function will allocate a struct for you and put it in *dst
 * @param src   Pointer to the source AVDictionary struct to copy items from.
 * @param flags Flags to use when setting entries in *dst
 *
 * @return 0 on success, negative AVERROR code on failure. If dst was allocated
 *           by this function, callers should free the associated memory.
 */
func AvDictCopy(dst **CAVDictionary, src *CAVDictionary, flags int) int {
	return int(C.av_dict_copy((**C.AVDictionary)(unsafe.Pointer(dst)), (*C.AVDictionary)(src), C.int(flags)))
}

/**
 * Free all the memory allocated for an AVDictionary struct
 * and all keys and values.
 */
func AvDictFree(m **CAVDictionary) {
	C.av_dict_free((**C.AVDictionary)(unsafe.Pointer(m)))
}

/**
 * Get dictionary entries as a string.
 *
 * Create a string containing dictionary's entries.
 * Such string may be passed back to av_dict_parse_string().
 * @note String is escaped with backslashes ('\').
 *
 * @warning Separators cannot be neither '\\' nor '\0'. They also cannot be the same.
 *
 * @param[in]  m             The dictionary
 * @param[out] buffer        Pointer to buffer that will be allocated with string containg entries.
 *                           Buffer must be freed by the caller when is no longer needed.
 * @param[in]  key_val_sep   Character used to separate key from value
 * @param[in]  pairs_sep     Character used to separate two pairs from each other
 *
 * @return                   >= 0 on success, negative on error
 */
func AvDictGetString(m *CAVDictionary, buffer unsafe.Pointer, keyValSep byte, pairsSep byte) int {
	return int(C.av_dict_get_string((*C.AVDictionary)(m), (**C.char)(buffer), C.char(keyValSep), C.char(pairsSep)))
}

/**
 * @}
 */

//  #endif /* AVUTIL_DICT_H */
