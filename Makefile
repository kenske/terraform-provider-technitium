default: fmt lint install docs server

build:
	go build -v ./...

install: build
	go install -v ./...

lint:
	docker run --rm -v $$(pwd):/app -w /app golangci/golangci-lint:v2.1.2 golangci-lint run

## Generate docs
docs:
	cd tools; go generate ./...

fmt:
	gofmt -s -w -e .

test:
	go test -v -cover -timeout=120s -parallel=10 ./...

testacc:
	TF_ACC=1 go test -v -cover -timeout 120m ./...

.PHONY: fmt lint test testacc build install docs
