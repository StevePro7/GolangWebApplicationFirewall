package main

import "fmt"

func main() {
	reqHeaders := make(map[string]string)
	reqHeaders["User-Agent"] = "Microsoft Internet Explorer"
	reqHeaders["content-type"] = "application/xml"

	fmt.Println(reqHeaders)
}
