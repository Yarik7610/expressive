package parser

import (
	"testing"

	"github.com/Yarik7610/expressive/lexer"
	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	nonPanicTests := []struct {
		Name string
		In   []Node
		Out  float64
	}{
		{
			Name: "1 + 3.",
			In: []Node{
				&BinaryNode{
					Token: lexer.Token{Type: lexer.TOKEN_PLUS, Raw: "+"},
					Left:  &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "1"}},
					Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "3."}},
				},
			},
			Out: 4,
		},
		{
			Name: ".1 - -.3",
			In: []Node{
				&BinaryNode{
					Token: lexer.Token{Type: lexer.TOKEN_MINUS, Raw: "-"},
					Left:  &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: ".1"}},
					Right: &UnaryNode{
						Token: lexer.Token{Type: lexer.TOKEN_MINUS, Raw: "-"},
						Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: ".3"}},
					},
				},
			},
			Out: 0.4,
		},
		{
			Name: "2 ^ 1e+1",
			In: []Node{
				&BinaryNode{
					Token: lexer.Token{Type: lexer.TOKEN_CARET, Raw: "^"},
					Left:  &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "2"}},
					Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "1e+1"}},
				},
			},
			Out: 1024,
		},
		{
			Name: "2 * 1e-1_0 ^ 2",
			In: []Node{
				&BinaryNode{
					Token: lexer.Token{Type: lexer.TOKEN_ASTERISK, Raw: "*"},
					Left:  &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "2"}},
					Right: &BinaryNode{
						Token: lexer.Token{Type: lexer.TOKEN_CARET, Raw: "^"},
						Left:  &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "1e-1_0"}},
						Right: &NumberNode{Token: lexer.Token{Type: lexer.TOKEN_NUMBER, Raw: "2"}},
					},
				},
			},
			Out: 2.0000000000000002e-20,
		},
	}

	for _, test := range nonPanicTests {
		t.Run(test.Name, func(t *testing.T) {
			out := Eval(test.In)
			assert.Equal(t, test.Out, out)
		})
	}
}
