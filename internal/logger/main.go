package logger

import (
	"io"
	"log"
	"os"
	"time"
)

const (
	Info    string = "info"
	Debug   string = "debug"
	Warning string = "warning"
	Error   string = "error"
	Fatal   string = "fatal"
)

type Logger struct {
	writeToFile bool
	level       string
	logger      *log.Logger
}

func NewLogger(level string, writeToFile bool) *Logger {
	flags := log.LstdFlags | log.Lshortfile
	logger := &Logger{level: level, writeToFile: writeToFile}
	logger.logger = log.New(os.Stdout, "", flags)
	if writeToFile {
		logger.setWriteToFile()
	}
	return logger
}

func (l *Logger) getTimestamp() string {
	return time.Now().Format("2006-01-02T15:04:05-0700")
}

func (l *Logger) Info(message string) {
	if l.level == Info {
		l.logger.SetPrefix("INFO : " + l.getTimestamp())
		l.logger.Println(message)
	}
}

func (l *Logger) Debug(message string) {
	if l.level == Debug {
		l.logger.SetPrefix("DEBUG : " + l.getTimestamp())
		l.logger.Println(message)
	}
}

func (l *Logger) Warning(message string) {
	if l.level == Warning {
		l.logger.SetPrefix("WARN : " + l.getTimestamp())
		l.logger.Println(message)
	}
}

func (l *Logger) Error(message string) {
	if l.level == Error {
		l.logger.SetPrefix("ERROR : " + l.getTimestamp())
		l.logger.Println(message)
	}
}

func (l *Logger) Fatal(message string) {
	if l.level == Fatal {
		l.logger.SetPrefix("FATAL : " + l.getTimestamp())
		l.logger.Fatal(message)
	}
}

func (l *Logger) setWriteToFile() {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		l.Error("Unable to open log file")
	}
	defer file.Close()

	mw := io.MultiWriter(os.Stdout, file)
	l.logger.SetOutput(mw)
}
