package main

import (
	"fmt"
	"github.com/cheyuexian/go-excise/wire/wire2/config"
)

func main() {
	s ,e := InitializeClient(config.Config{})
	fmt.Println(s,e)
}
