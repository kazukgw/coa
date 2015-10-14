package web

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/kazukgw/coa"
)

type GinHandlerBuilder struct {
	NewContext func(*gin.Context, coa.ActionGroup) Context
}

func (ab *GinHandlerBuilder) Build(zeroActionGroup interface{}) func(*gin.Context) {
	actionType := reflect.TypeOf(zeroActionGroup)
	if _, ok := reflect.New(actionType).Interface().(coa.ActionGroup); !ok {
		panic(actionType.String() + " dose not implement Action interface")
	}

	return func(c *gin.Context) {
		ag := reflect.New(actionType).Interface().(coa.ActionGroup)
		ctx := ab.NewContext(c, ag)
		coa.Exec(ag, ctx)
	}
}
