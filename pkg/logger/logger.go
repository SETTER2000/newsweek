package logger

import (
	"github.com/setter2008/newsweek/config"
	"log/slog"
	"os"
)

func NewLogger(config *config.LogConfig) *slog.Logger {
	var levelMap = map[int]slog.Level{
		0: slog.LevelDebug,
		1: slog.LevelInfo,
		2: slog.LevelWarn,
		3: slog.LevelError,
	}
	var programLevel = new(slog.LevelVar)
	// Устанавливаем уровень по умолчанию (можно из конфига)
	programLevel.Set(levelMap[config.Level])

	var handler slog.Handler

	if config.Format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: programLevel,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: programLevel,
		})
	}
	logger := slog.New(handler)
	return logger
}
