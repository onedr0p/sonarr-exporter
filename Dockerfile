FROM golang:1.14-alpine as build

ENV GO111MODULE=on \
    CGO_ENABLED=0

RUN apk add --no-cache curl ca-certificates git alpine-sdk upx

WORKDIR /go/src/github.com/onedr0p/sonarr-exporter
COPY . .

RUN export GOOS=$(echo ${TARGETPLATFORM} | cut -d / -f1) \
    && export GOARCH=$(echo ${TARGETPLATFORM} | cut -d / -f2) \
    && GOARM=$(echo ${TARGETPLATFORM} | cut -d / -f3); export GOARM=${GOARM:1} \
    && go mod download \
    && go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o sonarr-exporter ./cmd/sonarr-exporter/ \
    && upx -f --brute sonarr-exporter \
    && chmod +x sonarr-exporter

FROM alpine:3.11
RUN apk add --no-cache ca-certificates tini curl
COPY --from=build /go/src/github.com/onedr0p/sonarr-exporter/sonarr-exporter /usr/local/bin/sonarr-exporter
ENTRYPOINT ["/sbin/tini", "--", "sonarr-exporter"]
