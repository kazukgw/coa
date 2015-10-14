package web

import (
	"net/http"

	"github.com/kazukgw/coa"
)

type Context interface {
	coa.Context
	ResponseWriter() http.ResponseWriter
	Request() *http.Request
	Logger() Logger
}

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warning(...interface{})
	Warningf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Panic(...interface{})
	Panicf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}
