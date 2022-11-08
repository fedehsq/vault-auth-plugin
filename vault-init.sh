#!/bin/sh
echo "BUILDING THE PLUGIN..."
go build -o vault/plugins/auth-plugin vault/cmd/main.go
echo "STARTING VAULT SERVER..."
vault server -dev -dev-root-token-id root -dev-plugin-dir=./vault/plugins -dev-listen-address $(ipconfig getifaddr en0):8200