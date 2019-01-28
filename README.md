# Go-libRT

Golang 版本的公共运行时框架，包括

- [X] 基本的命令行
- [X] 基于MesagePack 的 RPC 机制
- [X] 日志机制
- [ ] 内置了本地数据队列
- [ ] ECMA / JS 解释器 （暂缓）


## Build

```bash
export GO111MODULE=on
go mod vendor
go build
```

## TODO

- [X] 测试日志能否记录
- [ ] 提供远程日志记录服务
    + [X] 引入 RPC 机制
    + 引入 本地缓存
    + [X] 引入 创建日志专用的命令
- [ ] 提供远程日志汇总服务，可显示节点名称
