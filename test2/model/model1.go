package model

import "fmt"

type T1 struct{
	T11
	b int
}
type  T11 struct {
	a int
}
func (self *T11) t11(){
	fmt.Println(self.a)
}
