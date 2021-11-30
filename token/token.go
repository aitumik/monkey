package token

const (
    ILLEGAL   = "ILLEGAL"
    EOF       = "EOF"

    // identifiers + literals
    IDENT     = "IDENT"
    INT       = "INT"

    // operators - +,-,*,/
    ASSIGN    = "="
    PLUS      = "+"

    // delimeters
    COMMA     = ","
    SEMICOLON = ";"

    LPAREN    = "("
    RPAREN    = ")"

    LBRACE    = "{"
    RBRACE    = "}"

    // keywords
    FUNCTION  = "FUNCTION"
    LET       = "LET"
)


type TokenType string

var keywords = map[string]TokenType{
  "fn" : FUNCTION,
  "let" : LET,
}

type Token struct {
  Type    TokenType
  Literal string
}

func LookupIdent(ident string) TokenType {
  if tok,ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
