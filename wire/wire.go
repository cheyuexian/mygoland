//+build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(msg string) Event{
	wire.Build(NewGreeter1,NewMessage1,NewEvent1)
	return Event{}
}

