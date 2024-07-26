#!/bin/bash

# This shell is executed before docker build.

#goreleaser release
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./qbit-auto-limit .
docker build -t xiaoyi510/qbit-auto-limit:v0.0.7 .
