package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/hecatoncheir/Broker"
	"github.com/hecatoncheir/Configuration"
	"github.com/hecatoncheir/Logger"
	"github.com/hecatoncheir/Loguna/logToFileWriter"
	"io/ioutil"
)

func TestCanWriteLogsDataToFile(test *testing.T) {
	config := configuration.New()

	bro := broker.New(config.APIVersion, config.ServiceName)
	err := bro.Connect(config.Development.Broker.Host, config.Development.Broker.Port)
	if err != nil {
		test.Fatal(err)
	}

	logWriter := logToFileWriter.New(config.Development.LogFilePath)

	go func() {
		logStartMessage := logToFileWriter.LogData{
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
			data := logToFileWriter.LogData{}
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

	logData := logger.LogData{Message: "test log data", Time: time.Now().UTC()}
	encodedLogData, err := json.Marshal(logData)
	if err != nil {
		test.Error(err.Error())
	}

	brokerEventData := broker.EventData{Message: "test", Data: string(encodedLogData)}

	go bro.WriteToTopic(config.Development.LogunaTopic, brokerEventData)

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
