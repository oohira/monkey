package ast

import "github.com/oohira/monkey/token"

// Node is the interface that represents AST node.
type Node interface {
	TokenLiteral() string
}

// Statement is the interface that represents a statement in AST.
type Statement interface {
	Node
	statementNode()
}

// Expression is the interface that represents an expression in AST.
type Expression interface {
	Node
	expressionNode()
}

// Program represents a Monkey program.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns a text representation of the program.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement represents a let statement.
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

// TokenLiteral returns a text representation of the let statement.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {
}

// ReturnStatement represents a return statement.
type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

// TokenLiteral returns a text representation of the return statement.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) statementNode() {
}

// Identifier represents an identifier.
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

// TokenLiteral returns a text representation of the identifier.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {
}
