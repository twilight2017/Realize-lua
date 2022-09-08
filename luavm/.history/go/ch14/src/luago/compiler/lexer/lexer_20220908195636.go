package lexer

import "bytes"
import "fmt"
import "regexp"
import "strconv"
import "strings"

var reOpeningLongBracket = regexp.MustCompile(`^\[=*\[`)

type Lexer struct {
	chunk     string //源代码
	chunkName string //源文件名
	line      int    //当前行号
}

func NewLexer(chunk, chunkName string) *Lexer {
	return &Lexer{chunk, chunkName, 1}
}

func (self *Lexer) NextToken() (line, kind int, token string) {
	switch self.chunk[0] {
	case ';':
		self.next(1)
		return self.line, TOKEN_SEP_SEMI, ";"
	case ',':
		self.next(1)
		return self.line, TOKEN_SEP_COMMA, ","
	case '(':
		self.next(1)
		return self.line, TOKEN_SEP_LPAREN, "("
	case ')':
		self.next(1)
		return self.line, TOKEN_SEP_RPAREN, ")"
	case ']':
		self.next(1)
		return self.line, TOKEN_SEP_RBRACK, "]"
	case '{':
		self.next(1)
		return self.line, TOKEN_SEP_LCURLY, "{"
	case '}':
		self.next(1)
		return self.line, TOKEN_SEP_RCURLY, "}"
	case '+':
		self.next(1)
		return self.line, TOKEN_OP_ADD, "+"
	case '-':
		self.next(1)
		return self.line, TOKEN_OP_SUB, "-"
	case '*':
		self.next(1)
		return self.line, TOKEN_OP_MUL, "*"
	case '^':
		self.next(1)
		return self.line, TOKEN_OP_POW, "^"
	case '%':
		self.next(1)
		return self.line, TOKEN_OP_MOD, "%"
	case '&':
		self.next(1)
		return self.line, TOKEN_OP_BAND, "&"
	case '|':
		self.next(1)
		return self.line, TOKEN_OP_BOR, "|"
	case '#':
		self.next(1)
		return self.line, TOKEN_OP_LEN, "#"
	case ':':
		if self.test("::") {
			self.next(2)
			return self.line, TOKEN_SEP_LABEL, "::"
		} else {
			self.next(1)
			return self.line, TOKEN_SEP_COLON, ":"
		}
	case '/':
		if self.test("//") {
			self.next(2)
			return self.line, TOKEN_OP_IDIV, "//"
		} else {
			self.next(1)
			return self.line, TOKEN_OP_DIV, "/"
		}
	case '~':
		if self.test("~=") {
			self.next(2)
			return self.line, TOKEN_OP_NE, "~="
		} else {
			self.next(1)
			return self.line, TOKEN_OP_WAVE, "~"
		}
	case '=':
		if self.test("==") {
			self.next(2)
			return self.line, TOKEN_OP_EQ, "=="
		} else {
			self.next(1)
			return self.line, TOKEN_OP_ASSIGN, "="
		}
	case '<':
		if self.test("<=") {
			self.next(2)
			return self.line, TOKEN_OP_LE, "<="
		} else if self.test("<<") {
			self.next(2)
			return self.line, TOKEN_OP_SHL, "<<"
		} else {
			self.next(1)
			return self.line, TOKEN_OP_LT, "<"
		}
	case '>':
		if self.test(">=") {
			self.next(2)
			return self.line, TOKEN_OP_GE, ">="
		} else if self.test(">>") {
			self.next(2)
			return self.line, TOKEN_OP_SHR, ">>"
		} else {
			self.next(1)
			return self.line, TOKEN_OP_GT, ">"
		}
	case '.':
		if self.test("..."){
			self.next(3)
			return self.line, TOKEN_VARARG, "..."
		}else if self.test(".."){
			self.next(2)
			return self.line, TOKEN_OP_CONCAT, ".."
		}

}

func (self *Lexer) skipWhiteSpace() {
	for len(self.chunk) > 0 {
		if self.test("--") {
			self.skipComment()
		} else if self.test("\r\n") || self.test("\n\r") {
			self.next(2)
			self.line += 1
		} else if isNewLine(self.chunk[0]) {
			self.next(1)
			self.line += 1
		} else if isWhiteSpace(self.chunk[0]) {
			self.next(1)
		} else {
			break
		}
	}
}

func (self *Lexer) test(s string) bool {
	return strings.HasPrefix(self.chunk, s)
}

//跳过n个字符
func (self *Lexer) next(n int) {
	self.chunk = self.chunk[n:]
}

//判断字符是否是空白字符
func isWhiteSpace(c byte) bool {
	switch c {
	case '\t', '\n', '\v', 'f', '\r', ' ':
		return true
	}
	return false
}

//判断字符是否是回车或者换行符
func isNewLine(c byte) bool {
	return c == '\r' || c == '\n'
}

func (self *Lexer) skipComment() {
	self.next(2)
	if self.test("[") {
		//long comment
		if reOpeningLongBracket.FindString(self.chunk) != "" {
			self.scanLongString()
			return
		}
	}
	//short comment,短字符串换行符前面的部分都要跳过
	for len(self.chunk) > 0 && !isNewLine(self.chunk[0]) {
		self.next(1)
	}
}
