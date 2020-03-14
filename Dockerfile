FROM golang:1.13-alpine as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

# hadolint ignore=DL3018
RUN apk add --no-cache curl ca-certificates git

WORKDIR /go/src/github.com/onedr0p/sonarr-exporter
COPY . .
RUN chmod +x build.sh \
    && sh build.sh

FROM alpine:3.11

# hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates tini curl

COPY --from=build /go/src/github.com/onedr0p/sonarr-exporter/sonarr-exporter /usr/local/bin/sonarr-exporter
RUN chmod +x /usr/local/bin/sonarr-exporter

ENTRYPOINT ["/sbin/tini", "--", "sonarr-exporter"]