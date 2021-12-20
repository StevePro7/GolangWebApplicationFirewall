package main

import (
	"fmt"
	"waftesting/waf"
)

func main() {
	//fmt.Println("hello")

	payload := "ModSecurity: Warning. Matched \"Operator `Rx' with parameter `(?i:(?:\\b(?:(?:s(?:elect\\b.{1,100}?\\b(?:(?:(?:length|count)\\b.{1,100}?|.*?\\bdump\\b.*)\\bfrom|to(?:p\\b.{1,100}?\\bfrom|_(?:numbe|cha)r)|(?:from\\b.{1,100}?\\bwher|data_typ)e|instr)|ys_context)|in(?:to\\b\\W* (304 characters omitted)' against variable `ARGS:artist' (Value: `0 div 1 union#foo*/*bar\\x0d\\x0aselect#foo\\x0d\\x0a1,2,current_user' ) [file \"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf\"] [line \"1102\"] [id \"942480\"] [rev \"\"] [msg \"SQL Injection Attack\"] [data \"Matched Data: union#foo*/*bar\\x0d\\x0aselect found within ARGS:artist: 0 div 1 union#foo*/*bar\\x0d\\x0aselect#foo\\x0d\\x0a1,2,current_user\"] [severity \"2\"] [ver \"OWASP_CRS/3.3.2\"] [maturity \"0\"] [accuracy \"0\"] [tag \"application-multi\"] [tag \"language-multi\"] [tag \"platform-multi\"] [tag \"attack-sqli\"] [tag \"OWASP_CRS\"] [tag \"capec/1000/152/248/66\"] [tag \"PCI/6.5.2\"] [tag \"paranoia-level/2\"] [hostname \"echo-a\"] [uri \"/test/artists.php\"] [unique_id \"16395648888.540748\"] [ref \"o8,23v30,53t:urlDecodeUni\"]"
	//data := waf.Parser(payload)
	//fmt.Println(data)

	dictionary := waf.Parser(payload)
	fmt.Println(dictionary["file"])
	fmt.Println(dictionary["line"])
	fmt.Println(dictionary["id"])
	//fmt.Println(dictionary["msg"])
	fmt.Println(dictionary[waf.ParserMsg])
	fmt.Println(dictionary["uri"])
	fmt.Println(dictionary["data"])

	//	x := 1
	//	y := 2
	//	res := waf.GoAdd(x, y)
	//	fmt.Println(res)
}
