package service

import (
	"FINE/internal/model"
	"context"
	"errors"
)

var sExample *example

func init() {
	sExample = &example{}
}

type example struct {
}

func Example() *example {
	if sExample == nil {
		panic("service forget init?")
	}
	return sExample
}

func (e *example) SayHello(ctx context.Context, in *model.HelloIn) (*model.HelloOut, error) {
	if in.Name != "RectLight" {
		return nil, errors.New("name is not RectLight")
	}
	return &model.HelloOut{Message: "Hello, " + in.Name}, nil
}
