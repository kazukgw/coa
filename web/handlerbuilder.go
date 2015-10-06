package web

import (
	"net/http"
	"reflect"

	"github.com/kazukgw/coa"
)

type HandlerBuilder struct {
	NewContext func(http.ResponseWriter, *http.Request, coa.ActionGroup) coa.Context
}

func (ab *HandlerBuilder) Build(zeroActionGroup interface{}) func(http.ResponseWriter, *http.Request) {
	actionType := reflect.TypeOf(zeroActionGroup)
	if _, ok := reflect.New(actionType).Interface().(coa.ActionGroup); !ok {
		panic(actionType.String() + " dose not implement coa.ActionGroup interface")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ag := reflect.New(actionType).Interface().(coa.ActionGroup)
		coa.Exec(ab.NewContext(w, r, ag))
	}
}
