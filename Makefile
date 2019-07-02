test:
	go test -race -cover ./...

# Quick test
qt:
	go test ./...

bt:
	go test -bench='.' ./...

init:
	go get -u github.com/stretchr/testify
