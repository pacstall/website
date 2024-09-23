package log

import (
	"fmt"
	glog "log"
	"os"
	"strings"

	"github.com/fatih/color"
	"pacstall.dev/webserver/pkg/common/config"
)

var (
	logInfo  = color.CyanString("INFO")
	logError = color.RedString("ERROR")
	logDebug = color.GreenString("DEBUG")
	logTrace = color.HiGreenString("TRACE")
	logFatal = color.New(color.BgHiRed, color.FgBlack).Sprintf("FATAL")
	logWarn  = color.YellowString("WARN")
)

type tLogLevel uint8

const (
	_LEVEL_TRACE tLogLevel = iota
	_LEVEL_DEBUG
	_LEVEL_INFO
	_LEVEL_WARN
	_LEVEL_ERROR
	_LEVEL_FATAL
)

var logLevels = map[string]tLogLevel{
	"TRACE": _LEVEL_TRACE,
	"DEBUG": _LEVEL_DEBUG,
	"INFO":  _LEVEL_INFO,
	"WARN":  _LEVEL_WARN,
	"ERROR": _LEVEL_ERROR,
	"FATAL": _LEVEL_FATAL,
}

func getLogLevel() tLogLevel {
	l, ok := logLevels[strings.ToUpper(config.LogLevel)]
	if !ok {
		doLog(logWarn, "unknown log level '%s'. defaulting to INFO", config.LogLevel)
		return _LEVEL_INFO
	}

	return l
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
	if getLogLevel() > _LEVEL_INFO {
		return
	}

	doLog(logInfo, message, args...)
}

func Error(message string, args ...any) {
	if getLogLevel() > _LEVEL_ERROR {
		return
	}

	doLog(logError, message, args...)
}

func Fatal(message string, args ...any) {
	if getLogLevel() > _LEVEL_FATAL {
		return
	}

	doLog(logFatal, message, args...)
}

func Warn(message string, args ...any) {
	if getLogLevel() > _LEVEL_WARN {
		return
	}

	doLog(logWarn, message, args...)
}

func Debug(message string, args ...any) {
	if getLogLevel() > _LEVEL_DEBUG {
		return
	}

	doLog(logDebug, message, args...)
}

func Trace(message string, args ...any) {
	if getLogLevel() > _LEVEL_TRACE {
		return
	}

	doLog(logTrace, message, args...)
}
