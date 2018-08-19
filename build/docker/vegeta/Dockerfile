FROM golang:alpine

RUN apk --no-cache add bash git && \
    go get -u github.com/tsenart/vegeta

COPY loadtest.sh .

ENTRYPOINT ["tail", "-f", "/dev/null"]
