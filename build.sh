#! /bin/bash
# Copyright 2018 Kuei-chun Chen. All rights reserved.

# dep init
# dep ensure
export version="0.1.0"
mkdir -p build
# env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$version" -o build/argos-linux-x64 argos.go
env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$version" -o build/argos-osx-x64 argos.go
# env GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$version" -o build/argos-win-x64.exe argos.go
