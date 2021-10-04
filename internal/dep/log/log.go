package log

import (
	"go.uber.org/zap"
	"recruiting_system/internal/dep/config"
)

var s *service

type service struct {
	*zap.Logger        `wire:"-"`
	*zap.SugaredLogger `wire:"-"`
	config.Log
}

func init() {
	s = new(service)
	s.Log = config.ProvideLog()

	var cfg = zap.Config{
		Level:             s.level(),
		Development:       s.Log.Development,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     s.customEncoderConfig(),
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields:     nil,
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	if s.Log.LogRotate == true {
		logger = logger.WithOptions(s.logRotateOptions())
	}

	defer logger.Sync()

	s.SugaredLogger = logger.Sugar()
	s.Logger = logger

	s.Logger.Info("logService init successfully")
}