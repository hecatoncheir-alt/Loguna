package main

import (
	"github.com/hecatoncheir/Loguna/broker"
	"github.com/hecatoncheir/Loguna/configuration"
	"github.com/hecatoncheir/Loguna/logger"
	"os"
	"testing"
	"time"
	"encoding/json"
)

func TestCanWriteLogsDataToFile(test *testing.T) {
	conf := configuration.New()

	bro := broker.New()

	bro.ListenTopic(conf.Development.LogunaTopic, conf.APIVersion)

	_, err := bro.ListenTopic(
		conf.Development.LogunaTopic, conf.APIVersion)
	if err != nil {
		test.Fatal(err)
	}

	logWriter := logger.New(conf.Development.LogFilePath)

	go SubscribeLoggerToChannelOfTopic(logWriter, conf.Development.LogunaTopic, conf.APIVersion)

	logData := logger.LogData{ApiVersion: "1.0.0", Message: "test"}
	message, err:= json.Marshal(logData)
	if err != nil {
		test.Fatal(err)
	}

	err = bro.WriteToTopic(conf.Development.LogunaTopic, string(message))
	if err != nil {
		test.Fatal(err)
	}

	time.Sleep(time.Second * 2)

	_, err = os.Stat(conf.Development.LogFilePath)

	if os.IsNotExist(err) {
		test.Fatal()
	}
}
