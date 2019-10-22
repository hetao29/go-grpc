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
genproto:
	find proto_src -name "*.proto" | xargs -I {} protoc --include_imports \
               --include_source_info \
               --descriptor_set_out=reservation_service_definition.pb \
	       --proto_path=/Users/hetal/grpc/third_party/googleapis \
	       --proto_path=proto_src \
	       --grpc-gateway_out=logtostderr=true:proto \
	       --plugin=proto-google-common-protos --go_out=plugins=grpc:proto \
	       "{}"
test:
	curl -v http://127.0.0.1:50001/v1/user/profile/get
	curl -v http://127.0.0.1:50001/v1/user/login -X POST -d '{"name":"x","password":"P"}'
