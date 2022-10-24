package logger

import (
	"log"
	"os"
)

func initLogger() *log.Logger {
	flags := log.LstdFlags | log.Lshortfile
	logger := log.New(os.Stdout, "", flags)
	return logger
}

var Logger = initLogger()

func Debug(s string, i ...interface{}) {
	Logger.Println("[DEBUG] "+s, i)
}

func Info(s string, i ...interface{}) {
	Logger.Println("[INFO] "+s, i)
}

func Warn(s string, i ...interface{}) {
	Logger.Println("[WARN] "+s, i)
}

func Error(s string, i ...interface{}) {
	Logger.Println("[ERROR] "+s, i)
}

func Fatal(s string, i ...interface{}) {
	Logger.Fatal("[FATAL] "+s, i)
}

func Panic(s string, i ...interface{}) {
	Logger.Panic("[PANIC] "+s, i)
}
