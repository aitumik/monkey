package parser

import (
	"github.com/aitumik/monkey/ast"
	"github.com/aitumik/monkey/lexer"
	"github.com/aitumik/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}
