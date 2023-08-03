#!/bin/bash

OPENAPI_SPEC="http://api.istreamplanet.com/openapi.json"

ENV=$1
if [[ $ENV == "stage" ]]; then
  OPENAPI_SPEC="http://stage.api.istreamplanet.com/openapi.json"
fi

echo -e "Generating SDK against ${OPENAPI_SPEC}\n"

SCRIPT_DIR=${PWD}

# Recreate isp directory to ensure clean build
rm -rf isp
mkdir isp

# Copy required files to isp directory
cp ./prerequisites/.openapi-generator-ignore ./isp/.openapi-generator-ignore
cp ./prerequisites/convenience._go ./isp/convenience.go
docker build -t generate-sdk . --no-cache --build-arg OPENAPI_SPEC="${OPENAPI_SPEC}"
docker run --rm -it -v ${SCRIPT_DIR}/isp:/go-sdk/isp generate-sdk
cp ./prerequisites/client._go ./isp/client.go
