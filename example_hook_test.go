package coa_test

import (
	"fmt"

	"github.com/kazukgw/coa"
)

type HookGroup struct {
	ActionPrint1
	ActionPrint2
	coa.DoSelf
	ActionPrint3
}

func (ag *HookGroup) PreExec(ctx coa.Context) error {
	fmt.Println("PreExec!")
	return nil
}

func (ag *HookGroup) PostExec(ctx coa.Context) error {
	fmt.Println("PostExec!")
	return nil
}

func (ag *HookGroup) Do(ctx coa.Context) error {
	fmt.Println("do")
	return nil
}

func (ag *HookGroup) HandleError(ctx coa.Context, err error) error {
	// handl error
	return nil
}

func Example_hook() {
	ag := &HookGroup{}
	ctx := &Context{ag}
	coa.Exec(ag, ctx)
	// Output:
	// PreExec!
	// 1
	// 2
	// do
	// 3
	// PostExec!
}
