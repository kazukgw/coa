package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kazukgw/coa"
	"github.com/zenazn/goji/web"
)

type Context struct {
	*web.C
	Res  http.ResponseWriter
	Req  *http.Request
	AG   coa.ActionGroup
	logr *Logger
}

func NewContext(c *web.C, w http.ResponseWriter, r *http.Request, act coa.ActionGroup) coa.Context {
	return &Context{c, w, r, act, NewLogger()}
}

func (c *Context) ResponseWriter() http.ResponseWriter {
	return c.Res
}

func (c *Context) Request() *http.Request {
	return c.Req
}

func (c *Context) ActionGroup() coa.ActionGroup {
	return c.AG
}

func (c *Context) Logger() coa.Logger {
	return c.logr
}

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	l := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)
	return &Logger{l}
}

func (l *Logger) Info(v ...interface{}) {
	l.logger.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.logger.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.logger.Println(v...)
}

func (l *Logger) Panic(v ...interface{}) {
	l.logger.Panicln(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Fatalln(v...)
}
