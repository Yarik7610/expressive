package main

import (
	"os"
	"strings"

	"github.com/Yarik7610/expressive/lexer"
)

func main() {
	if len(os.Args) != 2 {
		panic("missing input expression")
	}
	input := os.Args[1]

	l := lexer.NewLexer(strings.NewReader(input))

	tokens := l.Lex()
	lexer.PrintTokens(tokens)
}
