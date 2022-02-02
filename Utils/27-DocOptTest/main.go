package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"waftesting/waf"
)

const usage = `Dikastes - the decider.
Usage:
  dikastes server [options]
  dikastes client <namespace> <account> [--method <method>] [options]
Options:
  <namespace>            Service account namespace.
  <account>              Service account name.
  -h --help              Show this screen.
  -l --listen <port>     Unix domain socket path [default: /var/run/dikastes/dikastes.sock]
  -d --dial <target>     Target to dial. [default: localhost:50051]
  -r --rules <target>    Directory where WAF rules are stored.
  --debug                Log at Debug level.`

var VERSION string

func main() {

	VERSION = "1.1"
	arguments, _ := docopt.ParseArgs(usage, nil, VERSION)
	fmt.Println("stevepro beg")
	fmt.Println(arguments)

	if arguments["server"].(bool) {
		fmt.Println("server now")
	}

	rulesetArgument := arguments["--rules"]

	// Check if WAF should be enabled first before proceeding...
	err := waf.CheckRulesSetExists(rulesetArgument)
	if err != nil {
		//log.Errorf("WAF Core Rules Set check: '%s'", err.Error())
		fmt.Println("bob")
	}

	//test := arguments["--rules"]
	//if test == nil {
	//	fmt.Println("NO rules test!")
	//}
	//if test != nil {
	//	fmt.Println("testing OK..!!")
	//
	//	rulesetDirectory := test.(string)
	//	if len(rulesetDirectory) > 0 {
	//		fmt.Printf("rules directory!! : '%v'\n", rulesetDirectory)
	//	}
	//}

	//dial := arguments["--dial"].(string)
	//fmt.Println(dial)

	fmt.Println("stevepro end")
}
