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

# Fix date-time fields: replace format: date-time with format: int64 and type: string with type: integer
echo -e "${BLUE}🔧 Converting date-time fields to int64 format...${NC}"

awk '
NR == 1 { prev = $0; next }
{
  if (prev ~ /type: string/ && $0 ~ /format: date-time/) {
    gsub(/type: string/, "type: integer", prev)
    gsub(/format: date-time/, "format: int64")
  } else if (prev ~ /format: date-time/ && $0 ~ /type: string/) {
    gsub(/format: date-time/, "format: int64", prev)
    gsub(/type: string/, "type: integer")
  } else if (prev ~ /format: date-time/) {
    gsub(/format: date-time/, "format: int64", prev)
  }
  print prev
  prev = $0
}
END {
  if (prev ~ /format: date-time/) {
    gsub(/format: date-time/, "format: int64", prev)
  }
  print prev
}
' "$FIXED_OPENAPI_PATH" > "$FIXED_OPENAPI_PATH.tmp" && mv "$FIXED_OPENAPI_PATH.tmp" "$FIXED_OPENAPI_PATH"

echo -e "${GREEN}✅ Date-time fields converted to int64 format successfully${NC}"

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

# Fix InputQueryFilterOperation to accept generic objects (map[string]interface{})
echo -e "${BLUE}🔧 Converting InputQueryFilterOperation to generic object type...${NC}"

# Use awk to replace the entire InputQueryFilterOperation definition
awk '
  BEGIN {in_filter=0; skip_lines=0}
  /^    InputQueryFilterOperation:/ {
    in_filter=1
    print "    InputQueryFilterOperation:"
    print "      type: object"
    print "      additionalProperties: true"
    print "      example:"
    print "        _name: filter"
    print "        _eq:"
    print "          _field: \"severity\""
    print "          _value: 2"
    print "      required:"
    print "      - _name"
    next
  }
  in_filter==1 && /^    [^ ]/ {in_filter=0}
  in_filter==1 {next}
  {print}
' "$FIXED_OPENAPI_PATH" > "$FIXED_OPENAPI_PATH.tmp" && mv "$FIXED_OPENAPI_PATH.tmp" "$FIXED_OPENAPI_PATH"

echo -e "${GREEN}✅ InputQueryFilterOperation converted to generic object type!${NC}"


# Robustly inject extraData field into InputQueryPagingOperation
echo -e "${BLUE}🔧 Forcing extraData field into InputQueryPagingOperation...${NC}"
awk '
  BEGIN {in_paging=0}
  /^    InputQueryPagingOperation:/ {in_paging=1; print; next}
  in_paging==1 && /^      properties:/ {print; print "        extraData:"; print "          type: array"; print "          items:"; print "            type: string"; next}
  in_paging==1 && /^    [^ ]/ {in_paging=0}
  {print}
' "$FIXED_OPENAPI_PATH" > "$FIXED_OPENAPI_PATH.tmp" && mv "$FIXED_OPENAPI_PATH.tmp" "$FIXED_OPENAPI_PATH"
echo -e "${GREEN}✅ extraData field forcibly injected into InputQueryPagingOperation${NC}"

echo -e "${BLUE}🎉 OpenAPI specification parsing and modification completed successfully!${NC}"
