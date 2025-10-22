FIXED_OPENAPI_PATH="/workspace/tmp/thehive_openapi_fixed.yaml"
CLIENT_PATH="/workspace/thehive"
OPENAPI_GITIGNORE_PATH="/workspace/.openapi-generator-ignore"
export GO_POST_PROCESS_FILE="/scripts/postprocess_client.sh"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color


echo -e "${BLUE}🔧 Generating Go client from OpenAPI specification...${NC}"
rm -rf "$CLIENT_PATH"/*
cp "$OPENAPI_GITIGNORE_PATH" "$CLIENT_PATH"

echo -e "${GREEN}✅ .openapi-generator-ignore copied successfully${NC}"

echo -e "${GREEN}✅ Starting client generation...${NC}"

/usr/local/bin/docker-entrypoint.sh generate \
  -i "$FIXED_OPENAPI_PATH" \
  -g go \
  --git-user-id StrangeBee \
  --git-repo-id TheHive4Go \
  -o "$CLIENT_PATH" \
  --additional-properties=packageName=thehive,enumClassPrefix=true,packageVersion=1.0.0,withGoMod=false,isGoSubmodule=true,hideGenerationTimestamp=true \
  --name-mappings _id=UnderscoreId,_type=UnderscoreType,_createdBy=UnderscoreCreatedBy,_createdAt=UnderscoreCreatedAt,_updatedBy=UnderscoreUpdatedBy,_updatedAt=UnderscoreUpdatedAt \
  --openapi-normalizer KEEP_ONLY_FIRST_TAG_IN_OPERATION=true \
  --enable-post-process-file

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Go client generated successfully at ${CLIENT_PATH}${NC}"
else
    echo -e "${RED}❌ Failed to generate Go client${NC}"
    exit 1
fi