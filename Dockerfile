FROM golang:1.10.3-alpine3.8 AS builder
ADD . /go/src/github.com/rms1000watt/hello-world-quic
WORKDIR /go/src/github.com/rms1000watt/hello-world-quic
RUN apk add -U ca-certificates git && \
    go get -u github.com/kardianos/govendor && \
    echo "Running: govendor sync" && \
    govendor sync && \
    echo "Running: go test ./..." && \
    go test ./... && \
    echo "Running: go build" && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo -ldflags '-w -extldflags=-static' -o /hello-world-quic

FROM scratch
COPY --from=builder /hello-world-quic /hello-world-quic
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/hello-world-quic"]
