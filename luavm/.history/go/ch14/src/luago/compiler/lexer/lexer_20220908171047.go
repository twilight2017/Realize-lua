package lexer

import "bytes"
import "fmt"
import "regexp"
import "strconv"
import "strings"

type Lexer struct {
	chunk     string //源代码
	chunkName string //源文件名
	line      int    //当前行号
}

func NewLexer(chunk, chunkName string) *Lexer {

}
