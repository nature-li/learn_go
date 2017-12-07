#!/usr/bin/env bash
export GOPATH=$(cd $(dirname $0) && pwd)
echo ${GOPATH}

# compile
go install main

# run
${GOPATH}/bin/main