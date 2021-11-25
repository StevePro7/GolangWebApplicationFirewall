package waf

// #include "waf.h"
import "C"

func MyGoAdd(x, y int) int {
	return x + y - 1
}

func MyCGoAdd(x, y int) int {
	return int(C.TheCAdd(x, y))
}
