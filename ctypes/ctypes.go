// Share C types between multiple packages
package ctypes

/*
#include <stdint.h>
*/
import "C"
import "unsafe"

type Char C.char

type UChar C.uchar

type Int C.int

type UInt C.uint

type Int8 C.int8_t

type UInt8 C.uint8_t

type Int16 C.int16_t

type UInt16 C.uint16_t

type Int32 C.int32_t

type UInt32 C.uint32_t

type Int64 C.int64_t

type UInt64 C.uint64_t

type SizeT C.size_t

type CFunc unsafe.Pointer
