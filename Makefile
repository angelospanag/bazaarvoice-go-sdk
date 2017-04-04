.PHONY: test build run clean format imports vet
TESTTARGET=$$(go list ./... | grep -v /vendor/)
VETTARGET=$$(go list ./... | grep -v /vendor/)
IMPORTSTARGET=$$(find . -type f -name '*.go' -not -path "./vendor/*")
APPBIN=$(shell basename $(PWD))

test: imports
	@go test -timeout=5s -cover -race $(TESTTARGET)

build: imports
	@go build

clean:
	@go clean

format:
	@gofmt -s -w .

vet:
	@go vet $(VETTARGET)

imports:
	@goimports -w $(IMPORTSTARGET)

run: build
	@./$(APPBIN)

gendoc: aglio -i api-description.apib -o index.html
