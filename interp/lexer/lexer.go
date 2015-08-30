package lexer

import (
	"bytes"
	"io"
	"log"
)

const nilRune = rune(-1)

type lexer struct {
	name       string
	start      Position
	pos        Position
	input      io.RuneScanner
	state      stateFn
	leftDelim  rune
	rightDelim rune
	tokens     chan token
	buf        *bytes.Buffer
}

type stateFn func(*lexer) stateFn

// New returns new lexer struct instance
func New(r io.RuneScanner) (l *lexer) {

	l = &lexer{
		name:   "",
		input:  r,
		state:  nil,
		tokens: make(chan token),
		buf:    new(bytes.Buffer),
	}

	go l.run()

	return
}

/*
================================================================================
Lexer APIs
================================================================================
*/

// NextToken returns next obtaining token
// This API is blocking api.
func (l *lexer) NextToken() token {

	token := <-l.tokens

	return token
}

// GetTokenCh returns token incoming channel
func (l *lexer) GetTokenCh() <-chan token {

	return l.tokens
}

/*
================================================================================
Lexer private methods
================================================================================
*/

func (l *lexer) run() {

	for l.state = lex; l.state != nil; {
		l.state = l.state(l)
	}
}

func (l *lexer) emit(t tokenType) {

	val := l.buf.String()

	if val != "<nil>" {
		l.tokens <- token{
			Typ: t,
			Pos: l.start,
			Val: val,
		}
	}

	l.buf.Reset()
}

func (l *lexer) peek() rune {

	var (
		r   rune  = l.next()
		err error = l.input.UnreadRune()
	)

	if err != nil {
		return nilRune
	}

	l.pos -= 1

	return r
}

func (l *lexer) next() rune {

	var (
		r   rune
		err error
	)

	if r, _, err = l.input.ReadRune(); err != nil {
		return nilRune
	}

	l.pos += 1

	return r
}

/*
================================================================================
lexer functions
================================================================================
*/

func lex(l *lexer) stateFn {

	switch r := l.peek(); {

	case r == '(':
		l.emit(TLeftParen)
		l.next()

	case r == ')':
		l.emit(TRightParen)
		l.next()

	case r == '{':
		l.emit(TLeftSquareBracket)
		l.next()

	case r == '}':
		l.emit(TRightSquareBracket)
		l.next()

	case r == ':':
		l.emit(TDblColon)
		l.next()

	case r == ';':
		l.emit(TSemiColon)
		l.next()

	case r == '.':
		return lexDotPrefix

	case r == '"' || r == '`':
		l.leftDelim = r
		return lexString

	case isNumber(r):
		return lexNumber

	case isSpace(r):
		return lexSpace

	case isLetter(r):
		return lexIdentifier

	case r == nilRune:
		l.emit(TEOF)
		return nil

	default:
		log.Printf("no matching: %s", r)

	}

	return lex
}

func lexIdentifier(l *lexer) stateFn {

	for r := l.peek(); isLetter(r); r = l.peek() {
		l.buf.WriteRune(r)
		l.next()
	}

	l.emit(TIdentifier)

	return lex
}

func lexSpace(l *lexer) stateFn {

	for isSpace(l.peek()) {
		l.next()
	}

	l.emit(TSpace)

	return lex
}

func lexString(l *lexer) stateFn {

	// skip the left delimiter
	l.next()

	for r := l.peek(); r != l.leftDelim; r = l.peek() {
		l.buf.WriteRune(r)
		l.next()
	}

	// skip the right delimiter
	l.next()

	l.emit(TString)

	return lex
}

func lexNumber(l *lexer) stateFn {

	for r := l.peek(); isNumber(r) || r == '.'; r = l.peek() {
		l.buf.WriteRune(r)
		l.next()
	}

	l.emit(TNumber)

	return lex
}

func lexDotPrefix(l *lexer) stateFn {

	// skip the first dot '.'
	l.next()

	switch r := l.peek(); {

	case r == 's':
		l.emit(TDotS)
		l.next()

	case r == nilRune:
		l.emit(TDot)
	}

	return lex
}

/*
================================================================================
Rune check functions
================================================================================
*/

func isNumber(r rune) bool {

	return '0' <= r && r <= '9'
}

func isLetter(r rune) bool {

	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '-' || r == '_' || r == '?' || r == '!'
}

func isSpace(r rune) bool {

	return r == ' ' || r == '\t' || r == '\n' || r == '\v' || r == '\f' || r == '\r'
}
