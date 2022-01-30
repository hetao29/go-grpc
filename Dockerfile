FROM golang:1.17-alpine as builder
LABEL maintainer="hetao<hetao@hetao.name>"
LABEL version="1.0"

WORKDIR /data/go-grpc/

COPY . ./

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \ 
	&& apk update && apk add git tree \
	&& tree -L 4 && export GOPROXY=https://goproxy.cn && cd server && go build -v -o ../bin/server . \
	&& rm -rf /var/lib/apk/*
RUN tree /data/go-grpc -L 2

FROM alpine:3.15 as prod

RUN apk --no-cache add ca-certificates tree

WORKDIR /data/go-grpc/

COPY --from=builder /data/go-grpc/bin ./bin/

COPY etc /data/go-grpc/etc/
RUN tree /data/go-grpc -L 2

HEALTHCHECK --interval=5s --timeout=5s --retries=3 \
    CMD ps aux | grep "server" | grep -v "grep" > /dev/null; if [ 0 != $? ]; then exit 1; fi

EXPOSE 80/tcp
EXPOSE 81/tcp

CMD ["/data/go-grpc/bin/server"]
