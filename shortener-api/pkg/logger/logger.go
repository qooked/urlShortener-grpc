package logger

import "log/slog"

type Logger struct {
	logs map[string]interface{}
}

func NewLogger() *Logger {
	return &Logger{
		logs: make(map[string]interface{}),
	}
}

func (l *Logger) Add(key string, value interface{}) {
	l.logs[key] = value
}

func (l *Logger) Info(article string) {
	slog.Info(article, l.logs)
}

func (l *Logger) Error(err error) {
	slog.Error(err.Error(), l.logs)
}
