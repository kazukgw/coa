coa
===============

coa = **Co**mposable **A**ction

Overview
--------------

Package coa provides the executor and interfaces for composable action.

Example
--------------

```go
package main

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

// ActionPrint1 implements coa.Action interface.
type ActionPrint1 struct {
}

func (t1 ActionPrint1) Do(ctx coa.Context) error {
	fmt.Println(1)
	return nil
}

// ActionPrint2 implements coa.Action interface.
type ActionPrint2 struct {
}

func (t2 ActionPrint2) Do(ctx coa.Context) error {
	fmt.Println(2)
	return nil
}

// ActionGroup includes ActionPrint1 and ActionPrint2.
type ActionGroup struct {
	ActionPrint1
	coa.DoSelf
	ActionPrint2
}

func (ag *ActionGroup) Do(ctx coa.Context) error {
	fmt.Println("do")
	return nil
}

func (ag *ActionGroup) HandleError(ctx coa.Context, err error) error {
	return nil
}

func main() {
	ag := &ActionGroup{}
	ctx := &Context{ag}
	coa.Exec(ag, ctx)
	// Output:
	// 1
	// do
	// 2
}
```

### In a web app

```go
// this is a reusable Action
type AuthUserByHeaderToken struct {
	CurrentUser *User
}

func (a *AuthUserByHeaderToken) Do(ctx coa.Context) error {
	cctx := ctx.(CustomContext)
	tkn := cctx.Request().Header["X-Auth-Token"]
	a.CurrentUser = FindUserByAuthToken(tkn)
	if a.CurrentUser == nil {
		  return errors.New("user not found")
	}
	return nil
}

// this is a reusable Action
type RenderJSON struct {
	Data interface{}
}

func (a *MarshalJson) Do(ctx coa.Context) error {
	cctx := ctx.(CustomContext)
	w := cctx.ResponseWriter()
	if err := json.NewEncoder(w).Encode(a.Data); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	return nil
}

// ActionGroup
type GetUser struct {
	AuthUserByHeaderToken
	coa.DoSelf
	RenderJSON
	YourErrorHandler
}

func (ag *GetUser) Do(ctx coa.Context) {
	ag.Data := ag.CurrentUser
}
```

