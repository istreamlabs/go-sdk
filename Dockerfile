FROM node:latest

WORKDIR /go-sdk

COPY openapi-merge.json openapi-merge.json

ARG SLATES_SPEC
ARG CHANNELAPI_SPEC

ENV SLATES_SPEC=${SLATES_SPEC}
ENV CHANNELAPI_SPEC=${CHANNELAPI_SPEC}

RUN echo ${CHANNELAPI_SPEC}

RUN npm install curl

# Grab ChannelAPI and Slates openapi specs and merge them into "output.json"
# config for the openapi-merge-cli is in `openapi-merge.json`
RUN curl ${CHANNELAPI_SPEC} > channelapi.json 
RUN curl ${SLATES_SPEC} > slates.json
RUN ["npx", "openapi-merge-cli"]

FROM openapitools/openapi-generator-cli:v6.6.0

WORKDIR /go-sdk

COPY .generator.yaml .generator.yaml
COPY templates templates
COPY --from=0 /go-sdk/output.json output.json

CMD ["sh", "-c", "java -jar /opt/openapi-generator/modules/openapi-generator-cli/target/openapi-generator-cli.jar generate -c .generator.yaml -i output.json -g go -o isp --skip-validate-spec --git-user-id=istreamlabs --git-repo-id=go-sdk"]