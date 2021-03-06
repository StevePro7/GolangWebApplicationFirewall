package waf

import (
	"testing"
)

func TestParser_DetectedSQLiUsingLibinjection(t *testing.T) {

	payload := "ModSecurity: Warning. detected SQLi using libinjection. [file \"/etc/waf/custom-REQUEST-942-APPLICATION-ATTACK-SQLI.conf\"] [line \"45\"] [id \"942100\"] [rev \"\"] [msg \"\"] [data \"\"] [severity \"0\"] [ver \"OWASP_CRS/3.3.2\"] [maturity \"0\"] [accuracy \"0\"] [hostname \"http://localhost\"] [uri \"/test/artists.php\"] [unique_id \"7ce62288-d6dd-4be0-8b31-ae27876aeeea\"] [ref \"v30,53\"]"
	dictionary := Parser(payload)

	assert(t, dictionary, ParserFile, "/etc/waf/custom-REQUEST-942-APPLICATION-ATTACK-SQLI.conf")
	assert(t, dictionary, ParserLine, "45")
	assert(t, dictionary, ParserId, "942100")
	assert(t, dictionary, ParserMsg, "Warning. detected SQLi using libinjection.")
	assert(t, dictionary, ParserData, "")
	assert(t, dictionary, ParserSeverity, "0")
	assert(t, dictionary, ParserVersion, "OWASP_CRS/3.3.2")
	assert(t, dictionary, ParserHostname, "http://localhost")
	assert(t, dictionary, ParserUri, "/test/artists.php")
	assert(t, dictionary, ParserUniqueId, "7ce62288-d6dd-4be0-8b31-ae27876aeeea")
}

func TestParser_MatchedData1UE1WithinArgs(t *testing.T) {
	payload := "ModSecurity: Warning. detected SQLi using libinjection. [file \"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf\"] [line \"45\"] [id \"942100\"] [rev \"\"] [msg \"SQL Injection Attack Detected via libinjection\"] [data \"Matched Data: 1UE1 found within ARGS:artist: 0 div 1 union#foo*/*bar\\x0d\\x0aselect#foo\\x0d\\x0a1,2,current_user\"] [severity \"2\"] [ver \"OWASP_CRS/3.3.2\"] [maturity \"0\"] [accuracy \"0\"] [hostname \"echo-a\"] [uri \"/test/artists.php\"] [unique_id \"7ce62288-d6dd-4be0-8b31-ae27876aeeea\"] [ref \"v30,53\"]"
	dictionary := Parser(payload)

	assert(t, dictionary, ParserFile, "/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf")
	assert(t, dictionary, ParserLine, "45")
	assert(t, dictionary, ParserId, "942100")
	assert(t, dictionary, ParserMsg, "SQL Injection Attack Detected via libinjection")
	assert(t, dictionary, ParserData, `Matched Data: 1UE1 found within ARGS:artist: 0 div 1 union#foo*/*bar\x0d\x0aselect#foo\x0d\x0a1,2,current_user`)
	assert(t, dictionary, ParserSeverity, "2")
	assert(t, dictionary, ParserVersion, "OWASP_CRS/3.3.2")
	assert(t, dictionary, ParserHostname, "echo-a")
	assert(t, dictionary, ParserUri, "/test/artists.php")
	assert(t, dictionary, ParserUniqueId, "7ce62288-d6dd-4be0-8b31-ae27876aeeea")
}

