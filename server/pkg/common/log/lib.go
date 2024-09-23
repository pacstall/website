package log

import (
	"fmt"
	glog "log"
	"os"

	"github.com/fatih/color"
)

var (
	logInfo  = color.CyanString("INFO")
	logError = color.RedString("ERROR")
	logDebug = color.GreenString("DEBUG")
	logFatal = color.New(color.BgHiRed, color.FgBlack).Sprintf("FATAL")
	logWarn  = color.YellowString("WARN")
)

type tLogLevel uint8

var Level = struct {
	Info  tLogLevel
	Error tLogLevel
	Debug tLogLevel
	Fatal tLogLevel
	Warn  tLogLevel
}{
	Debug: 0,
	Info:  1,
	Warn:  2,
	Error: 3,
	Fatal: 4,
}

var logLevel = Level.Debug

func SetLogLevel(level tLogLevel) {
	logLevel = level
}

var logger = glog.New(os.Stdout, "", glog.Ldate|glog.Ltime)

func doLog(level, message string, args ...any) {
	msg := fmt.Sprintf("%s: %s\n", level, fmt.Sprintf(message, args...))

	if level == logFatal {
		logger.Fatal(msg)
	} else {
		logger.Print(msg)
	}
}

func Info(message string, args ...any) {
	if logLevel > Level.Info {
		return
	}

	doLog(logInfo, message, args...)
}

func Error(message string, args ...any) {
	if logLevel > Level.Error {
		return
	}

	doLog(logError, message, args...)
}

func Fatal(message string, args ...any) {
	if logLevel > Level.Fatal {
		return
	}

	doLog(logFatal, message, args...)
}

func Warn(message string, args ...any) {
	if logLevel > Level.Warn {
		return
	}

	doLog(logWarn, message, args...)
}

func Debug(message string, args ...any) {
	if logLevel > Level.Debug {
		return
	}

	doLog(logDebug, message, args...)
}
