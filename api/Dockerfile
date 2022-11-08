# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Federico Bernacca"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

# Setup folders
RUN mkdir /api
WORKDIR /api

# Copy the source from the current directory to the working Directory inside the container
COPY . .
ADD .env.docker .env

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Run the vault
CMD go run cmd/main.go