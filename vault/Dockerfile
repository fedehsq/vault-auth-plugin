# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Federico Bernacca"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

# Install vault
RUN apk add --no-cache vault libcap
RUN setcap cap_ipc_lock= /usr/sbin/vault

# Setup folders
RUN mkdir /vault
WORKDIR /vault

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the plugin
RUN go build -o plugins/auth-plugin cmd/main.go

# Run the vault
CMD vault server -dev -dev-root-token-id root -dev-plugin-dir=./plugins -dev-listen-address 0.0.0.0:8200