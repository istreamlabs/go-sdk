#!/bin/bash

GENERATOR_IMAGE="openapitools/openapi-generator-cli:v7.21.0"

API="${1-isp}"
ENV="${2-prod}"

OPENAPI_SPEC=""
if [[ "$API" == "isp" ]]; then
  OPENAPI_SPEC="https://api.istreamplanet.com/openapi-3.0.json"
  if [[ $ENV == "stage" ]]; then
    OPENAPI_SPEC="https://stage.api.istreamplanet.com/openapi-3.0.json"
  fi
elif [[ "$API" == "isp-slate" ]]; then
  OPENAPI_SPEC="https://api.istreamplanet.com/docs/slates/openapi-3.0.json"
  if [[ $ENV == "stage" ]]; then
    OPENAPI_SPEC="https://stage.api.istreamplanet.com/docs/slates/openapi-3.0.json"
  elif [[ $ENV == "int" ]]; then
    OPENAPI_SPEC="https://int.api.istreamplanet.com/docs/slates/openapi-3.0.json"
  fi
elif [[ "$API" == "isp-lifecycle" ]]; then
  OPENAPI_SPEC="https://api.istreamplanet.com/state/openapi-3.0.json"
  if [[ $ENV == "stage" ]]; then
    OPENAPI_SPEC="https://stage.api.istreamplanet.com/state/openapi-3.0.json"
  fi
else
  >&2 echo "Unrecognized api $API. Valid options are: isp, isp-slate, isp-lifecycle"
  >&2 echo ""
  exit 1
fi

echo -e "Generating ${API} SDK against ${ENV} (${OPENAPI_SPEC})\n"

SCRIPT_DIR=${PWD}

mkdir -p "$SCRIPT_DIR/spec"
SPEC_FILE="spec/$API.yaml"

echo "Writing $OPENAPI_SPEC to $SPEC_FILE"
curl -s "$OPENAPI_SPEC" | yq -P -o=yaml '.' > "$SCRIPT_DIR/$SPEC_FILE"

if [[ "$API" == "isp" ]]; then
  yq -i '.info.version = "0.0.0"' "$SCRIPT_DIR/$SPEC_FILE"
  # The upstream spec defines components.schemas.BasicAuth, which collides with the
  # generator's built-in BasicAuth in configuration.go. The schema is not referenced
  # anywhere (basic_auth fields are inline); drop it so we do not emit model_basic_auth.go.
  yq -i 'del(.components.schemas.BasicAuth)' "$SCRIPT_DIR/$SPEC_FILE"
else
  yq -i '.info.version = "1.0.0"' "$SCRIPT_DIR/$SPEC_FILE"
fi


# Recreate directory to ensure clean build
rm -rf "${API}"
mkdir "${API}"

if [[ "$API" == "isp" ]]; then
  # Copy required files to directory
  cp ./prerequisites/.openapi-generator-ignore.isp ./${API}/.openapi-generator-ignore
  cp ./prerequisites/convenience._go ./${API}/convenience.go
  cp ./prerequisites/closeable_transport._go ./${API}/closeable_transport.go
  cp ./prerequisites/isp_client._go ./${API}/client.go
else
  # Copy required files to directory
  cp ./prerequisites/.openapi-generator-ignore ./${API}/.openapi-generator-ignore
  cp ./prerequisites/convenience._go ./${API}/convenience.go
fi



# Generate the SDK
docker run --rm \
  --user $(id -u) \
  -v ${SCRIPT_DIR}:/go-sdk \
  "${GENERATOR_IMAGE}" generate \
    -g go -c "go-sdk/.generator.yaml" \
    -i "go-sdk/${SPEC_FILE}" \
    -o go-sdk/${API}

