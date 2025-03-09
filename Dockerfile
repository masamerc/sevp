# syntax=docker/dockerfile:1

# This Dockerfile is for integration testing
FROM golang:1.23-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o target/sevp

FROM alpine:latest
WORKDIR /root/app
# copy over the binary
COPY --from=builder /app/target/sevp .
# copy over the config file
COPY aws_config.test /root/.aws/config
COPY sevp_config.test.toml /root/.config/sevp.toml
CMD ["./sevp"]
