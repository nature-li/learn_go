#!/usr/bin/env bash
set -x

export GOPATH=$(cd $(dirname $0) && pwd)
echo ${GOPATH}

# compile
cd src/main
go install main

cd ${GOPATH}

# run
${GOPATH}/bin/main


