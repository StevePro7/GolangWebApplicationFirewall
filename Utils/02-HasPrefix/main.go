package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("start")
	filename := "..2021_12_13_11_41_42.859633307"
	//test :=strings.HasPrefix("..2021", filename)
	test := strings.HasPrefix(filename, "..")
	if test {
		fmt.Println("test success")
	}

	fmt.Println(filename)

	fmt.Println(strings.HasPrefix("my string", "prefix")) // false
	fmt.Println(strings.HasPrefix("mystring", "my"))      // true
	fmt.Println(strings.HasPrefix("mystring", ""))        // true

	fmt.Println("-end-")
}
