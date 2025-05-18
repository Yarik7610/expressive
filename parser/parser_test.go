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
			Name: "No tokens",
			In:   []lexer.Token{},
			Out:  []Node{},
		},
	}

	for _, test := range nonPanicTests {
		t.Run(test.Name, func(t *testing.T) {
			out := NewParser(test.In).Parse()
			assert.EqualValues(t, test.Out, out)
		})
	}
}
