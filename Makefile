# 使用方法，如果要传递版本信息，用下面的方法，如果没有传递，默认优先从version.txt中获取，然后是git tag:
# make VERSION=xxx cmd
ROOT_DIR :=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
VERSION_FILE := version.txt
DEFAULT_VERSION := latest

# 自动获取最大 git tag（语义化排序）
LATEST_GIT_TAG := $(shell git tag --sort=-v:refname | head -n 1)

# 判断 VERSION 是否传入
ifeq ($(origin VERSION), undefined)
    # 尝试从 version.txt 读取
    ifeq ($(wildcard $(VERSION_FILE)),)
        # 如果 git tag 存在，使用它；否则 fallback 到 default
        ifneq ($(strip $(LATEST_GIT_TAG)),)
            VERSION := $(LATEST_GIT_TAG)
        else
            VERSION := $(DEFAULT_VERSION)
        endif
    else
        VERSION := $(shell cat $(VERSION_FILE))
    endif

endif
all:
	@echo "Please specify the cmd!";

test_build:
	#正常编译
	export GOPROXY=https://goproxy.cn && go build -v -ldflags "-X main.version=$(VERSION) -X main.build=`date -u +%Y-%m-%d%H:%M:%S`" -o bin/server server/main.go
	export GOPROXY=https://goproxy.cn && go build -v -ldflags "-X main.version=$(VERSION) -X main.build=`date -u +%Y-%m-%d%H:%M:%S`" -o bin/client client/client_go/main.go
test_server:	
	./bin/server
test_client:
	./bin/client

genprotodoc:
	sudo rm -rf ./proto/docs/md/*
	sudo rm -rf ./proto/docs/html/*
	find proto/src -name "*.proto" | xargs -I {} sudo docker run --rm -v ${ROOT_DIR}:${ROOT_DIR} -w ${ROOT_DIR} hetao29/docker-protoc:latest protoc \
	       --proto_path=proto/src \
	       --proto_path=proto/lib \
	       --doc_out=./proto/doc \
	       --doc_opt=html,"{}".html \
	       "{}"
	find proto/src -name "*.proto" | xargs -I {} sudo docker run --rm -v ${ROOT_DIR}:${ROOT_DIR} -w ${ROOT_DIR} hetao29/docker-protoc:latest protoc \
	       --proto_path=proto/src \
	       --proto_path=proto/lib \
	       --doc_out=./proto/doc \
	       --doc_opt=markdown,"{}".md \
	       "{}"
genproto:
	find proto/out/go/ -name "*.go*" | xargs sudo rm -rf
	find proto/src -name "*.proto" | xargs -I {} sudo docker run --rm -v ${ROOT_DIR}:${ROOT_DIR} -w ${ROOT_DIR} hetao29/docker-protoc:latest protoc \
		--proto_path=proto/src \
		--proto_path=proto/lib  \
		--grpc-gateway_out=logtostderr=true:proto/out/go \
		--plugin=proto-google-common-protos --go_out=proto/out/go --go-grpc_out=proto/out/go \
		"{}"
genjavaproto:
	sudo rm -rf proto/out/java/*
	find proto/src -name "*.proto" | xargs -I {} sudo docker run --rm -v ${ROOT_DIR}:${ROOT_DIR} -w ${ROOT_DIR} hetao29/docker-protoc:latest protoc \
		--proto_path=proto/src \
		--proto_path=proto/lib \
		--plugin=proto-google-common-protos \
		--java_out=proto/out/java \
		--grpc-java_out=proto/out/java \
		--plugin=protoc-gen-grpc-java=/usr/bin/protoc-gen-grpc-java \
		"{}"
genphpproto:
	sudo rm -rf proto/out/php/*
	find proto/src -name "*.proto" | xargs -I {} sudo docker run --rm -v ${ROOT_DIR}:${ROOT_DIR} -w ${ROOT_DIR} hetao29/docker-protoc:latest protoc \
		--php_out=proto/out/php \
		--grpc_out=proto/out/php \
		--proto_path=proto/lib \
		--proto_path=proto/src \
		--plugin=protoc-gen-grpc=/usr/bin/grpc_php_plugin \
		"{}"

##docker 和 k8s 相关
dockerbuild:
	#可选!!!
	#生成 docker image
	@echo "Docker build with VERSION=$(VERSION)"
	@echo $(VERSION) > $(VERSION_FILE)
	sudo DOCKER_BUILDKIT=1 docker build --build-arg VERSION=$(VERSION) --progress=plain . -t hetao29/go-grpc:$(VERSION)
dockerpush:
	@echo "Docker push with VERSION=$(VERSION)"
	@echo $(VERSION) > $(VERSION_FILE)
	sudo docker tag hetao29/go-grpc:$(VERSION) hetao29/go-grpc:$(VERSION)
	sudo docker push hetao29/go-grpc:$(VERSION)
lint:
	#可选!!!
	find . -name "*.go" | xargs -I {} golint "{}"
fmt:
	find modules -name "*.go" | xargs -I {} go fmt "{}"
