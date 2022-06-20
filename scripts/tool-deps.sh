#!/bin/sh

# Migration tool v4.15.2
# https://github.com/golang-migrate/migrate

curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-arm64.tar.gz | tar xvz -C ./bin
rm ./bin/README.md
rm ./bin/LICENSE