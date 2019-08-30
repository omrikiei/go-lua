package lexer

import (
	"github.com/omrikiei/go-lua/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	inputs := []string{
		"=+(){},;",
		"function add(a, b);\n" +
			"local sum = a + b;-- this is a comment that should be ignored\n" +
			"return sum; --[[ this is a multiline\n" +
		" comment that should be ignored... ]]-- end\n",
		"function a(b);\nb = !5",
	}
	tests := [][]struct {
		expectedType token.Type
		expectedLiteral string
	} {
		{
			{token.ASSIGN, "="},
			{ token.PLUS, "+"},
			{ token.LPAREN, "("},
			{ token.RPAREN, ")"},
			{ token.LBRACE, "{"},
			{ token.RBRACE, "}"},
			{ token.COMMA, ","},
			{ token.SEMICOLON, ";"},
		},
		{
			{ token.FUNCTION, "function"},
			{ token.IDENT, "add"},
			{ token.LPAREN, "("},
			{ token.IDENT, "a"},
			{token.COMMA, ","},
			{token.IDENT, "b"},
			{ token.RPAREN, ")"},
			{ token.SEMICOLON, ";"},
			{ token.LOCAL, "local"},
			{token.IDENT, "sum"},
			{token.ASSIGN, "="},
			{token.IDENT, "a"},
			{token.PLUS, "+"},
			{token.IDENT, "b"},
			{ token.SEMICOLON, ";"},
			{ token.RETURN, "return"},
			{token.IDENT, "sum"},
			{ token.SEMICOLON, ";"},
			{token.END, "end"},
		},
		{
			{token.FUNCTION, "function"},
			{token.IDENT, "a"},
			{ token.LPAREN, "("},
			{ token.IDENT, "b"},
			{ token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{ token.IDENT, "b"},
			{token.ASSIGN, "="},
			{token.ILLEGAL, "!"},
			{token.NUMBER, "5"},
		},
	}
	for j, input := range inputs {
		l := New(input)
		for i,tt := range tests[j] {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d][%d] - %q - tokentype wrong. expected=%d, got=%d, literal=%q", j, i, l.ch,
					tt.expectedType, tok.Type, tt.expectedLiteral)
			}
			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d][%d] - literal wrong. expected=%q, got=%q", j, i, tt.expectedLiteral,
					tok.Literal)
			}
		}
	}
}