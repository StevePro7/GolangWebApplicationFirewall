package callback

// #include "calladd.h"
// #include "funcwrap.h"
import "C"
import (
	"fmt"
)

func Run() {
	fmt.Println("Go Run beg")
	C.pass_GoAdd()
	fmt.Println("Go Run end")
}

//export GoAdd
func GoAdd(a, b int) int {
	fmt.Printf("Go GoAdd beg (x,y)=(%d,%d)", a, b)
	fmt.Println()
	sum := a + b
	fmt.Printf("Go GoAdd end [%d]", sum)
	fmt.Println()
	return sum
}
