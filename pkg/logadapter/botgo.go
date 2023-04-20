package logadapter

import (
	"github.com/go-kratos/kratos/v2/log"
)

type BotgoLogger struct {
	lg *log.Helper
}

func NewBotgoLogger(lg *log.Helper) *BotgoLogger {
	return &BotgoLogger{lg: lg}
}

func (l *BotgoLogger) Debug(v ...interface{}) {
	l.lg.Debug(v...)
}

func (l *BotgoLogger) Info(v ...interface{}) {
	l.lg.Info(v...)
}

func (l *BotgoLogger) Error(v ...interface{}) {
	l.lg.Error(v...)
}

func (l *BotgoLogger) Debugf(format string, v ...interface{}) {
	l.lg.Debugf(format, v...)
}

func (l *BotgoLogger) Infof(format string, v ...interface{}) {
	l.lg.Infof(format, v...)
}

func (l *BotgoLogger) Errorf(format string, v ...interface{}) {
	l.lg.Errorf(format, v...)
}

func (l *BotgoLogger) Warn(v ...any) {
	l.lg.Warn(v...)
}

func (l *BotgoLogger) Warnf(format string, v ...any) {
	l.lg.Warnf(format, v...)
}

func (l *BotgoLogger) Sync() error {
	// won't need sync
	return nil
}
