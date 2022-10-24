GOARCH = amd64

UNAME = $(shell uname -s)

ifndef OS
	ifeq ($(UNAME), Linux)
		OS = linux
	else ifeq ($(UNAME), Darwin)
		OS = darwin
	endif
endif

.DEFAULT_GOAL := all

all: fmt build start

build:
	GOOS=$(OS) GOARCH="$(GOARCH)" go build -o vault/plugins/auth-plugin vault/cmd/main.go

start:
	vault server -dev -dev-root-token-id root -dev-plugin-dir=./vault/plugins -dev-listen-address $(shell ipconfig getifaddr en0):8200

vault-setup:
	vault auth enable -path=auth-plugin auth-plugin
	vault secrets enable ssh
	vault write ssh/roles/otp_key_role \
    key_type=otp \
    default_user=vagrant \
    cidr_list=0.0.0.0/0
	vault policy write bh-policy ./vault/bh-policy.hcl
	vault policy write user-policy ./vault/user-policy.hcl

clean:
	rm -f ./vault/plugins/auth-plugin

fmt:
	go fmt $$(go list ./...)

.PHONY: build clean fmt start enable