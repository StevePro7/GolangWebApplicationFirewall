package waf

import (
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

const DELIM = " "
const NumElements = 2

const ParserMsg = "msg"

func Parser(payload string) map[string]string {

	dictionary := make(map[string]string)
	//regex := regexp.MustCompile(`\[([^\[\]]*)\]`)
	regex := regexp.MustCompile(`\[([^\[\]]*)]`)

	submatchall := regex.FindAllString(payload, -1)
	for _, element := range submatchall {
		element = strings.Trim(element, "[")
		element = strings.Trim(element, "]")

		splitN := strings.SplitAfterN(element, DELIM, NumElements)
		if len(splitN) != NumElements {
			log.Errorf("WAF Parsing payload '%s' for element '%s' cannot be split into key / value", payload, element)
		}

		key := strings.Trim(splitN[0], DELIM)
		value := splitN[1]

		dictionary[key] = value
	}

	return dictionary
}
