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
		panic(err)
	}

	err = SubscribeLoggerToChannelOfTopic(logWriter,
		config.Production.LogunaTopic,
		config.APIVersion)

	if err != nil {
		panic(err)
	}
}

func SubscribeLoggerToChannelOfTopic(logWriter *logger.LogWriter,
	topic, channel string) error {

	bro := broker.New()
	topicEvents, err := bro.ListenTopic(topic, channel)
	if err != nil {
		return err
	}

	for event := range topicEvents {
		data := logger.LogData{}
		err = json.Unmarshal(event, &data)
		if err != nil {
			println(err)
		}

		if data.ApiVersion == channel {
			err = logWriter.Write(data)
			if err != nil {
				println(err)
			}
		}
	}

	return nil
}
