package internal

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ToInt(str string) int {
	die := func(err error, message string, args ...any) {
		if err != nil {
			fmt.Printf(message, args...)
			panic(err)
		}
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		die(err, "could not convert '%s' to int\n", str)
	}

	return num
}

func ToBool(str string) bool {
	return str == "true" || str == "1"
}

type formatter[T any] interface {
	Format(string) T
}

type stringFormatter struct{}

func (f stringFormatter) Format(str string) string {
	return str
}

type intFormatter struct{}

func (f intFormatter) Format(str string) int {
	return ToInt(str)
}

type boolFormatter struct{}

func (f boolFormatter) Format(str string) bool {
	return ToBool(str)
}

func censorWhenPrivate[T any](key string, value T) string {
	if strings.Contains(key, "SECRET") || strings.Contains(key, "PRIVATE") || strings.Contains(key, "PASSWORD") {
		return "**redacted**"
	}

	return fmt.Sprintf("%#v", value)
}

func GetEnvVar[T any](key string, format formatter[T]) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("could not find environment variable '%s'", key))
	}

	out := format.Format(val)

	log.Default().Printf("using %v = %v [from env]\n", key, censorWhenPrivate(key, out))
	return out
}

func GetEnvOrDefault[T any](key string, defaultValue T, format formatter[T]) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Default().Printf("using %v = %v [default]\n", key, censorWhenPrivate(key, defaultValue))
		return defaultValue
	}

	out := format.Format(val)

	log.Default().Printf("using %v = %v [from env]\n", key, censorWhenPrivate(key, out))
	return out
}

var _stringFormatter formatter[string] = stringFormatter{}
var _intFormatter formatter[int] = intFormatter{}
var _boolFormatter formatter[bool] = boolFormatter{}

func GetEnvString(key string) string {
	return GetEnvVar(key, _stringFormatter)
}

func GetEnvStringOrDefault(key string, defaultValue string) string {
	return GetEnvOrDefault(key, defaultValue, _stringFormatter)
}

func GetEnvInt(key string) int {
	return GetEnvVar(key, _intFormatter)
}

func GetEnvIntOrDefault(key string, defaultValue int) int {
	return GetEnvOrDefault(key, defaultValue, _intFormatter)
}

func GetEnvBool(key string) bool {
	return GetEnvVar(key, _boolFormatter)
}

func GetEnvBoolOrDefault(key string, defaultValue bool) bool {
	return GetEnvOrDefault(key, defaultValue, _boolFormatter)
}
