package app

import (
	"github.com/sirupsen/logrus"
)

// Context represent a context for application.
type Context struct {
	Logger     logrus.FieldLogger
	LanguageID string
	RequestID  string
	Version    string
	IP         string
	HTTPPrefix string
}

// WithLogger add logger for context.
func (ctx *Context) WithLogger(logger logrus.FieldLogger) *Context {
	ret := *ctx
	ret.Logger = logger
	return &ret
}
