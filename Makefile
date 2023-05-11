linter:
	golangci-lint run

coverage:
	go test -cover -coverprofile='coverage.out' ./...
	go tool cover -html='coverage.out'