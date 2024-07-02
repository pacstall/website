package config

import (
	"fmt"
	"os"
	"strconv"
)

func toInt(str string) int {
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

func toBool(str string) bool {
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
	return toInt(str)
}

type boolFormatter struct{}

func (f boolFormatter) Format(str string) bool {
	return toBool(str)
}

func getEnvVar[T any](key string, format formatter[T]) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		if os.Getenv("GO_ENV") == "test" {
			fmt.Printf("Running in test mode. Using value '0' for '%s'\n", key)
			return format.Format("0")
		}

		panic(fmt.Sprintf("could not find environment variable '%s'", key))
	}

	return format.Format(val)
}

func getEnvOrDefault[T any](key string, defaultValue T, format formatter[T]) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return format.Format(val)
}

var _stringFormatter formatter[string] = stringFormatter{}
var _intFormatter formatter[int] = intFormatter{}
var _boolFormatter formatter[bool] = boolFormatter{}

func getEnvString(key string) string {
	return getEnvVar(key, _stringFormatter)
}

func getEnvInt(key string) int {
	return getEnvVar(key, _intFormatter)
}

func getEnvBool(key string) bool {
	return getEnvVar(key, _boolFormatter)
}

func getEnvBoolOrDefault(key string, defaultValue bool) bool {
	return getEnvOrDefault(key, defaultValue, _boolFormatter)
}
