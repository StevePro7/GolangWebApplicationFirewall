package main

import "fmt"

func main() {
	for i := 17; i <= 21; i++ {
		go func() {
			go func(i int) {
				apiVersion := fmt.Sprintf("v1.%d", i)
				fmt.Println(apiVersion)
			}(i)
		}()
	}
	fmt.Println("hello")
}
