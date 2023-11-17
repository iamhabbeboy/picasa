build:
	go build -o picasa && cp picasa /usr/local/bin

test:
	go test -v ./...

lint:
	golangci-lint run