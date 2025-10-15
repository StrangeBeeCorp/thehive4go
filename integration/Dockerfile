# Integration test runner with Go test
FROM golang:1.23.2-alpine

RUN apk add --no-cache git ca-certificates curl

WORKDIR /app

# Copy everything from the parent directory
COPY . .

# Download dependencies for the main module
RUN go mod tidy

# Run integration tests
CMD ["go", "test", "-v", "-timeout=10m", "./integration"]