func TestParser_MatchedOperatorRxWithParameterAgainstArgs(t *testing.T) {

	payload := "ModSecurity: Warning. Matched \"Operator `Rx' with parameter `(?:/\\*!?|\\*/|[';]--|--[\\s\\r\\n\\v\\f]|--[^-]*?-|[^&-]#.*?[\\s\\r\\n\\v\\f]|;?\\\\x00)' against variable `ARGS:artist' (Value: `0 div 1 union#foo*/*bar\\x0d\\x0aselect#foo\\x0d\\x0a1,2,current_user' ) [file \"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf\"] [line \"1188\"] [id \"942440\"] [rev \"\"] [msg \"SQL Comment Sequence Detected\"] [data \"Matched Data: n#foo*/*bar\\x0d found within ARGS:artist: 0 div 1 union#foo*/*bar\\x0d\\x0aselect#foo\\x0d\\x0a1,2,current_user\"] [severity \"2\"] [ver \"OWASP_CRS/3.3.2\"] [maturity \"0\"] [accuracy \"0\"] [tag \"application-multi\"] [tag \"language-multi\"] [tag \"platform-multi\"] [tag \"attack-sqli\"] [tag \"OWASP_CRS\"] [tag \"capec/1000/152/248/66\"] [tag \"PCI/6.5.2\"] [tag \"paranoia-level/2\"] [hostname \"echo-a\"] [uri \"/test/artists.php\"] [unique_id \"7ce62288-d6dd-4be0-8b31-ae27876aeeea\"] [ref \"o12,12o30,6v30,53t:urlDecodeUni\"]"
	dictionary := Parser(payload)

	assert(t, dictionary, ParserFile, "/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf")
	assert(t, dictionary, ParserLine, "1188")
	assert(t, dictionary, ParserId, "942440")
	assert(t, dictionary, ParserMsg, "SQL Comment Sequence Detected")
	assert(t, dictionary, ParserData, `Matched Data: n#foo*/*bar\x0d found within ARGS:artist: 0 div 1 union#foo*/*bar\x0d\x0aselect#foo\x0d\x0a1,2,current_user`)
	assert(t, dictionary, ParserSeverity, "2")
	assert(t, dictionary, ParserVersion, "OWASP_CRS/3.3.2")
	assert(t, dictionary, ParserHostname, "echo-a")
	assert(t, dictionary, ParserUri, "/test/artists.php")
	assert(t, dictionary, ParserUniqueId, "7ce62288-d6dd-4be0-8b31-ae27876aeeea")
}

func TestParser_MatchedOperatorRxWithParameterAgainstVariable(t *testing.T) {

	payload := "ModSecurity: Warning. Matched \"Operator `Rx' with parameter `(?i:(?:\\b(?:(?:s(?:elect\\b.{1,100}?\\b(?:(?:(?:length|count)\\b.{1,100}?|.*?\\bdump\\b.*)\\bfrom|to(?:p\\b.{1,100}?\\bfrom|_(?:numbe|cha)r)|(?:from\\b.{1,100}?\\bwher|data_typ)e|instr)|ys_context)|in(?:to\\b\\W* (304 characters omitted)' against variable `ARGS:artist' (Value: `0 div 1 union#foo*/*bar\\x0d\\x0aselect#foo\\x0d\\x0a1,2,current_user' ) [file \"/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf\"] [line \"1102\"] [id \"942480\"] [rev \"\"] [msg \"SQL Injection Attack\"] [data \"Matched Data: union#foo*/*bar\\x0d\\x0aselect found within ARGS:artist: 0 div 1 union#foo*/*bar\\x0d\\x0aselect#foo\\x0d\\x0a1,2,current_user\"] [severity \"2\"] [ver \"OWASP_CRS/3.3.2\"] [maturity \"0\"] [accuracy \"0\"] [tag \"application-multi\"] [tag \"language-multi\"] [tag \"platform-multi\"] [tag \"attack-sqli\"] [tag \"OWASP_CRS\"] [tag \"capec/1000/152/248/66\"] [tag \"PCI/6.5.2\"] [tag \"paranoia-level/2\"] [hostname \"echo-a\"] [uri \"/test/artists.php\"] [unique_id \"7ce62288-d6dd-4be0-8b31-ae27876aeeea\"] [ref \"o8,23v30,53t:urlDecodeUni\"]"
	dictionary := Parser(payload)

	assert(t, dictionary, ParserFile, "/etc/waf/REQUEST-942-APPLICATION-ATTACK-SQLI.conf")
	assert(t, dictionary, ParserLine, "1102")
	assert(t, dictionary, ParserId, "942480")
	assert(t, dictionary, ParserMsg, "SQL Injection Attack")
	assert(t, dictionary, ParserData, `Matched Data: union#foo*/*bar\x0d\x0aselect found within ARGS:artist: 0 div 1 union#foo*/*bar\x0d\x0aselect#foo\x0d\x0a1,2,current_user`)
	assert(t, dictionary, ParserSeverity, "2")
	assert(t, dictionary, ParserVersion, "OWASP_CRS/3.3.2")
	assert(t, dictionary, ParserHostname, "echo-a")
	assert(t, dictionary, ParserUri, "/test/artists.php")
	assert(t, dictionary, ParserUniqueId, "7ce62288-d6dd-4be0-8b31-ae27876aeeea")
}

func assert(t *testing.T, dictionary map[string]string, key, expectValue string) {
	actualValue := dictionary[key]
	if expectValue != actualValue {
		t.Errorf("Payload[\"%s\"] expect: \"%s\" actual: \"%s\"", key, expectValue, actualValue)
	}
}
