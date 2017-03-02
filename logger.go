package gologging

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

const (
	debug = "debug"
	info  = "info"
	warn  = "warn"
	error = "error"
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

func SetOutput(file string) {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic("Could not be possible to create file at " + file + "\nerr: " + err.Error())
	}

	log.SetOutput(logFile)
}

func (l logger) Debug(message string, args ...interface{}) {
	l.logMessage(debug, message, args...)
}
func (l logger) Info(message string, args ...interface{}) {
	l.logMessage(info, message, args...)
}
func (l logger) Warn(message string, args ...interface{}) {
	l.logMessage(warn, message, args...)
}
func (l logger) Error(message string, args ...interface{}) {
	l.logMessage(error, message, args...)

}
func (l logger) Fatal(message string, args ...interface{}) {
	l.logMessage(fatal, message, args...)
}

func (l logger) logMessage(level string, message string, args ...interface{}) {
	logMes := fmt.Sprintf("%s||%s||%s||%s", time.Now().Format(time.RFC3339), l.owner, level, message)
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
