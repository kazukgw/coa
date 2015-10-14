package coa_test

import (
	"fmt"

	"github.com/kazukgw/coa"
)

type NestedGroup struct {
	ActionPrint1
	ActionPrint2
	coa.DoSelf
	ActionPrint3
}

func (ag *NestedGroup) Do(ctx coa.Context) error {
	fmt.Println("nested")
	return nil
}

func (ag *NestedGroup) HandleError(ctx coa.Context, err error) error {
	// handl error
	return nil
}

type NestGroup struct {
	NestedGroup
	ActionPrint1
	coa.DoSelf
	ActionPrint2
}

func (ag *NestGroup) Do(ctx coa.Context) error {
	fmt.Println("do")
	return nil
}

func (ag *NestGroup) HandleError(ctx coa.Context, err error) error {
	// handl error
	return nil
}

func Example_nestedAction() {
	ag := &NestGroup{}
	ctx := &Context{ag}
	coa.Exec(ag, ctx)
	// Output:
	// 1
	// 2
	// nested
	// 3
	// 1
	// do
	// 2
}
