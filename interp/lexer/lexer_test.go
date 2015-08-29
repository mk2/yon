package lexer_test

import (
	"bytes"
	"github.com/mk2/yon/interp/lexer"
	"log"
	"testing"
)

func TestLexer_new(t *testing.T) {

	l := lexer.New(bytes.NewBufferString("a b:; 123 1.23 `test test`"))

	for tkn := l.NextToken(); tkn.GetType() != "eof"; tkn = l.NextToken() {

		if tkn.Typ != lexer.TSpace {
			log.Printf("token value: %s", tkn.Val)
		}
	}
}
