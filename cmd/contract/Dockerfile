FROM golang:1.13-alpine as builder

COPY ./ /Project

ENV GOPROXY https://goproxy.cn,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
	apk add --no-cache git gcc ca-certificates musl-dev linux-headers && \
	cd /Project/cmd/contract && go build -o contract ./

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
	apk add --no-cache ca-certificates && \
	apk add --no-cache tzdata && \
	mkdir /Project/ && mkdir /Project/logs && mkdir /Project/config

COPY --from=builder /Project/cmd/contract/contract /Project/

WORKDIR /Project/

ENTRYPOINT ["./contract"]
