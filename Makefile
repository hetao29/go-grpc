init:
	#protoc -I=proto --go_out=proto_out --php_out=proto_out --js_out=proto_out --java_out=proto_out proto/helloworld.proto
	#protoc -I=proto --go_out=proto_out proto/helloworld.proto
	protoc --go_out=plugins=grpc:proto_out proto/helloworld.proto
build:
	cd server && go build -o ../bin/server .
	cd client && go build -o ../bin/client .
start:	
	./bin/test
