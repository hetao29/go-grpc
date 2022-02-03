ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
all:
	@echo "Please specify the cmd!";

test_build:
	#正常编译
	export GOPROXY=https://goproxy.cn && cd server && go build -v -ldflags "-X main.version=1.0.0 -X main.build=`date -u +%Y-%m-%d%H:%M:%S`" -o ../bin/server .
	export GOPROXY=https://goproxy.cn && cd client && go build -v -ldflags "-X main.version=1.0.0 -X main.build=`date -u +%Y-%m-%d%H:%M:%S`" -o ../bin/client .
test_server:	
	./bin/server
test_client:
	./bin/client

initjava:
	#可选!!!
	# 要生成 protobuf 的java 语言才需要安装，执行前，需要先 make initprotoc，不然编译不通过
	git submodule init
	git submodule update grpc/grpc-java
	#https://github.com/grpc/grpc-java/tree/master/compiler
	REALPREFIX=$(realpath .) && cd grpc/grpc-java/compiler && export CXXFLAGS="-I$$REALPREFIX/grpc/tmp/include/" LDFLAGS="-L$$REALPREFIX/grpc/tmp/lib/" && echo $$REALPREFIX && ../gradlew java_pluginExecutable
	#cd grpc/grpc-java/compiler && export CXXFLAGS="-I/root/go-grpc/grpc/tmp/include/" LDFLAGS="-L/root/go-grpc/grpc/tmp/lib" && ../gradlew java_pluginExecutable

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
		--plugin=proto-google-common-protos --go_out=plugins=grpc:proto/out/go \
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
	sudo DOCKER_BUILDKIT=1 docker build --progress=plain . -t hetao29/go-grpc:1.0.1
lint:
	#可选!!!
	find . -name "*.go" | xargs -I {} golint "{}"
fmt:
	find modules -name "*.go" | xargs -I {} go fmt "{}"
