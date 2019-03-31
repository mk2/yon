.PHONY: clean dist release build build-release update-deps get-deps test-interp test

default: build test

GO=go

dist: build

release: release-build

build:
	cd cmd; go build -o yon; mv yon ../

release-build:
	cd cmd; go build -o yon -tags=interp_release; mv yon ../

test: test-interp

test-interp:
	cd interp; go test

restore:
	go mod download

clean:
	rm -f yon
