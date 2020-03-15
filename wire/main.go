package main

import "fmt"

type Message struct{
	msg string
}

type Greeter struct {
	Message Message
}

type Event struct{
	Greeter Greeter
}
func NewMessage1(msg string) Message{
	return Message{msg:msg}
}

func NewGreeter1(m Message) Greeter{
	return Greeter{Message:m}
}
func NewEvent1(g Greeter) Event  {
	return Event{Greeter:g}
}
func (g Greeter) Greet() Message{
	return g.Message
}

func (e Event) Start(){
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	//message := NewMessage("hello world")
	//greeter := NewGreeter(message)
	//event := NewEvent(greeter)
	//event.Start()


	event := InitializeEvent("hello_world")
	event.Start()

}
