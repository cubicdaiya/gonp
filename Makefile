all: strdiff intdiff unistrdiff uniintdiff unifilediff strpatch

strdiff: *.go examples/strdiff.go
	go build -o $@ examples/strdiff.go

intdiff: *.go examples/intdiff.go
	go build -o $@ examples/intdiff.go

unistrdiff: *.go examples/unistrdiff.go
	go build -o $@ examples/unistrdiff.go

uniintdiff: *.go examples/uniintdiff.go
	go build -o $@ examples/uniintdiff.go

unifilediff: *.go examples/unifilediff.go
	go build -o $@ examples/unifilediff.go

strpatch: *.go examples/strpatch.go
	go build -o $@ examples/strpatch.go

fmt:
	go fmt ./...

check:
	go test -v .

clean:
	rm -f strdiff intdiff unistrdiff uniintdiff unifilediff strpatch

.PHONY: all
