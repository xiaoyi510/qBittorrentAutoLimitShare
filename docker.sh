#!/bin/bash

# This shell is executed before docker build.

goreleaser release
docker build -t xiaoyi510/qbit-auto-limit:v0.0.7 .


