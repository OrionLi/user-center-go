# Go语言示例项目：用户中心

在生产环境中，我们可以通过一个统一的用户服务来实现用户信息的远程调用

该项目使用的是 sqlite 数据库，所以你无需额外其他操作便可以运行该示例项目

通过启动 `main.go` 运行服务端，`client.go` 模拟了 grpc 客户端的调用

文件通过以下命令行指令编译

```bash
protoc -I proto/ --go_out=plugins=grpc:proto proto/user.proto
```

当然，该项目已提供了编译好的代码，在 `proto/userpb` 文件夹下

