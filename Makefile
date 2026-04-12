all: build

test:
	go test -v -cover ./...

lint:
	golangci-lint run ./...

run:
	go run .

build:
	go build -o ethiocal .

clean:
	rm -f ethiocal

package:
	fyne package -release

.PHONY: all test lint run build clean package
