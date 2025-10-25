package logging

import (
    "go.uber.org/zap"
)

func Init(debug bool) {
    var logger *zap.Logger
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
