package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"luthor/token"
)

type result struct {
	s string
	t token.Token
}

func TestLuthor(t *testing.T) {

	var index int

	expectedResults := []result{
		{s: "car", t: token.VARIABLE},
		{s: " ", t: token.WS},
		{s: "=", t: token.EQ},
		{s: " ", t: token.WS},
		{s: "123.456", t: token.NUMBER},
		{s: " ", t: token.WS},
		{s: "+", t: token.ADD},
		{s: " ", t: token.WS},
		{s: "bus", t: token.VARIABLE},
		{s: string(eof), t: token.EOF},
	}

	reader := strings.NewReader("car = 123.456 + bus")

	l := NewLuthor(reader)

	for {
		expected := expectedResults[index]

		tok, str := l.Lex()

		assert.Equal(t, expected.s, str, "invalid string")
		assert.Equal(t, expected.t.String(), tok.String(), "invalid token")

		index++

		if tok == token.EOF || tok == token.ILLEGAL {
			break
		}
	}
}
