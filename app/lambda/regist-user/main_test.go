package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestRegistUser(t *testing.T) {
	request := &Request{
		Email: "sample@test.com",
		Name:  "test1",
	}
	jsonReq, _ := json.Marshal(request)
	req := events.APIGatewayProxyRequest{
		Body: string(jsonReq),
	}

	ctx := context.Background()
	res, err := HandleRequest(ctx, req)
	if err != nil {
		t.Fatalf("err is {%v}", err)
	}
	fmt.Println(res)
}
