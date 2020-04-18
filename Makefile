REPOSITORY := github.com/patrislav/go-katas

.PHONY: bin/wordchain
bin/wordchain:
	go build -o bin/wordchain $(REPOSITORY)/cmd/wordchain

.PHONY: test
test:
	go test ./...

.PHONY: bench
bench:
	go test ./... -bench=.
