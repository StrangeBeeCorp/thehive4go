FIXED_OPENAPI_PATH="/workspace/tmp/thehive_openapi_fixed.yaml"
CLIENT_PATH="/workspace/thehive"
CONFIG_PATH="/workspace/oapi-codegen.yaml"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🔧 Generating Go client from OpenAPI specification using oapi-codegen...${NC}"

# Clean up old generated files
rm -rf "$CLIENT_PATH"/*

echo -e "${GREEN}✅ Starting client generation with oapi-codegen...${NC}"

# Generate client using oapi-codegen
oapi-codegen --config "$CONFIG_PATH" "$FIXED_OPENAPI_PATH"

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Go client generated successfully${NC}"
    
    # Apply post-processing to generated files
    echo -e "${BLUE}🔧 Post-processing generated client...${NC}"
    /scripts/postprocess_client.sh "$CLIENT_PATH/client.go"
    echo -e "${GREEN}✅ Post-processing completed${NC}"
else
    echo -e "${RED}❌ Failed to generate Go client${NC}"
    exit 1
fi