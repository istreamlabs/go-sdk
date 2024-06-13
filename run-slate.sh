#!/bin/bash

OPENAPI_SPEC="http://api.istreamplanet.com/docs/slates/openapi.json"

ENV=$1
if [[ $ENV == "stage" ]]; then
  OPENAPI_SPEC="http://stage.api.istreamplanet.com/docs/slates/openapi.json"
fi

echo -e "Generating SDK against ${OPENAPI_SPEC}\n"

SCRIPT_DIR=${PWD}

# Recreate isp-slates directory to ensure clean build
rm -rf isp-slates
mkdir isp-slates

# Copy required files to isp directory
cp ./prerequisites/.openapi-generator-ignore ./isp-slates/.openapi-generator-ignore
cp ./prerequisites/convenience._go ./isp-slates/convenience.go
cp ./prerequisites/slates_client._go ./isp-slates/client.go

# build sdk generation image
docker build -t generate-sdk . \
  --no-cache \
  --build-arg OPENAPI_SPEC="${OPENAPI_SPEC}" \
  --build-arg OUT=isp-slates

# generate sdk
docker run --rm \
  -u "$(id -u):$(id -g)" \
  -v ${SCRIPT_DIR}/isp-slates:/go-sdk/isp-slates \
  generate-sdk

if [[ "$GITHUB_ACTIONS" = "true" ]]; then
  # Logicless templates are dumping extra quotes around enum values, so we've
  # added some padding chars and now need to replace the padding + quotes.
  # This is preferable to writing an entire huge Java project for a `trim`
  # function in the template. I hate this. 🫣
  sed -i -E 's/@@@@"([^"]+)"@@@@/\1/g' ./isp/*.go

  # Logicless templates are dumping `example:"null"` on every field, so we've
  # got to remove those.
  sed -i -E 's/ example:"null"//g' ./isp/*.go
else
  sed -i '' -E 's/@@@@"([^"]+)"@@@@/\1/g' ./isp/*.go
  sed -i '' -E 's/ example:"null"//g' ./isp/*.go
fi
