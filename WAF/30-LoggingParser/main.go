package main

import (
	"fmt"
	"waftesting/waf"
)

func main() {
	//fmt.Println("hello")

	payload := "ModSecurity: Warning. detected SQLi using libinjection. [file \"/etc/waf/custom-REQUEST-942-APPLICATION-ATTACK-SQLI.conf\"] [line \"45\"] [id \"942100\"] [rev \"\"] [msg \"SQL Injection Attack Detected via libinjection\"] [data \"Matched Data: 1UE1 found within ARGS:artist: 0 div 1 union#foo*/*bar\\x0d\\x0aselect#foo\\x0d\\x0a1,2,current_user\"] [severity \"2\"] [ver \"OWASP_CRS/3.3.2\"] [maturity \"0\"] [accuracy \"0\"] [hostname \"echo-a\"] [uri \"/test/artists.php\"] [unique_id \"163956574873.676390\"] [ref \"v30,53\"]"
	//data := waf.Parser(payload)
	//fmt.Println(data)

	dictionary := waf.Parser(payload)
	fmt.Println(dictionary["file"])
	fmt.Println(dictionary["line"])
	fmt.Println(dictionary["id"])
	fmt.Println(dictionary["msg"])
	fmt.Println(dictionary["uri"])
	fmt.Println(dictionary["data"])

	//	x := 1
	//	y := 2
	//	res := waf.GoAdd(x, y)
	//	fmt.Println(res)
}
