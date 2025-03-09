# syntax=docker/dockerfile:1

# This Dockerfile is for integration testing
FROM golang:1.23-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN mkdir -p /root/.aws
RUN mkdir -p /root/.config

# Test case 1
RUN cp ./tests/aws_config.test /root/.aws/config
RUN cp ./tests/sevp_config.test.toml /root/.config/sevp.toml
RUN go test -v ./src/...

# Test case 2
RUN cp ./tests/aws_config.test /root/.aws/config
RUN cp ./tests/sevp_config.test.toml /root/.config/sevp.toml
# Failing test
RUN aldflaksfljalfj

