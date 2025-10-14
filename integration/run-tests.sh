#!/bin/bash

set -e

echo "🚀 Starting TheHive Integration Tests"

# Function to cleanup on exit
cleanup() {
    echo "🧹 Cleaning up..."
    docker-compose down
}
trap cleanup EXIT

# Start TheHive and dependencies
echo "📦 Starting TheHive services..."
docker-compose up -d thehive elasticsearch cassandra minio

# Wait for services to be healthy
echo "⏳ Waiting for services to be ready..."
timeout=300  # 5 minutes
elapsed=0
while ! docker-compose exec -T thehive curl -s -f http://localhost:9000/api/v1/status/public > /dev/null 2>&1; do
    if [ $elapsed -ge $timeout ]; then
        echo "❌ Timeout waiting for TheHive to be ready"
        docker-compose logs thehive
        exit 1
    fi
    echo "  Still waiting for TheHive... (${elapsed}s/${timeout}s)"
    sleep 5
    elapsed=$((elapsed + 5))
done

echo "✅ TheHive is ready!"

# Run the tests
echo "🧪 Running integration tests..."
docker-compose up --build --exit-code-from integration-tests integration-tests

echo "🎉 Integration tests completed successfully!"