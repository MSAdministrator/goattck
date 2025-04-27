package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type LogLevel int

const (
	Info LogLevel = iota
	Debug
	Warning
	Error
	Fatal
)

var levelNames = map[LogLevel]string{
	Info:    "INFO",
	Debug:   "DEBUG",
	Warning: "WARN",
	Error:   "ERROR",
	Fatal:   "FATAL",
}

// ParseLogLevel converts a string to a LogLevel.
func ParseLogLevel(level string) LogLevel {
	level = strings.ToLower(strings.TrimSpace(level))
	switch level {
	case "info":
		return Info
	case "debug":
		return Debug
	case "warn", "warning":
		return Warning
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Info // Default to Info if the input is invalid.
	}
}

// Logger struct with configurable output destination.
type Logger struct {
	level  LogLevel
	logger *log.Logger
}

// NewLogger initializes a logger
func NewLogger(level LogLevel) *Logger {
	flags := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stdout, "", flags)

	return &Logger{level: level, logger: logger}
}

// formatMessage creates a log message with a timestamp and level.
func (l *Logger) formatMessage(level LogLevel, message string, a ...any) string {
	timestamp := time.Now().Format("2006-01-02T15:04:05-0700")
	prefix := fmt.Sprintf("[%s] %s: ", levelNames[level], timestamp)
	return prefix + fmt.Sprintf(message, a...)
}

// logMessage writes a formatted message to the log.
func (l *Logger) logMessage(level LogLevel, message string, a ...any) {
	if level >= l.level {
		l.logger.Output(3, l.formatMessage(level, message, a...))
	}
}

// Logging methods.
func (l *Logger) Info(message string, a ...any)    { l.logMessage(Info, message, a...) }
func (l *Logger) Debug(message string, a ...any)   { l.logMessage(Debug, message, a...) }
func (l *Logger) Warning(message string, a ...any) { l.logMessage(Warning, message, a...) }
func (l *Logger) Error(message string, a ...any)   { l.logMessage(Error, message, a...) }
func (l *Logger) Fatal(message string, a ...any) {
	l.logMessage(Fatal, message, a...)
	os.Exit(1)
}
