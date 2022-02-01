all:
	@echo "Please specify the cmd!";
start:
	docker stack deploy --with-registry-auth -c docker-compose.yml com_mxiqi_microservice_main
stop:
	docker stack rm com_mxiqi_microservice_main
test_build:
	#正常编译
	export GOPROXY=https://goproxy.cn && cd server && go build -v -ldflags "-X main.version=1.0.0 -X main.build=`date -u +%Y-%m-%d%H:%M:%S`" -o ../bin/server .
	export GOPROXY=https://goproxy.cn && cd client && go build -v -ldflags "-X main.version=1.0.0 -X main.build=`date -u +%Y-%m-%d%H:%M:%S`" -o ../bin/client .
test_server:	
	./bin/server
test_client:
	./bin/client

###下面是和 protobuf 协议相关的操作，除非更新 proto，一般情况不需要安装
initprotoc:
	#安装protobuf 相关工具，用来生成 proto 目录里的代码，只有需要更新协议的才需要安装
	#https://github.com/grpc-ecosystem/grpc-gateway
	#ubuntu18+
	sudo apt-get install autoconf automake libtool curl protobuf-compiler-grpc golang-grpc-gateway golang-google-genproto-dev golang-google-grpc-dev
	wget https://github.com/googleapis/googleapis/archive/master.zip -O googleapis-master.zip && unzip googleapis-master.zip
	#ubnutn14
	#sudo apt-get install autoconf automake libtool curl
	#export GOPROXY=https://goproxy.cn && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	#export GOPROXY=https://goproxy.cn && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	#export GOPROXY=https://goproxy.cn && go get -u github.com/golang/protobuf/protoc-gen-go
	#git submodule init
	#git submodule update grpc/grpc
	#cd grpc/grpc && git submodule init && git submodule update && make
	#git submodule update grpc/protobuf
	#REALPREFIX=$(realpath .) && cd grpc/protobuf && git submodule init && git submodule update && ./autogen.sh && ./configure --prefix=$$REALPREFIX/grpc/tmp && make && make check && sudo make install && sudo ldconfig
initjava:
	#可选!!!
	# 要生成 protobuf 的java 语言才需要安装，执行前，需要先 make initprotoc，不然编译不通过
	git submodule init
	git submodule update grpc/grpc-java
	#https://github.com/grpc/grpc-java/tree/master/compiler
	REALPREFIX=$(realpath .) && cd grpc/grpc-java/compiler && export CXXFLAGS="-I$$REALPREFIX/grpc/tmp/include/" LDFLAGS="-L$$REALPREFIX/grpc/tmp/lib/" && echo $$REALPREFIX && ../gradlew java_pluginExecutable
	#cd grpc/grpc-java/compiler && export CXXFLAGS="-I/root/go-grpc/grpc/tmp/include/" LDFLAGS="-L/root/go-grpc/grpc/tmp/lib" && ../gradlew java_pluginExecutable

genprotodoc:
	#可选!!!
	#https://github.com/pseudomuto/protoc-gen-doc
	export GOPROXY=https://goproxy.cn && go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
	find proto_src -name "*.proto" | xargs -I {} protoc \
		   --proto_path=googleapis-master/ \
	       --proto_path=proto_src \
	       --plugin=protoc-gen-doc=`which protoc-gen-doc`\
	       --doc_out=./proto_doc \
	       --doc_opt=html,"{}".html \
	       --grpc-gateway_out=logtostderr=true:proto \
	       --plugin=proto-google-common-protos --go_out=plugins=grpc:proto \
	       "{}"
genproto:
	#可选!!!
	find proto_src -name "*.proto" | xargs -I {} protoc \
			--proto_path=googleapis/ \
	       --proto_path=proto_src \
	       --grpc-gateway_out=logtostderr=true:proto \
	       --plugin=proto-google-common-protos --go_out=plugins=grpc:proto \
	       "{}"
genjavaproto:
	find proto_src -name "*.proto" | xargs -I {} protoc \
		   --proto_path=googleapis-master/ \
	       --proto_path=proto_src \
	       --plugin=proto-google-common-protos \
		   --java_out=proto/java \
		   --grpc-java_out=proto/java \
		   --plugin=protoc-gen-grpc-java=grpc/grpc-java/compiler/build/exe/java_plugin/protoc-gen-grpc-java \
	       "{}"
genphpproto:
	find proto_src -name "*.proto" | xargs -I {} protoc \
		--php_out=proto/php \
		--grpc_out=proto/php \
		--proto_path=googleapis/ \
		--proto_path=proto_src \
		--plugin=protoc-gen-grpc=/usr/bin/grpc_php_plugin \
		"{}"

##docker 和 k8s 相关
dockerbuild:
	#可选!!!
	#生成 docker image
	DOCKER_BUILDKIT=1 docker build --progress=plain . -t hetao29/go-grpc:1.0.1
lint:
	#可选!!!
	find . -name "*.go" | xargs -I {} golint "{}"
fmt:
	find modules -name "*.go" | xargs -I {} go fmt "{}"
