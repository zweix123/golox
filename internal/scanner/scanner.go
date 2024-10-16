package scanner

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	source  string
	start   int // 被扫描的词素的第一个字符的偏移量
	current int // 当前正在处理的字符的偏移量
	line    int

	tokens []Token // 维护结果
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() ([]Token, error) {
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()
		if err != nil {
			return nil, fmt.Errorf("scan token error: %w", err)
		}
	}
	s.tokens = append(s.tokens, Token{Type: EOF, Line: s.line})
	return s.tokens, nil
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) addToken(tokenType TokenType, literal any) {
	s.tokens = append(s.tokens, Token{
		Type:    tokenType,
		Lexeme:  s.source[s.start:s.current],
		Literal: literal,
		Line:    s.line,
	})
}

func (s *Scanner) scanToken() error { // 每次扫描, 返回一个token
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LeftParen, nil)
	case ')':
		s.addToken(RightParen, nil)
	case '{':
		s.addToken(LeftBrace, nil)
	case '}':
		s.addToken(RightBrace, nil)
	case ',':
		s.addToken(Comma, nil)
	case '.':
		s.addToken(Dot, nil)
	case '-':
		s.addToken(Minus, nil)
	case '+':
		s.addToken(Plus, nil)
	case '*':
		s.addToken(Star, nil)
	// 除法比较特殊
	case ';':
		s.addToken(Semicolon, nil)

	// 这里是一种最短匹配, 即看很小的几位就确定其含义, 这对符号来说是可以的,
	// 但是对于标识符来说, 就不行, 标识符是最长匹配, 即一个完成的token读完, 再决定.
	case '!':
		if s.match('=') {
			s.addToken(BangEqual, nil)
		} else {
			s.addToken(Bang, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(EqualEqual, nil)
		} else {
			s.addToken(Equal, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(LessEqual, nil)
		} else {
			s.addToken(Less, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(GreaterEqual, nil)
		} else {
			s.addToken(Greater, nil)
		}

	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(Slash, nil)
		}

	// TODO: 添加对/**/注释的支持

	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++

	case '"':
		err := s.string()
		if err != nil {
			return fmt.Errorf("string error: %w", err)
		}

	default:
		if isDigit(c) {
			err := s.number()
			if err != nil {
				return fmt.Errorf("number error: %w", err)
			}
		} else if isAlpha(c) {
			s.identifier()
		} else {
			return fmt.Errorf("unexpected character: %c", c)
		}

	}
	return nil
}

/*
 * 三个动作
 * 1. 前进
 * 2. 先看再决策是否前进
 * 3. 只看不前进
 * s.currentu是指向当前处理的一下个的
 */

func (s *Scanner) advance() byte { // 获取下一个字符
	//! 到末尾了怎么办?使用了字符串末尾的'\000'了么?
	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) match(expected byte) bool { // 匹配并(假如匹配成功则)前进, 相当于先试再前进
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

/*
 * lookahead前瞻: 语法需要前瞻的字符越少, scanner越快
 */

func (s *Scanner) peek() byte { // 获取下一个字符(而不前进)
	if s.isAtEnd() {
		return '\000'
	}

	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return '\000'
	}
	return s.source[s.current+1]
}

/*
 * 字符串
 */

func (s *Scanner) string() error {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' { // Lox支持多行字符串
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		return fmt.Errorf("unterminated string")
	}

	s.advance() // the closing "

	s.addToken(String, s.source[s.start+1:s.current-1])
	return nil
}

func (s *Scanner) number() error {
	for isDigit(s.peek()) {
		s.advance()
	}

	// Look for a fractional part.
	if s.peek() == '.' && isDigit(s.peekNext()) { // 家涮Lox将万物视为对象的, 怎么处理123.sqrt()这样的语法?
		s.advance() // consume the "."

		for isDigit(s.peek()) {
			s.advance()
		}
	}
	f, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		return fmt.Errorf("parse float error: %w", err)
	}
	s.addToken(Number, f)
	return nil
}

func (s *Scanner) identifier() {
	for isAlpha(s.peek()) || isDigit(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	tokenType, ok := keywords[text]
	if !ok {
		tokenType = Identifier // 假如不在关键字(保留子)中, 则是标识符
	}
	s.addToken(tokenType, nil)
}
