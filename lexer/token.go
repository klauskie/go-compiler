package lexer

import (
	"../aux/constant"
	"fmt"
)

type Token struct {
	Type uint8
	Word string
}

type TokenList struct {
	Elements []Token
}

func (list *TokenList) Add(t Token) {
	list.Elements = append(list.Elements, t)
}

func (list *TokenList) Print() {
	for i, token := range list.Elements {
		fmt.Printf("i: %d, %s\n", i, token.ToString())
	}
}

func (t *Token) ToString() string {
	return fmt.Sprintf("Type: %s, Word: %s", t.getTypeToString(), t.Word)
}

func (t *Token) getTypeToString() string {
	var stype string
	switch t.Type {
	case constant.S_SUM:
		stype = "+"
		break
	case constant.S_SUBTRACT:
		stype = "-"
		break
	case constant.S_ASTERISK:
		stype = "*"
		break
	case constant.S_FORWARD_SLASH:
		stype = "/"
		break
	case constant.S_LESS:
		stype = "<"
		break
	case constant.S_LESS_EQUAL:
		stype = "<="
		break
	case constant.S_MORE:
		stype = ">"
		break
	case constant.S_MORE_EQUAL:
		stype = ">="
		break
	case constant.S_EQUAL_EQUAL:
		stype = "=="
		break
	case constant.S_NOT_EQUAL:
		stype = "!="
		break
	case constant.S_EQUAL:
		stype = "="
		break
	case constant.S_SEMICOLON:
		stype = ";"
		break
	case constant.S_COMMA:
		stype = ","
		break
	case constant.S_OPEN_PARENTHESIS:
		stype = "("
		break
	case constant.S_CLOSE_PARENTHESIS:
		stype = ")"
		break
	case constant.S_OPEN_SQR_BRACKET:
		stype = "["
		break
	case constant.S_CLOSE_SQR_BRACKET:
		stype = "]"
		break
	case constant.S_OPEN_CURLY_BRACKET:
		stype = "{"
		break
	case constant.S_CLOSE_CURLY_BRACKET:
		stype = "}"
		break
	case constant.S_OPEN_COMMENT_BLOCK:
		stype = "/*"
		break
	case constant.S_CLOSE_COMMENT_BLOCK:
		stype = "*/"
		break
	case constant.K_IF_ID:
		stype = "if"
		break
	case constant.K_ELSE_ID:
		stype = "else"
		break
	case constant.K_INT_ID:
		stype = "int"
		break
	case constant.K_RETURN_ID:
		stype = "return"
		break
	case constant.K_VOID_ID:
		stype = "void"
		break
	case constant.K_WHILE_ID:
		stype = "while"
		break
	case constant.K_INPUT_ID:
		stype = "input"
		break
	case constant.K_OUTPUT_ID:
		stype = "output"
		break
	case constant.H_WORD:
		stype = "word"
		break
	case constant.H_NUMBER:
		stype = "number"
		break
	case constant.D_SPACE:
		stype = "space"
		break
	case constant.D_NEWLINE:
		stype = "newline"
		break
	case constant.D_TAB:
		stype = "tab"
		break
	}
	return stype
}
