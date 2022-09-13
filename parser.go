package rich

import (
	"errors"
)

type TokenType int

const (
	STRING TokenType = iota + 1
	OPEN_TAG
	CLOSE_TAG
)

type Token struct {
	Type    TokenType
	Literal string
}

func tokenize(s string) []Token {
	begin := 0
	end := 1
	var prev rune
	var tokens []Token
	for _, c := range s {
		if prev != '\\' {
			switch c {
			case '[', ']':
				lit := s[begin : end-1]
				if len(lit) != 0 {
					tokens = append(tokens, Token{Type: STRING, Literal: lit})
				}
				tokenType := OPEN_TAG
				if c == ']' {
					tokenType = CLOSE_TAG
				}
				tokens = append(tokens, Token{Type: tokenType, Literal: s[end-1 : end]})
				begin = end
			}
		}
		end += 1
		prev = c
	}
	lit := s[begin : end-1]
	if len(lit) != 0 {
		tokens = append(tokens, Token{Type: STRING, Literal: lit})
	}
	return tokens
}

type Part interface {
	part()
}

type OpeningTag string

func (p OpeningTag) String() string {
	return "[" + string(p) + "]"
}

func (p OpeningTag) part() {}

type ClosingTag string

func (p ClosingTag) String() string {
	return "[/" + string(p) + "]"
}

func (p ClosingTag) part() {}

type Text string

func (p Text) part() {}

func parse(tokens []Token) ([]Part, error) {
	var parts []Part
	inTag := false
	for _, token := range tokens {
		switch token.Type {
		case STRING:
			if inTag {
				if token.Literal[0] == '/' {

					parts = append(parts, ClosingTag(token.Literal[1:]))
				} else {
					parts = append(parts, OpeningTag(token.Literal))

				}
			} else {
				parts = append(parts, Text(token.Literal))
			}
		case OPEN_TAG:
			if inTag {
				return nil, errors.New("can't open tag inside another tag")
			} else {
				inTag = true
			}
		case CLOSE_TAG:
			if inTag {
				inTag = false
			} else {
				return nil, errors.New("can't close tag outside of tag")
			}
		}
	}
	return parts, nil
}

func parseString(s string) ([]Part, error) {
	tokens := tokenize(s)
	return parse(tokens)
}
