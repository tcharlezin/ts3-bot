#!/bin/bash

env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o xClaimedBot ./cmd/api
docker build --platform linux/x86_64 -f Dockerfile -t tcharlezin/xclaimedbot:latest .
docker push tcharlezin/xclaimedbot:latest
rm xClaimedBot