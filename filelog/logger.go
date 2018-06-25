package filelog

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type LogData struct {
	APIVersion, Message, Service string
	Time                         time.Time
}

type LogWriter struct {
	PathOfLogFile string
	LogFile       *os.File
}

func New(pathOfLogFile string) *LogWriter {
	logger := LogWriter{PathOfLogFile: pathOfLogFile}
	return &logger
}

var (
	ErrLogFileCanNotBeCreate = errors.New("log file can't be create")
	ErrLogFileCanNotBeOpen   = errors.New("log file can't be open")
	ErrLogDataWithoutTime    = errors.New("log data without time")
	ErrLogDataCanNotBeWrite  = errors.New("log data can't be write")
	ErrLogDataCanNotBeSave   = errors.New("log data can't be save")
)

func (logger *LogWriter) Write(logData LogData) error {
	if logData.Time.IsZero() {
		return ErrLogDataWithoutTime
	}

	fileInfo, err := os.Stat(logger.PathOfLogFile)

	if os.IsNotExist(err) {
		logger.LogFile, err = os.Create(logger.PathOfLogFile)
		if err != nil {
			return ErrLogFileCanNotBeCreate
		}
	} else {
		logger.LogFile, err = os.OpenFile(logger.PathOfLogFile, os.O_RDWR, 6044)
		if err != nil {
			return ErrLogFileCanNotBeOpen
		}
	}

	logTimeFormat := time.RFC1123

	logForWrite := fmt.Sprintf(
		"%v - %v\n", logData.Time.Format(logTimeFormat), logData.Message)

	if fileInfo != nil {
		last := fileInfo.Size()
		_, err = logger.LogFile.WriteAt([]byte(logForWrite), last)
		if err != nil {
			return ErrLogDataCanNotBeWrite
		}
	} else {
		_, err = logger.LogFile.Write([]byte(logForWrite))
		if err != nil {
			return ErrLogDataCanNotBeWrite
		}
	}

	err = logger.LogFile.Close()
	if err != nil {
		return ErrLogDataCanNotBeSave
	}

	return nil
}
