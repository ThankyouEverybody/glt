# glt

`glt` 全名 `go-leo-tech`, 从工作中汇总出高效的工具集合

[English](./README_EN.md)


## v0.1.0 [2023-08-01]
### cache
- [x] `SafeMemoryCache`：基于内存的安全高效缓存。
### task
- [x] `Worker`：高效的任务队列，通过chan解耦，生产者-消费者模型，对外开放等待完成接口。
- [x] `Worker Group`：基于`Worker `扩展出高效的`Worker Group` ，对外开放hash分类和等待接口。
- [x] `Delay`： 添加延迟任务，在当前版本中添加太多延迟任务可能会导致内存溢出错误，这与添加数量有关。这是因为当协程太多时，每个协程都会创建一个计时器，这会消耗大量的内存资源。
### util
- [x] `String`：转换字符串。
- [x] `IsImplements`：使用反射确定接口类型。

## v0.1.1 [2023-08-02]
尝试解决go get 模块找不到的问题。

## v0.1.2 [2023-08-06]

### container/queue

- [x] `PriorityQueue`：优先级队列，大根堆与小根堆根据`compare`决定，支持泛型。

### container/set

- [x] `Set`：借用map数据结构封装的api，支持泛型。

### task

- [x] `DelayPool`：使用`PriorityQueue`数据结构和`Worker`实现延时任务池，避免`tiem.After`同时创建过多chan

### cache

- [x] `SafeMemoryCache`：支持泛型。

## TODO

- [x] `Delay`：当添加过多延迟任务时，`time.After`可能会导致内存溢出的问题。[在v0.1.2中`DelayPool`进行了优化]
- [ ] `Set`：不支持并发操作。