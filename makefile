run:
	PORT=7890 go run cmd/main.go
test:
	go test -v ./...
remove-cache-test:
	go clean -testcache