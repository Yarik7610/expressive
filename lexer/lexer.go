package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

type Lexer struct {
	scanner bufio.Reader
	cur     rune
}

func NewLexer(reader io.Reader) *Lexer {
	l := &Lexer{scanner: *bufio.NewReader(reader)}
	//advance on init so l.cur wasn't 0 (EOF)
	l.advance()
	return l
}

func (l *Lexer) Lex() []Token {
	tokens := make([]Token, 0)

	for l.cur != 0 {
		tokenType := TOKEN_UNKNOWN

		switch l.cur {
		case '#':
			for l.cur != '\n' && l.cur != 0 {
				l.advance()
			}
			continue
		case ' ', '\t', '\n', '\r':
			l.advance()
			continue
		case '+':
			tokenType = TOKEN_PLUS
		case '-':
			tokenType = TOKEN_MINUS
		case '/':
			tokenType = TOKEN_SLASH
		case '*':
			tokenType = TOKEN_ASTERISK
		case '(':
			tokenType = TOKEN_BRACE_LEFT
		case ')':
			tokenType = TOKEN_BRACE_RIGHT
		default:
			if (l.cur >= '0' && l.cur <= '9') || l.cur == '.' {
				tokens = append(tokens, l.number())
				continue
			}
		}

		if tokenType == TOKEN_UNKNOWN {
			panic(fmt.Sprintf("lexer detected unknown token: %q", l.cur))
		}

		tokens = append(tokens, Token{tokenType, string(l.cur)})
		l.advance()
	}

	tokens = append(tokens, Token{TOKEN_EOF, "TOKEN_EOF"})

	return tokens
}

func (l *Lexer) advance() {
	r, _, err := l.scanner.ReadRune()
	if err != nil {
		//If err, we end advancing like if it was an EOF
		l.cur = 0
		return
	}
	l.cur = r
}

func (l *Lexer) number() Token {
	var b bytes.Buffer

	expCount := 0
	dotCount := 0

	for (l.cur >= '0' && l.cur <= '9') || l.cur == 'e' || l.cur == '_' || l.cur == '.' {
		if l.cur == 'e' {
			expCount++
		} else if l.cur == '.' {
			dotCount++
		}
		if expCount > 1 || dotCount > 1 {
			panic(fmt.Sprintf("lexer: %q was detected in number more than once", l.cur))
		}

		b.WriteRune(l.cur)
		l.advance()
	}

	raw := b.Bytes()
	r, _ := utf8.DecodeLastRune(raw)
	if r == '_' || r == 'e' {
		panic(fmt.Sprintf("lexer: %q must separate successive digits", r))
	}

	return Token{TOKEN_NUMBER, b.String()}
}

func PrintTokens(tokens []Token) {
	fmt.Printf("%5s | %20s | %20s\n", "index", "type", "raw")
	for i, token := range tokens {
		fmt.Printf("%5d | %20s | %20s\n", i, TOKENS[token.Type], token.Raw)
	}
}
