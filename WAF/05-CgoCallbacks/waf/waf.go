package waf

// #include "waf.h"
import "C"
import "log"

func InitializeModSecurity() {
	log.Println("GoInitializeModSecurity beg")
	C.InitializeModSecurity()
	log.Println("GoInitializeModSecurity end")
}
