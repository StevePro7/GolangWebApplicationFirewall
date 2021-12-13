package waf

// #include "waf.h"
import "C"
import "log"

func InitializeModSecurity() {
	log.Println("GoInitializeModSecurity beg")
	C.InitializeModSecurity()
	log.Println("GoInitializeModSecurity end")
}

func Run() {
	log.Println("GoRUN beg")
	C.pass_GoAdd()
	log.Println("GoRUN end")
}

//export GoAdd
func GoAdd(a, b int) int {
	return a + b
}
