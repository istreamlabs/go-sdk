#!/bin/sh

# Regenerate the SDK from the latest OpenAPI.

set -e

if (echo '5.99.99'; openapi-generator version) | sort -V | tail -1 | grep -q '5.99.99'; then
	echo 'Error: openapi-generator 6.0.0 or newer required'
	exit 1
fi

# KRT 20210823 - the uncommented command only appears to work on Mac; switch commented commands for Ubuntu.
# sudo openapi-generator-cli generate -c .generator.yaml -i http://api.istreamplanet.com/openapi.json -g go -o isp --skip-validate-spec
openapi-generator generate -c .generator.yaml -i http://api.istreamplanet.com/openapi.json -g go -o isp --skip-validate-spec
