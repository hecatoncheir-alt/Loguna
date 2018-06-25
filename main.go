package main

import (
	"encoding/json"
	"time"

	"github.com/hecatoncheir/Broker"
	"github.com/hecatoncheir/Configuration"
	"github.com/hecatoncheir/Loguna/filelog"
)

func main() {
	config := configuration.New()

	bro := broker.New(config.APIVersion, config.ServiceName)
	err := bro.Connect(config.Production.EventBus.Host, config.Production.EventBus.Port)
	if err != nil {
		panic(err.Error())
	}

	logWriter := filelog.New(config.Production.LogFilePath)
	defer logWriter.LogFile.Close()

	logStartMessage := filelog.LogData{
		Time:    time.Now(),
		Message: "Prepare log session"}

	err = logWriter.Write(logStartMessage)
	if err != nil {
		panic(err.Error())
	}

	for eventData := range bro.InputChannel {
		if err != nil {
			println(err.Error())
		}

		if eventData.APIVersion == config.APIVersion {
			logData := filelog.LogData{}
			err = json.Unmarshal([]byte(eventData.Data), &logData)
			if err != nil {
				println(err.Error())
			}

			err = logWriter.Write(logData)
			if err != nil {
				println(err.Error())
			}
		}
	}
}
