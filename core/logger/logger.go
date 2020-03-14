package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"snappy-api/core/config"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	Fatal = errorType(iota)
	Info
	Debug
	Error
)

type handleError struct {
	ErrorType                errorType
	OriginalError            error
	LogFilePath, LogFileName string
	LogFileSize              int64
}

var (
	instance *handleError
	padLock  sync.Mutex
)

type errorType uint

func initLogger(filePath string, fileName string, logFileSize int64) *handleError {
	h := &handleError{}
	h.LogFilePath = filePath
	h.LogFileName = fileName
	h.LogFileSize = logFileSize
	return h
}

func GetLogInstance(filePath string, fileName string) *handleError {
	padLock.Lock()
	cfg := config.LogConfigs()

	if fileName == "" {
		fileName = cfg[config.LOG_FILE_NAME]
	}
	if filePath == "" {
		filePath = cfg[config.LOG_FILE_PATH]
	}
	if instance == nil {
		fileSize, _ := strconv.ParseInt(cfg[config.LOG_FILE_SIZE], 10, 32)
		instance = initLogger(filePath, fileName, fileSize)
	}
	padLock.Unlock()

	return instance
}

func (h *handleError) writeErrorLog(err error, errorType errorType) {
	currentTime := time.Now()
	str := []byte(getErrorType(errorType) + "" + currentTime.Format("2006.01.02 15:04:05") + " -- " + err.Error() + "\n")
	writeLog(h, str)
}

func (h *handleError) writeTextLog(text string, errorType errorType) {
	currentTime := time.Now()
	str := []byte(getErrorType(errorType) + "" + currentTime.Format("2006.01.02 15:04:05") + " -- " + text + "\n")
	writeLog(h, str)
}

func writeLog(h *handleError, str []byte) {

	if er := os.MkdirAll(h.LogFilePath, 0777); er != nil {
		fmt.Println(er)
	}
	logPath := h.LogFilePath + h.createLogFile() + ".txt"

	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, e := f.Write(str); e != nil {
		fmt.Println(e)
	}

}

func getErrorType(errorType errorType) string {
	if errorType == Fatal {
		return "FATAL : "
	} else if errorType == Info {
		return "INFO : "
	} else if errorType == Debug {
		return "DEBUG : "
	}
	return "INFO : "
}

func (h *handleError) Error(err error) {
	h.writeErrorLog(err, Error)
}

func (h *handleError) Fatal(err error) {
	h.writeErrorLog(err, Fatal)
}

func (h *handleError) Debug(err error) {
	h.writeErrorLog(err, Debug)
}

func (h *handleError) Info(err error) {
	h.writeErrorLog(err, Info)
}

func (h *handleError) TextError(text string) {
	h.writeTextLog(text, Error)
}

func (h *handleError) TextFatal(text string) {
	h.writeTextLog(text, Fatal)
}

func (h *handleError) TextDebug(text string) {
	h.writeTextLog(text, Debug)
}

func (h *handleError) TextInfo(text string) {
	h.writeTextLog(text, Info)
}

func (h *handleError) createLogFile() string {
	files, err := ioutil.ReadDir(h.LogFilePath)
	if err != nil {
		fmt.Println(err)
	}
	tmpLogFile := ""
	for _, f := range files {
		if f.Size() < 10000000 {
			tmpLogFile = strings.Replace(f.Name(), ".txt", "", -1)
		}
	}

	if tmpLogFile == "" {
		tmpLogFile = h.LogFileName
		if len(files) != 0 {
			tmpLogFile = tmpLogFile + strconv.Itoa(len(files))
		}
	}
	return tmpLogFile
}
