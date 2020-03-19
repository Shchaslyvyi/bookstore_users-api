package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log system: Uber Zap used
var (
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

// Info is a function to build the logger Info message
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

// Error is a function to build the logger Error message
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error: ", err))
	log.Info(msg, tags...)
	log.Sync()
}
