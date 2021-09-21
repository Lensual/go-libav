package advance

import (
	"reflect"
	"unsafe"
)

//映射C的内存
func UnsafePtr2ByteSlice(ptr unsafe.Pointer, len int) []byte {
	var slice []byte
	h := (*reflect.SliceHeader)((unsafe.Pointer(&slice)))
	h.Cap = len
	h.Len = len
	h.Data = uintptr(ptr)
	return slice
}
