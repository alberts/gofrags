package sse42

//#include <string.h>
//#cgo linux CFLAGS: -O3 -mavx
import "C"
import "unsafe"

func Memcmp(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	if len(b1) == 0 {
		return true
	}
	return C.memcmp(unsafe.Pointer(&b1[0]), unsafe.Pointer(&b2[0]), C.size_t(len(b1))) == 0
}

func Strncmp(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	if len(b1) == 0 {
		return true
	}
	return C.strncmp((*C.char)(unsafe.Pointer(&b1[0])), (*C.char)(unsafe.Pointer(&b2[0])), C.size_t(len(b1))) == 0
}

func Memchr(s []byte, c byte) int {
	if len(s) == 0 {
		return -1
	}
	s0 := unsafe.Pointer(&s[0])
	p := C.memchr(s0, C.int(c), C.size_t(len(s)))
	if p == nil {
		return -1
	}
	return int(uintptr(p) - uintptr(s0))
}

func builtinCopy(dst, src []byte) int {
	return copy(dst, src)
}

func Memmove(dst, src []byte) int {
	n := len(dst)
	if len(src) < len(dst) {
		n = len(src)
	}
	if n == 0 {
		return 0
	}
	C.memmove(unsafe.Pointer(&dst[0]), unsafe.Pointer(&src[0]), C.size_t(n))
	return n
}
