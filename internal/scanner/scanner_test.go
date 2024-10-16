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
