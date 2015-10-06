package main

import (
	"encoding/json"

	"github.com/kazukgw/coa"
)

type ErrorHandler struct {
}

type ErrorJSON struct {
	Err string `json:"error"`
}

func (eh ErrorHandler) HandleError(ctx coa.Context, err error) {
	ctx.Logger().Error(err)
	w := ctx.ResponseWriter()
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(ErrorJSON{err.Error()})
}
