init:
	find proto_src -name "*.proto" | xargs -I {} protoc --include_imports \
               --include_source_info \
               --descriptor_set_out=reservation_service_definition.pb \
	       --proto_path=/Users/hetal/grpc/third_party/googleapis \
	       --proto_path=proto_src \
	       --plugin=proto-google-common-protos --go_out=plugins=grpc:proto \
	       "{}"
build:
	export GOPROXY=https://goproxy.cn && cd server && go build -v -o ../bin/server .
	export GOPROXY=https://goproxy.cn && cd client && go build -o ../bin/client .
start:	
	./bin/test
