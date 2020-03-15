
//+build wireinject

package main

import (
	"context"
	"github.com/cheyuexian/go-excise/wire/wire1/foobarbaz"
	. "github.com/google/wire"
)

func InitializeBaz(ctx context.Context) (foobarbaz.Baz, error) {
	Build(foobarbaz.SuperSet)
	return foobarbaz.Baz{}, nil
}