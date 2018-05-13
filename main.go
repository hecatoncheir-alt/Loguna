package main

import (
	"fmt"
	"github.com/hecatoncheir/loguna/broker"
	"github.com/hecatoncheir/loguna/configuration"
	"github.com/hecatoncheir/loguna/logger"
	"time"
)

func main() {
	logFilePath := "log"
	logWriter := logger.New(logFilePath)
	defer logWriter.LogFile.Close()

	logStartMessage := logger.LogData{
		Time:    time.Now(),
		Message: "Prepare log session"}

	err := logWriter.Write(logStartMessage)
	if err != nil {
		panic(err)
	}

	config := configuration.New()

	bro := broker.New()
	topicEvents, err := bro.ListenTopic(
		config.Production.LogunaTopic, config.APIVersion)
	if err != nil {
		panic(err)
	}

	for event := range topicEvents {
		/// if api version is actual
		fmt.Println(string(event))

	}
}
