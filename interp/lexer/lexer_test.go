package lexer_test

import (
	"bytes"
	"testing"

	"github.com/mk2/yon/interp/kit"
	"github.com/mk2/yon/interp/lexer"
	"github.com/mk2/yon/interp/token"
)

func TestLexer_new(t *testing.T) {

	l := lexer.New(bytes.NewBufferString("a 'b `test` 123 1.23 test {}"))

	var cnt = 0
	for tkn := l.NextToken(); tkn.GetType() != token.TEOF; tkn = l.NextToken() {

		if tkn.GetType() == token.TSpace {
			continue
		}
		kit.Printf("token type: %v token value: %s", tkn.GetType(), tkn.GetVal())

		switch cnt {
		case 0:
			assertTokenType(t, tkn.GetType(), token.TIdentifier)
		case 1:
			assertTokenType(t, tkn.GetType(), token.TIdentifier)
		case 2:
			assertTokenType(t, tkn.GetType(), token.TString)
		case 3:
			assertTokenType(t, tkn.GetType(), token.TNumber)
		case 4:
			assertTokenType(t, tkn.GetType(), token.TNumber)
		case 5:
			assertTokenType(t, tkn.GetType(), token.TIdentifier)
		case 6:
			assertTokenType(t, tkn.GetType(), token.TLeftBrace)
		case 7:
			assertTokenType(t, tkn.GetType(), token.TRightBrace)
		}

		cnt++
	}
}

func assertTokenType(t *testing.T, r kit.TokenType, a kit.TokenType) {

	if r != a {
		t.Errorf("[Invalid token type] real: %d assumed: %d", r, a)
	}
}
