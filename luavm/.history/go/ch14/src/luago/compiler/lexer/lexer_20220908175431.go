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
	return &Lexer{chunk, chunkName, 1}
}

func (self *Lexer) NextToken() (line, kind int, token string) {
	self.skipWhiteSpace()
	if len(self.chunk) == 0 {
		//返回表示分析结束的特殊Token
		return self.line, TOKEN_EOF, "EOF"
	}
	switch self.chunk[0] {
	case ';':
		self.next(1)
		return self.line, TOKEN_SEP_SEMI, ""
	case ',':
		self.next(1)
		return self.line, TOKEN_SEP_COMMA, ""
	}
}

func skipWhiteSpace(){
	for len(self.chunk) > 0{
		if self.test("--"){
			self.skipComment()
		}else if self.test("\r\n") || self.test("\n\r"){
			self.next(2)
			self.line += 1
		}
		else if isNewLine(self.chunk[0]){
			self.next(1)
			self.line += 1
		}else if isWhiteSpace(self.chunk[0]){
			self.next(1)
		}
	}
}
