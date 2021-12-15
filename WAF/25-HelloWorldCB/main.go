package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"waftesting/checker"
)

func main() {
	log.Println("start")
	checker.Check()
	mytime := time.Now().Unix()
	mytext := strconv.FormatInt(mytime, 10)
	log.Println(mytime)
	log.Println(mytext)

	fmt.Printf("t1: %T\n", mytime)
	fmt.Printf("t2: %T\n", mytext)
	log.Println("-end-")
}
