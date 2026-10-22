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

# Fix Header securityScheme: openapi-generator needs an apiKey with a name and a
# location, and the spec declares an `http` scheme. Rewrites the whole block so it
# does not depend on which of `name`, `in` and `scheme` the spec version lists
# (6.0 lists no `name`).
awk '
  function flush() { if (in_hdr) { print "      name: X-USER"; print "      in: header"; in_hdr = 0 } }
  /^  securitySchemes:$/ { in_sec = 1; print; next }
  in_sec && /^[^ ]/ { flush(); in_sec = 0 }
  in_sec && /^  [^ ]/ { flush(); in_sec = 0 }
  in_sec && /^    [^ ]/ { flush(); in_hdr = ($0 == "    Header:"); print; next }
  in_hdr && /^      type: / { print "      type: apiKey"; next }
  in_hdr && /^      (scheme|name|in): / { next }
  { print }
  END { flush() }
' "$FIXED_OPENAPI_PATH" > "$FIXED_OPENAPI_PATH.tmp" && mv "$FIXED_OPENAPI_PATH.tmp" "$FIXED_OPENAPI_PATH"
# Strict mode: the generator rejects a Header scheme without name and location,
# but only after it wiped the client. Fail here instead.
if ! awk '
  /^  securitySchemes:$/ { in_sec = 1; next }
  in_sec && /^    Header:$/ { in_hdr = 1; next }
  in_hdr && /^    [^ ]/ { in_hdr = 0 }
  in_hdr && $0 == "      type: apiKey" { t = 1 }
  in_hdr && $0 == "      name: X-USER" { n = 1 }
  in_hdr && $0 == "      in: header"   { i = 1 }
  END { exit (t && n && i) ? 0 : 1 }
' "$FIXED_OPENAPI_PATH"; then
  echo -e "${RED}❌ Strict check failed: the Header securityScheme is not an apiKey named X-USER in a header.${NC}" >&2
  echo -e "${RED}   The awk no longer matches the spec's securitySchemes layout. Re-audit it before shipping.${NC}" >&2
  exit 1
fi
echo -e "${GREEN}✅ Header securityScheme fixed successfully${NC}"

# Flatten the inline oneOf of the `sort` query parameter (List actions scoped to a
# case or alert, 6.0): openapi-generator names it OneOfFieldDirectionmapstring and
# never declares it, so the client does not compile. The items become untyped
# (interface{}): a field name, `+field`, or a {field: direction} map all fit.
awk '
  /^      - name: / { in_sort = ($0 == "      - name: sort"); skipping = 0 }
  in_sort && /^                oneOf:$/ { skipping = 1; next }
  skipping && /^                [ -]/ { next }
  { skipping = 0; print }
' "$FIXED_OPENAPI_PATH" > "$FIXED_OPENAPI_PATH.tmp" && mv "$FIXED_OPENAPI_PATH.tmp" "$FIXED_OPENAPI_PATH"
if awk '
  /^      - name: / { in_sort = ($0 == "      - name: sort") }
  in_sort && /oneOf:/ { found = 1 }
  END { exit found ? 0 : 1 }
' "$FIXED_OPENAPI_PATH"; then
  echo -e "${RED}❌ Strict check failed: a sort query parameter still declares an inline oneOf, which generates an undeclared Go type.${NC}" >&2
  exit 1
fi
echo -e "${GREEN}✅ Inline oneOf of the sort query parameter flattened${NC}"

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

