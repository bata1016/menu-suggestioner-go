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

type PutUserParams struct {
	Email string
	Name  string
}

// Put は、新しいユーザーを登録します。
func (u *User) Put(params *PutUserParams) error {
	ddb := ddb.NewDynamoDB()
	_, err := ddb.PutItem(params)
	if err != nil {
		return err
	}
	return nil
}
