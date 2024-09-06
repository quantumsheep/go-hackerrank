.PHONY: generate
generate:
	go run ./cmd/generate > api.go
	go fmt api.go
