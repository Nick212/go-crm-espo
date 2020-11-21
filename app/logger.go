package app

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	logrusinstanceonce sync.Once
	logrusinstance     *logrus.Logger
)

// GlobalLogger return a global logger (for logs outside some context).
func GlobalLogger() *logrus.Logger {
	logrusinstanceonce.Do(func() {
		logrusinstance = NewLogger()
	})
	return logrusinstance
}

// NewLogger create a new instace of logrus logger.
func NewLogger() *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	})
	return l
}
