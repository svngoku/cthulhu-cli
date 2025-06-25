package utils

import (
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func SetupLogger(verbose bool) {
	var l *zap.Logger
	if verbose {
		l, _ = zap.NewDevelopment()
	} else {
		l, _ = zap.NewProduction()
	}
	log = l.Sugar()
}

func Log() *zap.SugaredLogger { return log }
