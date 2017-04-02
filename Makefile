strdiff: *.go examples/strdiff.go
	go build -o $@ examples/strdiff.go

fmt:
	go fmt ./...

check:
	go test -v ./...
