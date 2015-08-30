package token

import (
	"fmt"

	"github.com/mk2/yon/interp/kit"
)

const (
	TError kit.TokenType = iota
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

type Token struct {
	Typ kit.TokenType
	Pos kit.Position
	Val string
}

func (t Token) GetType() kit.TokenType {

	return t.Typ
}

func (t Token) GetPos() kit.Position {

	return t.Pos
}

func (t Token) GetVal() string {

	return t.Val
}

func (t Token) String() string {

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
