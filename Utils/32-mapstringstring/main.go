package main

import (
	"fmt"
	"waftesting/waf"
)

func unpackDictionary(dictionary map[string]string) ([]string, []string) {

	var keys []string
	var vals []string

	for key, val := range dictionary {
		keys = append(keys, key)
		vals = append(vals, val)
	}

	return keys, vals
}

func main() {
	reqHeaders := make(map[string]string)
	reqHeaders["key"] = "val"
	reqHeaders["User-Agent"] = "Microsoft Internet Explorer"
	reqHeaders["content-type"] = "application/xml"

	fmt.Println(reqHeaders)
	reqHeaderKeys, reqHeaderVals := unpackDictionary(reqHeaders)

	//reqBody := "{\"key\":\"val\"}";
	reqBody := "{stevepro}"
	bob := waf.LoadModSecurityCoreRuleSet(reqHeaderKeys, reqHeaderVals, reqBody)
	fmt.Println(bob)
}
