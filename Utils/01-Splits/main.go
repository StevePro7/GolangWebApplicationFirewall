package main

import (
	"fmt"
	"strings"
)

func splitInput(input, delim, defaultLeft, defaultRight string) (actualLeft, actualRight string) {
	splitN := strings.SplitN(input, delim, 2)
	length := len(splitN)

	actualLeft = defaultLeft
	actualRight = defaultRight

	if length == 1 && len(splitN[0]) > 0 {
		actualLeft = splitN[0]
	}
	if length == 2 && len(splitN[1]) > 0 {
		actualRight = splitN[1]
	}

	return actualLeft, actualRight
}

func main() {

	//tempProtocol :=""
	//tempProtocol := "HTTP/2.0"
	//tempProtocol := "HTTP"
	tempProtocol := "/3.0"

	httpProtocol, httpVersion := splitInput(tempProtocol, "/", "HTTP", "1.1")
	fmt.Println(httpProtocol)
	fmt.Println(httpVersion)
}
