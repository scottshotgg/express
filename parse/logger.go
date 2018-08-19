package parse

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// FIXME: see if we can change this in the binary as a compile time flag
	ExpressDebug = "EXPR_DEBUG"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	err    error

	zapConfig = zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "lvl",
			NameKey:        "name",
			CallerKey:      "call",
			MessageKey:     "msg",
			StacktraceKey:  "stack",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
)

func (p *Parser) InitLogger() error {
	// FIXME: for now just check 'true' for now
	if os.Getenv(ExpressDebug) == "true" {
		zapConfig.Development = true
	}

	logger, err = zapConfig.Build()
	if err != nil {
		return err
	}

	// Use a sugared logger; slower but has print/f/ln which makes it more versatile and readable
	sugar = logger.Sugar()

	return nil
}
