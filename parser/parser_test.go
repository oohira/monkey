package parser

import (
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
