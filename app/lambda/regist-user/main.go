package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/aws/aws-lambda-go/events"
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

// HandleRequest は、lambdaのリクエストハンドラ関数です。
func HandleRequest(ctx context.Context, apiRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var request Request
	err := json.Unmarshal([]byte(apiRequest.Body), &request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Got Error in Unmarshal Request. Detail: %v", err.Error()),
			StatusCode: 500,
		}, err
	}
	user := models.NewUser()
	uuid, _ := uuid.NewRandom()
	params := &models.PutUserParams{
		Uuid:      uuid.String(),
		Email:     request.Email,
		Name:      request.Name,
		CreatedAt: time.Now().UTC().String(),
	}
	err = user.Put(params)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Got Error in Regist User Request. Detail: %v", err.Error()),
			StatusCode: 500,
		}, err
	}
	return events.APIGatewayProxyResponse{
		Body:       "Success Create New User",
		StatusCode: 200,
	}, err
}
