#!/bin/sh

set -o errexit
set -o nounset

export CGO_ENABLED=0
export GO111MODULE=on
export GOFLAGS="${GOFLAGS:-} -mod=${MOD}"

echo "Running tests:"
go test -installsuffix "static" "$@"
echo
