package logadapter

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type BotgoLogger struct {
	log logx.Logger
}

func NewBotgoLogger() *BotgoLogger {
	return &BotgoLogger{log: logx.WithContext(context.Background())}
}

func (l *BotgoLogger) Debug(v ...interface{}) {
	l.log.Debug(v...)
}

func (l *BotgoLogger) Info(v ...interface{}) {
	l.log.Info(v...)
}

func (l *BotgoLogger) Error(v ...interface{}) {
	l.log.Error(v...)
}

func (l *BotgoLogger) Debugf(format string, v ...interface{}) {
	l.log.Debugf(format, v...)
}

func (l *BotgoLogger) Infof(format string, v ...interface{}) {
	l.log.Infof(format, v...)
}

func (l *BotgoLogger) Errorf(format string, v ...interface{}) {
	l.log.Errorf(format, v...)
}

func (l *BotgoLogger) Warn(v ...any) {
	warns := append([]any{"Warning"}, v...)
	l.log.Error(warns...)
}

func (l *BotgoLogger) Warnf(format string, v ...any) {
	l.log.Errorf("Warning "+format, v...)
}

func (l *BotgoLogger) Sync() error {
	return nil
}
