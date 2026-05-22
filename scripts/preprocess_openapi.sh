TEMP_BASE_OPENAPI_PATH="/workspace/tmp/thehive_openapi.yaml"
FIXED_OPENAPI_PATH="/workspace/tmp/thehive_openapi_fixed.yaml"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🔧 Parsing and modifying OpenAPI specification...${NC}"

cp "$TEMP_BASE_OPENAPI_PATH" "$FIXED_OPENAPI_PATH"

# Fix BasicAuth schema name to BasicAuthCredentials (avoids conflict with Go http library)
# Use pattern-based replacement (no hardcoded line numbers) to handle any spec version
sed -i 's|^\(    \)BasicAuth:$|\1BasicAuthCredentials:|' "$FIXED_OPENAPI_PATH"
sed -i "s|#/components/schemas/BasicAuth'|#/components/schemas/BasicAuthCredentials'|g" "$FIXED_OPENAPI_PATH"
sed -i "s|#/components/schemas/BasicAuth\"|#/components/schemas/BasicAuthCredentials\"|g" "$FIXED_OPENAPI_PATH"
echo -e "${GREEN}✅ BasicAuth fix applied successfully${NC}"

# Replace "Api Key:" with "ApiKey:" (naming convention fix)
sed -i 's/Api Key:/ApiKey:/g' "$FIXED_OPENAPI_PATH"
echo -e "${GREEN}✅ Api Key fix applied successfully${NC}"

# Fix Header securityScheme - change from http to apiKey and fix name
sed -i '/  securitySchemes:/,/^[^ ]/ {
  /    Header:/,/^    [^ ]/ {
    s/      type: http/      type: apiKey/
    s/      name: Header/      name: X-USER/
    /      scheme: basic/d
    /      scheme: bearer/d
    /      name: X-USER/a\
      in: header
  }
}' "$FIXED_OPENAPI_PATH"
echo -e "${GREEN}✅ Header securityScheme fixed successfully${NC}"

# Fix InputQueryFilterOperation to accept generic objects (map[string]interface{})
echo -e "${BLUE}🔧 Converting InputQueryFilterOperation to generic object type...${NC}"
awk '
  BEGIN {in_filter=0}
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

# Relax 5.6-only required fields so the generated SDK accepts payloads from
# 5.5.X too. These fields are present on 5.6 responses but absent on 5.5,
# and no consumer of the SDK actually reads them. Dropping them from each
# schema's `required:` list turns their Go field type from `T` into `*T`
# (and skips the runtime requiredProperties guard) — a nil-check at the
# call site if a future feature needs them, vs total decode failure today.
#
# Each awk pass scopes to one `    ModelName:` block (up to the next
# top-level model `    [A-Z]...:`) and drops the matching `      - field`
# entries from the `required:` array.
echo -e "${BLUE}🔧 Relaxing 5.6-only required fields for 5.5/5.6 cross-version SDK...${NC}"
awk '
  function in_block() { return scope != "" }
  BEGIN { scope = "" }
  /^    OutputComment:$/        { scope = "OutputComment"; print; next }
  /^    OutputObservable:$/     { scope = "OutputObservable"; print; next }
  /^    OutputAttachment:$/     { scope = "OutputAttachment"; print; next }
  /^    OutputProfile:$/        { scope = "OutputProfile"; print; next }
  /^    OutputPublicStatus:$/   { scope = "OutputPublicStatus"; print; next }
  in_block() && /^    [A-Z][a-zA-Z_]*:$/ { scope = ""; print; next }
  in_block() && scope == "OutputComment"      && /^      - external$/      { next }
  in_block() && scope == "OutputObservable"   && /^      - external$/      { next }
  in_block() && scope == "OutputAttachment"   && /^      - external$/      { next }
  in_block() && scope == "OutputProfile"      && /^      - type$/          { next }
  in_block() && scope == "OutputProfile"      && /^      - forExternal$/   { next }
  in_block() && scope == "OutputPublicStatus" && /^      - imports$/       { next }
  { print }
' "$FIXED_OPENAPI_PATH" > "$FIXED_OPENAPI_PATH.tmp" && mv "$FIXED_OPENAPI_PATH.tmp" "$FIXED_OPENAPI_PATH"
echo -e "${GREEN}✅ 5.6-only required fields relaxed (Output{Comment,Observable,Attachment,Profile,PublicStatus})${NC}"

# Strict mode: assert each targeted entry is actually gone from its schema's
# required: list. If a future spec rename or removal silently defeats the
# awk match, fail preprocess instead of shipping a SDK that's strict-
# required-again on 5.5 payloads. The relax must have a visible effect or
# we want to know about it.
relax_assert_removed() {
  local schema=$1 field=$2
  awk -v s="    ${schema}:" -v f="      - ${field}" '
    $0 == s        { in_block = 1; next }
    /^    [A-Z]/   { in_block = 0 }
    in_block && $0 == f { found = 1 }
    END { exit found ? 1 : 0 }
  ' "$FIXED_OPENAPI_PATH" && return 0
  echo -e "${RED}❌ Strict check failed: '${field}' is still listed in ${schema}.required after the relax awk.${NC}" >&2
  echo -e "${RED}   The awk pattern stopped matching (spec format may have changed). Re-audit the relax block before shipping.${NC}" >&2
  exit 1
}

relax_assert_removed OutputComment      external
relax_assert_removed OutputObservable   external
relax_assert_removed OutputAttachment   external
relax_assert_removed OutputProfile      type
relax_assert_removed OutputProfile      forExternal
relax_assert_removed OutputPublicStatus imports
echo -e "${GREEN}✅ Strict check: all 6 relax targets confirmed removed${NC}"

# Inject extraData field into InputQueryPagingOperation
echo -e "${BLUE}🔧 Forcing extraData field into InputQueryPagingOperation...${NC}"
awk '
  BEGIN {in_paging=0}
  /^    InputQueryPagingOperation:/ {in_paging=1; print; next}
  in_paging==1 && /^      properties:/ {print; print "        extraData:"; print "          type: array"; print "          items:"; print "            type: string"; next}
  in_paging==1 && /^    [^ ]/ {in_paging=0}
  {print}
' "$FIXED_OPENAPI_PATH" > "$FIXED_OPENAPI_PATH.tmp" && mv "$FIXED_OPENAPI_PATH.tmp" "$FIXED_OPENAPI_PATH"
echo -e "${GREEN}✅ extraData field injected into InputQueryPagingOperation${NC}"

echo -e "${BLUE}🎉 OpenAPI specification parsing and modification completed successfully!${NC}"
