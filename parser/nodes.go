package parser

import (
	"fmt"
	"strings"

	"github.com/Yarik7610/expressive/lexer"
)

type Node interface {
	String(spaceCount int) string
}

type NumberNode struct {
	lexer.Token
}

func (nn *NumberNode) String(spaceCount int) string {
	spaceString := strings.Repeat(" ", spaceCount)
	return fmt.Sprint(spaceString, nn.Raw)
}

type BinaryNode struct {
	lexer.Token
	Left  Node
	Right Node
}

func (bn *BinaryNode) String(spaceCount int) string {
	spaceString := strings.Repeat(" ", spaceCount)
	return fmt.Sprint(spaceString, bn.Raw, "\n", spaceString, bn.Left.String(spaceCount+1), "\n", spaceString, bn.Right.String(spaceCount+1))
}

type UnaryNode struct {
	lexer.Token
	Right Node
}

func (un *UnaryNode) String(spaceCount int) string {
	spaceString := strings.Repeat(" ", spaceCount)
	return fmt.Sprint(spaceString, un.Raw, "\n", spaceString, un.Right.String(spaceCount+1))
}
