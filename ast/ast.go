package ast

import (
	"bytes"

	"github.com/oohira/monkey/token"
)

// Node is the interface that represents AST node.
type Node interface {
	TokenLiteral() string
	String() string
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

// TokenLiteral returns the first token literal of the program.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String returns a text representation of the program.
func (p *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}

// LetStatement represents a let statement.
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

// TokenLiteral returns the first token literal of the let statement.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String returns a text representation of the let statement.
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

func (ls *LetStatement) statementNode() {
}

// ReturnStatement represents a return statement.
type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

// TokenLiteral returns the first token literal of the return statement.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String returns a text representation of the return statement.
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) statementNode() {
}

// ExpressionStatement represents a expression statement.
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

// TokenLiteral returns the first token literal of the expression statement.
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String returns a text representation of the expression statement.
func (es *ExpressionStatement) String() string {
	var out bytes.Buffer
	if es.Expression != nil {
		out.WriteString(es.Expression.String())
	}
	return out.String()
}

func (es *ExpressionStatement) statementNode() {
}

// PrefixExpression represents an expression with prefix operator.
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// TokenLiteral returns the first token literal of the prefix expression.
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String returns a text representation of the prefix expression.
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) expressionNode() {
}

// Identifier represents an identifier.
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

// TokenLiteral returns the token literal of the identifier.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String returns a text representation of the identifier.
func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {
}

// IntegerLiteral represents an integer literal.
type IntegerLiteral struct {
	Token token.Token // token.INT
	Value int64
}

// TokenLiteral returns the token literal of the integer.
func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

// String returns a text representation of the integer.
func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) expressionNode() {
}
