package lexer

import (
  "testing"
  "github.com/aitumik/monkey/token"
)

func TestNextToken(t *testing.T) {
  input := 
  `
  let a = 5;
  `
  /**
    let b = 6;

    let add = fn(x,y) {
      x + y;
    };

    let result = add(a,b)
  `
  **/
  tests :=[]struct{
    expectedType     token.TokenType
    expectedLiteral  string
  }{
    {token.LET,"let"},
    {token.IDENT, "a"},
    {token.ASSIGN,"="},
    {token.INT,"5"},
    {token.SEMICOLON,";"},
    /**
    {token.LET,"let"},
    {token.IDENT, "b"},
    {token.ASSIGN,"="},
    {token.INT,"6"},
    {token.COMMA,";"},

    {token.LET,"let"},
    {token.IDENT, "add"},
    {token.ASSIGN,"="},
    {token.FUNCTION,"fn"},
    {token.LPAREN,"("},
    {token.IDENT,"x"},
    {token.COMMA,","},
    {token.RPAREN,")"},
    {token.LBRACE,"{"},
    {token.IDENT, "x"},
    {token.PLUS, "+"},
    {token.IDENT, "y"},
    {token.SEMICOLON, ";"},
    {token.RBRACE,"}"},
    {token.SEMICOLON, ";"},

    {token.LET,"let"},
    {token.IDENT, "result"},
    {token.ASSIGN,"="},
    {token.IDENT,"add"},
    {token.LPAREN,"("},
    {token.IDENT,"a"},
    {token.COMMA,","},
    {token.IDENT,"b"},
    {token.RPAREN,")"},
    {token.COMMA,";"},
    **/
    {token.EOF, ""},
  }

  l := New(input)

  for i,tt := range tests {
    tok := l.NextToken()
    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - literal wrong. expected=%s, got=%q",i,tt.expectedLiteral,tok.Literal)
    }
  }
}

