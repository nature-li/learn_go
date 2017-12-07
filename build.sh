#!/usr/bin/env bash
# open debug flag
set -x

# set GOPATH
export GOPATH=$(cd $(dirname $0) && pwd)

# compile
rm -rf ${GOPATH}/bin
rm -rf ${GOPATH}/pkg
$(cd src && go install main)

# run
${GOPATH}/bin/main
