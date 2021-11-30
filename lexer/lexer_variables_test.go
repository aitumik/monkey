package lexer

import (
  "testing"
)

type TestCase struct {
  expectedType   token.Type
  expectedLiteral token.Literal
}

func TestVariables(t *testing.T) {
  input := 
  `
  let a = 5;
  let b = 6;
  `

  l := New(input)

  tests :=[]TestCase{
    {token.LET,"let"},
    {token.IDENT,"a"},
    {token.ASSIGN,"="},
    {token.INT,"5"},
    {token.SEMICOLON,";"},

    {token.LET,"let"},
    {token.IDENT,"b"},
    {token.ASSIGN,"="},
    {token.INT,"6"},
    {token.SEMICOLON,";"},
  }

  for i,tt := range tests {
  }
}
