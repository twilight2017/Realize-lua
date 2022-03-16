package main

import (
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		proto := Undump(data)
		list(proto)
	}
}

func list(f *Prototype) {
	PrintHeader(f)
	PrintCode(f)
	PrintDetail(f)
	for _, p := range f.Protos {
		list(p)
	}
}
