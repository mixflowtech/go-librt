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


### Changes from Keybase

1. move VDebugLog from libkb to logger.