package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kazukgw/coa"
	coaweb "github.com/kazukgw/coa/web"
	"github.com/zenazn/goji/web"
)

type Context struct {
	*web.C
	responseWriter http.ResponseWriter
	request        *http.Request
	actionGroup    coa.ActionGroup
	logr           *Logger
}

func NewContext(c *web.C, w http.ResponseWriter, r *http.Request, act coa.ActionGroup) coaweb.Context {
	return &Context{c, w, r, act, NewLogger()}
}

func (c *Context) ResponseWriter() http.ResponseWriter {
	return c.responseWriter
}

func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) ActionGroup() coa.ActionGroup {
	return c.actionGroup
}

func (c *Context) Logger() coaweb.Logger {
	return c.logr
}

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	l := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)
	return &Logger{l}
}

func (l *Logger) Debug(v ...interface{}) {
	l.logger.Println(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.logger.Println(v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.logger.Println(v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.logger.Println(v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *Logger) Panic(v ...interface{}) {
	l.logger.Panicln(v...)
}
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Fatalln(v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}
