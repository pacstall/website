package log

import (
	"io/ioutil"
	"os"
	"reflect"
)

type logLevels struct {
	Debug Level
	Info  Level
	Warn  Level
	Error Level
}

type Level uint8

func NewLogLevel(level string) Level {
	var strToLogLevel = map[string]Level{
		"debug": LogLevels.Debug,
		"info":  LogLevels.Info,
		"warn":  LogLevels.Warn,
		"error": LogLevels.Error,
	}

	if level, ok := strToLogLevel[level]; ok {
		return level
	}

	Error.Fatalf("Invalid log level: '%s'. Accepted values are %#v.\n", level, reflect.ValueOf(strToLogLevel).MapKeys())
	return LogLevels.Debug
}

var LogLevels = logLevels{
	Debug: Level(0),
	Info:  Level(1),
	Warn:  Level(2),
	Error: Level(3),
}

func setLogLevel(level Level) {
	if level == LogLevels.Debug {
		Debug.SetOutput(os.Stdout)
		Info.SetOutput(os.Stdout)
		Warn.SetOutput(os.Stdout)
		Error.SetOutput(os.Stdout)
	} else if level == LogLevels.Info {
		Debug.SetOutput(ioutil.Discard)
		Info.SetOutput(os.Stdout)
		Warn.SetOutput(os.Stdout)
		Error.SetOutput(os.Stdout)
	} else if level == LogLevels.Warn {
		Debug.SetOutput(ioutil.Discard)
		Info.SetOutput(ioutil.Discard)
		Warn.SetOutput(os.Stdout)
		Error.SetOutput(os.Stdout)
	} else if level == LogLevels.Error {
		Debug.SetOutput(ioutil.Discard)
		Info.SetOutput(ioutil.Discard)
		Warn.SetOutput(ioutil.Discard)
		Error.SetOutput(os.Stdout)
	}
}
