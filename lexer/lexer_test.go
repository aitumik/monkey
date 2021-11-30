package lexer

import (
  "testing"
  "github.com/aitumik/monkey/token"
)

type TestCase struct {
    expectedType  token.TokenType
    expectedLiteral string
}

func TestNextToken(t *testing.T) {
    input := `=+(){},;`

    l := New(input)

    tests := []TestCase{
        {token.ASSIGN, "="},
        {token.PLUS,"+"},
        {token.LPAREN,"("},
        {token.RPAREN,")"},
        {token.LBRACE,"{"},
        {token.RBRACE,"}"},
        {token.COMMA,","},
        {token.SEMICOLON,";"},
        {token.EOF,""},
    }

    for index,testcase := range tests {
        tok := l.NextToken()

        if tok.Type != testcase.expectedType {
            t.Fatalf("test %d => %q != %q",index,testcase.expectedType,tok.Type)
        }

        if tok.Literal != testcase.expectedLiteral {
            t.Fatalf("%q != %q",testcase.expectedLiteral,tok.Literal)
        }
    }
}

