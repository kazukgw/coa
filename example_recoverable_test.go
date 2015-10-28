package coa_test

import (
	"fmt"

	"github.com/kazukgw/coa"
)

type RecoverableGroup struct {
	ActionPrint1
	ActionPrint2
	coa.DoSelf
	ActionPrint3
}

func (ag *RecoverableGroup) Recover(r interface{}) {
	fmt.Println(r)
}

func (ag *RecoverableGroup) Do(ctx coa.Context) error {
	panic("recover")
}

func (ag *RecoverableGroup) HandleError(ctx coa.Context, err error) error {
	// handl error
	return nil
}

func Example_recover() {
	ag := &RecoverableGroup{}
	ctx := &Context{ag}
	coa.Exec(ag, ctx)
	// Output:
	// 1
	// 2
	// recover
}
