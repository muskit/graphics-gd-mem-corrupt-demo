#!/usr/bin/env bash

cd "$(dirname "$0")"

cd src
go get -u
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ../linux_amd64.so -buildmode=c-shared
GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o ../library.wasm

cd ..
cp library.wasm ../builds/web/
