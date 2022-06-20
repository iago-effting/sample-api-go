#!/bin/sh

go run ./cmd/migrations/main.go db create_go "$@"