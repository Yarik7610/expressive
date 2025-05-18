package parser

import (
	"fmt"

	"github.com/Yarik7610/expressive/lexer"
)

// EBNF grammar:
// <expr> ::= <term>
// <term> ::= <factor> (("+" | "-") <factor>)*
// <factor> ::= <unary> (("*" | "/") <unary>)*
// <unary> ::= "-"? <unary> | <primary>
// <primary> ::= NUMBER | "(" <expr> ")"

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func NewParser(tokens []lexer.Token) *Parser {
	p := Parser{tokens, 0}
	return &p
}

func (p *Parser) Parse() []Node {
	nodes := make([]Node, 0)

	for !p.isEnd() {
		nodes = append(nodes, p.parseExpr())
	}

	return nodes
}

func (p *Parser) parseExpr() Node {
	return p.parseTerm()
}

func (p *Parser) parseTerm() Node {
	lhs := p.parseFactor()

	for p.match(lexer.TOKEN_PLUS, lexer.TOKEN_MINUS) {
		op := p.previous()
		rhs := p.parseFactor()
		lhs = &BinaryNode{Token: op, Left: lhs, Right: rhs}
	}

	return lhs
}

func (p *Parser) parseFactor() Node {
	lhs := p.parseUnary()

	for p.match(lexer.TOKEN_ASTERISK, lexer.TOKEN_SLASH) {
		op := p.previous()
		rhs := p.parseUnary()
		lhs = &BinaryNode{Token: op, Left: lhs, Right: rhs}
	}

	return lhs
}

func (p *Parser) parseUnary() Node {
	if p.match(lexer.TOKEN_MINUS) {
		return &UnaryNode{Token: p.previous(), Right: p.parseUnary()}
	}

	return p.parsePrimary()
}

func (p *Parser) parsePrimary() Node {
	if p.match(lexer.TOKEN_NUMBER) {
		return &NumberNode{p.previous()}
	}

	if p.match(lexer.TOKEN_BRACE_LEFT) {
		node := p.parseExpr()
		p.require(lexer.TOKEN_BRACE_RIGHT, "expected ')'")
		return node
	}

	panic("parser: expected number or expression")
}

func (p *Parser) match(tokenTypes ...int) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) require(tokenType int, errorMessage string) {
	if p.check(tokenType) {
		p.advance()
		return
	}

	panic(fmt.Sprintf("parser: %s", errorMessage))
}

func (p *Parser) check(tokenType int) bool {
	return p.peek().Type == tokenType
}

func (p *Parser) advance() lexer.Token {
	if !p.isEnd() {
		p.pos++
	}
	return p.previous()
}

func (p *Parser) peek() lexer.Token {
	return p.tokens[p.pos]
}

func (p *Parser) previous() lexer.Token {
	return p.tokens[p.pos-1]
}

func (p *Parser) isEnd() bool {
	return p.pos == len(p.tokens) || p.peek().Type == lexer.TOKEN_EOF
}

func PrintNodes(nodes []Node) {
	for _, node := range nodes {
		fmt.Println(node.String(0))
	}
}
