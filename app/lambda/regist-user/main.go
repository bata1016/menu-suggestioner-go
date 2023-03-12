package main

import (
	"context"
	"github.com/google/uuid"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	models "github.com/module/menu-suggestioner-go/app/models"
)

// Request は、lambdaハンドラのリクエストです。
type Request struct {
	Uuid  string `json:"uuid"`
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

// HandleRequest は、lambdaのリクエストハンドラ関数です。
func HandleRequest(ctx context.Context, request *Request) (*Response, error) {
	user := models.NewUser()
	uuid, _ := uuid.NewRandom()
	params := &models.PutUserParams{
		Uuid:      uuid.String(),
		Email:     request.Email,
		Name:      request.Name,
		CreatedAt: time.Now().UTC().String(),
	}
	err := user.Put(params)
	if err != nil {
		return nil, err
	}
	return &Response{
		Message: "新しいユーザーの登録に成功しました。",
	}, nil
}
