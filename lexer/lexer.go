package lexer

import (
  "fmt"
  "github.com/aitumik/monkey/token"
)

// lexer struct
type Lexer struct {
  input    string
  position int
  readPos  int
  ch       byte
}

// method on struct lexer
func (l *Lexer) readChar() {
  //if finished set it to EOF 
  fmt.Println("Trying to read char")
  if l.readPos >= len(l.input) {
    l.ch = 0
  } else {
    // else get the last read pos and set the run
    // to ch
    l.ch = l.input[l.readPos]
  }

  l.position = l.readPos
  l.readPos += 1
}

// method NextToken() returns a token
func (l *Lexer) NextToken() token.Token {
  var tok token.Token
  switch l.ch {
    case '=' :
      // makes the (token,lexeme)
      tok = newToken(token.ASSIGN,l.ch)
    case ';':
      tok = newToken(token.SEMICOLON,l.ch)
    case ',':
      tok = newToken(token.COMMA,l.ch)
    case '+':
      tok = newToken(token.PLUS,l.ch)
    case '(':
      tok = newToken(token.LPAREN,l.ch)
    case ')':
      tok = newToken(token.RPAREN,l.ch)
    case '{':
      tok = newToken(token.LBRACE,l.ch)
    case '}':
      tok = newToken(token.RBRACE,l.ch)
    case 0:
      tok.Literal = ""
      tok.Type    = token.EOF
  }

  l.readChar()
  return tok
}

func newToken(tokenType token.TokenType,ch byte) (tok token.Token) {
  tok = token.Token{Type: tokenType,Literal: string(ch)}
  fmt.Println("Creating new token : ",ch,tokenType)
  return
}

func New(input string) *Lexer {
  l := &Lexer{input : input}
  l.readChar()
  return l
}
