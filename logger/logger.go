package logger

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	logLevel := logrus.DebugLevel

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	writers := make([]io.Writer, 0, 2)
	writers = append(writers, os.Stdout)

	mw := io.MultiWriter(writers...)
	logrus.SetOutput(mw)
	logrus.SetLevel(logLevel)
}

func Get() *logrus.Entry {
	return logrus.NewEntry(logrus.StandardLogger())
}
