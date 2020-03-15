
// +build linux,cgo darwin,cgo
package main

import (
	"flag"
	"fmt"
	"sync"
)
func test(){
	fmt.Println("123")
	//test1()
}
var (
	 ss  = flag.String("conf","defal","hha ")
)

type A struct {
	pool *sync.Pool
}

func main() {
	a := &A{}
	a.pool.New = func() interface{} {
		fmt.Println("hello")
		return "123"
	}
	aa := a.pool.Get()
	fmt.Println("aaa ",aa)
	a.pool.Put("1")

	aa = a.pool.Get()
	fmt.Println("aaaa ",aa)
}
