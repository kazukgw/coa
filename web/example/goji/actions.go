package main

import (
	"encoding/json"
	"errors"

	"github.com/kazukgw/coa"
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
	jsonStr, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	w := ctx.ResponseWriter()
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
	r := ctx.Request()
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
	act := ctx.ActionGroup()
	if hasP, ok := act.(HasParams); ok {
		r := ctx.Request()
		return json.NewDecoder(r.Body).Decode(hasP.Params())
	}
	return nil
}
