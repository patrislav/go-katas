REPOSITORY := github.com/patrislav/go-katas

all: test bin/wordchain bin/karatechop

.PHONY: bin/wordchain
bin/wordchain:
	go build -o bin/wordchain $(REPOSITORY)/cmd/wordchain

.PHONY: bin/karatechop
bin/karatechop:
	go build -o bin/karatechop $(REPOSITORY)/cmd/karatechop

.PHONY: bin/numgen
bin/numgen:
	go build -o bin/numgen $(REPOSITORY)/cmd/numgen

.PHONY: test
test:
	go test ./...

.PHONY: bench
bench:
	go test ./... -bench=.
