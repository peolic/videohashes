update:
	go get github.com/stashapp/stash@develop && \
	go mod tidy

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
