.PHONY: fmt test race vet build check prove clean

fmt:
	gofmt -w cmd internal

test:
	go test ./...

race:
	go test -race ./...

vet:
	go vet ./...

build:
	mkdir -p bin
	go build -o bin/edge-openai ./cmd/edge-openai

check: fmt test race vet build

prove: build
	./bin/edge-openai prove

clean:
	rm -rf bin results
