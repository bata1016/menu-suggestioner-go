package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	models "github.com/module/menu-suggestioner-go/app/models"
)

// Request は、lambdaハンドラのリクエストです。
type Request struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Response は、lambdaハンドラのレスポンスです。
type Response struct {
	Message string
}

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest 、lambdaのリクエストハンドラ関数です。
func HandleRequest(request *Request, ctx context.Context) (*Response, error) {
	user := models.NewUser()
	params := &models.PutUserParams{
		Email: request.Email,
		Name:  request.Name,
	}
	err := user.Put(params)
	if err != nil {
		return nil, err
	}
	return &Response{
		Message: "新しいユーザーの登録に成功しました。",
	}, nil
}
