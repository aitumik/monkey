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

func TestVariables(t *testing.T) {
    input := `let five = 5
    let ten = 10;
    
    let add = fn(x,y) {
        x + y;
    }

    let result = add(five,ten);
    `

    l := New(input)

    tests := []TestCase{
    {token.LET, "let"},
           {token.IDENT, "five"},
           {token.ASSIGN, "="},
           {token.INT, "5"},
           {token.SEMICOLON, ";"},
           {token.LET, "let"},
           {token.IDENT, "ten"},
           {token.ASSIGN, "="},
           {token.INT, "10"},
           {token.SEMICOLON, ";"},
           {token.LET, "let"},
           {token.IDENT, "add"},
           {token.ASSIGN, "="},
           {token.FUNCTION, "fn"},
           {token.LPAREN, "("},
           {token.IDENT, "x"},
           {token.COMMA, ","},
           {token.IDENT, "y"},
           {token.RPAREN, ")"},
           {token.LBRACE, "{"},
           {token.IDENT, "x"},
           {token.PLUS, "+"},
           {token.IDENT, "y"},
           {token.SEMICOLON, ";"},
           {token.RBRACE, "}"},
           {token.SEMICOLON, ";"},
           {token.LET, "let"},
           {token.IDENT, "result"},
           {token.ASSIGN, "="},
           {token.IDENT, "add"},
           {token.LPAREN, "("},
           {token.IDENT, "five"},
           {token.COMMA, ","},
           {token.IDENT, "ten"},
           {token.RPAREN, ")"},
           {token.SEMICOLON, ";"},
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

