TEMP_BASE_OPENAPI_PATH="/workspace/tmp/thehive_openapi.yaml"
FIXED_OPENAPI_PATH="/workspace/tmp/thehive_openapi_fixed.yaml"

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🔧 Parsing and modifying OpenAPI specification...${NC}"

cp "$TEMP_BASE_OPENAPI_PATH" "$FIXED_OPENAPI_PATH"

# Fix BasicAuth to BasicAuthCredentials
sed -i '27744s/^    BasicAuth:$/    BasicAuthCredentials:/' "$FIXED_OPENAPI_PATH"
sed -i "27733s|#/components/schemas/BasicAuth|#/components/schemas/BasicAuthCredentials|" "$FIXED_OPENAPI_PATH"
sed -i "27740s|basic: '#/components/schemas/BasicAuth'|basic: '#/components/schemas/BasicAuthCredentials'|" "$FIXED_OPENAPI_PATH"
echo -e "${GREEN}✅ BasicAuth fix applied successfully${NC}"

# Replace "Api Key:" with "ApiKey:"
sed -i 's/Api Key:/ApiKey:/g' "$FIXED_OPENAPI_PATH"
echo -e "${GREEN}✅ Api Key fix applied successfully${NC}"

# Remove ProxyConfig default section
sed -i '/    ProxyConfig:/,/^    [^ ]/ { /      default:/,/        state: disabled/d; }' "$FIXED_OPENAPI_PATH"
echo -e "${GREEN}✅ ProxyConfig default section removed successfully${NC}"

# Fix Header securityScheme - change from http to apiKey and fix name
sed -i '/  securitySchemes:/,/^[^ ]/ {
  /    Header:/,/^    [^ ]/ {
    s/      type: http/      type: apiKey/
    s/      name: Header/      name: X-USER/
    /      scheme: bearer/d
    /      name: X-USER/a\
      in: header
  }
}' "$FIXED_OPENAPI_PATH"
echo -e "${GREEN}✅ Header securityScheme fixed successfully${NC}"

echo -e "${BLUE}🎉 OpenAPI specification parsing and modification completed successfully!${NC}"
