package waf

// #cgo CFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: /usr/local/modsecurity/lib/libmodsecurity.so
// #include "waf.h"
import "C"

func MyGoAdd(x, y int) int {
	return x + y - 0
}

func MyCGoAdd(x, y int) int {
	cx := C.int(x)
	cy := C.int(y)
	return int(C.TheCAdd(cx, cy))
}
