# Logger Sub System

在 KeyBase 的开发中，Logger 是最先开发的模块。

提供了 SetExternalHandler 机制，允许将日志数据汇总，并发往别处。

目前，Keybase 代码基中，还未发现日志服务的相关代码。

## Usage

```go
    // "test" as the module name
    log := logger.New("test")   
    // "fancy", the ouput format, it's default when debug is T
    // the last string param , will used as filename
    log.Configure("fancy", true, "")    
    log.Info("KBFS version %s","1.0")
```

## CLI

日志服务相关的子命令分两种

### *Client*

> $(CMD) log send -d --level <debug|info|notice|warn|error|critical|fatal> $message $remote_host

如果没有给出 remote host, 则使用本地默认的日志服务，进行输出。

### *Server*

> $(CMD) log serve --listen $laddr

启动日志采集服务，默认端口为 13399，默认监听 0.0.0.0 (理论上应该是 localhost，但是考虑到本身就是日志服务...)

目前直接记入文件

### *Control*

> $(CMD) log watch

TBD: 实时监听日志

> $(CMD) log query --since $offset $N

TBD: 从查询日志

<hr>

### Changes from Keybase

1. move VDebugLog from libkb to logger.