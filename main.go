package main

import (
  "fmt"
  "github.com/aitumik/monkey/lexer"
)

func lexString() {
  input := `let a = 5;`

  l := lexer.New(input)
  tok := l.NextToken()
  fmt.Println(tok.Type)
  tok = l.NextToken()
  fmt.Println(tok.Type)
  tok = l.NextToken()
  fmt.Println(tok.Type)
}

func main() {
  fmt.Println("Monkey language")
  lexString()
}
