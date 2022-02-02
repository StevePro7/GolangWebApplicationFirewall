package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
)

const usage = `Dikastes - the decider.
Usage:
  naval_fate ship new <name>...
  naval_fate ship <name> move <x> <y> [--speed=<kn>]
  naval_fate ship shoot <x> <y>
  naval_fate mine (set|remove) <x> <y> [--moored|--drifting]
  naval_fate -h | --help
  naval_fate --version`

var VERSION string

func main() {

	VERSION = "2.0"
	arguments, _ := docopt.ParseArgs(usage, nil, VERSION)
	fmt.Println(arguments)
}
