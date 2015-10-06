package main

import (
	"errors"
)

type User struct {
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	AuthToken string `json:"auth_token"`
}

var user1 = &User{"user01", "User One", "token-01"}
var user2 = &User{"user02", "User Two", "token-02"}
var user3 = &User{"user03", "User One", "token-03"}

var tokenMap = map[string]*User{
	user1.AuthToken: user1,
	user2.AuthToken: user2,
	user3.AuthToken: user3,
}

var idMap = map[string]*User{
	user1.UserID: user1,
	user2.UserID: user2,
	user3.UserID: user3,
}

func FindUserByAuthToken(token string) *User {
	return tokenMap[token]
}

func FindUserByUserID(id string) *User {
	return idMap[id]
}

func SaveUser(u *User) error {
	if u == nil {
		return errors.New("user is nil")
	}
	updateUser := idMap[u.UserID]
	if updateUser == nil {
		return errors.New("user not found")
	}
	updateUser.Name = u.Name
	updateUser.AuthToken = u.AuthToken
	return nil
}
