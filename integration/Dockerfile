# Simple integration test runner
FROM golang:1.18-alpine

RUN apk add --no-cache git ca-certificates curl

WORKDIR /app

# Copy everything from the build context (parent directory)
COPY . .

# Run the test directly
CMD ["go", "run", "integration/cmd/test/main.go"]
