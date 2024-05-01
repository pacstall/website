package log

import (
	"fmt"
	glog "log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"pacstall.dev/webserver/config"
)

var (
	logInfo  = color.CyanString("INFO")
	logError = color.RedString("ERROR")
	logDebug = color.GreenString("DEBUG")
	logFatal = color.New(color.BgHiRed, color.FgBlack).Sprintf("FATAL")
	logWarn  = color.YellowString("WARN")
)

const (
	logDiscordError  = "❌ Error ❌"
	logDiscordWarn   = "⚠️ Warning"
	logDiscordFatal  = "💀☢️💥 Fatal 🪦⚰️🧟‍♂️"
	logDiscordNotify = "📢 Notification"
)

var discordConfiguration = config.DiscordConfiguration{
	Enabled: false,
}

func Init(config config.DiscordConfiguration) {
	discordConfiguration = config
}

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
	go sendDiscordMessage(true, logDiscordError, message, args...)
}

func Fatal(message string, args ...any) {
	if logLevel > Level.Fatal {
		return
	}

	sendDiscordMessage(true, logDiscordFatal, message, args...)
	doLog(logFatal, message, args...)
}

func Warn(message string, args ...any) {
	if logLevel > Level.Warn {
		return
	}

	doLog(logWarn, message, args...)
	go sendDiscordMessage(true, logDiscordWarn, message, args...)
}

func Debug(message string, args ...any) {
	if logLevel > Level.Debug {
		return
	}

	doLog(logDebug, message, args...)
}

func Notify(message string, args ...any) {
	go sendDiscordMessage(false, logDiscordNotify, message, args...)
}

func NotifyCustom(level, message string, args ...any) {
	go sendDiscordMessage(false, level, message, args...)
}

func sendDiscordMessage(tag bool, level, message string, args ...any) {
	if !discordConfiguration.Enabled {
		return
	}

	msg := fmt.Sprintf("Webserver - %s: %s\n", level, fmt.Sprintf(message, args...))
	if tag {
		msg = fmt.Sprintf("%s %s", discordConfiguration.Tags, msg)
	}

	_, err := discordClient.ChannelMessageSend(
		discordConfiguration.ChannelID,
		msg,
	)

	if err != nil {
		panic(fmt.Sprintf("failed to send discord message\n%v", err))
	}
}

var discordClient = func() *discordgo.Session {
	if discordConfiguration.Enabled {
		return connect(discordConfiguration.Token)
	}

	return nil
}()
