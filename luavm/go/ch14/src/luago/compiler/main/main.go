package main

import "fmt"
import "io/ioutil"
import "os"
import "lexer"

func main() {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		testLexer(string(data), os.Args[1])
	}
}

func testLexer(chunk, chunkName string){
	lexer := NextLexer(chunk, chunkName)
	for{
		line, kind, token := lexer.NextToken()
		fmt.Printf("[%2d] [%-10s] %s\n",
		    line, kindToCategory(kind), token
		)
		if kind ==TOKEN_EOF{
			break
		}
	}
}
