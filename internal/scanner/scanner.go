package scanner

import "fmt"

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

	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++

	default:
		return fmt.Errorf("unexpected character: %c", c)
	}
	return nil
}

func (s *Scanner) advance() byte { // 获取下一个字符
	//! 到末尾了怎么办?
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
