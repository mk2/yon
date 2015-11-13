.PHONY: clean get-deps test-interp test

default: test

GO=go

dist: build

build:
	cd cmd; go build -o yon; mv yon ../

test: test-interp

test-interp:
	cd interp; make

get-deps:
	go get github.com/Scalingo/codegangsta-cli
	go get github.com/fatih/color
	go get github.com/mattn/go-isatty
	go get github.com/shiena/ansicolor

clean:
	rm -f yon
