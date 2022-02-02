package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"log"
)

const usage = `RandX
Returns random numbers and/or letters. 
Usage:
  randx  numbers  --count <count> [--range <range>...] [--verbose]
  randx  letters  --count <count> [--lang <lang>] [--verbose]
  randx -h | --help
  randx --version
Required options:
  -c --count <count>	    A count of random numbers or letters. 
Options:
  -l --lang <lang>      A language. Optional. Default: en.
  -r --range <range>    A range. Optional. Default: 0-100.
  -h --help             Shows this screen.
  --verbose             Shows details.
  --version             Shows version.`

var VERSION string

func main() {

	VERSION = "3.0"
	arguments, _ := docopt.ParseArgs(usage, nil, VERSION)
	fmt.Println("stevepro beg")
	fmt.Println(arguments)

	//if arguments["server"].(bool) {
	//	fmt.Println("server now")
	//}
	//
	////rulesetDirectory := arguments["--rules"].(string)
	////fmt.Printf("rules directory : '%v'\n", rulesetDirectory)
	//
	//if arguments["--debug"].(bool) {
	//	fmt.Println("debugging...!!")
	//}

	test := arguments["--lang"]
	if test == nil {
		fmt.Println("NO test!")
	}
	if test != nil {
		fmt.Println("testing OK..!!")

		rulesetDirectory := arguments["--lang"].(string)
		if len(rulesetDirectory) > 0 {
			fmt.Printf("rules directory!! : '%v'\n", rulesetDirectory)
		}
	}

	var config struct {
		Numbers bool     `docopt:"numbers"`
		Letters bool     `docopt:"letters"`
		Count   int      `docopt:"-c,--count"`
		Lang    string   `docopt:"-l,--lang"`
		Range   []string `docopt:"-r,--range"`
		Verbose bool     `docopt:"--verbose"`
	}

	err := arguments.Bind(&config)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("--lang:", config.Lang)
	//dial := arguments["--dial"].(string)
	//fmt.Println(dial)

	fmt.Println("stevepro end")
}
