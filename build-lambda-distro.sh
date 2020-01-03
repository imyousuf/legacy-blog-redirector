#!/bin/sh
set -eu
rm -rfv main main.zip
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
