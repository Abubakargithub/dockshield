package logging

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func Init(debug bool) {
	var err error
	if debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	zap.L().Info("Logger initialized")
}

func Sync() {
	_ = zap.L().Sync()
}
