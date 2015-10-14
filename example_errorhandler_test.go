package coa_test

import (
	"errors"
	"fmt"

	"github.com/kazukgw/coa"
)

type ActionError struct {
}

func (a *ActionError) Do(ctx coa.Context) error {
	return errors.New("error!")
}

type ErrorGroup struct {
	ActionPrint1
	ActionPrint2
	ActionError
	coa.DoSelf
	ActionPrint3
}

func (ag *ErrorGroup) Do(ctx coa.Context) error {
	fmt.Println("do")
	return nil
}

func (ag *ErrorGroup) HandleError(ctx coa.Context, err error) error {
	fmt.Println(err.Error())
	return err
}

func Example_errorHandle() {
	ag := &ErrorGroup{}
	ctx := &Context{ag}
	coa.Exec(ag, ctx)
	// Output:
	// 1
	// 2
	// error!
}
