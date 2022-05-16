#!/bin/sh

export CGO_ENABLED=0
export GO111MODULE=on
export GOFLAGS="${GOFLAGS:-}"

echo "Running golangci-lint: "
ERRS=$(golangci-lint --go=1.18 run "$@" 2>&1 || true)
if [ -n "${ERRS}" ]; then
    echo "FAIL"
    echo "${ERRS}"
    echo
    exit 1
fi
echo "PASS"
echo
