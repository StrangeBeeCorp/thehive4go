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

# Fix timestamp fields: replace format: datetime_ms with format: int64
sed -i 's/format: datetime_ms/format: int64/g' "$FIXED_OPENAPI_PATH"
echo -e "${GREEN}✅ Timestamp format fix applied successfully (datetime_ms -> int64)${NC}"

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

# Fix InputQueryFilterOperation to include all filter operations
echo -e "${BLUE}🔧 Fixing InputQueryFilterOperation to support all filter operations...${NC}"

# Create temporary files for our operation
INPUT_QUERY_FILTER_START=$(grep -n "^    InputQueryFilterOperation:" "$FIXED_OPENAPI_PATH" | cut -d: -f1)
INPUT_QUERY_GENERIC_START=$(grep -n "^    InputQueryGenericOperation:" "$FIXED_OPENAPI_PATH" | cut -d: -f1)

if [ -z "$INPUT_QUERY_FILTER_START" ] || [ -z "$INPUT_QUERY_GENERIC_START" ]; then
    echo -e "${RED}❌ Could not find InputQueryFilterOperation or InputQueryGenericOperation sections${NC}"
    exit 1
fi

# Calculate the line where the InputQueryFilterOperation section ends
END_LINE=$((INPUT_QUERY_GENERIC_START - 1))

# Create the new complete InputQueryFilterOperation definition
cat > /tmp/new_filter_operation.yaml << 'EOF'
    InputQueryFilterOperation:
      example:
        _name: _name
        _eq: "{}"
      properties:
        _eq:
          type: object
        _and:
          items:
            $ref: "#/components/schemas/Filter"
          type: array
        _or:
          items:
            $ref: "#/components/schemas/Filter"
          type: array
        _not:
          $ref: "#/components/schemas/Filter"
        _any:
          type: object
        _like:
          $ref: "#/components/schemas/FieldValue"
        _gt:
          $ref: "#/components/schemas/FieldValue"
        _gte:
          $ref: "#/components/schemas/FieldValue"
        _lt:
          $ref: "#/components/schemas/FieldValue"
        _lte:
          $ref: "#/components/schemas/FieldValue"
        _ne:
          $ref: "#/components/schemas/FieldValue"
        _in:
          $ref: "#/components/schemas/FieldValue"
        _between:
          $ref: "#/components/schemas/FieldValue"
        _contains:
          $ref: "#/components/schemas/FieldValue"
        _endsWith:
          $ref: "#/components/schemas/FieldValue"
        _startsWith:
          $ref: "#/components/schemas/FieldValue"
        _match:
          $ref: "#/components/schemas/FieldValue"
        _id:
          type: string
        _name:
          type: string
      required:
      - _name
EOF

# Create temporary files to hold the parts of the original file
head -n $((INPUT_QUERY_FILTER_START - 1)) "$FIXED_OPENAPI_PATH" > /tmp/before_filter.yaml
tail -n +$INPUT_QUERY_GENERIC_START "$FIXED_OPENAPI_PATH" > /tmp/after_filter.yaml

# Combine the parts with our new filter operation
cat /tmp/before_filter.yaml /tmp/new_filter_operation.yaml /tmp/after_filter.yaml > "$FIXED_OPENAPI_PATH"

# Clean up temporary files
rm -f /tmp/before_filter.yaml /tmp/after_filter.yaml /tmp/new_filter_operation.yaml

echo -e "${GREEN}✅ InputQueryFilterOperation fixed - now supports all filter operations!${NC}"

echo -e "${BLUE}🎉 OpenAPI specification parsing and modification completed successfully!${NC}"
