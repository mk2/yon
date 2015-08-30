package lexer

import "fmt"

type tokenType int
type Position int

const (
	TError tokenType = iota
	TEOF
	TSpace
	TIdentifier
	TNumber
	TDot
	TDotS
	TString
	TDblColon
	TSemiColon
	TLeftParen
	TRightParen
	TLeftSquareBracket
	TRightSquareBracket
)

type token struct {
	Typ tokenType
	Pos Position
	Val string
}

func (t token) GetType() string {

	typStr := "undefined"

	switch t.Typ {

	case TEOF:
		typStr = "eof"

	case TSpace:
		typStr = "space"

	case TIdentifier:
		typStr = "identifier"

	case TNumber:
		typStr = "number"

	case TString:
		typStr = "string"

	case TDblColon:
		typStr = "symbol-double-colon"

	case TSemiColon:
		typStr = "symbol-semi-colon"

	}

	return typStr
}

func (t token) String() string {

	switch t.Typ {

	case TIdentifier:
		return t.Val

	case TNumber:
		return fmt.Sprintf("%lf", t.Val)

	case TString:
		return t.Val

	}

	return ""
}
