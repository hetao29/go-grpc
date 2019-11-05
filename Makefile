build:
	export GOPROXY=https://goproxy.cn && cd server && go build -v -o ../bin/server .
	export GOPROXY=https://goproxy.cn && cd client && go build -o ../bin/client .
start:	
	./bin/test
init:
	#https://github.com/grpc-ecosystem/grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
genprotodoc:
	#https://github.com/pseudomuto/protoc-gen-doc
	find proto_src -name "*.proto" | xargs -I {} protoc \
	       --proto_path=/Users/hetal/grpc/third_party/googleapis \
	       --proto_path=proto_src \
	       --plugin=protoc-gen-doc=`which protoc-gen-doc`\
	       --doc_out=./proto_doc \
	       --doc_opt=html,"{}".html \
	       --grpc-gateway_out=logtostderr=true:proto \
	       --plugin=proto-google-common-protos --go_out=plugins=grpc:proto \
	       "{}"
genproto:
	find proto_src -name "*.proto" | xargs -I {} protoc \
	       --proto_path=/Users/hetal/grpc/third_party/googleapis \
	       --proto_path=proto_src \
	       --grpc-gateway_out=logtostderr=true:proto \
	       --plugin=proto-google-common-protos --go_out=plugins=grpc:proto \
	       "{}"
genjavaproto:
	find proto_src -name "*.proto" | xargs -I {} protoc \
	       --proto_path=/home/hetao/grpc/third_party/googleapis \
	       --proto_path=proto_src \
	       --plugin=proto-google-common-protos \
		   --java_out=proto/java \
		   --grpc-java_out=proto/java \
		   --plugin=protoc-gen-grpc-java=/home/hetao/grpc-java/compiler/build/exe/java_plugin/protoc-gen-grpc-java \
	       "{}"
genphpproto:
	#https://github.com/grpc/grpc/tree/master/src/php
	find proto_src -name "*.proto" | xargs -I {} protoc \
		--php_out=proto/php \
		--grpc_out=proto/php \
		--proto_path=/Users/hetal/grpc/third_party/googleapis \
		--proto_path=proto_src \
		--plugin=protoc-gen-grpc=/Users/hetal/grpc/bins/opt/grpc_php_plugin \
		"{}"

test:
	curl -v http://127.0.0.1:50001/v1/user/profile/get
	curl -v http://127.0.0.1:50001/v1/user/login -X POST -d '{"name":"x","password":"P"}'
lint:
	golint modules/...
	golint server/*
	golint client/*
