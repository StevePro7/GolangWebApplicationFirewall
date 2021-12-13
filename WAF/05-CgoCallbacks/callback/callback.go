package callback

// #include "calladd.h"
// #include "funcwrap.h"
import "C"

func Run() {
	// call the wrapper
	C.pass_GoAdd()
}

//export GoAdd
func GoAdd(a, b int) int {
	return a + b
}
