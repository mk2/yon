.PHONY: clean dist release build build-release update-deps get-deps test-interp test

default: build test

GO=go

dist: build

release: release-build

build:
	cd cmd; $(GO) build -o yon; mv yon ../

release-build:
	cd cmd; $(GO) build -o yon -tags=interp_release; mv yon ../

test: test-interp

test-interp:
	cd interp; make

restore:
	$(GO) mod download

clean:
	rm -f yon
