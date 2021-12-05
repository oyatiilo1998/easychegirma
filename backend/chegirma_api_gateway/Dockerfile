FROM golang:1.16 as builder

RUN mkdir -p $GOPATH/src/chegirma_api_gateway
WORKDIR $GOPATH/src/chegirma_api_gateway

COPY . ./

RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/chegirma_api_gateway /
FROM alpine
COPY --from=builder chegirma_api_gateway .
ENTRYPOINT [ "/chegirma_api_gateway" ]