package main

import (
	"fmt"
	"strings"
)

func main() {

	//socket := "127.0.0.1:80"
	socket := "127.0.0.1"
	split := strings.SplitN(socket, ":", 2)
	fmt.Println(len(split))
	fmt.Println(split[0])
	fmt.Println(split[1])
}
