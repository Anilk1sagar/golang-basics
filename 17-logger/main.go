package main

import (
	"fmt"
	"os"

	"gitlab.com/anilk1sagar/go_basic_practice/17-logger/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	logger.Log().Errorln("Helloooooooooo")

	if err := do(); err != nil {

		fmt.Println("Error is: ", os.Stderr, err.Error())
		os.Exit(1)
	}
}

func do() error {

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := config.Build()
	if err != nil {
		return err
	}

	logger.Debug("test")
	logger.Info("test")
	logger.Warn("test")
	logger.Error("test")
	// logger.Panic("test")

	return nil
}
