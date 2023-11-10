#!/bin/bash

# This shell is executed before docker build.

gorelease release
docker build -t xiaoyi510/qbit-auto-limit:v0.0.6 .


