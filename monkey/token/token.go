// token/token.go

package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string //便利だからstring型に
}

const (
	ILLEGAL = "ILLEGAL" //トークンが未知であるときに使う
	EOF     = "EOF"     //ファイル終端、構文解析器にここで終了を教える

	//識別子+リテラル
	IDENT  = "IDENT" //add, foobar, x, y
	INT    = "INT"   //123
	STRING = "STRING"

	//演算子
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	AND = "&&"
	OR  = "||"

	//デリミタ
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	WHILE    = "WHILE"

	LBRACKET = "["
	RBRACKET = "]"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"while":  WHILE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
