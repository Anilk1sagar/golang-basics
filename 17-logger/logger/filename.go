package logger

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

type Hook struct {
	Field        string
	Skip         int
	levels       []logrus.Level
	SkipPrefixes []string
	Formatter    func(file, function string, line int) string
}

func (hook *Hook) Levels() []logrus.Level {
	return hook.levels
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data[hook.Field] = hook.Formatter(hook.findCaller())
	return nil
}

func (hook *Hook) findCaller() (string, string, int) {

	pc, file, line, _ := runtime.Caller(1)

	fName := runtime.FuncForPC(pc).Name()

	return file, fName, line

}

func NewHook(levels ...logrus.Level) *Hook {
	hook := Hook{
		Field:        "File",
		Skip:         5,
		levels:       levels,
		SkipPrefixes: []string{"logrus/", "logrus@"},
		Formatter: func(file, function string, line int) string {
			return fmt.Sprintf("%s:%d", file, line)
		},
	}
	if len(hook.levels) == 0 {
		hook.levels = logrus.AllLevels
	}

	return &hook
}
