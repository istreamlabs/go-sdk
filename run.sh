#!/bin/bash

GENERATOR_IMAGE="openapitools/openapi-generator-cli:v6.6.0"

API="${1-isp}"
ENV="${2-prod}"

OPENAPI_SPEC=""
if [[ "$API" == "isp" ]]; then
  OPENAPI_SPEC="https://api.istreamplanet.com/openapi-3.0.json"
  if [[ $ENV == "stage" ]]; then
    OPENAPI_SPEC="https://stage.api.istreamplanet.com/openapi-3.0.json"
  fi
elif [[ "$API" == "isp-slate" ]]; then
  OPENAPI_SPEC="https://api.istreamplanet.com/docs/slates/openapi-3.0.json"
  if [[ $ENV == "stage" ]]; then
    OPENAPI_SPEC="https://stage.api.istreamplanet.com/docs/slates/openapi-3.0.json"
  elif [[ $ENV == "int" ]]; then
    OPENAPI_SPEC="https://int.api.istreamplanet.com/docs/slates/openapi-3.0.json"
  fi
elif [[ "$API" == "isp-lifecycle" ]]; then
  OPENAPI_SPEC="https://api.istreamplanet.com/state/openapi-3.0.json"
  if [[ $ENV == "stage" ]]; then
    OPENAPI_SPEC="https://stage.api.istreamplanet.com/state/openapi-3.0.json"
  fi
else
  >&2 echo "Unrecognized api $API. Valid options are: isp, isp-slate, isp-lifecycle"
  >&2 echo ""
  exit 1
fi

echo -e "Generating ${API} SDK against ${ENV} (${OPENAPI_SPEC})\n"

SCRIPT_DIR=${PWD}

mkdir -p "$SCRIPT_DIR/spec"
SPEC_FILE="spec/$API.yaml"

echo "Writing $OPENAPI_SPEC to $SPEC_FILE"
curl -s "$OPENAPI_SPEC" | yq -P -o=yaml '.' > "$SCRIPT_DIR/$SPEC_FILE"

if [[ "$API" == "isp" ]]; then
  yq -i '.info.version = "0.0.0"' "$SCRIPT_DIR/$SPEC_FILE"
  # The upstream spec defines components.schemas.BasicAuth, which collides with the
  # generator's built-in BasicAuth in configuration.go. The schema is not referenced
  # anywhere (basic_auth fields are inline); drop it so we do not emit model_basic_auth.go.
  yq -i 'del(.components.schemas.BasicAuth)' "$SCRIPT_DIR/$SPEC_FILE"
else
  yq -i '.info.version = "1.0.0"' "$SCRIPT_DIR/$SPEC_FILE"
fi


# Recreate directory to ensure clean build
rm -rf "${API}"
mkdir "${API}"

# Copy required files to directory
cp ./prerequisites/.openapi-generator-ignore ./${API}/.openapi-generator-ignore
cp ./prerequisites/convenience._go ./${API}/convenience.go
cp ./prerequisites/${API}_client._go ./${API}/client.go

# Generate the SDK
docker run --rm \
  --user $(id -u) \
  -v ${SCRIPT_DIR}:/go-sdk \
  "${GENERATOR_IMAGE}" generate \
    -g go -c "go-sdk/.generator.yaml" \
    -i "go-sdk/${SPEC_FILE}" \
    -o go-sdk/${API}

# Logicless templates are dumping extra quotes around enum values, so we've
# added some padding chars and now need to replace the padding + quotes.
# This is preferable to writing an entire huge Java project for a `trim`
# function in the template. I hate this. 🫣
sed -i.bak -E 's/@@@@"([^"]+)"@@@@/\1/g' ./${API}/*.go

# Logicless templates are dumping `example:"null"` on every field, so we've
# got to remove those.
sed -i.bak -E 's/ example:"null"//g' ./${API}/*.go

# Strip leading/trailing regex delimiters added by generator (e.g., "/^...$/")
# Restrict to struct tag context (inside backticks) to avoid accidental matches.
# Note: Use POSIX ERE (no \b). Tested with BSD sed (macOS).
sed -i.bak -E 's/(`[^`]*pattern:")\/([^"]*)\/(")/\1\2\3/g' ./${API}/*.go

# OpenAPI Generator emits PatchOrgChannelRequestPublishingSrtPublicationsInnerVideoEncodersInner
# for srt_publications[].video_encoders but never generates that model. The spec uses
# components/schemas/SrtPublicationEncoder (same as audio_encoders).
sed -i.bak -E 's/PatchOrgChannelRequestPublishingSrtPublicationsInnerVideoEncodersInner/SrtPublicationEncoder/g' ./${API}/*.go

# Correct an error in the unit tests
sed -i.bak -E 's,"github.com/istreamlabs/go-sdk/v2/isp","github.com/istreamlabs/go-sdk/v2/isp-slate",g' ./isp-slate/**/*.go
sed -i.bak -E 's,"github.com/istreamlabs/go-sdk/v2/isp","github.com/istreamlabs/go-sdk/v2/isp-lifecycle",g' ./isp-lifecycle/**/*.go

# Cleanup all sed backups
find . -name '*.bak' -delete

echo ""
echo "git status -s"
git status -s
