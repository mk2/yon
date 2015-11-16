.PHONY: clean dist release build build-release update-deps get-deps test-interp test

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

save-deps:
	find ./interp -type d | xargs -L 1 bash -c 'cd "$0" && godep save'
	find ./repl -type d | xargs -L 1 bash -c 'cd "$0" && godep save'
	find ./cmd -type d | xargs -L 1 bash -c 'cd "$0" && godep save'

update-deps:
	find ./interp -type d | xargs -L 1 bash -c 'cd "$0" && godep update'
	find ./repl -type d | xargs -L 1 bash -c 'cd "$0" && godep update'
	find ./cmd -type d | xargs -L 1 bash -c 'cd "$0" && godep update'

restore-deps:
	find ./interp -type d | xargs -L 1 bash -c 'cd "$0" && godep restore'
	find ./repl -type d | xargs -L 1 bash -c 'cd "$0" && godep restore'
	find ./cmd -type d | xargs -L 1 bash -c 'cd "$0" && godep restore'

clean:
	rm -f yon
