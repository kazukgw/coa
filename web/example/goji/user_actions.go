package main

import "github.com/kazukgw/coa"

// Fieldの上から順にAction interface を実装するものが実行される
type GetUser struct {
	AuthenticateUserByHeaderToken
	coa.DoSelf // 自分を実行するときはDoSelfをおく
	ResultActionHandler

	ErrorHandler
}

func (act *GetUser) Do(ctx coa.Context) error {
	act.ResultAction = &ResultJSON{
		Data: struct {
			Response string `json:"response"`
			*User    `json:"data"`
		}{"success", act.CurrentUser},
		Code: 200,
	}
	return nil
}

type UpdateUser struct {
	ParamsUnmarshaler
	AuthenticateUserByHeaderToken
	coa.DoSelf
	ResultActionHandler

	ErrorHandler
	ParamUser *User
}

func (act *UpdateUser) Params() interface{} {
	act.ParamUser = &User{}
	return act.ParamUser
}

func (act *UpdateUser) Do(ctx coa.Context) error {
	currentUser := act.CurrentUser
	act.ParamUser.UserID = currentUser.UserID
	if err := SaveUser(act.ParamUser); err != nil {
		return err
	}
	act.ResultAction = &ResultJSON{
		Data: struct {
			Response string `json:"response"`
		}{"success"},
		Code: 200,
	}
	return nil
}
