package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

//var log logrus.Logger

func main() {
	log.Info("beg")

	jsonStr := []string{"1", "2", "3"}
	jsonObj, _ := json.Marshal(jsonStr)
	//fmt.Println(string(jsonObj))

	log.WithFields(log.Fields{
		"json": &jsonObj,
	}).Info("message")

	reqMethod := "GET"
	log.WithFields(log.Fields{
		"mymethod": reqMethod,
	}).Info("Check start")

	//fmt.Println("hello")
	log.Info("end")
}
