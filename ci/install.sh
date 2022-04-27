#!/bin/bash

sudo gem install coveralls-lcov
go install github.com/jandelgado/gcov2lcov@latest
go get -u golang.org/x/lint/golint
