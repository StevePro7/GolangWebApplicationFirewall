package main

import (
	"log"
)

// #include "modsec.c"
import "C"

func initModSec() {
	log.Println("initModSec start")
	C.MyInit()
	log.Println("initModSec -end-")
}

func main() {
	log.Println("start")
	initModSec()
	log.Println("-end-")
}