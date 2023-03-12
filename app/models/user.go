package models

import (
	ddb "github.com/module/menu-suggestioner-go/app/aws/dynamoDB"
)

// User は、ユーザーです。
type User struct {
	Name  string
	Email string
}

// NewUser は、Userを生成する。
func NewUser() *User {
	return &User{
		Name:  "",
		Email: "",
	}
}

// PutUserParams は、ユーザー登録のリクエストパラメーターです。
type PutUserParams struct {
	Uuid      string `json:"uuid"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

// Put は、新しいユーザーを登録します。
func (u *User) Put(params *PutUserParams) error {
	ddb := ddb.NewDynamoDB("users")
	_, err := ddb.PutItem(params)
	if err != nil {
		return err
	}
	return nil
}
