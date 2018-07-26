# Hello World QUIC

## Introduction

Quic Server & Client

## Contents

- [Install](#install)
- [Build](#build)
- [Run](#run)
- [cURL](#curl)
- [Chrome](#chrome)

## Install

```bash
go get -u github.com/kardianos/govendor
govendor sync
```

## Build

```bash
./build.sh
```

## Run

```bash
docker-compose up -d
```

## cURL

```bash
go run main.go curl --cacert certs/ca.crt https://localhost:7100/
```

## Chrome

These aren't working.. not sure why not (the `curl` works above)

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --user-data-dir=/tmp/chrome \
  --no-proxy-server \
  --enable-quic \
  --origin-to-force-quic-on=localhost:7100 \
  --allow-insecure-localhost \
  https://localhost:7100/

/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --user-data-dir=/tmp/chrome \
  --no-proxy-server \
  --enable-quic \
  --origin-to-force-quic-on=www.example.com:443 \
  --host-resolver-rules='MAP www.example.com:443 localhost:7100' \
  https://www.example.com/

/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --user-data-dir=/tmp/chrome \
  --no-proxy-server \
  --enable-quic \
  --origin-to-force-quic-on=www.example.com:443 \
  --host-resolver-rules='MAP www.example.com:443 127.0.0.1:7100' \
  https://www.example.com/
```