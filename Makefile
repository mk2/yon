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

ensure-deps:
	dep ensure

clean:
	rm -f yon
