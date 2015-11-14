.PHONY: clean dist release build build-release get-deps test-interp test

default: test

GO=go

dist: build

release: release-build

build:
	cd cmd; go build -o yon -installsuffix .; mv yon ../

release-build:
	cd cmd; go build -o yon -installsuffix . -tags=interp_release; mv yon ../

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
