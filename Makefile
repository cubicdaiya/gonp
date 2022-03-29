all: *.go examples/*.go
	go build -o strdiff examples/strdiff.go
	go build -o intdiff examples/intdiff.go
	go build -o unistrdiff examples/unistrdiff.go
	go build -o uniintdiff examples/uniintdiff.go
	go build -o strpatch examples/strpatch.go

strdiff: *.go
	go build -o $@ examples/strdiff.go

intdiff: *.go examples/intdiff.go
	go build -o $@ examples/intdiff.go

unistrdiff: *.go
	go build -o $@ examples/unistrdiff.go

uniintdiff: *.go
	go build -o $@ examples/uniintdiff.go

strpatch: *.go
	go build -o $@ examples/strpatch.go

fmt:
	go fmt ./...

check:
	go test -v .

clean:
	rm -f strdiff intdiff unistrdiff uniintdiff strpatch

.PHONY: all
