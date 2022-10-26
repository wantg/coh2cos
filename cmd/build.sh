#!/usr/bin/env bash

# ./cmd/build.sh -debug

clear
wails build $1

./cmd/migrate.sh
mv ./build/cohcos-init.db ./build/bin/cohcos.db