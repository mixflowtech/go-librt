# Framed MessagePack RPC

FMP 是 Keybase 开发的基于 MessagePack 的RPC调用机制，相比 gRPC 用起来相对麻烦些。但是代码量比较小，项目相对可控。

The protocol files are defined using Avro IDL in avdl

更新 avdl 协议，必须在 Linux / MacOS 环境下进行，具体可参考 engine/genprotocol/README.md

便易起见，生成的这部分代码也提交进入代码仓库。需要注意，每次生成工作结束后，应单独更新。
