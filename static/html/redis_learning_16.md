# reids学习16 关于日志

### 基本经验：

1. 用列表存储最新出现的日志。

2. 根据日志中消息出现的频率用有序集合存储。


### 适合放进Redis的数据

#### 计数器



