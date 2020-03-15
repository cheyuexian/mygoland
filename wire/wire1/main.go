package main

import (
	"context"
	"fmt"
)


func main() {
	baz , err := InitializeBaz(context.Background())
	fmt.Println(baz,err)
}
