#!/bin/sh

# Regenerate the SDK from the latest OpenAPI. Note that at the time of writing
# this requires some hand-crafted edits which should be automated in the future.

set -e

# KRT 20210823 - the uncommented command only appears to work on Mac; switch commented commands for Ubuntu.
# sudo openapi-generator-cli generate -c .generator.yaml -i http://api.istreamplanet.com/openapi.json -g go -o isp --skip-validate-spec
openapi-generator generate -c .generator.yaml -i http://api.istreamplanet.com/openapi.json -g go -o isp --skip-validate-spec
