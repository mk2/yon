.PHONY: test text-lexer test-interp

default: test

GO=go

test: test-lexer test-parser test-interp

test-lexer:
	cd lexer; $(GO) test

test-parser:
	cd parser; $(GO) test

test-interp:
	$(GO) test
