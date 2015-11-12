.PHONY: clean get-deps test

default: test

test:
	# invoke interpreter test suite
	cd interp; make

get-deps:
	go get -u github.com/Scalingo/codegangsta-cli
	go get -u github.com/fatih/color
	go get -u github.com/mattn/go-isatty
	go get -u github.com/shiena/ansicolor