# Logicless templates are dumping extra quotes around enum values, so we've
# added some padding chars and now need to replace the padding + quotes.
# This is preferable to writing an entire huge Java project for a `trim`
# function in the template. I hate this. 🫣
sed -i.bak -E 's/@@@@"([^"]+)"@@@@/\1/g' ./${API}/*.go

# Logicless templates are dumping `example:"null"` on every field, so we've
# got to remove those.
sed -i.bak -E 's/ example:"null"//g' ./${API}/*.go

# Strip leading/trailing regex delimiters added by generator (e.g., "/^...$/")
# Restrict to struct tag context (inside backticks) to avoid accidental matches.
# Note: Use POSIX ERE (no \b). Tested with BSD sed (macOS).
sed -i.bak -E 's/(`[^`]*pattern:")\/([^"]*)\/(")/\1\2\3/g' ./${API}/*.go

# The spec's srt_publications video_encoders and audio_encoders both use the
# same shape as components/schemas/SrtPublicationEncoder, but huma
# non-deterministically inlines one while $ref-ing the other. The generator
# therefore synthesises a duplicate model for whichever is inlined. Drop both
# possible synthetic files and rewrite references to the canonical type.
if [[ "$API" == "isp" ]]; then
  rm -f ./${API}/model_patch_org_channel_request_publishing_srt_publications_inner_video_encoders_inner.go
  rm -f ./${API}/model_patch_org_channel_request_publishing_srt_publications_inner_audio_encoders_inner.go
  sed -i.bak -E 's/PatchOrgChannelRequestPublishingSrtPublicationsInnerVideoEncodersInner/SrtPublicationEncoder/g' ./${API}/*.go
  sed -i.bak -E 's/PatchOrgChannelRequestPublishingSrtPublicationsInnerAudioEncodersInner/SrtPublicationEncoder/g' ./${API}/*.go

  # The upstream spec non-deterministically inlines OriginManifestDefaults on
  # manifest_defaults, fallback_defaults, or alternate_manifest_defaults (huma
  # emits $ref for some and inline schemas for others, varying between runs).
  # Drop whichever synthetic model files the generator created and rewrite all
  # possible synthetic type names back to the canonical OriginManifestDefaults.
  rm -f ./${API}/model_publication_origin_manifest_defaults.go
  rm -f ./${API}/model_publication_origin_fallback_defaults.go
  rm -f ./${API}/model_patch_org_channel_request_publishing_publications_inner_origin_manifest_defaults.go
  rm -f ./${API}/model_patch_org_channel_request_publishing_publications_inner_origin_fallback_defaults.go
  rm -f ./${API}/model_patch_org_channel_request_publishing_publications_inner_origin_alternate_manifest_defaults_value.go
  rm -f ./${API}/model_channeldochuma_origin_manifest_defaults.go
  rm -f ./${API}/model_channeldochuma_origin_fallback_defaults.go
  rm -f ./${API}/model_channeldochuma_origin_alternate_manifest_defaults_value.go
  sed -i.bak -E 's/PublicationOriginManifestDefaults/OriginManifestDefaults/g' ./${API}/*.go
  sed -i.bak -E 's/PublicationOriginFallbackDefaults/OriginManifestDefaults/g' ./${API}/*.go
  sed -i.bak -E 's/PatchOrgChannelRequestPublishingPublicationsInnerOriginManifestDefaults/OriginManifestDefaults/g' ./${API}/*.go
  sed -i.bak -E 's/PatchOrgChannelRequestPublishingPublicationsInnerOriginFallbackDefaults/OriginManifestDefaults/g' ./${API}/*.go
  sed -i.bak -E 's/PatchOrgChannelRequestPublishingPublicationsInnerOriginAlternateManifestDefaultsValue/OriginManifestDefaults/g' ./${API}/*.go
  sed -i.bak -E 's/ChanneldochumaOriginManifestDefaults/OriginManifestDefaults/g' ./${API}/*.go
  sed -i.bak -E 's/ChanneldochumaOriginFallbackDefaults/OriginManifestDefaults/g' ./${API}/*.go
  sed -i.bak -E 's/ChanneldochumaOriginAlternateManifestDefaultsValue/OriginManifestDefaults/g' ./${API}/*.go

  # Scrub deleted synthetic files from the .openapi-generator/FILES manifest
  # so regeneration does not produce spurious diffs in this file.
  FILES_LIST="./${API}/.openapi-generator/FILES"
  sed -i.bak \
    -e '/model_patch_org_channel_request_publishing_srt_publications_inner_video_encoders_inner\.go/d' \
    -e '/model_patch_org_channel_request_publishing_srt_publications_inner_audio_encoders_inner\.go/d' \
    -e '/model_publication_origin_manifest_defaults\.go/d' \
    -e '/model_publication_origin_fallback_defaults\.go/d' \
    -e '/model_patch_org_channel_request_publishing_publications_inner_origin_manifest_defaults\.go/d' \
    -e '/model_patch_org_channel_request_publishing_publications_inner_origin_fallback_defaults\.go/d' \
    -e '/model_patch_org_channel_request_publishing_publications_inner_origin_alternate_manifest_defaults_value\.go/d' \
    -e '/model_channeldochuma_origin_manifest_defaults\.go/d' \
    -e '/model_channeldochuma_origin_fallback_defaults\.go/d' \
    -e '/model_channeldochuma_origin_alternate_manifest_defaults_value\.go/d' \
    "$FILES_LIST"
fi

# Correct an error in the unit tests
sed -i.bak -E 's,"github.com/istreamlabs/go-sdk/v2/isp","github.com/istreamlabs/go-sdk/v2/isp-slate",g' ./isp-slate/**/*.go
sed -i.bak -E 's,"github.com/istreamlabs/go-sdk/v2/isp","github.com/istreamlabs/go-sdk/v2/isp-lifecycle",g' ./isp-lifecycle/**/*.go

# Cleanup all sed backups
find . -name '*.bak' -delete

# Our custom model_simple.mustache never references "bytes" or "fmt", but the
# stock model.mustache pessimistically imports them for any model with required
# fields, so generated output ships with unused imports that break `go build`.
# goimports drops unused imports and tidies formatting; this is what the
# generator's own GO_POST_PROCESS_FILE info log nudges you toward.
# Version is pinned so the toolchain is reproducible across machines / CI.
go run golang.org/x/tools/cmd/goimports@v0.24.0 -w ./${API}

echo ""
echo "git status -s"
git status -s
