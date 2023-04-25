FROM golang:1.20

# Set destination for COPY
WORKDIR /app

COPY go.mod go.sum ./
ADD ./cmd ./cmd
RUN go mod download

RUN dir -s

 RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/web_api
