package main

import (
	"fmt"
)

func unpackDictionary(dictionary map[string]string) ([]string, []string) {

	keys := []string{}
	values := []string{}

	for key, value := range dictionary {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}

func main() {
	reqHeaders := make(map[string]string)
	reqHeaders["key"] = "val"
	reqHeaders["User-Agent"] = "Microsoft Internet Explorer"
	reqHeaders["content-type"] = "application/xml"

	fmt.Println(reqHeaders)
	keys, values := unpackDictionary(reqHeaders)
	fmt.Println(len(keys))
	fmt.Println(len(values))
}
