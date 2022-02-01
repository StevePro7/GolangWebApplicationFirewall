package main

import (
	"fmt"
	"waftesting/cgoexample"
)

func main() {
	f, e := cgoexample.Open("readme", "qq")
	defer f.Close()

	if e != nil {
		fmt.Println("Error: ", e)
	}
}
