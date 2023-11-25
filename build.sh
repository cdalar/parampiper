#!/bin/bash
go build -ldflags="-w -s -X 'github.com/cdalar/parampiper/cmd.Version=$(git rev-parse HEAD | cut -c1-7)' \
-X 'github.com/cdalar/parampiper/cmd.BuildTime=$(date -u '+%Y-%m-%d %H:%M:%S')' \
-X 'github.com/cdalar/parampiper/cmd.GoVersion=$(go version)'"
