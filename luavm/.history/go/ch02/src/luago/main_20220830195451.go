package main

import "fmt"
import "io/ioutil"
import "os"
import "luago/binchunk"

func list(f *binchunk.Prototype) {
	printHeader(f)
	printCode(f)
	printDetail(f)
}
func main() {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		proto := binchunk.Undump(data)
		list(proto)
	}
}