# Relax required fields that only newer TheHive versions send, so the generated
# SDK accepts payloads from older ones: external, type, forExternal and imports
# (5.6, absent on 5.5), extraData on actions, isCaseSensitive on observable types
# and hidden on tags (5.8), requiresImpact on case statuses (6.0). They are absent
# on the versions before, and no consumer of the SDK actually reads them. Dropping them from each
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
  /^    OutputAction:$/         { scope = "OutputAction"; print; next }
  /^    OutputCaseStatus:$/     { scope = "OutputCaseStatus"; print; next }
  /^    OutputObservableType:$/ { scope = "OutputObservableType"; print; next }
  /^    OutputTag:$/            { scope = "OutputTag"; print; next }
  in_block() && /^    [A-Z][a-zA-Z_]*:$/ { scope = ""; print; next }
  in_block() && scope == "OutputComment"      && /^      - external$/      { next }
  in_block() && scope == "OutputObservable"   && /^      - external$/      { next }
  in_block() && scope == "OutputAttachment"   && /^      - external$/      { next }
  in_block() && scope == "OutputProfile"      && /^      - type$/          { next }
  in_block() && scope == "OutputProfile"      && /^      - forExternal$/   { next }
  in_block() && scope == "OutputPublicStatus" && /^      - imports$/       { next }
  in_block() && scope == "OutputAction"       && /^      - extraData$/     { next }
  in_block() && scope == "OutputCaseStatus"   && /^      - requiresImpact$/ { next }
  in_block() && scope == "OutputObservableType" && /^      - isCaseSensitive$/ { next }
  in_block() && scope == "OutputTag"          && /^      - hidden$/        { next }
  { print }
' "$FIXED_OPENAPI_PATH" > "$FIXED_OPENAPI_PATH.tmp" && mv "$FIXED_OPENAPI_PATH.tmp" "$FIXED_OPENAPI_PATH"
echo -e "${GREEN}✅ Version-specific required fields relaxed (Output{Comment,Observable,Attachment,Profile,PublicStatus,Action,CaseStatus,ObservableType,Tag})${NC}"

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
relax_assert_removed OutputAction       extraData
relax_assert_removed OutputCaseStatus   requiresImpact
relax_assert_removed OutputObservableType isCaseSensitive
relax_assert_removed OutputTag          hidden
echo -e "${GREEN}✅ Strict check: all 10 relax targets confirmed removed${NC}"

# Inject extraData field into InputQueryPagingOperation, unless the spec declares
# it (6.0 does): a second `extraData` key is a duplicate the generator rejects.
paging_has_extra_data() {
  awk '
    /^    InputQueryPagingOperation:$/ { in_paging = 1; next }
    in_paging && /^    [^ ]/ { in_paging = 0 }
    in_paging && /^        extraData:$/ { n++ }
    END { print n + 0 }
  ' "$FIXED_OPENAPI_PATH"
}
if [ "$(paging_has_extra_data)" = "0" ]; then
  echo -e "${BLUE}🔧 Forcing extraData field into InputQueryPagingOperation...${NC}"
  awk '
    BEGIN {in_paging=0}
    /^    InputQueryPagingOperation:/ {in_paging=1; print; next}
    in_paging==1 && /^      properties:/ {print; print "        extraData:"; print "          type: array"; print "          items:"; print "            type: string"; next}
    in_paging==1 && /^    [^ ]/ {in_paging=0}
    {print}
  ' "$FIXED_OPENAPI_PATH" > "$FIXED_OPENAPI_PATH.tmp" && mv "$FIXED_OPENAPI_PATH.tmp" "$FIXED_OPENAPI_PATH"
  echo -e "${GREEN}✅ extraData field injected into InputQueryPagingOperation${NC}"
else
  echo -e "${GREEN}✅ InputQueryPagingOperation already declares extraData${NC}"
fi
if [ "$(paging_has_extra_data)" != "1" ]; then
  echo -e "${RED}❌ Strict check failed: InputQueryPagingOperation must declare extraData exactly once.${NC}" >&2
  exit 1
fi

