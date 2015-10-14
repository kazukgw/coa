package coa_test

import (
	"fmt"

	"github.com/kazukgw/coa"
)

type Context struct {
	actionGroup coa.ActionGroup
}

func (c *Context) ActionGroup() coa.ActionGroup {
	return c.actionGroup
}

type ActionPrint1 struct {
}

func (t1 ActionPrint1) Do(ctx coa.Context) error {
	fmt.Println(1)
	return nil
}

type ActionPrint2 struct {
}

func (t2 ActionPrint2) Do(ctx coa.Context) error {
	fmt.Println(2)
	return nil
}

type ActionPrint3 struct {
}

func (t3 ActionPrint3) Do(ctx coa.Context) error {
	fmt.Println(3)
	return nil
}

type ActionGroup struct {
	ActionPrint1
	ActionPrint2
	coa.DoSelf
	ActionPrint3
}

func (ag *ActionGroup) Do(ctx coa.Context) error {
	fmt.Println("do")
	return nil
}

func (ag *ActionGroup) HandleError(ctx coa.Context, err error) error {
	// handle error
	return nil
}

func Example() {
	ag := &ActionGroup{}
	ctx := &Context{&ActionGroup{}}
	coa.Exec(ag, ctx)
	// Output:
	// 1
	// 2
	// do
	// 3
}
