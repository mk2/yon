package lexer_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/mk2/yon/interp/lexer"
	"github.com/mk2/yon/interp/token"
)

func TestLexer_new(t *testing.T) {

	l := lexer.New(bytes.NewBufferString("a b:; 123 1.23 `test test`"))

	for tkn := l.NextToken(); tkn.GetType() != token.TEOF; tkn = l.NextToken() {

		if tkn.GetType() != token.TSpace {
			log.Printf("token value: %s", tkn.GetVal())
		}
	}
}
