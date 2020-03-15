//+build wireinject

package main

import (
	. "github.com/cheyuexian/go-excise/wire/wire2/config"
	"github.com/google/wire"
)


func InitializeClient(config Config) (*Service, error) { // <-- 第二个参数设置成error
	wire.Build(NewService,NewAPIClient )
	return nil,nil
}


