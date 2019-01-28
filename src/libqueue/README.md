# libQueue

设计目标：

1. 高性能的持久化磁盘队列，采用 Memory + mmap 双实现，基于 commitlog / write a head log 机制，对队列数据进行处理。
2. 提供事务一致性，当属于某个事务的写入失败时，整个事务回滚。
3. 对于失败的事务，提供回调接口
4. 事务执行失败的回调接口，应允许程序检视数据，决定是否重试
5. 应加入 master / leader 选举机制，允许 multi-node 构成集群写入（接口上考虑）
6. 写入的数据采用 FlexBuffer / FlatBuffer

实际：

1. 对 bbolt 数据块进行二次封装，使用自增id，模拟磁盘队列。
2. 采用 MessagePack 编码 写入