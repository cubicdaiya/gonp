TARGETS = strdiff intdiff unistrdiff uniintdiff unifilediff strpatch filepatch

all: $(TARGETS)

strdiff: *.go examples/strdiff.go
	go build -o $@ examples/$@.go

intdiff: *.go examples/intdiff.go
	go build -o $@ examples/$@.go

unistrdiff: *.go examples/unistrdiff.go
	go build -o $@ examples/$@.go

uniintdiff: *.go examples/uniintdiff.go
	go build -o $@ examples/$@.go

unifilediff: *.go examples/unifilediff.go examples/util.go
	go build -o $@ examples/$@.go examples/util.go

strpatch: *.go examples/strpatch.go
	go build -o $@ examples/$@.go

filepatch: *.go examples/filepatch.go examples/util.go
	go build -o $@ examples/$@.go examples/util.go

fmt:
	go fmt ./...

check:
	go test -v .

bench:
	go test -bench . -benchmem

clean:
	rm -f $(TARGETS)

.PHONY: all
