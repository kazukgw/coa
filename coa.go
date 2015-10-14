// Package coa provides the action executor and interfaces for composable action.
package coa

import (
	"reflect"
	"sync"
)

type Action interface {
	Do(Context) error
}

type ActionGroup interface {
	Action
	HandleError(Context, error) error
}

// PreExec() will be executed before every action excute.
type HasPreExec interface {
	ActionGroup
	PreExec(Context) error
}

// PostExec() will be executed after every action excute.
// But if error handler of the ActionGroup return error,
// PostExec will not be excuted.
type HasPostExec interface {
	ActionGroup
	PostExec(Context) error
}

// If a ActionGroup implements this interface,
// every action will be executed in parallel.
type Parallelable interface {
	ActionGroup
	HandleErrorParallel(Context, error)
	Error() error
}

// If a ActionGroup implements this interface,
// Repeat() will be executed after every action execute. And if Repeat() return
// true, every action will be executed again.
type Repeatable interface {
	ActionGroup
	Repeat() bool
}

// Context is the interface that wrap the ActionGroup Getter function.
type Context interface {
	ActionGroup() ActionGroup
}

// If A ActionGroup has DoSelf as field,
// the Exec function execute Do() of the ActionGroup.
type DoSelf struct{}

// exec Execute Actions in the ActionGroup with Context.
func Exec(ag ActionGroup, ctx Context) error {
	agValue := reflect.ValueOf(ag).Elem()

	paralellableAG, parallelable := ag.(Parallelable)
	repeatableAG, repeatable := ag.(Repeatable)

	if pe, ok := ag.(HasPreExec); ok {
		if err := pe.PreExec(ctx); err != nil {
			if err := pe.HandleError(ctx, err); err != nil {
				return err
			}
		}
	}

	for {
		if parallelable {
			var wg sync.WaitGroup
			for i := 0; i < agValue.NumField(); i++ {
				wg.Add(1)
				fieldIndex := i
				go func() {
					defer wg.Done()
					field := agValue.FieldByIndex([]int{fieldIndex}).Addr().Interface()
					if err := execField(field, ctx, ag); err != nil {
						paralellableAG.HandleErrorParallel(ctx, err)
					}
				}()
			}
			wg.Wait()
			if err := paralellableAG.Error(); err != nil {
				if err := ag.HandleError(ctx, err); err != nil {
					return err
				}
			}
		} else {
			for i := 0; i < agValue.NumField(); i++ {
				field := agValue.FieldByIndex([]int{i}).Addr().Interface()
				if err := execField(field, ctx, ag); err != nil {
					if err := ag.HandleError(ctx, err); err != nil {
						return err
					}
				}
			}
		}

		if !repeatable || !repeatableAG.Repeat() {
			break
		}
	}

	if pe, ok := ag.(HasPostExec); ok {
		if err := pe.PostExec(ctx); err != nil {
			pe.HandleError(ctx, err)
			return err
		}
	}
	return nil
}

func execField(field interface{}, ctx Context, ag ActionGroup) error {
	switch f := field.(type) {
	case ActionGroup:
		return Exec(f, ctx)
	case Action:
		return f.Do(ctx)
	case *DoSelf:
		return ag.Do(ctx)
	}
	return nil
}
