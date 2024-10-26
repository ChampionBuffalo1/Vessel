package log

import (
	"log/slog"
	"os"
	"sync"
)

var (
	once   sync.Once
	logger *slog.Logger
)

const LOG_FILE = "vessel.log"

func init() {
	_ = InitLogger()
}

func InitLogger() *slog.Logger {
	once.Do(func() {
		file, err := os.OpenFile(LOG_FILE, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			file, err = os.Create(LOG_FILE)
			if err != nil {
				panic(err)
			}
		}
		level := slog.LevelInfo
		if env := os.Getenv("NO_DEBUG"); env == "" {
			level = slog.LevelDebug
		}

		handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
			Level: level,
		})
		logger = slog.New(handler)
	})
	return logger
}

func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}

func Warn(msg string, args ...any) {
	logger.Warn(msg, args...)
}
