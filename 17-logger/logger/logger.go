package logger

import (
	"os"

	"github.com/x-cray/logrus-prefixed-formatter"

	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
)

func Log() *logrus.Logger {

	logger := logrus.New()

	logger.Out = ansicolor.NewAnsiColorWriter(os.Stdout)
	logger.Level = logrus.DebugLevel
	logger.SetReportCaller(true)
	logger.ReportCaller = true

	// Formater
	logger.Formatter = &prefixed.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
		ForceColors:     true,
	}

	filenameHook := NewHook()
	filenameHook.Field = "File"

	logger.AddHook(filenameHook)

	// logger.Info("Success message")
	// logger.Warning("Warning message")
	// logger.Error("Error message")

	return logger
}
