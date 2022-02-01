package waf

// #include "waf.h"
import "C"
import "fmt"

func StevePro() {

	fmt.Println("beg")
	_, err := C.Adriana()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("end")
}
