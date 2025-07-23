#first build
FROM golang:1.24-alpine3.21 as builder
LABEL maintainer="hetao<hetao@hetao.name>"
LABEL version="1.0"

WORKDIR /data/go-grpc/
ENV GOPROXY=https://goproxy.cn,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \ 
	&& apk update && apk add git tree \
	&& tree -L 4

COPY . ./

ARG VERSION
ENV APP_VERSION=$VERSION
RUN echo "Building image with version: $VERSION"
RUN	--mount=type=cache,target=/root/.cache/go-build \
	--mount=type=cache,target=/go/pkg/mod \
	go build -v -ldflags "-X main.version=$VERSION -X main.build=`date -u +%Y-%m-%d%H:%M:%S`" -o bin/server ./server

RUN rm -rf /var/lib/apk/*

RUN tree /data/go-grpc -L 2

FROM alpine:3.21 as prod

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \ 
	&& apk update && apk add ca-certificates

WORKDIR /data/go-grpc/

COPY --from=builder /data/go-grpc/bin ./bin/

COPY etc /data/go-grpc/etc/

HEALTHCHECK --interval=5s --timeout=5s --retries=3 \
    CMD ps aux | grep "server" | grep -v "grep" > /dev/null; if [ 0 != $? ]; then exit 1; fi

EXPOSE 80/tcp
EXPOSE 81/tcp

CMD ["/data/go-grpc/bin/server"]
