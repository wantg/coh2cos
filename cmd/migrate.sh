#!/usr/bin/env bash

clear
rm -rf ./build/cohcos-init.db
go run ./cmd/migrate.go