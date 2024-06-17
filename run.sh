#!/bin/bash

GENERATOR_IMAGE="openapitools/openapi-generator-cli:v6.6.0"

API="${1-isp}"
ENV="${2-prod}"

OPENAPI_SPEC=""
if [[ "$API" == "isp" ]]; then
  OPENAPI_SPEC="http://api.istreamplanet.com/openapi.json"
  if [[ $ENV == "stage" ]]; then
    OPENAPI_SPEC="http://stage.api.istreamplanet.com/openapi.json"
  fi
elif [[ "$API" == "isp-slate" ]]; then
  OPENAPI_SPEC="http://api.istreamplanet.com/docs/slates/openapi.json"
  if [[ $ENV == "stage" ]]; then
    OPENAPI_SPEC="http://stage.api.istreamplanet.com/docs/slates/openapi.json"
  fi
else
  >&2 echo "Unrecognized api $API. Valid options are: isp, isp-slate"
  >&2 echo ""
  exit 1
fi

echo -e "Generating ${API} SDK against ${ENV} (${OPENAPI_SPEC})\n"

SCRIPT_DIR=${PWD}

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
    -i "${OPENAPI_SPEC}" \
    -o go-sdk/${API}

# Logicless templates are dumping extra quotes around enum values, so we've
# added some padding chars and now need to replace the padding + quotes.
# This is preferable to writing an entire huge Java project for a `trim`
# function in the template. I hate this. 🫣
sed -i.bak -E 's/@@@@"([^"]+)"@@@@/\1/g' ./${API}/*.go

# Logicless templates are dumping `example:"null"` on every field, so we've
# got to remove those.
sed -i.bak -E 's/ example:"null"//g' ./${API}/*.go

# Correct an error in the unit tests
sed -i.bak -E 's,"github.com/istreamlabs/go-sdk/isp","github.com/istreamlabs/go-sdk/isp-slate",g' ./isp-slate/**/*.go

# Cleanup all sed backups
find . -name '*.bak' -delete
