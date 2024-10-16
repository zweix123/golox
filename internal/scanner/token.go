package scanner

import "fmt"

type TokenType string

const (
	// Single-character tokens.
	LeftParen  TokenType = "("
	RightParen TokenType = ")"
	LeftBrace  TokenType = "{"
	RightBrace TokenType = "}"
	Comma      TokenType = ","
	Dot        TokenType = "."
	Minus      TokenType = "-"
	Plus       TokenType = "+"
	Semicolon  TokenType = ";"
	Slash      TokenType = "/"
	Star       TokenType = "*"

	// One or two character tokens.
	Bang         TokenType = "!"
	BangEqual    TokenType = "!="
	Equal        TokenType = "="
	EqualEqual   TokenType = "=="
	Greater      TokenType = ">"
	GreaterEqual TokenType = ">="
	Less         TokenType = "<"
	LessEqual    TokenType = "<="

	// Literals.
	Identifier TokenType = "Identifier" // [a-zA-Z_][a-zA-Z0-9_]*
	String     TokenType = "String"
	Number     TokenType = "Number"

	// Keywords.
	And    TokenType = "and"
	Class  TokenType = "class"
	Else   TokenType = "else"
	False  TokenType = "false"
	Fun    TokenType = "fun"
	For    TokenType = "for"
	If     TokenType = "if"
	Nil    TokenType = "nil"
	Or     TokenType = "or"
	Print  TokenType = "print"
	Return TokenType = "return"
	Super  TokenType = "super"
	This   TokenType = "this"
	True   TokenType = "true"
	Var    TokenType = "var"
	While  TokenType = "while"

	EOF TokenType = "EOF"
)

type Token struct {
	Type    TokenType // token类型
	Lexeme  string    // 对应词素
	Literal any       // 字面量
	Line    int       // 所在行数
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) *Token {
	return &Token{
		Type:    tokenType,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("[%s %s %v]", t.Type, t.Lexeme, t.Literal)
}
