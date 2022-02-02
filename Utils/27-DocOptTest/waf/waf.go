package waf

import "fmt"

func CheckRulesSetExists(rulesetArgument interface{}) error {

	//test := arguments["--rules"]
	if rulesetArgument == nil {
		fmt.Println("NO args test!")
	}
	if rulesetArgument != nil {
		fmt.Println("testing OK..!!")

		rulesetDirectory := rulesetArgument.(string)
		if len(rulesetDirectory) > 0 {
			fmt.Printf("rules directory!! : '%v'\n", rulesetDirectory)
		}
	}

	return nil
}
