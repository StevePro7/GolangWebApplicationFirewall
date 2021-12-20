package waf

import (
	"fmt"
	"regexp"
	"strings"
)

const DELIM = " "

func Parser(payload string) map[string]string {

	dict := make(map[string]string)

	//str1 := "ModSecurity: Warning. detected SQLi using libinjection. [file \"/etc/waf/custom-REQUEST-942-APPLICATION-ATTACK-SQLI.conf\"] [line \"45\"] [id \"942100\"] [rev \"\"] [msg \"SQL Injection Attack Detected via libinjection\"] [data \"Matched Data: 1UE1 found within ARGS:artist: 0 div 1 union#foo*/*bar\\x0d\\x0aselect#foo\\x0d\\x0a1,2,current_user\"] [severity \"2\"] [ver \"OWASP_CRS/3.3.2\"] [maturity \"0\"] [accuracy \"0\"] [hostname \"echo-a\"] [uri \"/test/artists.php\"] [unique_id \"163956574873.676390\"] [ref \"v30,53\"]"
	//re := regexp.MustCompile(`\[([^\[\]]*)\]`)
	re := regexp.MustCompile(`\[([^\[\]]*)\]`)
	//fmt.Printf("Pattern: %v\n", re.String())      // print pattern
	//fmt.Println("Matched:", re.MatchString(str1)) // true

	//fmt.Println("\nText between square brackets:")
	submatchall := re.FindAllString(payload, -1)
	for _, element := range submatchall {
		element = strings.Trim(element, "[")
		element = strings.Trim(element, "]")
		//fmt.Println(element)
		data := strings.SplitAfterN(element, " ", 2)
		leng := len(data)
		if leng != 2 {
			fmt.Println("Yikes!")
		}
		key := strings.Trim(data[0], " ")
		val := data[1]
		//fmt.Println(leng)
		dict[key] = val
	}

	//dict["file"] = "/tmp"
	//dict["rule"] = "942100"
	return dict
}
