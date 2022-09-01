package main

import "fmt"
import "luago/api"
import "luago/state"

func main() {
	ls := state.New()
	printStack(ls)
	ls.PushBoolean(true)
	printStack(ls)
	ls.PushInteger(10)
	printStack(ls)
	ls.PushNil()
	printStack(ls)
	ls.PushString("Hello")
	printStack(ls)
	ls.Pushvalue(-1)
	printStack(ls)
	ls.Replace(3)
	printStack(ls)
	ls.SetTop(6)
	printStack(ls)
	ls.Remove(-3)
	printStack(ls)
	ls.SetTop(-5)
	printStack(ls)
}
