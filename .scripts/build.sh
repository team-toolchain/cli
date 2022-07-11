#!/bin/bash

go mod download
rm -rf ./dist
gox -output "./dist/tt" -osarch=linux/amd64 ./cmd/tt
