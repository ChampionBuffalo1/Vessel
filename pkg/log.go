package pkg

import (
	"log/slog"
	"os"
	"sync"
)

var (
	once   sync.Once
	logger *slog.Logger
)

type LogLeveler struct {
}

func (l *LogLeveler) Level() slog.Level {
	env := os.Getenv("DISABLE_DEBUG_LOG")
	if env == "" {
		return slog.LevelDebug
	}
	return slog.LevelInfo
}

const LOG_FILE = "vessel.log"

func NewLogger() *slog.Logger {
	once.Do(func() {
		file, err := os.Open(LOG_FILE)
		if err != nil {
			file, err = os.Create(LOG_FILE)
			if err != nil {
				panic(err)
			}
		}
		handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
			Level: &LogLeveler{},
		})
		logger = slog.New(handler)
	})
	return logger

}
