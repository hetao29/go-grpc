# 说明

包含 grpc 的示例
1. 有自动的命令，除了 go get相关的工具外，grpc/protobuf 相关的都安装在 grpc/tmp 目录下，不影响系统
2. grpc 生成 php,golang,java 相关代码
3. 支持 grpc 的 restful 的 gateway 的示例
4. 包括 mysql(grom),redis 的操作
5. 规划好了各模块，可以按对应目录结构进行进一步开发

## 目录说明

1. grpc目录，包含官方的 grpc 和 grpc-java 的外部库
2. client_xxx目录，是各种语言的客户端示例代码
3. proto_src目录，proto 的定义文件目录
4. proto_doc目录，proto 的 文档目录，自动生成
5. proto目录，是自动生成的 proto 的目录，不需要手动修改
6. server/client目录，golang 版本的 server 和 client
7. modules目录，是用户自定义目录，用于实现各种功能
8. db目录，是测试的sql文件，可以导入本地 mysql 库

## Makefile 说明

### 编译
```bash
make build
```

### 运行
```bash
make start
```

### 生成 grpc和protoc的依赖，（可选，如果在生成 protobuf 对应的go/php 文件）
这个命令，会在 tmp 目录下生成protocbuf 相关工具，
```bash
make initprotoc
```

### 生成 java依赖，（可选，根据 protobuf 生成对应的 java 文件）
```bash
make initjava
```

### 根据 protobuf生成 go 文件
```bash
make genproto
```

### 根据 protobuf生成 java 文件

```bash
make genjavaproto
```


### 根据 protobuf生成 php 文件

```bash
make genphpproto
```
