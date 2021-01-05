#!/bin/sh

# Regenerate the SDK from the latest OpenAPI. Note that at the time of writing
# this requires some hand-crafted edits which should be automated in the future.

openapi-generator generate -c .generator.yaml -i http://stage.api.istreamplanet.com/openapi.json -g go -o isp --skip-validate-spec
