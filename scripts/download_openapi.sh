OPENAPI_SPEC_URL="https://docs.strangebee.com/thehive/api-docs/docs.yaml"
TEMP_OPENAPI_PATH="/workspace/tmp/thehive_openapi.yaml"
mkdir -p /workspace/tmp

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Download OpenAPI spec
echo -e "${BLUE}📥 Downloading OpenAPI specification...${NC}"
if curl -sSL "${OPENAPI_SPEC_URL}" -o "${TEMP_OPENAPI_PATH}"; then
    echo -e "${GREEN}✅ OpenAPI spec downloaded successfully${NC}"
else
    echo -e "${RED}❌ Failed to download OpenAPI spec from ${OPENAPI_SPEC_URL}${NC}"
    exit 1
fi