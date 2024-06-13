# iStreamPlanet SDK for Go

This is the iStreamPlanet SDK for the Go programming language. It is generated from the [iStreamPlanet OpenAPI](https://api.istreamplanet.com/openapi.json) with [OpenAPI Generator](https://openapi-generator.tech/).

See https://istreamlabs.github.io/docs/sdks/golang

# How to run

## Prerequisites

A clean directory is required to ensure defunct files are removed. However, some files do need to be kept to ensure the SDK is generated properly.

This is automatically handled when using `./run.sh isp`

## Requirements

- Docker

Generated SDK will be outputted to the `isp` folder.

```sh
# By default, it generates against prod
./run.sh isp
./run.sh isp-slates

# Generate against stage
./run.sh isp stage
./run.sh isp-slates stage
```

## Verify SDK generated successfully

```sh
# There should be no build errors
go build ./...

# Running the example with valid client id, client secret, and org should be successful.
CLIENT_ID="<YOUR_CLIENT_ID>" CLIENT_SECRET="<YOUR_CLIENT_SECRET>" ORG="<ORG>" go run ./example

# Or use a JWT you already have:
AUTH_HEADER=$(restish auth-header isp) ORG="<ORG>" go run ./example
```
