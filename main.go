package main

import (
	"encoding/json"
	"time"

	"github.com/hecatoncheir/Broker"
	"github.com/hecatoncheir/Configuration"
	"github.com/hecatoncheir/Loguna/logToFileWriter"
)

func main() {
	config := configuration.New()

	bro := broker.New(config.APIVersion, config.ServiceName)
	err := bro.Connect(config.Production.Broker.Host, config.Production.Broker.Port)
	if err != nil {
		panic(err.Error())
	}

	logWriter := logToFileWriter.New(config.Production.LogFilePath)
	defer logWriter.LogFile.Close()

	logStartMessage := logToFileWriter.LogData{
		Time:    time.Now(),
		Message: "Prepare log session"}

	err = logWriter.Write(logStartMessage)
	if err != nil {
		panic(err.Error())
	}

	topicEvents, err := bro.ListenTopic(config.Production.LogunaTopic, config.APIVersion)
	if err != nil {
		panic(err.Error())
	}

	for event := range topicEvents {
		eventData := broker.EventData{}
		err = json.Unmarshal(event, &eventData)
		if err != nil {
			println(err.Error())
		}

		if eventData.APIVersion == config.APIVersion {
			logData := logToFileWriter.LogData{}
			json.Unmarshal([]byte(eventData.Data), &logData)

			err = logWriter.Write(logData)
			if err != nil {
				println(err.Error())
			}
		}
	}
}
