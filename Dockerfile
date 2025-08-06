FROM golang:1.23 AS builder

WORKDIR /some
# Copy dependency files first (for caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o weather-app main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /some/weather-app . 
# Run the binary
CMD ["./weather-app"]