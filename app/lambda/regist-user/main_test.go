package main

import (
	"context"
	"fmt"
	"testing"
)

func TestRegistUser(t *testing.T) {
	request := &Request{
		Email: "sample@test.com",
		Name:  "test1",
	}
	ctx := context.Background()
	res, err := HandleRequest(ctx, request)
	if err != nil {
		t.Fatalf("err is {%v}", err)
	}
	fmt.Println(res)
}
