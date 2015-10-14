package main

import (
	"encoding/json"

	"github.com/kazukgw/coa"
	"github.com/kazukgw/coa/web"
)

type ErrorHandler struct {
}

type ErrorJSON struct {
	Err string `json:"error"`
}

func (eh ErrorHandler) HandleError(ctx coa.Context, err error) error {
	wctx := ctx.(web.Context)
	wctx.Logger().Error(err)
	w := wctx.ResponseWriter()
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(ErrorJSON{err.Error()})
	return err
}
