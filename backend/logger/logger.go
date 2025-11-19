// Package logger represents a generic logging interface

package logger

import "os"

// Log is a package level variable, every program should access logging function through "Log"
var Log Logger
var logFileH *os.File

// Logger represent common interface for logging functions
type Logger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	Tracef(format string, args ...interface{})
	Trace(args ...interface{})
}

// SetLogger is the setter for log variable, it should be the only way to assign value to log.
//
// Parameters:
// - newLogger - Logger instance
//
// Returns:
// - N/A
func SetLogger(newLogger Logger) {
	Log = newLogger
}

// OpenFile opnes the log file for writting.
//
// Parameters:
// - fileName - string representing the log file name.
//
// Returns:
// - A pointer to log file.
// - An error if opening log file fails.
func OpenFile(fileName string) (*os.File, error) {
	logFileH, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return logFileH, err
}

// Close closes the log file.
//
// Parameters:
// - N/A
//
// Returns:
// - N/A
func Close() {
	if logFileH != nil {
		logFileH.Close()
	}
}
