main: build

update:
	go get github.com/stashapp/stash@develop && \
	go mod tidy

# videohashes

windows-amd64:
	GOOS=windows GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/windows-amd64/videohashes.exe ./cmd/videohashes

linux-amd64:
	GOOS=linux GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/linux-amd64/videohashes ./cmd/videohashes

macos-amd64:
	GOOS=darwin GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/macos-amd64/videohashes ./cmd/videohashes

windows-arm64:
	GOOS=windows GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/windows-arm64/videohashes.exe ./cmd/videohashes

linux-arm64:
	GOOS=linux GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/linux-arm64/videohashes ./cmd/videohashes

macos-arm64:
	GOOS=darwin GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/macos-arm64/videohashes ./cmd/videohashes

windows-arm:
	GOOS=windows GOARCH=arm EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/windows-arm/videohashes.exe ./cmd/videohashes

linux-arm:
	GOOS=linux GOARCH=arm EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/linux-arm/videohashes ./cmd/videohashes


build: windows-amd64 linux-amd64 macos-amd64 windows-arm64 linux-arm64 macos-arm64 windows-arm linux-arm

# duration

duration-windows-amd64:
	GOOS=windows GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/windows-amd64/duration.exe ./cmd/duration

duration-linux-amd64:
	GOOS=linux GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/linux-amd64/duration ./cmd/duration

duration-macos-amd64:
	GOOS=darwin GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/macos-amd64/duration ./cmd/duration

duration-windows-arm64:
	GOOS=windows GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/windows-arm64/duration.exe ./cmd/duration

duration-linux-arm64:
	GOOS=linux GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/linux-arm64/duration ./cmd/duration

duration-macos-arm64:
	GOOS=darwin GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/macos-arm64/duration ./cmd/duration

duration-windows-arm:
	GOOS=windows GOARCH=arm EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/windows-arm/duration.exe ./cmd/duration

duration-linux-arm:
	GOOS=linux GOARCH=arm EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/linux-arm/duration ./cmd/duration

duration-build: duration-windows-amd64 duration-linux-amd64 duration-macos-amd64 duration-windows-arm64 duration-linux-arm64 duration-macos-arm64 duration-windows-arm duration-linux-arm

# phashcompare

phashcompare-windows-amd64:
	GOOS=windows GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/windows-amd64/phashcompare.exe ./cmd/phashcompare

phashcompare-linux-amd64:
	GOOS=linux GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/linux-amd64/phashcompare ./cmd/phashcompare

phashcompare-macos-amd64:
	GOOS=darwin GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/macos-amd64/phashcompare ./cmd/phashcompare

phashcompare-windows-arm64:
	GOOS=windows GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/windows-arm64/phashcompare.exe ./cmd/phashcompare

phashcompare-linux-arm64:
	GOOS=linux GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/linux-arm64/phashcompare ./cmd/phashcompare

phashcompare-macos-arm64:
	GOOS=darwin GOARCH=arm64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/macos-arm64/phashcompare ./cmd/phashcompare


phashcompare-windows-arm:
	GOOS=windows GOARCH=arm EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/windows-arm/phashcompare.exe ./cmd/phashcompare

phashcompare-linux-arm:
	GOOS=linux GOARCH=arm EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/linux-arm/phashcompare ./cmd/phashcompare


phashcompare-build: phashcompare-windows-amd64 phashcompare-linux-amd64 phashcompare-macos-amd64 phashcompare-windows-arm64 phashcompare-linux-arm64 phashcompare-macos-arm64 phashcompare-windows-arm phashcompare-linux-arm
