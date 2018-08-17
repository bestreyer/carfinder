FROM golang:alpine

RUN apk --no-cache add git make && \
    go get github.com/golang/dep/cmd/dep && \
    go get github.com/oxequa/realize

RUN mkdir -p /go/src/github.com/bestreyer/carfinder/
WORKDIR /go/src/github.com/bestreyer/carfinder/

COPY entrypoint.sh /bin/entrypoint.sh
RUN chmod +x /bin/entrypoint.sh

VOLUME ["/go/src/github.com/bestreyer/carfinder"]
EXPOSE 8888
ENTRYPOINT ["/bin/entrypoint.sh"]
