package main

import (
	log "github.com/sirupsen/logrus"
)

//var log logrus.Logger

func main() {
	log.Info("beg")

	jsonStr := []string{"1", "2", "3"}
	log.WithFields(log.Fields{
		"json": jsonStr,
	}).Info("message")

	reqMethod := "GET"
	log.WithFields(log.Fields{
		"mymethod": reqMethod,
	}).Info("Check start")
	//fmt.Println("hello")

	log.Info("end!!")
}
