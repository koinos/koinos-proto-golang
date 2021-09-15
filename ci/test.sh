#!/bin/bash

set -e
set -x

go test -v github.com/koinos/koinos-proto-golang/koinos/canonical -coverprofile=./build/canonical.out -coverpkg=./koinos/canonical
gcov2lcov -infile=./build/canonical.out -outfile=./build/canonical.info

golint -set_exit_status ./...
