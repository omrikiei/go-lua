package lexer

import (
	"github.com/omrikiei/go-lua/token"
	"unicode"
)

var SimpleSymbols = map[byte]token.Type{
	';': token.SEMICOLON,
	'(': token.LPAREN,
	')': token.RPAREN,
	',': token.COMMA,
	'+': token.PLUS,
	'-': token.MINUS,
	'/': token.DIV,
	'*': token.MUL,
	'%': token.MOD,
	'^': token.CARET,
	'{': token.LBRACE,
	'}': token.RBRACE,
	'[': token.LBRACKET,
	']': token.RBRACKET,
	'#': token.HASH,
	':': token.COLON,
	'\\': token.ESCAPE,
	'>': token.GT,
	'<': token.LT,
}

var Reserved = map[string]token.Type {
	"=": token.ASSIGN,
	"==": token.EQ,
	"~=": token.NE,
	"<=": token.LE,
	">=": token.GE,
	".": token.DOT,
	"..": token.TWO_DOT,
	"...": token.THREE_DOT,
	"function": token.FUNCTION,
	"and": token.AND,
	"end": token.END,
	"in": token.IN,
	"or": token.OR,
	"repeat": token.REPEAT,
	"break": token.BREAK,
	"false": token.FALSE,
	"local": token.LOCAL,
	"return": token.RETURN,
	"do": token.DO,
	"for": token.FOR,
	"nil": token.NIL,
	"then": token.THEN,
	"else": token.ELSE,
	"not": token.NOT,
	"true": token.TRUE,
	"elseif": token.ELSEIF,
	"if": token.IF,
	"util": token.UNTIL,
	"while": token.WHILE,
}

type Lexer struct {
	input string
	line int
	position int // pointer to current char
	readPosition int //current reading position
	ch byte //current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() *token.Token {
	l.skipWhitespaceAndComments()
	var tok *token.Token
	if tType, ok := SimpleSymbols[l.ch]; ok {
		// It's a simple symbol
		tok = token.NewToken(tType, string(l.ch))
		l.readChar()
	} else {
		identifier := l.readIdentifier()
		if tokenType, ok := Reserved[identifier]; ok {
			// It's a reserved keyword
			tok = token.NewToken(tokenType, identifier)
		} else if len(identifier) > 1 && (identifier[0] == '\'' || identifier[0] == '"') &&
			identifier[0] == identifier[len(identifier)-1] {
			// It's a string def
			tok = token.NewToken(token.STRING, identifier[1:len(identifier)-1])
		} else if '0' <= identifier[0] && identifier[0] <= '9' {
			// It's a Number def
			tok = token.NewToken(token.NUMBER, identifier)
		} else if unicode.IsLetter(rune(identifier[0])) || identifier[0] == '_' {
			tok = token.NewToken(token.IDENT, identifier)
		} else if l.ch == 0 {
			tok = token.NewToken(token.EOF, "")
		} else {
			// Illegal token
			tok = token.NewToken(token.ILLEGAL, string(l.ch))
			l.readChar()
		}
	}
	return tok
}

func (l *Lexer) skipWhitespaceAndComments() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
	if l.ch == '-' && l.input[l.readPosition] == '-' {
		// It's a comment
		if l.position+4 <= len(l.input) && l.input[l.position:l.position+4] == "--[[" {
			// Multiline comment
			for l.ch != 0 && l.input[l.position-4:l.position] != "]]--" {
				l.readChar()
			}
		} else {
			// Regular comment
			for l.ch != 0 && l.ch != '\n' && l.ch != '\r' {
				l.readChar()
			}
		}
		l.skipWhitespaceAndComments()
	}
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	if unicode.IsLetter(rune(l.ch)) {
		for unicode.IsLetter(rune(l.ch)) || l.ch == '_' || unicode.IsNumber(rune(l.ch)) || l.ch == '-' {
			l.readChar()
		}
		return l.input[position:l.position]
	}
	if unicode.IsDigit(rune(l.ch)) {
		for unicode.IsDigit(rune(l.ch)) || l.ch == '.' {
			l.readChar()
		}
		return l.input[position:l.position]
	}
	switch l.ch {
	case '.':
		for l.ch == '.' {
			l.readChar()
		}
		return l.input[position:l.position]
	case '=':
		return l.readOperator()
	case '~':
		return l.readOperator()
	case '>':
		return l.readOperator()
	case '<':
		return l.readOperator()
	case '\\':
		l.readChar()
		s := l.input[position:l.position]
		l.readChar()
		return s
	case '"':
		for l.ch != '"' {
			l.readChar()
		}
		return l.input[position:l.position]
	case '\'':
		for l.ch != '\'' {
			l.readChar()
		}
		return l.input[position:l.position]
	default:
		return string(l.input[position])
	}
}

func (l *Lexer) readOperator() string {
	position := l.position
	l.readChar()
	if l.ch != '=' {
		return string(l.input[position])
	} else {
		s := l.input[position:l.position]
		l.readChar()
		return s
	}
}