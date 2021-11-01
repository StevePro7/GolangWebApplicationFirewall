package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("start")

	//root := "/etc/config"
	root := "./config/"
	items, _ := ioutil.ReadDir(root)
	for _, item := range items {
		if item.IsDir() {
			subFolder := root + item.Name() + "/"
			subitems, _ := ioutil.ReadDir(subFolder)
			for _, subitem := range subitems {
				if !subitem.IsDir() {
					// handle file there
					subitemName := subFolder + subitem.Name()
					//fmt.Println(item.Name() + "/" + subitemName)
					fmt.Println(subitemName)
				}
			}
		} else {
			// handle file there
			//fileName := item.Name()
			fileName := root + item.Name()
			fmt.Println(fileName)
		}
	}

	fmt.Println("-end-")
}
