package log

import (
	"github.com/Sirupsen/logrus"
)

var Logger = logrus.New()

func SetLevel(level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Warnf("Could not parse log level, defaulting to 'info': %s", err)
		lvl = logrus.InfoLevel
	}
	Logger.Level = lvl
}

func SetFormatter(format string) {
	switch format {
	case "json":
		Logger.Formatter = &logrus.JSONFormatter{}
	default:
		Logger.Formatter = &logrus.TextFormatter{}
	}
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Infof(fmt string, args ...interface{}) {
	Logger.Infof(fmt, args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Errorf(fmt string, args ...interface{}) {
	Logger.Errorf(fmt, args...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

func Fatalf(fmt string, args ...interface{}) {
	Logger.Fatalf(fmt, args...)
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func Debugf(fmt string, args ...interface{}) {
	Logger.Debugf(fmt, args...)
}

func Panic(args ...interface{}) {
	Logger.Panic(args...)
}

func Panicf(fmt string, args ...interface{}) {
	Logger.Panicf(fmt, args...)
}
