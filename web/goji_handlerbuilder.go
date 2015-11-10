package web

import (
	"net/http"
	"reflect"

	"github.com/kazukgw/coa"
	"github.com/zenazn/goji/web"
)

type GojiHandlerBuilder struct {
	NewContext func(*web.C, http.ResponseWriter, *http.Request, coa.ActionGroup) coa.Context
}

func (ab *GojiHandlerBuilder) Build(zeroActionGroup interface{}) func(web.C, http.ResponseWriter, *http.Request) {
	actionType := reflect.TypeOf(zeroActionGroup)
	if _, ok := reflect.New(actionType).Interface().(coa.ActionGroup); !ok {
		panic(actionType.String() + " dose not implement Action interface")
	}

	return func(c web.C, w http.ResponseWriter, r *http.Request) {
		ag := reflect.New(actionType).Interface().(coa.ActionGroup)
		ctx := ab.NewContext(&c, w, r, ag)
		coa.Exec(ag, ctx)
	}
}
