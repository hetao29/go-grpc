init:
	#protoc --proto_path=/Users/hetal/grpc/third_party/googleapis --proto_path=. --plugin=proto-google-common-protos --go_out=plugins=grpc:proto_out proto/helloworld.proto
	protoc --include_imports \
               --include_source_info \
               --descriptor_set_out=reservation_service_definition.pb --proto_path=/Users/hetal/grpc/third_party/googleapis --proto_path=. --plugin=proto-google-common-protos --go_out=plugins=grpc:proto_out proto/helloworld.proto
	#https://github.com/grpc/grpc/tree/master/src/php
	#protoc --php_out=proto_out --grpc_out=proto_out --plugin=protoc-gen-grpc=/Users/hetal/grpc/bins/opt/grpc_php_plugin proto/helloworld.proto
build:
	cd server && go build -o ../bin/server .
	cd client && go build -o ../bin/client .
start:	
	./bin/test
