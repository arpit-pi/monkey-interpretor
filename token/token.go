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
)

type Token struct {
	Type TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
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