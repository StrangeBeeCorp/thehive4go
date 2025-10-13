#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🚀 Running full OpenAPI client generation pipeline...${NC}"

# Step 1: Download OpenAPI spec
echo -e "${BLUE}Step 1: Downloading OpenAPI specification...${NC}"
/scripts/download_openapi.sh
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Failed to download OpenAPI spec${NC}"
    exit 1
fi

# Step 2: Preprocess OpenAPI spec
echo -e "${BLUE}Step 2: Preprocessing OpenAPI specification...${NC}"
/scripts/preprocess_openapi.sh
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Failed to preprocess OpenAPI spec${NC}"
    exit 1
fi

# Step 3: Generate client code
echo -e "${BLUE}Step 3: Generating Go client code...${NC}"
/scripts/generate_client.sh
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Failed to generate client code${NC}"
    exit 1
fi

echo -e "${GREEN}🎉 Full pipeline completed successfully!${NC}"