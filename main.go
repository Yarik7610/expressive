package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/Yarik7610/expressive/lexer"
	"github.com/Yarik7610/expressive/parser"
)

func proccessFile(file *os.File) {
	var b bytes.Buffer

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			b.WriteString(line + "\n")
		} else {
			b.WriteString(fmt.Sprintf("%f\n", proccessString(line)))
		}
	}

	outputFile, err := os.Create("output.txt")
	if err != nil {
		panic(fmt.Sprintf("error creating output file: %s", err))
	}
	defer outputFile.Close()

	_, err = outputFile.Write(b.Bytes())
	if err != nil {
		panic(fmt.Sprintf("error writing to output file: %s", err))
	}
}

func proccessString(input string) float64 {
	l := lexer.NewLexer(strings.NewReader(input))
	tokens := l.Lex()

	p := parser.NewParser(tokens)
	nodes := p.Parse()

	return parser.Eval(nodes)
}

func main() {
	if len(os.Args) != 2 {
		panic("missing input expression")
	}
	input := os.Args[1]

	if file, err := os.Open(input); err == nil {
		defer file.Close()
		proccessFile(file)
	} else {
		fmt.Println(proccessString(input))
	}
}
