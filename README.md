# 说明

包含 grpc 的示例
1. grpc 生成 php,golang,java 相关代码
2. 支持 grpc 的 restful 的 gateway 的示例
3. 包括 mysql(grom),redis 的操作
4. 规划好了各模块，可以按对应目录结构进行进一步开发

## 目录说明

1. server，golang 版本的 server
2. client，golang 版本的 client
3. client/xxx，是各种语言的客户端示例代码
4. modules，是用户自定义目录，用于实现各种功能
5. db，是测试的sql文件，可以导入本地 mysql 库
6. proto/src，proto 的定义文件目录
7. proto/doc，proto 的 文档目录，自动生成
8. proto/out，是自动生成的 proto 的目录，不需要手动修改

## Makefile 说明

### 编译
```bash
make build
```

### 运行
```bash
make start
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
