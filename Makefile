test:
	go test -race -cover ./...

# Quick test
qt:
	go test ./...

init:
	go get -u github.com/stretchr/testify
