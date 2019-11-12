# 说明

包含 grpc 的示例
1. grpc 生成 php,golang,java 相关代码
2. 支持 grpc 的 restful 的 gateway 的示例
3. 包括 mysql(grom),redis 的操作
4. 规划好了各模块，可以按对应目录结构进行进一步开发

## 目录说明

1. grpc目录，包含官方的 grpc 和 grpc-java 的外部库
2. client_xxx目录，是各种语言的客户端示例代码
3. proto_src目录，proto 的定义文件目录
4. proto_doc目录，proto 的 文档目录，自动生成
5. proto目录，是自动生成的 proto 的目录，不需要手动修改
6. server/client目录，golang 版本的 server 和 client
7. modules目录，是用户自定义目录，用于实现各种功能

## Makefile 说明

```bash
make build
```


## 运行
```bash
make start
```
