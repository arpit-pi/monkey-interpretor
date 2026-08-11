/*Example Monkey code
  let five = 5;
  let ten = 10;

  let add = fn(x,y) {
    x + y;
  };

  let result = add(five, ten);
  
*/

package token

type TokenType string

// TokenTypes 
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	IDENT   = "IDENT" // x, y
	INT     = "INT" 
	
	// Operators
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	MUL     = "*"
	DIV     = "/"
	BANG    = "!"

	//COMPARATORS
	LT      = "<"
	GT      = ">"
	EQ      = "=="
	NOT_EQ  = "!="

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

type Token struct {
	Type TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"if": IF,
	"else": ELSE,
	"true": TRUE,
	"false": FALSE,
	"return": RETURN,
}

func LookupIdent(iden string) TokenType {
	// check if map contains the iden,
	// if true, return the keyword type
	// else it is an identifier 
	if tok, ok := keywords[iden]; ok {
		return tok
	}

	return IDENT
}