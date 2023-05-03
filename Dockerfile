FROM golang:1.20 AS builder

# Set destination for COPY
WORKDIR /app

COPY go.mod go.sum ./
ADD ./cmd ./cmd
RUN go mod download

RUN dir -s

FROM builder AS web_api

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/web_api


FROM builder AS qa

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
