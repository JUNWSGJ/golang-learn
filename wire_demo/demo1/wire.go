//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

func InitializeEvent1(msg string) Event {
	wire.Build(NewMessage, NewGreeter, NewEvent)
	return Event{} //返回值没有实际意义，只需符合函数签名即可
}

// EventSet Event通常是一起使用的一个集合，使用wire.NewSet进行组合
var EventSet = wire.NewSet(NewEvent, NewMessage, NewGreeter)

func InitializeEvent2(msg string) Event {
	wire.Build(EventSet)
	return Event{} //返回值没有实际意义，只需符合函数签名即可
}