# Put back GET /api/v1/connector/cortex/action/{entityType}/{entityId} (CortexAPI
# GetActionByEntity), from its 5.6.2 definition. The 6.0 documentation dropped it,
# but TheHive 6.0.0 still serves it and TheHiveMCP reads the actions it ran through
# it, with one client for every TheHive from 5.5 to 6.0. It goes in front of the
# `actions` path; the step fails when it changes nothing, so a spec that moves the
# anchor is audited instead of silently losing the operation. A spec that declares
# the path again needs nothing.
ACTION_BY_ENTITY_PATH="  /api/v1/connector/cortex/action/{entityType}/{entityId}:"
ACTION_BY_ENTITY_ANCHOR="  /api/v1/connector/cortex/actions:"
if grep -qxF "$ACTION_BY_ENTITY_PATH" "$FIXED_OPENAPI_PATH"; then
  echo -e "${GREEN}✅ The spec declares GET action by entity itself${NC}"
else
  echo -e "${BLUE}🔧 Putting GET action by entity back...${NC}"
  ACTION_BY_ENTITY_BLOCK="$FIXED_OPENAPI_PATH.block"
  cat > "$ACTION_BY_ENTITY_BLOCK" <<'ACTION_BY_ENTITY_EOF'
  /api/v1/connector/cortex/action/{entityType}/{entityId}:
    get:
      tags:
      - Cortex
      summary: Get action by entity
      description: |
        List the responder actions run on an entity, given its type and its id.
        Missing from the published 6.0 documentation, still served by TheHive 6.0.
      operationId: Get action by entity
      parameters:
      - name: entityType
        in: path
        required: true
        schema:
          type: string
      - name: entityId
        in: path
        required: true
        description: The _id of the entity or its 'name' (depends of the entity)
        schema:
          type: string
      responses:
        '200':
          description: Actions run on the entity
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/OutputAction'
        '400':
          description: BadRequest
          content:
            application/json:
              schema:
                type: object
                required:
                - type
                - message
                properties:
                  type:
                    type: string
                  message:
                    type: string
        '401':
          description: AuthenticationError
          content:
            application/json:
              schema:
                type: object
                required:
                - type
                - message
                properties:
                  type:
                    type: string
                  message:
                    type: string
        '403':
          description: AuthorizationError
          content:
            application/json:
              schema:
                type: object
                required:
                - type
                - message
                properties:
                  type:
                    type: string
                  message:
                    type: string
        '404':
          description: NotFoundError
          content:
            application/json:
              schema:
                type: object
                required:
                - type
                - message
                properties:
                  type:
                    type: string
                  message:
                    type: string
        '500':
          description: GenericError
          content:
            application/json:
              schema:
                type: object
                required:
                - type
                - message
                properties:
                  type:
                    type: string
                  message:
                    type: string
      security:
      - ApiKey: []
      - Basic: []
      x-TheHive-scopes:
      - name: TheHive
ACTION_BY_ENTITY_EOF
  awk -v block="$ACTION_BY_ENTITY_BLOCK" -v anchor="$ACTION_BY_ENTITY_ANCHOR" '
    $0 == anchor && !done { while ((getline line < block) > 0) print line; close(block); done = 1 }
    { print }
  ' "$FIXED_OPENAPI_PATH" > "$FIXED_OPENAPI_PATH.tmp"
  rm -f "$ACTION_BY_ENTITY_BLOCK"
  if cmp -s "$FIXED_OPENAPI_PATH.tmp" "$FIXED_OPENAPI_PATH"; then
    rm -f "$FIXED_OPENAPI_PATH.tmp"
    echo -e "${RED}❌ Strict check failed: the spec has no '${ACTION_BY_ENTITY_ANCHOR}' path to put GET action by entity in front of, so the step changed nothing.${NC}" >&2
    exit 1
  fi
  mv "$FIXED_OPENAPI_PATH.tmp" "$FIXED_OPENAPI_PATH"
  if [ "$(grep -cxF "$ACTION_BY_ENTITY_PATH" "$FIXED_OPENAPI_PATH")" != "1" ]; then
    echo -e "${RED}❌ Strict check failed: GET action by entity must be declared exactly once.${NC}" >&2
    exit 1
  fi
  echo -e "${GREEN}✅ GET action by entity put back${NC}"
fi

echo -e "${BLUE}🎉 OpenAPI specification parsing and modification completed successfully!${NC}"
