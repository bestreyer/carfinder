FROM golang:alpine

RUN apk --no-cache add git make && \
    go get github.com/golang/dep/cmd/dep

RUN mkdir -p /go/src/github.com/bestreyer/carfinder/
WORKDIR /go/src/github.com/bestreyer/carfinder/

ADD Gopkg.lock Gopkg.toml ./
RUN dep ensure --vendor-only

COPY ./ ./
RUN make build_production

FROM alpine:latest

RUN mkdir -p /opt
WORKDIR /opt

COPY --from=0 /go/src/github.com/bestreyer/carfinder/build/carfinder .
EXPOSE 80
ENTRYPOINT ["./carfinder", "server", "start", "-addr=0.0.0.0:80"]
