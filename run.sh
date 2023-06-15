#!/bin/bash

# get script dir
SCRIPT_DIR="${0%/*}"

# expand relative paths
SCRIPT_DIR=$(readlink -f $SCRIPT_DIR)

docker build -t generate-sdk .
docker run --rm -it -v ${SCRIPT_DIR}/isp:/go-sdk/isp generate-sdk
