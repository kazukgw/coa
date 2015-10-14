package coa_test

import (
	"errors"
	"fmt"
	"sync"

	"github.com/kazukgw/coa"
)

type ParallelableGroup struct {
	T1_1 ActionPrint1
	T2_1 ActionPrint2
	T3_1 ActionPrint3
	T1_2 ActionPrint1
	T2_2 ActionPrint2
	coa.DoSelf
	T3_2 ActionPrint3
	T1_3 ActionPrint1
	T2_3 ActionPrint2
	T3_3 ActionPrint3

	sync.Mutex
	Errors []error
}

func (ag *ParallelableGroup) Do(ctx coa.Context) error {
	fmt.Println("do")
	return nil
}

func (ag *ParallelableGroup) HandleError(ctx coa.Context, err error) error {
	// handle error
	return nil
}

func (ag *ParallelableGroup) HandleErrorParallel(ctx coa.Context, err error) {
	ag.Mutex.Lock()
	if err != nil {
		ag.Errors = append(ag.Errors)
	}
	ag.Mutex.Unlock()
}

func (ag *ParallelableGroup) Error() error {
	if len(ag.Errors) > 0 {
		return errors.New("some error")
	}
	return nil
}

func Example_parallel() {
	// ag := &ParallelableGroup{}
	// ctx := &Context{ag}
	// coa.Exec(ctx)
	// Output Not:
	// 1
	// 2
	// 3
	// 1
	// 2
	// do
	// 3
	// 1
	// 2
	// 3

	// Output Like:
	// 3
	// 1
	// do
	// 2
	// 2
	// 3
	// 1
	// 1
	// 2
	// 3
}
