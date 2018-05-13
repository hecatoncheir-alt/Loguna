package logger

import (
	"github.com/hecatoncheir/loguna/configuration"
	"os"
	"testing"
	"time"
)

func TestLoggerCanCheckTimeOfLogData(test *testing.T) {
	conf := configuration.New()

	logWriter := New(conf.Development.LogFilePath)
	logData := LogData{Message: "test message"}
	err := logWriter.Write(logData)
	if err != ErrLogDataWithoutTime {
		test.Fatal(err)
	}
}

func TestLoggerCanWriteLogData(test *testing.T) {
	conf := configuration.New()

	logWriter := New(conf.Development.LogFilePath)
	logData := LogData{Message: "test message", Time: time.Now().UTC()}
	err := logWriter.Write(logData)
	if err != nil {
		test.Fatal(err)
	}

	err = logWriter.LogFile.Close()
	if err != nil {
		test.Error(err)
	}

	err = os.Remove(conf.Development.LogFilePath)
	if err != nil {
		test.Error(err)
	}

}
