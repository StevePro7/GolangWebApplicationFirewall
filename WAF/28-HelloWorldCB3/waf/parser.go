package waf

import (
	"regexp"
	"strings"
)

const ParserDelim = " "
const ParserEmpty = ""
const ParserEscape = "\""
const ParserMatchAll = -1
const NumElements = 2

const ParserUniqueId = "unique_id"
const ParserFile = "file"
const ParserLine = "line"
const ParserId = "id"
const ParserMsg = "msg"
const ParserUri = "uri"
const ParserData = "data"

func Parser(payload string) map[string]string {

	dictionary := make(map[string]string)
	regex := regexp.MustCompile(`\[([^\[\]]*)]`)

	submatchall := regex.FindAllString(payload, ParserMatchAll)
	for _, element := range submatchall {
		element = strings.Trim(element, "[")
		element = strings.Trim(element, "]")

		// Record only entries in payload that conform to key / value.
		splitN := strings.SplitAfterN(element, ParserDelim, NumElements)
		if len(splitN) != NumElements {
			continue
		}

		key := strings.Trim(splitN[0], ParserDelim)
		value := strings.Replace(splitN[1], ParserEscape, ParserEmpty, ParserMatchAll)

		dictionary[key] = value
	}

	return dictionary
}
