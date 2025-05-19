# Expressive

Expressive is a calculator that can operate with many operands and operators.
It consists of 2 main parts: lexer and parser.

Lexer goes through input string and breaks it down into slice of tokens.

Parser takes slice of tokens, creates abstract syntax tree via recursive calls. The priority of operands that makes the recursive calls in right order is defined by EBNF grammar:

```
<expr> ::= <term>
<term> ::= <factor> (("+" | "-") <factor>)*
<factor> ::= <power> (("*" | "/" | "%") <power>)*
<power> ::= <unary> ("^" <unary>)*
<unary> ::= "-"? <unary> | <primary>
<primary> ::= NUMBER | "(" <expr> ")"
```

After creating the tree, i use usual recursive travese of it to evaluate final number.

## Operators

1. Plus (+)
2. Minus (-)
3. Float division (/)
4. Multiplication (\*)
5. Division by modulo (%)
6. Power (^)
7. Unary minus (-)

## Operands

All operands are automatically represented as float64.
Also, there is a support of:

1. Underscores in number (1_000_000)
2. Dots in number (3.141)
3. Exponential format (2e1, 2e+1, 2e-1)
4. Mixing first and third paragraph (2e1_0, 2e+1_0, 2e-1_0)
5. Mixing second and third paragraph, but no dots are allowed in power (3.141e2 is good, 3.141e2.2 is bad)

## Comments

You can write comments in input string, they start with `#` and nust be ended with newline character:

```test.txt
# Comment (must end with '\n')
5%2.5 (end expresssions with '\n' too)
(2-3) ^ 2
1_000 - 7
2*3+1
```

## Usage

You can start program by 2 variants:

```go
go run . "1+2"
```

or you can pass a file name, the result of all operations will be written to `output.txt`:

```go
go run . "test.txt"
```

## Afterword

For all the time i was doing programming, i didn't create a decent calculator, that at least can operate with more than 2 operands.
So, i decided to build it and practice new language that i'm currently learning :)
Credits to Xnacly and his post about his own calculator `https://xnacly.me/posts/2023/calculator-lexer/`, it inspired me to start this project.
