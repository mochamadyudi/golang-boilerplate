# syntax=docker/dockerfile:1
FROM golang:1.24 as builder

ARG GO_ENV=development
ENV GO_ENV $GO_ENV

# set working directory
WORKDIR /app

# COPY go.mod | go.sum
COPY go.mod go.sum ./

# download dependencies
RUN go mod download

# COPY source code
COPY . .

# build app
RUN if [ "$GO_ENV" = "development" ]; then \
    go build -p
