package waf

import (
	"regexp"
	"strings"
)

const ParserDelim = " "
const ParserEmpty = ""
const ParserEscape = "\""
const ParserPrefix = "ModSecurity: "
const ParserMatchAll = -1
const NumElements = 2

const ParserFile = "file"
const ParserLine = "line"
const ParserId = "id"
const ParserMsg = "msg"
const ParserData = "data"
const ParserSeverity = "severity"
const ParserVersion = "ver"
const ParserHostname = "hostname"
const ParserUri = "uri"
const ParserUniqueId = "unique_id"

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

	// Default MSG with ModSecurity preliminary text if payload MSG is empty.
	if len(dictionary[ParserMsg]) == 0 {
		index := strings.Index(payload, "[")
		msg := strings.TrimSpace(payload[:index])
		msg = strings.Replace(msg, ParserPrefix, ParserEmpty, ParserMatchAll)
		dictionary[ParserMsg] = msg
	}

	return dictionary
}
