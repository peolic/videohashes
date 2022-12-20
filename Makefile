update:
	go get github.com/stashapp/stash@develop && \
	go mod tidy

# videohashes

windows-amd64:
	GOOS=windows GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes-amd64.exe ./cmd/videohashes

linux-amd64:
	GOOS=linux GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes-amd64-linux ./cmd/videohashes

macos-amd64:
	GOOS=darwin GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes-amd64-macos ./cmd/videohashes

windows-arm64:
	GOOS=windows GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes-arm64.exe ./cmd/videohashes

linux-arm64:
	GOOS=linux GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes-arm64-linux ./cmd/videohashes

macos-arm64:
	GOOS=darwin GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes-arm64-macos ./cmd/videohashes

windows-arm:
	GOOS=windows GOARCH=arm EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes-arm.exe ./cmd/videohashes

linux-arm:
	GOOS=linux GOARCH=arm EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes-arm-linux ./cmd/videohashes


build: windows-amd64 linux-amd64 macos-amd64 windows-arm64 linux-arm64 macos-arm64 windows-arm linux-arm

# duration

duration-windows-amd64:
	GOOS=windows GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/duration-amd64.exe ./cmd/duration

duration-linux-amd64:
	GOOS=linux GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/duration-amd64-linux ./cmd/duration

duration-macos-amd64:
	GOOS=darwin GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/duration-amd64-macos ./cmd/duration

duration-windows-arm64:
	GOOS=windows GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/duration-arm64.exe ./cmd/duration

duration-linux-arm64:
	GOOS=linux GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/duration-arm64-linux ./cmd/duration

duration-macos-arm64:
	GOOS=darwin GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/duration-arm64-macos ./cmd/duration

duration-windows-arm:
	GOOS=windows GOARCH=arm EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/duration-arm.exe ./cmd/duration

duration-linux-arm:
	GOOS=linux GOARCH=arm EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/duration-arm-linux ./cmd/duration

duration-build: duration-windows-amd64 duration-linux-amd64 duration-macos-amd64 duration-windows-arm64 duration-linux-arm64 duration-macos-arm64 duration-windows-arm duration-linux-arm

# phashcompare

phashcompare-windows-amd64:
	GOOS=windows GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/phashcompare-amd64.exe ./cmd/phashcompare

phashcompare-linux-amd64:
	GOOS=linux GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/phashcompare-amd64-linux ./cmd/phashcompare

phashcompare-macos-amd64:
	GOOS=darwin GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/phashcompare-amd64-macos ./cmd/phashcompare

phashcompare-windows-arm64:
	GOOS=windows GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/phashcompare-arm64.exe ./cmd/phashcompare

phashcompare-linux-arm64:
	GOOS=linux GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/phashcompare-arm64-linux ./cmd/phashcompare

phashcompare-macos-arm64:
	GOOS=darwin GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/phashcompare-arm64-macos ./cmd/phashcompare


phashcompare-windows-arm:
	GOOS=windows GOARCH=arm EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/phashcompare-arm.exe ./cmd/phashcompare

phashcompare-linux-arm:
	GOOS=linux GOARCH=arm EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/phashcompare-arm-linux ./cmd/phashcompare


phashcompare-build: phashcompare-windows-amd64 phashcompare-linux-amd64 phashcompare-macos-amd64 phashcompare-windows-arm64 phashcompare-linux-arm64 phashcompare-macos-arm64 phashcompare-windows-arm phashcompare-linux-arm
