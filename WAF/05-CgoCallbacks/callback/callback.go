package callback

// #include "callback.h"
import "C"

func Run() {
	// call the wrapper
	//C.pass_GoAdd()
}

func GoAdd(x, y int) int {
	return x + y
}
