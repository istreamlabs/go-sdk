#!/bin/bash

OPENAPI_SPEC=""
API=$1
ENV=$2

MAIN_OPENAPI_SPEC="http://api.istreamplanet.com/openapi.json"
SLATE_OPENAPI_SPEC="http://api.istreamplanet.com/docs/slates/openapi.json"

if [[ $ENV == "stage" ]]; then
  MAIN_OPENAPI_SPEC="http://stage.api.istreamplanet.com/openapi.json"
  SLATE_OPENAPI_SPEC="http://stage.api.istreamplanet.com/docs/slates/openapi.json"
fi

if [[ "$API" == "isp" ]]; then
  OPENAPI_SPEC="${MAIN_OPENAPI_SPEC}"
elif [[ "$API" == "isp-slates" ]]; then
  OPENAPI_SPEC="${SLATE_OPENAPI_SPEC}"
else
  >&2 echo "Unrecognized api $API"
  exit 1
fi

echo -e "Generating ${API} SDK against ${OPENAPI_SPEC}\n"

SCRIPT_DIR=${PWD}

# Recreate directory to ensure clean build
rm -rf "${API}"
mkdir "${API}"

# Copy required files to directory
cp ./prerequisites/.openapi-generator-ignore ./${API}/.openapi-generator-ignore
cp ./prerequisites/convenience._go ./${API}/convenience.go
cp ./prerequisites/client._go ./${API}/client.go

# build sdk generation image
docker build -q -t generate-sdk . \
  --no-cache \
  --build-arg OPENAPI_SPEC="${OPENAPI_SPEC}" \
  --build-arg OUT=sdk

# generate sdk
docker run --rm \
  -u "$(id -u):$(id -g)" \
  -v ${SCRIPT_DIR}/${API}:/go-sdk/sdk \
  generate-sdk

if [[ "$GITHUB_ACTIONS" = "true" ]]; then
  # Logicless templates are dumping extra quotes around enum values, so we've
  # added some padding chars and now need to replace the padding + quotes.
  # This is preferable to writing an entire huge Java project for a `trim`
  # function in the template. I hate this. 🫣
  sed -i -E 's/@@@@"([^"]+)"@@@@/\1/g' ./${API}/*.go

  # Logicless templates are dumping `example:"null"` on every field, so we've
  # got to remove those.
  sed -i -E 's/ example:"null"//g' ./${API}/*.go
else
  sed -i '' -E 's/@@@@"([^"]+)"@@@@/\1/g' ./${API}/*.go
  sed -i '' -E 's/ example:"null"//g' ./${API}/*.go
fi
