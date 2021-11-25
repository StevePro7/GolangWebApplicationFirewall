package waf

// #include "waf.h"
import "C"

func MyGoAdd(x, y int) int {
	return x + y - 2
}

func MyCGoAdd(x, y int) int {
	cx := C.int(x)
	cy := C.int(y)
	return int(C.TheCAdd(cx, cy))
}
