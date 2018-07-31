package parser

import (
	"fmt"
	"testing"

	"github.com/oohira/monkey/ast"
	"github.com/oohira/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, test := range tests {
		stmt := program.Statements[i]
		if stmt.TokenLiteral() != "let" {
			t.Errorf("stmt[%d].TokenLiteral is not 'let'. got=%q", i, stmt.TokenLiteral())
		}
		letStmt, ok := stmt.(*ast.LetStatement)
		if !ok {
			t.Errorf("stmt[%d] is not *LetStatement. got=%T", i, stmt)
		}
		if letStmt.Name.Value != test.expectedIdentifier {
			t.Errorf("stmt[%d].Name.Value is not '%s'. got=%q",
				i, test.expectedIdentifier, letStmt.Name.Value)
		}
		if letStmt.Name.TokenLiteral() != test.expectedIdentifier {
			t.Errorf("stmt[%d].Name.TokenLiteral is not '%s'. got=%q",
				i, test.expectedIdentifier, letStmt.Name.TokenLiteral())
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for i, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt[%d] is not *ReturnStatement. got=%T", i, stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("stmt[%d].TokenLiteral is not 'return'. got=%q", i, stmt.TokenLiteral())
		}
	}
}

func TestPrefixExpressions(t *testing.T) {
	tests := []struct {
		input    string
		operator string
		value    int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("[%d] program has not enough statements. got=%d", i, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("[%d] program.Statements[0] is not *ExpressionStatement. got=%T",
				i, program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("[%d] exp is not *PrefixExpressionr. got=%T", i, stmt.Expression)
		}
		if exp.Operator != test.operator {
			t.Errorf("[%d] exp.Operator is not %s. got=%s", i, test.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, test.value) {
			return
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ExpressionStatement. got=%T",
			program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp is not *Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value is not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral is not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ExpressionStatement. got=%T",
			program.Statements[0])
	}
	testIntegerLiteral(t, stmt.Expression, 5)
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testIntegerLiteral(t *testing.T, exp ast.Expression, want int64) bool {
	literal, ok := exp.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp is not *IntegerLiteral. got=%T", exp)
		return false
	}
	if literal.Value != want {
		t.Errorf("literal.Value is not %d. got=%s", want, literal.Value)
		return false
	}
	if literal.TokenLiteral() != fmt.Sprintf("%d", want) {
		t.Errorf("literal.TokenLiteral is not %d. got=%s", want, literal.TokenLiteral())
		return false
	}
	return true
}
