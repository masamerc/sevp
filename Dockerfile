# syntax=docker/dockerfile:1

# This Dockerfile is for integration testing
FROM golang:1.23-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD ["sh", "tests/run_integration_tests.sh"]
