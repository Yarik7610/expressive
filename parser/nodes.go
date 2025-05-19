package parser

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Yarik7610/expressive/lexer"
)

type Node interface {
	Eval() float64
	String(spaceCount int) string
}

type NumberNode struct {
	lexer.Token
}

func (nn *NumberNode) String(spaceCount int) string {
	spaceString := strings.Repeat(" ", spaceCount)
	return fmt.Sprint(spaceString, nn.Raw)
}

func (nn *NumberNode) Eval() float64 {
	val, err := strconv.ParseFloat(nn.Raw, 64)
	if err != nil {
		panic(fmt.Sprintf("eval: number node error: %s", err))
	}
	return val
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

func (bn *BinaryNode) Eval() float64 {
	switch bn.Token.Type {
	case lexer.TOKEN_PLUS:
		return bn.Left.Eval() + bn.Right.Eval()
	case lexer.TOKEN_MINUS:
		return bn.Left.Eval() - bn.Right.Eval()
	case lexer.TOKEN_SLASH:
		return bn.Left.Eval() / bn.Right.Eval()
	case lexer.TOKEN_ASTERISK:
		return bn.Left.Eval() * bn.Right.Eval()
	case lexer.TOKEN_PERCENT:
		return math.Mod(bn.Left.Eval(), bn.Right.Eval())
	case lexer.TOKEN_CARET:
		return math.Pow(bn.Left.Eval(), bn.Right.Eval())
	default:
		panic("eval: binary node error: undefined operator")
	}
}

type UnaryNode struct {
	lexer.Token
	Right Node
}

func (un *UnaryNode) String(spaceCount int) string {
	spaceString := strings.Repeat(" ", spaceCount)
	return fmt.Sprint(spaceString, un.Raw, "\n", spaceString, un.Right.String(spaceCount+1))
}

func (un *UnaryNode) Eval() float64 {
	return -1 * un.Right.Eval()
}
