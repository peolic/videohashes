# Stash Video Hashes Generator

A small tool that generates Stash and StashBox compatible PHash and OSHash hashes.  
Generation uses the [stashapp/stash](https://github.com/stashapp/stash) implementation.

## [Download latest from releases](https://github.com/peolic/videohashes/releases/latest)

## Usage
```
$ ./videohashes.exe '/path/to/video file.mp4'
```

**Example output:**
```
Duration: 00:18:45 (1125)
PHash:    82d42996f6eb09d5
OSHash:   3e98ab590428139c
```


## Build
```sh
# Windows
GOOS=windows GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' go build videohashes.go
# Linux
GOOS=linux GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' go build videohashes.go -o videohashes-linux
# Mac OS
GOOS=darwin GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' go build videohashes.go -o videohashes-macos
```
