package filelog

import (
	"os"
	"testing"
	"time"

	"github.com/hecatoncheir/Configuration"
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

	err = os.Remove(conf.Development.LogFilePath)
	if err != nil {
		test.Error(err)
	}

}
