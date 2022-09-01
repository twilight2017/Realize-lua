package main

import "fmt"
import "luago/api"
import "luago/state"

func main() {
	ls := state.New()
	ls.PushBoolean(true)
	ls.PushInteger(10)
	ls.PushNil()
	ls.PushString("Hello")
	ls.Pushvalue(-1)
	ls.Replace(3)
	ls.SetTop(6)
	ls.Remove(-3)
	ls.SetTop(-5)
}
