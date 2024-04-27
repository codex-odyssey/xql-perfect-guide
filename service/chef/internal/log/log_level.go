package logging

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

const (
	debugStr = "debug"
	infoStr  = "info"
	warnStr  = "warn"
	errorStr = "error"
)

func getZapLogLevelFromEnv(logLevel string) zapcore.Level {
	switch strings.ToLower(logLevel) {
	case debugStr:
		return zapcore.DebugLevel
	case infoStr:
		return zapcore.InfoLevel
	case warnStr:
		return zapcore.WarnLevel
	case errorStr:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
