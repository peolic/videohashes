update:
	go get github.com/stashapp/stash@develop && \
	go mod tidy

# videohashes

windows:
	GOOS=windows GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes.exe ./cmd/videohashes

linux:
	GOOS=linux GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes-linux ./cmd/videohashes

macos:
	GOOS=darwin GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes-macos ./cmd/videohashes

build: windows linux macos

# duration

duration-windows:
	GOOS=windows GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/duration.exe ./cmd/duration

duration-linux:
	GOOS=linux GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/duration-linux ./cmd/duration

duration-macos:
	GOOS=darwin GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/duration-macos ./cmd/duration

duration-build: duration-windows duration-linux duration-macos

# phashcompare

phashcompare-windows:
	GOOS=windows GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/phashcompare.exe ./cmd/phashcompare

phashcompare-linux:
	GOOS=linux GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/phashcompare-linux ./cmd/phashcompare

phashcompare-macos:
	GOOS=darwin GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/phashcompare-macos ./cmd/phashcompare

phashcompare-build: phashcompare-windows phashcompare-linux phashcompare-macos
