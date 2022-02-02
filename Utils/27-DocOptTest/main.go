package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
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
  -d --dial <target>     Target to dial. [default: localhost:50052] 
  --debug                Log at Debug level.`

var VERSION string

func main() {

	VERSION = "2.0"
	arguments, _ := docopt.ParseArgs(usage, nil, VERSION)
	fmt.Println("stevepro beg")
	fmt.Println(arguments)

	if arguments["server"].(bool) {
		fmt.Println("server now")
	}

	//rulesetDirectory := arguments["--rules"].(string)
	//fmt.Printf("rules directory : '%v'\n", rulesetDirectory)

	if arguments["--debug"].(bool) {
		fmt.Println("debugging...!!")
	}

	//dial := arguments["--dial"].(string)
	//fmt.Println(dial)

	fmt.Println("stevepro end")
}
