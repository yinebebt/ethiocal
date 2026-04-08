test:
	go test -v -cover ./...

run:
	go run .

build:
	go build -o ethiocal .

.PHONY: test run build
