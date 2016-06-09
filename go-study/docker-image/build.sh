#!/bin/bash
apk update && apk add bash go && rm -rf /var/cache/apk/*
export GOPATH=/go

go build -o upccup -v
