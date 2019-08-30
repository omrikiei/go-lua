package token

const (
	ILLEGAL Type = iota
	EOF
	// Types
	IDENT
	NUMBER
	STRING
	TABLE
	FUNCTION
	USERDATA
	// Symbols
	ESCAPE
	ASSIGN
	PLUS
	MINUS
	MUL
	DIV
	MOD
	CARET
	COMMA
	HASH
	EQ
	NE
	LE
	GE
	LT
	GT
	SEMICOLON
	COLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET
	DOT
	TWO_DOT
	THREE_DOT
	DOUBLE_QUOTES
	QUOTES
	// Reserved keywords
	AND
	END
	IN
	REPEAT
	BREAK
	FALSE
	LOCAL
	RETURN
	DO
	FOR
	NIL
	THEN
	ELSE
	NOT
	TRUE
	ELSEIF
	IF
	OR
	UNTIL
	WHILE
)

type Type byte

type Token struct {
	Type    Type
	Literal string
}

func NewToken(t Type, l string) *Token {
	return &Token{
		Type:    t,
		Literal: l,
	}
}