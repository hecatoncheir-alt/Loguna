package main

import (
	"encoding/json"
	"github.com/hecatoncheir/Loguna/broker"
	"github.com/hecatoncheir/Loguna/configuration"
	"github.com/hecatoncheir/Loguna/logger"
	"time"
)

func main() {
	config := configuration.New()

	logWriter := logger.New(config.Production.LogFilePath)
	defer logWriter.LogFile.Close()

	logStartMessage := logger.LogData{
		Time:    time.Now(),
		Message: "Prepare log session"}

	err := logWriter.Write(logStartMessage)
	if err != nil {
		panic(err.Error())
	}

	bro := broker.New()
	err = bro.Connect(config.Production.Broker.Host, config.Production.Broker.Port)
	if err != nil {
		panic(err.Error())
	}
	topicEvents, err := bro.ListenTopic(config.Production.LogunaTopic, config.APIVersion)
	if err != nil {
		panic(err.Error())
	}

	for event := range topicEvents {
		data := logger.LogData{}
		err = json.Unmarshal(event, &data)
		if err != nil {
			println(err.Error())
		}

		if data.ApiVersion == config.APIVersion {
			err = logWriter.Write(data)
			if err != nil {
				println(err.Error())
			}
		}
	}
}
