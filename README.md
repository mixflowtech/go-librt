# Go-libRT

Golang 版本的公共运行时框架，包括

1. 基本的命令行
2. 基于MesagePack 的 RPC 机制
3. 日志机制
4. 内置了本地数据队列
5. JS 解释器


## Build

```bash
export GO111MODULE=on
go mod vendor
go build
```

## TODO

- [X] 测试日志能否记录
- [ ] 提供远程日志记录服务
    + 引入 RPC 机制
    + 引入 本地缓存
    + 引入 创建日志专用的命令
- [ ] 提供远程日志汇总服务，可显示节点名称
