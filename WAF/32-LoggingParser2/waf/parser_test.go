package waf

import (
	"fmt"
	"testing"
)

func TestParser_DetectedSQLiUsingLibinjection(t *testing.T) {
	payload := "ModSecurity: Warning. detected SQLi using libinjection. [file \"/etc/waf_custom/custom-REQUEST-942-APPLICATION-ATTACK-SQLI.conf\"] [line \"45\"] [id \"942100\"] [rev \"\"] [msg \"\"] [data \"\"] [severity \"0\"] [ver \"OWASP_CRS/3.3.2\"] [maturity \"0\"] [accuracy \"0\"] [hostname \"http://localhost\"] [uri \"/test/artists.php\"] [unique_id \"7ce62288-d6dd-4be0-8b31-ae27876aeeea\"] [ref \"v30,53\"]"
	dictionary := Parser(payload)

	assertion(t, "file", "/etc/waf_custom/custom-REQUEST-942-APPLICATION-ATTACK-SQLI.conf", dictionary)

	fmt.Println(dictionary)
}

func TestParser_MatchedData1UE1WithinArgs(t *testing.T) {
}

func TestParser_MatchedOperatorRxWithParameter(t *testing.T) {

}

func assertion(t *testing.T, key, expectValue string, dictionary map[string]string) {
	actualValue := dictionary[key]
	if expectValue != actualValue {
		t.Errorf("Payload[\"%s\"] expect: \"%s\" actual: \"%s\"", key, expectValue, actualValue)
	}
}
