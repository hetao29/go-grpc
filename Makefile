all:
	@echo "Please specify the cmd!";
initprotoc:
	#初始化，安装各种依赖
	#protoc 的相关初始工作
	#https://github.com/grpc-ecosystem/grpc-gateway
	export GOPROXY=https://goproxy.cn && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	export GOPROXY=https://goproxy.cn && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	export GOPROXY=https://goproxy.cn && go get -u github.com/golang/protobuf/protoc-gen-go
	git submodule init
	git submodule update grpc/grpc
	cd grpc/grpc && git submodule init && git submodule update && make
initjava:
	#java 语言需要安装的依赖，如果要用 java 语言开发，可能需要考虑
	git submodule init
	git submodule update grpc/grpc-java
	#https://github.com/grpc/grpc-java/tree/master/compiler
	cd grpc/grpc-java/compiler && ../gradlew java_pluginExecutable

dockerbuild:
	#生成 docker image
	docker build . -t hetao29/go-grpc:1.0.0
build:
	#正常编译
	export GOPROXY=https://goproxy.cn && cd server && go build -v -o ../bin/server .
	export GOPROXY=https://goproxy.cn && cd client && go build -o ../bin/client .
start:	
	./bin/test
genprotodoc:
	#https://github.com/pseudomuto/protoc-gen-doc
	find proto_src -name "*.proto" | xargs -I {} protoc \
	       --proto_path=grpc/grpc/third_party/googleapis \
	       --proto_path=proto_src \
	       --plugin=protoc-gen-doc=`which protoc-gen-doc`\
	       --doc_out=./proto_doc \
	       --doc_opt=html,"{}".html \
	       --grpc-gateway_out=logtostderr=true:proto \
	       --plugin=proto-google-common-protos --go_out=plugins=grpc:proto \
	       "{}"
genproto:
	find proto_src -name "*.proto" | xargs -I {} protoc \
	       --proto_path=grpc/grpc/third_party/googleapis \
	       --proto_path=proto_src \
	       --grpc-gateway_out=logtostderr=true:proto \
	       --plugin=proto-google-common-protos --go_out=plugins=grpc:proto \
	       "{}"
genjavaproto:
	find proto_src -name "*.proto" | xargs -I {} protoc \
	       --proto_path=grpc/grpc/third_party/googleapis \
	       --proto_path=proto_src \
	       --plugin=proto-google-common-protos \
		   --java_out=proto/java \
		   --grpc-java_out=proto/java \
		   --plugin=protoc-gen-grpc-java=grpc/grpc-java/compiler/build/exe/java_plugin/protoc-gen-grpc-java \
	       "{}"
genphpproto:
	#https://github.com/grpc/grpc/tree/master/src/php
	find proto_src -name "*.proto" | xargs -I {} protoc \
		--php_out=proto/php \
		--grpc_out=proto/php \
		--proto_path=grpc/grpc/third_party/googleapis \
		--proto_path=proto_src \
		--plugin=protoc-gen-grpc=grpc/grpc/bins/opt/grpc_php_plugin \
		"{}"

test:
	curl -v http://127.0.0.1:50001/v1/user/profile/get
	curl -v http://127.0.0.1:50001/v1/user/login -X POST -d '{"name":"x","password":"P"}'
lint:
	find . -name "*go" | xargs -I {} golint "{}"
	#golint modules/...
	#golint server/*
	#golint client/*
fmt:
	find . -name "*go" | xargs -I {} go fmt "{}"
