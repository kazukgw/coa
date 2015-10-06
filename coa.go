package coa

import (
	"net/http"
	"reflect"
	"sync"
)

type Action interface {
	Do(Context) error
}

type ActionGroup interface {
	Action
	HandleError(Context, error)
}

type HasPreExec interface {
	ActionGroup
	PreExec(Context) error
}

type HasPostExec interface {
	ActionGroup
	PostExec(Context) error
}

type Parallelable interface {
	ActionGroup
	HandleErrorParallel(Context, error)
	Error() error
}

type Repeatable interface {
	ActionGroup
	Repeate() bool
}

type Context interface {
	ResponseWriter() http.ResponseWriter
	Request() *http.Request
	ActionGroup() ActionGroup
	Logger() Logger
}

type Logger interface {
	Info(...interface{})
	Warning(...interface{})
	Error(...interface{})
	Panic(...interface{})
	Fatal(...interface{})
}

type DoSelf struct{}

func Exec(ctx Context) error {
	return exec(ctx, ctx.ActionGroup())
}

func exec(ctx Context, ag ActionGroup) error {
	agValue := reflect.ValueOf(ag).Elem()

	paralellableAG, parallelable := ag.(Parallelable)
	repeatableAG, repeatable := ag.(Repeatable)

	if pe, ok := ag.(HasPreExec); ok {
		if err := pe.PreExec(ctx); err != nil {
			pe.HandleError(ctx, err)
			return err
		}
	}

	for {
		if parallelable {
			var wg sync.WaitGroup
			for i := 0; i < agValue.NumField(); i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					field := agValue.FieldByIndex([]int{i}).Addr().Interface()
					if err := execField(field, ctx, ag); err != nil {
						paralellableAG.HandleErrorParallel(ctx, err)
					}
				}()
			}
			wg.Wait()
			if err := paralellableAG.Error(); err != nil {
				ag.HandleError(ctx, err)
				return err
			}
		} else {
			for i := 0; i < agValue.NumField(); i++ {
				field := agValue.FieldByIndex([]int{i}).Addr().Interface()
				if err := execField(field, ctx, ag); err != nil {
					ag.HandleError(ctx, err)
					return err
				}
			}
		}

		if !repeatable || !repeatableAG.Repeate() {
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
		return exec(ctx, f)
	case Action:
		return f.Do(ctx)
	case *DoSelf:
		return ag.Do(ctx)
	}
	return nil
}
