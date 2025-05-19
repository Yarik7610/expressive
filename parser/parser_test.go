package parser

import (
	"testing"

	"github.com/Yarik7610/expressive/lexer"
	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	nonPanicTests := []struct {
		Name string
		In   []lexer.Token
		Out  []Node
	}{
		{
			Name: "no tokens",
			In:   []lexer.Token{},
			Out:  []Node{},
		},
		{
			Name: "addition",
			In: []lexer.Token{
				{Type: lexer.TOKEN_NUMBER, Raw: "1"},
				{Type: lexer.TOKEN_PLUS, Raw: "+"},
				{Type: lexer.TOKEN_NUMBER, Raw: "3."},
				{Type: lexer.TOKEN_EOF, Raw: "TOKEN_EOF"},
			},
			Out: []Node{
				&BinaryNode{
					Token: lexer.Token{Type: lexer.TOKEN_PLUS, Raw: "+"},
					Left:  &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "1"}},
					Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "3."}},
				},
			},
		},
		{
			Name: "substraction and unary '-'",
			In: []lexer.Token{
				{Type: lexer.TOKEN_NUMBER, Raw: ".1"},
				{Type: lexer.TOKEN_MINUS, Raw: "-"},
				{Type: lexer.TOKEN_MINUS, Raw: "-"},
				{Type: lexer.TOKEN_NUMBER, Raw: ".3"},
				{Type: lexer.TOKEN_EOF, Raw: "TOKEN_EOF"},
			},
			Out: []Node{
				&BinaryNode{
					Token: lexer.Token{Type: lexer.TOKEN_MINUS, Raw: "-"},
					Left:  &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: ".1"}},
					Right: &UnaryNode{
						Token: lexer.Token{Type: lexer.TOKEN_MINUS, Raw: "-"},
						Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: ".3"}},
					},
				},
			},
		},
		{
			Name: "multiplication and substraction",
			In: []lexer.Token{
				{Type: lexer.TOKEN_MINUS, Raw: "-"},
				{Type: lexer.TOKEN_NUMBER, Raw: "2"},
				{Type: lexer.TOKEN_ASTERISK, Raw: "*"},
				{Type: lexer.TOKEN_MINUS, Raw: "-"},
				{Type: lexer.TOKEN_NUMBER, Raw: "3"},
				{Type: lexer.TOKEN_PLUS, Raw: "+"},
				{Type: lexer.TOKEN_MINUS, Raw: "-"},
				{Type: lexer.TOKEN_NUMBER, Raw: "4.2"},
				{Type: lexer.TOKEN_EOF, Raw: "TOKEN_EOF"},
			},
			Out: []Node{
				&BinaryNode{
					Token: lexer.Token{Type: lexer.TOKEN_PLUS, Raw: "+"},
					Left: &BinaryNode{
						Token: lexer.Token{Type: lexer.TOKEN_ASTERISK, Raw: "*"},
						Left: &UnaryNode{
							Token: lexer.Token{Type: lexer.TOKEN_MINUS, Raw: "-"},
							Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "2"}},
						},
						Right: &UnaryNode{
							Token: lexer.Token{Type: lexer.TOKEN_MINUS, Raw: "-"},
							Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "3"}},
						},
					},
					Right: &UnaryNode{
						Token: lexer.Token{Type: lexer.TOKEN_MINUS, Raw: "-"},
						Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "4.2"}},
					},
				},
			},
		},
		{
			Name: "division",
			In: []lexer.Token{
				{Type: lexer.TOKEN_NUMBER, Raw: "-2"},
				{Type: lexer.TOKEN_SLASH, Raw: "/"},
				{Type: lexer.TOKEN_NUMBER, Raw: "-3e5"},
				{Type: lexer.TOKEN_EOF, Raw: "TOKEN_EOF"},
			},
			Out: []Node{
				&BinaryNode{
					Token: lexer.Token{Type: lexer.TOKEN_SLASH, Raw: "/"},
					Left:  &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "-2"}},
					Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "-3e5"}},
				},
			},
		},
		{
			Name: "modulo",
			In: []lexer.Token{
				{Type: lexer.TOKEN_NUMBER, Raw: "2"},
				{Type: lexer.TOKEN_PERCENT, Raw: "%"},
				{Type: lexer.TOKEN_NUMBER, Raw: "3e5"},
				{Type: lexer.TOKEN_EOF, Raw: "TOKEN_EOF"},
			},
			Out: []Node{
				&BinaryNode{
					Token: lexer.Token{Type: lexer.TOKEN_PERCENT, Raw: "%"},
					Left:  &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "2"}},
					Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "3e5"}},
				},
			},
		},
		{
			Name: "pow and multiplication",
			In: []lexer.Token{
				{Type: lexer.TOKEN_NUMBER, Raw: "2"},
				{Type: lexer.TOKEN_ASTERISK, Raw: "*"},
				{Type: lexer.TOKEN_NUMBER, Raw: "4.2"},
				{Type: lexer.TOKEN_CARET, Raw: "^"},
				{Type: lexer.TOKEN_NUMBER, Raw: "3e1.1"},
				{Type: lexer.TOKEN_EOF, Raw: "TOKEN_EOF"},
			},
			Out: []Node{
				&BinaryNode{
					Token: lexer.Token{Type: lexer.TOKEN_ASTERISK, Raw: "*"},
					Left:  &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "2"}},
					Right: &BinaryNode{
						Token: lexer.Token{Type: lexer.TOKEN_CARET, Raw: "^"},
						Left:  &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "4.2"}},
						Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "3e1.1"}},
					},
				},
			},
		},
		{
			Name: "brackets",
			In: []lexer.Token{
				{Type: lexer.TOKEN_BRACE_LEFT, Raw: "("},
				{Type: lexer.TOKEN_NUMBER, Raw: "2"},
				{Type: lexer.TOKEN_PLUS, Raw: "+"},
				{Type: lexer.TOKEN_NUMBER, Raw: "1"},
				{Type: lexer.TOKEN_BRACE_RIGHT, Raw: ")"},
				{Type: lexer.TOKEN_ASTERISK, Raw: "*"},
				{Type: lexer.TOKEN_NUMBER, Raw: "4"},
				{Type: lexer.TOKEN_EOF, Raw: "TOKEN_EOF"},
			},
			Out: []Node{
				&BinaryNode{
					Token: lexer.Token{Type: lexer.TOKEN_ASTERISK, Raw: "*"},
					Left: &BinaryNode{
						Token: lexer.Token{Type: lexer.TOKEN_PLUS, Raw: "+"},
						Left:  &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "2"}},
						Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "1"}},
					},
					Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "4"}},
				},
			},
		},
	}

	panicTests := []struct {
		Name string
		In   []lexer.Token
	}{
		{
			Name: "unary '+'",
			In: []lexer.Token{
				{Type: lexer.TOKEN_NUMBER, Raw: ".1"},
				{Type: lexer.TOKEN_MINUS, Raw: "-"},
				{Type: lexer.TOKEN_PLUS, Raw: "+"},
				{Type: lexer.TOKEN_NUMBER, Raw: ".3e2"},
				{Type: lexer.TOKEN_EOF, Raw: "TOKEN_EOF"},
			},
		},
		{
			Name: "no closing bracket",
			In: []lexer.Token{
				{Type: lexer.TOKEN_BRACE_LEFT, Raw: "("},
				{Type: lexer.TOKEN_NUMBER, Raw: "-2"},
				{Type: lexer.TOKEN_PLUS, Raw: "+"},
				{Type: lexer.TOKEN_BRACE_LEFT, Raw: "("},
				{Type: lexer.TOKEN_NUMBER, Raw: "1"},
				{Type: lexer.TOKEN_BRACE_RIGHT, Raw: ")"},
				{Type: lexer.TOKEN_EOF, Raw: "TOKEN_EOF"},
			},
		},
		{
			Name: "no opening bracket",
			In: []lexer.Token{
				{Type: lexer.TOKEN_NUMBER, Raw: "-2"},
				{Type: lexer.TOKEN_PLUS, Raw: "+"},
				{Type: lexer.TOKEN_BRACE_LEFT, Raw: ")"},
				{Type: lexer.TOKEN_EOF, Raw: "TOKEN_EOF"},
			},
		},
		{
			Name: "no second operand",
			In: []lexer.Token{
				{Type: lexer.TOKEN_NUMBER, Raw: ".1"},
				{Type: lexer.TOKEN_MINUS, Raw: "^"},
				{Type: lexer.TOKEN_EOF, Raw: "TOKEN_EOF"},
			},
		},
	}

	for _, test := range nonPanicTests {
		t.Run(test.Name, func(t *testing.T) {
			out := NewParser(test.In).Parse()
			assert.EqualValues(t, test.Out, out)
		})
	}

	for _, test := range panicTests {
		t.Run(test.Name, func(t *testing.T) {
			assert.Panics(t, func() {
				NewParser(test.In).Parse()
			})
		})
	}
}
