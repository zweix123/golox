package scanner

import (
	"reflect"
	"testing"
)

func TestScanTokens(t *testing.T) {
	// 表测试法
	tests := []struct {
		input string
		want  []Token
	}{
		{input: "(", want: []Token{{Type: LeftParen, Lexeme: "(", Line: 1}, {Type: EOF, Line: 1}}},
		{input: ")", want: []Token{{Type: RightParen, Lexeme: ")", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "{", want: []Token{{Type: LeftBrace, Lexeme: "{", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "}", want: []Token{{Type: RightBrace, Lexeme: "}", Line: 1}, {Type: EOF, Line: 1}}},
		{input: ",", want: []Token{{Type: Comma, Lexeme: ",", Line: 1}, {Type: EOF, Line: 1}}},
		{input: ".", want: []Token{{Type: Dot, Lexeme: ".", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "-", want: []Token{{Type: Minus, Lexeme: "-", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "+", want: []Token{{Type: Plus, Lexeme: "+", Line: 1}, {Type: EOF, Line: 1}}},
		{input: ";", want: []Token{{Type: Semicolon, Lexeme: ";", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "*", want: []Token{{Type: Star, Lexeme: "*", Line: 1}, {Type: EOF, Line: 1}}},

		{input: "(){},.-+;*", want: []Token{{Type: LeftParen, Lexeme: "(", Line: 1}, {Type: RightParen, Lexeme: ")", Line: 1}, {Type: LeftBrace, Lexeme: "{", Line: 1}, {Type: RightBrace, Lexeme: "}", Line: 1}, {Type: Comma, Lexeme: ",", Line: 1}, {Type: Dot, Lexeme: ".", Line: 1}, {Type: Minus, Lexeme: "-", Line: 1}, {Type: Plus, Lexeme: "+", Line: 1}, {Type: Semicolon, Lexeme: ";", Line: 1}, {Type: Star, Lexeme: "*", Line: 1}, {Type: EOF, Line: 1}}},

		{input: "!", want: []Token{{Type: Bang, Lexeme: "!", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "!=", want: []Token{{Type: BangEqual, Lexeme: "!=", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "=", want: []Token{{Type: Equal, Lexeme: "=", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "==", want: []Token{{Type: EqualEqual, Lexeme: "==", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "<", want: []Token{{Type: Less, Lexeme: "<", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "<=", want: []Token{{Type: LessEqual, Lexeme: "<=", Line: 1}, {Type: EOF, Line: 1}}},
		{input: ">", want: []Token{{Type: Greater, Lexeme: ">", Line: 1}, {Type: EOF, Line: 1}}},
		{input: ">=", want: []Token{{Type: GreaterEqual, Lexeme: ">=", Line: 1}, {Type: EOF, Line: 1}}},

		{input: "=====", want: []Token{{Type: EqualEqual, Lexeme: "==", Line: 1}, {Type: EqualEqual, Lexeme: "==", Line: 1}, {Type: Equal, Lexeme: "=", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "!!===<<=>>=", want: []Token{{Type: Bang, Lexeme: "!", Line: 1}, {Type: BangEqual, Lexeme: "!=", Line: 1}, {Type: EqualEqual, Lexeme: "==", Line: 1}, {Type: Less, Lexeme: "<", Line: 1}, {Type: LessEqual, Lexeme: "<=", Line: 1}, {Type: Greater, Lexeme: ">", Line: 1}, {Type: GreaterEqual, Lexeme: ">=", Line: 1}, {Type: EOF, Line: 1}}},

		{input: "/", want: []Token{{Type: Slash, Lexeme: "/", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "//", want: []Token{{Type: EOF, Line: 1}}},
		{input: "//comment", want: []Token{{Type: EOF, Line: 1}}},
		{input: "//comment\n", want: []Token{{Type: EOF, Line: 2}}},
		{input: "//comment\n\n", want: []Token{{Type: EOF, Line: 3}}},
		{input: "//comment\n\n//comment", want: []Token{{Type: EOF, Line: 3}}},

		// 空白
		{input: " ", want: []Token{{Type: EOF, Line: 1}}},
		{input: "\r", want: []Token{{Type: EOF, Line: 1}}},
		{input: "\t", want: []Token{{Type: EOF, Line: 1}}},
		{input: "\n", want: []Token{{Type: EOF, Line: 2}}},

		// 字符串
		{input: "\"hello\"", want: []Token{{Type: String, Lexeme: "\"hello\"", Literal: "hello", Line: 1}, {Type: EOF, Line: 1}}},

		// 数字
		{input: "123", want: []Token{{Type: Number, Lexeme: "123", Literal: float64(123), Line: 1}, {Type: EOF, Line: 1}}},
		{input: "123.456", want: []Token{{Type: Number, Lexeme: "123.456", Literal: float64(123.456), Line: 1}, {Type: EOF, Line: 1}}},
		{input: "0.123", want: []Token{{Type: Number, Lexeme: "0.123", Literal: float64(0.123), Line: 1}, {Type: EOF, Line: 1}}},
		{input: ".456", want: []Token{{Type: Dot, Lexeme: ".", Line: 1}, {Type: Number, Lexeme: "456", Literal: float64(456), Line: 1}, {Type: EOF, Line: 1}}},
		{input: "123.", want: []Token{{Type: Number, Lexeme: "123", Literal: float64(123), Line: 1}, {Type: Dot, Lexeme: ".", Line: 1}, {Type: EOF, Line: 1}}},

		// 关键字
		{input: "and", want: []Token{{Type: And, Lexeme: "and", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "class", want: []Token{{Type: Class, Lexeme: "class", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "else", want: []Token{{Type: Else, Lexeme: "else", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "false", want: []Token{{Type: False, Lexeme: "false", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "for", want: []Token{{Type: For, Lexeme: "for", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "fun", want: []Token{{Type: Fun, Lexeme: "fun", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "if", want: []Token{{Type: If, Lexeme: "if", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "nil", want: []Token{{Type: Nil, Lexeme: "nil", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "or", want: []Token{{Type: Or, Lexeme: "or", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "print", want: []Token{{Type: Print, Lexeme: "print", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "return", want: []Token{{Type: Return, Lexeme: "return", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "super", want: []Token{{Type: Super, Lexeme: "super", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "this", want: []Token{{Type: This, Lexeme: "this", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "true", want: []Token{{Type: True, Lexeme: "true", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "var", want: []Token{{Type: Var, Lexeme: "var", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "while", want: []Token{{Type: While, Lexeme: "while", Line: 1}, {Type: EOF, Line: 1}}},

		// 标识符
		{input: "foo", want: []Token{{Type: Identifier, Lexeme: "foo", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "bar", want: []Token{{Type: Identifier, Lexeme: "bar", Line: 1}, {Type: EOF, Line: 1}}},
		{input: "baz", want: []Token{{Type: Identifier, Lexeme: "baz", Line: 1}, {Type: EOF, Line: 1}}},
	}
	for _, test := range tests {

		scanner := NewScanner(test.input)
		tokens, err := scanner.ScanTokens()
		if err != nil {
			t.Errorf("ScanTokens error: %v", err)
		}
		if !reflect.DeepEqual(tokens, test.want) {
			t.Errorf("ScanTokens(%q) = %v, want %v", test.input, tokens, test.want)
		}
	}
}
