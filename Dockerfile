FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY ./ ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /DunnoYT

EXPOSE 8080

CMD ["/DunnoYT"]
