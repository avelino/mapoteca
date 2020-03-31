package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
)

var logLevel logrus.Level

// New Log decorator
func New() *logrus.Logger {
	var l = logrus.New()

	caller := getCaller(2)
	filenameHook := newHook(caller)
	l.AddHook(filenameHook)

	l.Formatter.(*logrus.TextFormatter).FullTimestamp = true
	l.Formatter.(*logrus.TextFormatter).TimestampFormat = "2006-01-02 15:04:05"

	l.SetOutput(os.Stdout)

	return l
}

func getCaller(skip int) string {
	var projectName = "mapoteca"
	_, file, _, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}

	index := strings.LastIndex(file, projectName)

	if index == -1 {
		return ""
	}

	filePath := file[index+len(projectName)+1:]
	return filePath
}

type hook struct {
	Caller    string
	Formatter func(file string) string
}

func (h *hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *hook) Fire(entry *logrus.Entry) error {
	entry.Data["filename"] = h.Formatter(h.Caller)
	return nil
}

func newHook(caller string) *hook {
	h := hook{
		Caller: caller,
		Formatter: func(file string) string {
			return fmt.Sprintf("%s", file)
		},
	}

	return &h
}
