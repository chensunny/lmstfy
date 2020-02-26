#!/bin/bash
set -e -x

DIR=`go list ./... | grep -v client|awk '{printf $0" "}'`
go test -v -covermode=count -coverprofile=coverage.out ${DIR}
