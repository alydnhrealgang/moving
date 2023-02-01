package logs

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var Root = logrus.New()

func init() {
	Root.Out = os.Stdout
	Root.SetFormatter(&logrus.JSONFormatter{})
	useNoop, ok := os.LookupEnv("USE_NOOP_LOGGER")
	if ok && strings.Compare(useNoop, "true") == 0 {
		Root.Info("Use Noop Logger: true")
		Root.Out = &noopWriter{}
	}

	level, ok := os.LookupEnv("LOG_LEVEL")
	if ok {
		Root.Info("Use Log Level:", level)
		switch level {
		case "error":
			Root.Level = logrus.ErrorLevel
		case "warn":
			Root.Level = logrus.WarnLevel
		case "info":
			Root.Level = logrus.InfoLevel
		}
	}
}

type noopWriter struct {
}

func (nw noopWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
