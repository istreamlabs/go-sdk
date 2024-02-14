#!/bin/bash

CHANNELAPI_SPEC="https://api.istreamplanet.com/openapi.json"
SLATES_SPEC="https://api.istreamplanet.com/docs/slates/openapi.json"

ENV=$1
if [[ $ENV == "stage" ]]; then
  CHANNELAPI_SPEC="https://stage.api.istreamplanet.com/openapi.json"
  SLATES_SPEC="https://stage.api.istreamplanet.com/docs/slates/openapi.json"
fi

SCRIPT_DIR=${PWD}

# Recreate isp directory to ensure clean build
rm -rf isp
mkdir isp

# Copy required files to isp directory
cp ./prerequisites/.openapi-generator-ignore ./isp/.openapi-generator-ignore
cp ./prerequisites/convenience._go ./isp/convenience.go
cp ./prerequisites/client._go ./isp/client.go
docker build -t generate-sdk .  --no-cache --build-arg CHANNELAPI_SPEC="${CHANNELAPI_SPEC}" --build-arg SLATES_SPEC="${SLATES_SPEC}"
docker run --rm -v ${SCRIPT_DIR}/isp:/go-sdk/isp generate-sdk

# Logicless templates are dumping extra quotes around enum values, so we've
# added some padding chars and now need to replace the padding + quotes.
# This is preferable to writing an entire huge Java project for a `trim`
# function in the template. I hate this. 🫣
sed -i '' -E 's/@@@@"([^"]+)"@@@@/\1/g' ./isp/*.go

# Logicless templates are dumping `example:"null"` on every field, so we've
# got to remove those.
sed -i '' -E 's/ example:"null"//g' ./isp/*.go
