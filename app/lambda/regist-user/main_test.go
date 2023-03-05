package main

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestRegistUser(t *testing.T) {
	uuid, _ := uuid.NewRandom()
	request := &Request{
		Uuid:  uuid.String(),
		Email: "sample@test.com",
		Name:  "test1",
	}
	ctx := context.Background()
	res, err := HandleRequest(request, ctx)
	if err != nil {
		t.Fatalf("err is {%v}", err)
	}
	println(res)
}
