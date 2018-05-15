package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/hecatoncheir/Loguna/broker"
	"github.com/hecatoncheir/Loguna/configuration"
	"github.com/hecatoncheir/Loguna/logger"
	"io/ioutil"
)

func TestCanWriteLogsDataToFile(test *testing.T) {
	config := configuration.New()

	bro := broker.New()
	err := bro.Connect(config.Development.Broker.Host, config.Development.Broker.Port)
	if err != nil {
		test.Fatal(err)
	}

	logWriter := logger.New(config.Development.LogFilePath)

	go func() {
		logStartMessage := logger.LogData{
			Time:    time.Now().UTC(),
			Message: "Prepare log session"}

		err := logWriter.Write(logStartMessage)
		if err != nil {
			test.Fatal(err.Error())
		}

		topicEvents, err := bro.ListenTopic(config.Development.LogunaTopic, config.APIVersion)
		if err != nil {
			test.Fatal(err.Error())
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

			break
		}
	}()

	logData := logger.LogData{ApiVersion: "1.0.0", Message: "test", Time: time.Now().UTC()}

	go bro.WriteToTopic(config.Development.LogunaTopic, logData)

	time.Sleep(time.Second * 1)

	_, err = os.Stat(config.Development.LogFilePath)

	if os.IsNotExist(err) {
		test.Fatal(err)
	}

	logFile, err := ioutil.ReadFile(config.Development.LogFilePath)
	if err != nil {
		test.Error(err.Error())
	}

	if logFile == nil {
		test.Fail()
	}

	err = os.Remove(config.Development.LogFilePath)
	if err != nil {
		test.Error(err.Error())
	}
}
