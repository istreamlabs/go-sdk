#!/bin/sh

# Regenerate the SDK from the latest OpenAPI. Note that at the time of writing
# this requires some hand-crafted edits which should be automated in the future.

set -e

# KRT 20210823 - the uncommented command only appears to work on Mac; switch commented commands for Ubuntu.
# sudo openapi-generator-cli generate -c .generator.yaml -i http://api.istreamplanet.com/openapi.json -g go -o isp --skip-validate-spec
openapi-generator generate -c .generator.yaml -i http://api.istreamplanet.com/openapi.json -g go -o isp --skip-validate-spec

# Cleanup some files due to multiple tags generating duplicate API signatures.
# TODO: Remove once the OpenAPI is sanitized.
rm isp/api_signaling.go

sed -i.bak '/SignalingApi/d' isp/client.go
rm isp/client.go.bak
