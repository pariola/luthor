package main

import (
	"bufio"
	"bytes"
	"io"
	"unicode"

	"luthor/token"
)

const (
	eof = rune(0)
)

type Luthor struct {
	r *bufio.Reader
}

func NewLuthor(r io.Reader) *Luthor {
	return &Luthor{r: bufio.NewReader(r)}
}

func (l Luthor) unread() {
	_ = l.r.UnreadRune()
}

func (l Luthor) read() rune {

	r, _, err := l.r.ReadRune()

	if err != nil {
		return eof
	}

	return r
}

func (l Luthor) scanWhitespaces() (token.Token, string) {

	var b bytes.Buffer

	for {
		c := l.read()

		if !unicode.IsSpace(c) {
			l.unread()
			break
		}

		b.WriteRune(c)
	}

	return token.WS, b.String()
}

func (l Luthor) scanVariable() (token.Token, string) {

	var b bytes.Buffer

	for {
		c := l.read()

		if !unicode.IsLetter(c) {
			l.unread()
			break
		}

		b.WriteRune(c)
	}

	return token.VARIABLE, b.String()
}

func (l Luthor) scanNumber() (token.Token, string) {

	var b bytes.Buffer
	var hasPoint bool

	for {
		c := l.read()

		if !hasPoint && c == '.' {
			hasPoint = true
		} else if !unicode.IsNumber(c) {
			l.unread()
			break
		}

		b.WriteRune(c)
	}

	return token.NUMBER, b.String()
}

func (l Luthor) Lex() (token.Token, string) {

	c := l.read()
	str := string(c)

	if unicode.IsSpace(c) {
		l.unread()
		return l.scanWhitespaces()
	} else if unicode.IsNumber(c) {
		l.unread()
		return l.scanNumber()
	} else if unicode.IsLetter(c) {
		l.unread()
		return l.scanVariable()
	}

	switch c {
	case eof:
		return token.EOF, str
	case '=':
		return token.EQ, str
	case '+':
		return token.ADD, str
	case '-':
		return token.SUB, str
	case '*':
		return token.MUL, str
	case '/':
		return token.QUO, str
	}

	return token.ILLEGAL, str
}
