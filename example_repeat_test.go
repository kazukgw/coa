package coa_test

import (
	"fmt"

	"github.com/kazukgw/coa"
)

type RepeatableGroup struct {
	ActionPrint1
	ActionPrint2
	coa.DoSelf
	ActionPrint3

	Count int
}

func (ag *RepeatableGroup) Do(ctx coa.Context) error {
	fmt.Println("do")
	return nil
}

func (ag *RepeatableGroup) HandleError(ctx coa.Context, err error) error {
	return nil
}

func (ag *RepeatableGroup) Repeat() bool {
	ag.Count++
	return ag.Count < 3
}

func Example_repeat() {
	ag := &RepeatableGroup{}
	ctx := &Context{ag}
	coa.Exec(ag, ctx)
	// Output:
	// 1
	// 2
	// do
	// 3
	// 1
	// 2
	// do
	// 3
	// 1
	// 2
	// do
	// 3
}
