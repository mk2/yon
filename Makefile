.PHONY: clean get-deps test-interp test

default: test

GO=go

test: test-interp

test-interp:
	cd interp; make

get-deps:
	go get github.com/Scalingo/codegangsta-cli
	go get github.com/fatih/color
	go get github.com/mattn/go-isatty
	go get github.com/shiena/ansicolor
