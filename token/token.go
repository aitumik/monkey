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
    MINUS     = "-"
    BANG      = "!"
    ASTERISK  = "*"
    SLASH     = "/"

    LT = "<"
    GT = ">"

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
    TRUE      = "TRUE"
    FALSE     = "FALSE"
    RETURN    = "RETURN"
    IF        = "IF"
    ELSE      = "ELSE"

    // 2 char tokens
    EQ        = "=="
    NE        = "!="
)


type TokenType string

var keywords = map[string]TokenType{
  "fn" : FUNCTION,
  "let" : LET,
  "true" : TRUE,
  "false" : FALSE,
  "return" : RETURN,
  "if" : IF,
  "else" : ELSE,
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
