package lexer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	nonPanicTests := []struct {
		Name string
		In   string
		Out  []Token
	}{
		{
			Name: "empty input",
			In:   "",
			Out:  []Token{{TOKEN_EOF, "TOKEN_EOF"}},
		},
		{
			Name: "ignoring whitespaces",
			In:   "\r\t\n     ",
			Out:  []Token{{TOKEN_EOF, "TOKEN_EOF"}},
		},
		{
			Name: "ignoring comments",
			In:   "# this is a comment\n# this is a comment without a newline at the end",
			Out:  []Token{{TOKEN_EOF, "TOKEN_EOF"}},
		},
		{
			Name: "operator and separator tokens lex",
			In:   "+-/*^%()",
			Out: []Token{
				{TOKEN_PLUS, "+"},
				{TOKEN_MINUS, "-"},
				{TOKEN_SLASH, "/"},
				{TOKEN_ASTERISK, "*"},
				{TOKEN_CARET, "^"},
				{TOKEN_PERCENT, "%"},
				{TOKEN_BRACE_LEFT, "("},
				{TOKEN_BRACE_RIGHT, ")"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			Name: "number",
			In:   "123",
			Out:  []Token{{TOKEN_NUMBER, "123"}, {TOKEN_EOF, "TOKEN_EOF"}},
		},
		{
			Name: "underscore number",
			In:   "123_000",
			Out:  []Token{{TOKEN_NUMBER, "123_000"}, {TOKEN_EOF, "TOKEN_EOF"}},
		},
		{
			Name: "number with e and underscore in power",
			In:   "123e2_0",
			Out:  []Token{{TOKEN_NUMBER, "123e2_0"}, {TOKEN_EOF, "TOKEN_EOF"}},
		},
		{
			Name: "number with e and next '+'",
			In:   "123e+2",
			Out:  []Token{{TOKEN_NUMBER, "123e+2"}, {TOKEN_EOF, "TOKEN_EOF"}},
		},
		{
			Name: "number with e and next '-'",
			In:   "123e-2",
			Out:  []Token{{TOKEN_NUMBER, "123e-2"}, {TOKEN_EOF, "TOKEN_EOF"}},
		},
		{
			Name: "number with . at the start",
			In:   ".5",
			Out:  []Token{{TOKEN_NUMBER, ".5"}, {TOKEN_EOF, "TOKEN_EOF"}},
		},
		{
			Name: "number with . at the end",
			In:   "5.",
			Out:  []Token{{TOKEN_NUMBER, "5."}, {TOKEN_EOF, "TOKEN_EOF"}},
		},
	}

	panicTests := []struct {
		Name string
		In   string
	}{
		{
			Name: "expression with alphabet chars",
			In:   "df+123",
		},
		{
			Name: "number with underscore at the start",
			In:   "_123",
		},
		{
			Name: "number with underscore at the end",
			In:   "123_",
		},
		{
			Name: "number with e at the start",
			In:   "e123",
		},
		{
			Name: "number with e at the end",
			In:   "123e",
		},
		{
			Name: "number with e and no digits",
			In:   "123e.-2",
		},
		{
			Name: "number with e, '+' after and no digits",
			In:   "123e+.2",
		},
		{
			Name: "number with e, '-' after and no digits",
			In:   "123e-.2",
		},
		{
			Name: "number with e and '+' in wrong place",
			In:   "123e2+3",
		},
		{
			Name: "number with e and '-' in wrong place",
			In:   "123e23-",
		},
		{
			Name: "number with many e",
			In:   "12e3e2e",
		},
		{
			Name: "number with e and float power",
			In:   "123e2.2",
		},
		{
			Name: "number with many '.'",
			In:   ".52.3.",
		},
		{
			Name: "number with many adjacent '_'",
			In:   "5e2__0",
		},
		{
			Name: "number with many adjacent 'e'",
			In:   "5ee20",
		},
		{
			Name: "number with many adjacent 'e'",
			In:   "5..",
		},
	}

	for _, test := range nonPanicTests {
		t.Run(test.Name, func(t *testing.T) {
			r := strings.NewReader(test.In)
			out := NewLexer(r).Lex()
			assert.EqualValues(t, test.Out, out)
		})
	}

	for _, test := range panicTests {
		t.Run(test.Name, func(t *testing.T) {
			r := strings.NewReader(test.In)
			assert.Panics(t, func() {
				NewLexer(r).Lex()
			})
		})
	}
}
