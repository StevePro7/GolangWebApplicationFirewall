package waf

import (
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

const ParserDelim = " "
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
	//regex := regexp.MustCompile(`\[([^\[\]]*)\]`)
	regex := regexp.MustCompile(`\[([^\[\]]*)]`)

	submatchall := regex.FindAllString(payload, ParserMatchAll)
	for _, element := range submatchall {
		element = strings.Trim(element, "[")
		element = strings.Trim(element, "]")

		splitN := strings.SplitAfterN(element, ParserDelim, NumElements)
		if len(splitN) != NumElements {
			log.Errorf("WAF Parsing payload '%s' for element '%s' cannot be split into key/value pair", payload, element)
		}

		key := strings.Trim(splitN[0], ParserDelim)
		value := strings.Replace(splitN[1], ParserEscape, "", ParserMatchAll)

		dictionary[key] = value
	}

	return dictionary
}
