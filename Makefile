TARGETS = strdiff intdiff unistrdiff uniintdiff unifilediff strpatch

all: $(TARGETS)

strdiff: *.go examples/strdiff.go
	go build -o $@ examples/$@.go

intdiff: *.go examples/intdiff.go
	go build -o $@ examples/$@.go

unistrdiff: *.go examples/unistrdiff.go
	go build -o $@ examples/$@.go

uniintdiff: *.go examples/uniintdiff.go
	go build -o $@ examples/$@.go

unifilediff: *.go examples/unifilediff.go
	go build -o $@ examples/$@.go

strpatch: *.go examples/strpatch.go
	go build -o $@ examples/$@.go

fmt:
	go fmt ./...

check:
	go test -v .

clean:
	rm -f $(TARGETS)

.PHONY: all
