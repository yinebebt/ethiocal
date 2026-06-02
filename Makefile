all: build

# Build both binaries: GUI app and CLI/server tool.
build:
	go build -o ethiocal .
	CGO_ENABLED=0 go build -o ethiocal-cli ./cmd/ethiocal-cli

test:
	go test -v -cover ./...

lint:
	golangci-lint run ./...

# Run the GUI app.
run:
	go run .

# Run the CLI tool, e.g. make run-cli ARGS="convert gtoe 2024 9 11".
run-cli:
	go run ./cmd/ethiocal-cli $(ARGS)

clean:
	rm -f ethiocal ethiocal-cli

package:
	fyne package -release

# Slim arm64 APK. Resolves the NDK from ANDROID_NDK_HOME or the newest under $ANDROID_HOME/ndk.
package-android:
	@ndk="$${ANDROID_NDK_HOME:-$$(ls -d $$ANDROID_HOME/ndk/* 2>/dev/null | sort -V | tail -1)}"; \
	[ -n "$$ndk" ] || { echo "No NDK: set ANDROID_NDK_HOME or install one under $$ANDROID_HOME/ndk"; exit 1; }; \
	ANDROID_NDK_HOME="$$ndk" fyne package --os android/arm64

.PHONY: all build test lint run run-cli clean package package-android
