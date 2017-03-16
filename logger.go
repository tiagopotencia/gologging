package gologging

import (
	"fmt"
	"log"
	"os"
	"time"
	"sync"
)

type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

func init() {
	go func() {
		for {
			today := time.Now()
			to := today.Add(24 * time.Hour)
			to = time.Date(to.Year(), to.Month(), to.Day(), 0, 0, 0, 0, to.Location())
			diff := to.Sub(today) + (time.Duration(50) * time.Millisecond)
			time.Sleep(diff)
			SetOutputPerTime("")
		}

	}()
}

const (
	debug = "debug"
	info  = "info"
	warn  = "warn"
	Error = "error"
	fatal = "fatal"
)

type logger struct {
	owner string
}

func New(owner string) Logger {
	l := new(logger)
	l.owner = owner
	return l
}

func SetOutput(file string) error {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	log.SetOutput(logFile)

	return nil
}

var filename string
var lock = new(sync.RWMutex)

func SetOutputPerTime(fileNameWithoutExtension string) error {
	lock.Lock()
	defer lock.Unlock()
	if filename == "" {
		filename = fileNameWithoutExtension
	}
	now := time.Now()

	logName := filename + now.Format("2006-01-02") + ".log"

	logFile, err := os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	log.SetOutput(logFile)

	return nil
}

func (l logger) Debug(message string, args ...interface{}) {
	go l.logMessage(debug, message, time.Now(), args...)
	time.Sleep(5 * time.Microsecond)
}
func (l logger) Info(message string, args ...interface{}) {
	go l.logMessage(info, message, time.Now(), args...)
	time.Sleep(5 * time.Microsecond)
}
func (l logger) Warn(message string, args ...interface{}) {
	go l.logMessage(warn, message, time.Now(), args...)
	time.Sleep(5 * time.Microsecond)
}
func (l logger) Error(message string, args ...interface{}) {
	go l.logMessage(Error, message, time.Now(), args...)
	time.Sleep(5 * time.Microsecond)
}
func (l logger) Fatal(message string, args ...interface{}) {
	go l.logMessage(fatal, message, time.Now(), args...)
	time.Sleep(5 * time.Microsecond)
}

func (l logger) logMessage(level string, message string, t time.Time, args ...interface{}) {
	lock.Lock()
	defer lock.Unlock()
	logMes := fmt.Sprintf("%s||%s||%s||%s", t.Format(time.RFC3339), l.owner, level, message)
	if len(args) > 0 {
		logMes += "||["
		for i, arg := range args {
			logMes += fmt.Sprint(arg)
			if i < (len(args) - 1) {
				logMes += ","
			}
		}
		logMes += "]"
	}
	log.Println(logMes)
}
