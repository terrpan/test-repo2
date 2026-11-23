.PHONY: all fmt vet test build scan clean

all: fmt vet test build

fmt:
	@test -z "$$(gofmt -l .)" || { gofmt -l .; exit 1; }

vet:
	go vet ./...

test:
	go test -race ./...

build:
	CGO_ENABLED=0 go build -trimpath -o bin/server ./cmd/server

scan:
	trivy fs --scanners vuln,secret,misconfig --severity HIGH,CRITICAL --ignore-unfixed --exit-code 1 .

clean:
	rm -rf bin
