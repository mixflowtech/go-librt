# Note of GoConvey

1. 在测试代码中导入需要使用的 Package

```go
. "github.com/smartystreets/goconvey/convey"
```

2. 将待测试的代码，使用

```go
Convey("TestScanOldLogFiles, also a test of GoConvey.", t, func() {
......
})
```

包裹起来

3.使用 So 进行断言...