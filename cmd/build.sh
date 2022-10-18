#!/usr/bin/env bash

# ./cmd/build.sh -debug

clear
cp ./build/cohcos-init.db ./build/bin/cohcos.db

wails build $1