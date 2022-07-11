#!/bin/bash

go mod download
rm -rf ./dist
gox -output "./dist/{{.Dir}}_{{.OS}}_{{.Arch}}" -osarch=linux/amd64 ./cmd/tt
