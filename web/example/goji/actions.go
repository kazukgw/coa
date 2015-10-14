package main

import (
	"encoding/json"
	"errors"

	"github.com/kazukgw/coa"
	"github.com/kazukgw/coa/web"
)

type ResultActionHandler struct {
	ResultAction coa.Action
}

func (rh *ResultActionHandler) Do(ctx coa.Context) error {
	return rh.ResultAction.Do(ctx)
}

type ResultJSON struct {
	Data interface{}
	Code int
}

func (r *ResultJSON) Do(ctx coa.Context) error {
	wctx := ctx.(web.Context)

	jsonStr, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	w := wctx.ResponseWriter()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	_, err = w.Write(jsonStr)
	return err
}

var ErrUserNotFound = errors.New("user not found")

type AuthenticateUserByHeaderToken struct {
	CurrentUser *User
	AuthToken   string
}

func (auth *AuthenticateUserByHeaderToken) Do(ctx coa.Context) error {
	wctx := ctx.(web.Context)
	r := wctx.Request()
	auth.AuthToken = r.Header.Get("X-Auth-Token")
	u := FindUserByAuthToken(auth.AuthToken)
	if u == nil {
		return ErrUserNotFound
	}
	auth.CurrentUser = u
	return nil
}

type HasParams interface {
	Params() interface{}
}

type ParamsUnmarshaler struct {
}

func (pu *ParamsUnmarshaler) Do(ctx coa.Context) error {
	wctx := ctx.(web.Context)
	act := wctx.ActionGroup()
	if hasP, ok := act.(HasParams); ok {
		r := wctx.Request()
		return json.NewDecoder(r.Body).Decode(hasP.Params())
	}
	return nil
}
