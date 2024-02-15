FROM openapitools/openapi-generator-cli:v6.6.0

ARG OPENAPI_SPEC
ENV OPENAPI_SPEC=${OPENAPI_SPEC}

WORKDIR /go-sdk

COPY .generator.yaml .generator.yaml
COPY templates templates

CMD ["sh", "-c", "java -jar /opt/openapi-generator/modules/openapi-generator-cli/target/openapi-generator-cli.jar generate -c .generator.yaml -i ${OPENAPI_SPEC} -g go -o isp --skip-validate-spec --git-user-id=istreamlabs --git-repo-id=go-sdk"]
# use this for isp-slate api
# CMD ["sh", "-c", "java -jar /opt/openapi-generator/modules/openapi-generator-cli/target/openapi-generator-cli.jar generate -c .generator.yaml -i ${OPENAPI_SPEC} -g go -o isp-slate --skip-validate-spec --git-user-id=istreamlabs --git-repo-id=go-sdk"]