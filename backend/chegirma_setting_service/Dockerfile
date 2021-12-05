FROM golang:1.15 as builder

RUN mkdir -p $GOPATH/src/chegirma_setting_service
WORKDIR $GOPATH/src/chegirma_setting_service

COPY . ./

RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/chegirma_setting_service /
FROM alpine
COPY --from=builder chegirma_setting_service .
ENTRYPOINT [ "/chegirma_setting_service" ]