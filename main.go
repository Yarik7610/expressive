package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Yarik7610/expressive/lexer"
	"github.com/Yarik7610/expressive/parser"
)

func main() {
	if len(os.Args) != 2 {
		panic("missing input expression")
	}
	input := os.Args[1]

	l := lexer.NewLexer(strings.NewReader(input))
	tokens := l.Lex()

	p := parser.NewParser(tokens)
	nodes := p.Parse()

	fmt.Println(parser.Eval(nodes))
}
