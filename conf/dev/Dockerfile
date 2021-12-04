# Get Go image from DockerHub.
FROM golang:1.17 AS builder

# Set working directory.
WORKDIR /usr/app

# Copy dependency locks so we can cache.
COPY go.mod go.sum .

# Get all of our dependencies.
RUN go mod download

# Get air module for hot reload
RUN go get github.com/cosmtrek/air

# Copy all of our remaining application.
COPY . .
