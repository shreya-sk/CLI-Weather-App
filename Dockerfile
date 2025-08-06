FROM golang:1.23
WORKDIR /app

# Copy dependency files first (for caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN go build -o weather-app main.go

# Run the binary
CMD ["./weather-app"]