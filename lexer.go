package main

import (
	"strconv"
	"strings"
	"unicode"
)

const (
	_             = iota
	PLUS          // +
	MINUS         // -
	DIV           // /
	MUL           // *
	LPAREN        // (
	RPAREN        // )
	INTEGER       // variable type
	EOF           // end of file
	BEGIN         // reserved keyword
	END           // reserved keyword
	DOT           // .
	ASSIGN        // :=
	SEMI          // ;
	ID            // variable token type
	PROGRAM       // reserved keyword
	VAR           // reserved keyword
	COLON         // :
	COMMA         // ,
	REAL          // variable type
	INTEGER_CONST // 4,3, etc
	REAL_CONST    // 3.14 etc
	INTEGER_DIV   // for integer division (DIV keyword)
	FLOAT_DIV     // for float division ( forward slash / )
	PROCEDURE     // reserved keywoed
)

var RESERVED_KEYWORDS = map[string]Token{
	"PROGRAM": {PROGRAM, "PROGRAM"},
	"VAR":     {VAR, "VAR"},
	"DIV":     {INTEGER_DIV, "DIV"},
	"INTEGER": {INTEGER, "INTEGER"},
	"REAL":    {REAL, "REAL"},
	"BEGIN":   {BEGIN, "BEGIN"},
	"END":     {END, "END"},
	"PROCEDURE" : {PROCEDURE, "PROCEDURE"},
}

type Token struct {
	_type int
	val   any
}

type Lexer struct {
	text        string
	pos         int
	currentChar rune
}

func NewLexer(text string) *Lexer {
	return &Lexer{text: text, currentChar: rune(text[0])}
}

func (l *Lexer) advance() {
	l.pos++
	if l.pos >= len(l.text) {
		l.currentChar = 0
		return
	}
	l.currentChar = rune(l.text[l.pos])
}

func (l *Lexer) peek() rune {
	peek_pos := l.pos + 1
	if peek_pos >= len(l.text) {
		return 0
	}
	return rune(l.text[peek_pos])
}

func (l *Lexer) _id() Token {
	result := ""
	for unicode.IsDigit(l.currentChar) || unicode.IsLetter(l.currentChar) || l.currentChar == '_' {
		result += string(l.currentChar)
		l.advance()
	}
	token, ok := RESERVED_KEYWORDS[strings.ToUpper(result)]
	if !ok {
		token = Token{ID, result}
	}
	return token
}

func (l *Lexer) skipWhitespaces() {
	for l.currentChar == ' ' || l.currentChar == '\n' || l.currentChar == '\t' {
		l.advance()
	}
}

func (l *Lexer) skipComment() {
	for l.currentChar != '}' {
		l.advance()
	}
	l.advance()
}

func (l *Lexer) number() Token {
	temp := ""
	for unicode.IsDigit(l.currentChar) {
		temp += string(l.currentChar)
		l.advance()
	}
	if l.currentChar == '.' {
		temp += "."
		l.advance()

		for unicode.IsDigit(l.currentChar) {
			temp += string(l.currentChar)
			l.advance()
		}
		res, _ := strconv.ParseFloat(temp, 64)
		return Token{REAL_CONST, res}
	}
	res, _ := strconv.Atoi(temp)
	return Token{INTEGER_CONST, float64(res)}
}

func (l *Lexer) getNextToken() Token {
	for l.currentChar != 0 {
		if l.currentChar == ' ' || l.currentChar == '\n' || l.currentChar == '\t' {
			l.skipWhitespaces()
			continue
		}
		if unicode.IsDigit(l.currentChar) {
			return l.number()
		}
		if unicode.IsLetter(l.currentChar) || l.currentChar == '_' {
			return l._id()
		}
		if l.currentChar == ':' && l.peek() == '=' {
			l.advance()
			l.advance()
			return Token{ASSIGN, ":="}
		}
		switch l.currentChar {
		case ';':
			l.advance()
			return Token{SEMI, "-"}
		case '.':
			l.advance()
			return Token{DOT, "."}
		case '-':
			l.advance()
			return Token{MINUS, "-"}
		case '+':
			l.advance()
			return Token{PLUS, "+"}
		case '*':
			l.advance()
			return Token{MUL, "*"}
		case '(':
			l.advance()
			return Token{LPAREN, "("}
		case ')':
			l.advance()
			return Token{RPAREN, ")"}
		case '{':
			l.skipComment()
			continue
		case ':':
			l.advance()
			return Token{COLON, ':'}
		case ',':
			l.advance()
			return Token{COMMA, ','}
		case '/':
			l.advance()
			return Token{FLOAT_DIV, '/'}
		}
		panic("Unrecognized token" + string(l.currentChar))
	}
	return Token{EOF, nil}
}
