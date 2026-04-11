test:
	go test -v -cover ./...

lint:
	golangci-lint run ./...

run:
	go run .

build:
	go build -o ethiocal .

package:
	fyne package -release

.PHONY: test lint run build package
