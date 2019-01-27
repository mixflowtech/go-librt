# (Disk) Queue

内置一个基于 内存/磁盘 混合的数据队列，用于在多个虚拟机之间、多个任务之间传递数据。

- 任务有一个输入队列和多个输出队列
- 输出队列中，默认存在 stdout 和 stderr
- 执行引擎 Engine ，可以选择使用 In-memory Queue ，具体类型可以由最后一个任务指定。
- 队列中的数据 默认为 FlatBuffer 编码
- 队列中，数据处理方式为 CopyOnWrite，即如果不修改，则默认传入下一个组件
- 任务的入口为 tasklet 函数

QUAD
-----------

1. 向磁盘位置和Topic，写入文件
2. 从位置 和 Topic 读出文件
3. 
