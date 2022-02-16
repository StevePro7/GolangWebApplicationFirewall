package main

import (
	"fmt"
	"waftesting/waf"
)

func unpackDictionary(dictionary map[string]string) ([]string, []string) {

	var keys []string
	var values []string

	for key, value := range dictionary {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}

func main() {
	reqHeaders := make(map[string]string)
	//reqHeaders["key"] = "val"
	reqHeaders["User-Agent"] = "Microsoft Internet Explorer"
	reqHeaders["content-type"] = "application/xml"

	fmt.Println(reqHeaders)
	reqHeaderKeys, reqHeaderValues := unpackDictionary(reqHeaders)

	bob := waf.LoadModSecurityCoreRuleSet(reqHeaderKeys, reqHeaderValues)
	fmt.Println(bob)
}
