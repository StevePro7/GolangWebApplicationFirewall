package waf

import (
	"log"
	"os"
)

// Assume WAF is not enabled by default.
var wafIsEnabled = false

func PreFlightCheck(directory string) {

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Println("dir does not exist")
		return
	}

	log.Println("dir exists!")
	wafIsEnabled = true
}

// IsEnabled helper function used by client calling cod
func IsEnabled() bool {
	return wafIsEnabled
}
