FROM openapitools/openapi-generator-cli:v6.6.0

WORKDIR /go-sdk

COPY .generator.yaml .generator.yaml
COPY templates templates

CMD ["java", "-jar", "/opt/openapi-generator/modules/openapi-generator-cli/target/openapi-generator-cli.jar",  "generate", "-c", ".generator.yaml", "-i", "http://api.istreamplanet.com/openapi.json", "-g", "go", "-o", "isp", "--skip-validate-spec"]
