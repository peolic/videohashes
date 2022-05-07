update:
	go get github.com/stashapp/stash@develop && \
	go mod tidy

windows:
	GOOS=windows GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes.exe videohashes.go

linux:
	GOOS=linux GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes-linux videohashes.go

macos:
	GOOS=darwin GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
		go build -o dist/videohashes-macos videohashes.go

build: windows linux macos
