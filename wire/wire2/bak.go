// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/cheyuexian/go-excise/wire/wire2/config"
)

// Injectors from wire.go:

func InitializeClient(config2 config.Config) (*config.Service, error) {
	apiClient, err := config.NewAPIClient(config2)
	if err != nil {
		return nil, err
	}
	service := config.NewService(apiClient)
	return service, nil
}
