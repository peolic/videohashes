#!/bin/bash

GOOS=windows GOARCH=amd64 EXTRA_LDFLAGS='-extldflags=-static -s -w' \
    go build videohashes.go
