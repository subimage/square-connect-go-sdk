FROM alpine:latest

COPY square_connect_openapi.patch build/
COPY model_catalog_custom_attribute_value.patch build/
COPY model_catalog_item.patch build/
COPY model_catalog_item_variation.patch build/
COPY open-api-3_square build/

RUN apk add curl jq openjdk11 go && \
    cd build && \
    wget https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.51/swagger-codegen-cli-3.0.51.jar -O swagger-codegen-cli.jar && \
    cat open-api-3_square | jq '.data["attributes"]["json-spec"] | fromjson' > square-connect-openapi.json && \
    #patch < square_connect_openapi.patch && \
    mkdir square-connect-sdk && \
    java -jar ./swagger-codegen-cli.jar generate -i square-connect-openapi.json -l go -o square-connect-sdk && \
    cd square-connect-sdk
    #patch < ../model_catalog_custom_attribute_value.patch  && \
    #patch < ../model_catalog_item.patch && \
    #patch < ../model_catalog_item_variation.patch

# fix some issues with Swagger. TODO: file an issue
RUN cd build/square-connect-sdk && \
    gofmt -s -w . && \
    cd ../..

ENTRYPOINT ["/bin/sh"]
