FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
WORKDIR /go/src/slotsevo-admin
COPY server/ ./

RUN go env && go mod tidy && go build -o server .

FROM alpine:latest
LABEL MAINTAINER="0xl1352807232@gmaill.com"

WORKDIR /go/src/slotsevo-admin

COPY --from=0 /go/src/slotsevo-admin/server ./
COPY --from=0 /go/src/slotsevo-admin/config.yaml ./
COPY --from=0 /go/src/slotsevo-admin/resource ./resource

EXPOSE 8888

ENTRYPOINT ./server
