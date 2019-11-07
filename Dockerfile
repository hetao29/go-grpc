FROM golang:1.13-alpine as builder
LABEL maintainer="hetao<hetao@hetao.name>"
LABEL version="1.0"

WORKDIR /data/go-grpc/

COPY . ./

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \ 
	&& apk update && apk add git tree \
	&& tree -L 4 && export GOPROXY=https://goproxy.cn && cd server && go build -v -o ../bin/server . \
	&& rm -rf /var/lib/apk/*

FROM alpine:3.10 as prod

RUN apk --no-cache add ca-certificates

WORKDIR /data/go-grpc/

COPY --from=builder /data/go-grpc/bin/ ./

COPY etc /data/go-grpc/

HEALTHCHECK --interval=5m --timeout=3s CMD curl -f http://localhost/ || exit 1

EXPOSE 80/tcp

CMD ["/data/go-grpc/bin/server"]
