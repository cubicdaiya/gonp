all: *.go examples/*.go
	go build -o strdiff examples/strdiff.go
	go build -o intdiff examples/intdiff.go

strdiff: *.go
	go build -o $@ examples/strdiff.go

intdiff: *.go examples/intdiff.go
	go build -o $@ examples/intdiff.go

fmt:
	go fmt ./...

check:
	go test -v ./...

clean:
	rm -f strdiff intdiff

.PHONY: all
