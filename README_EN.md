# glt

`glt` full name `go-leo-tech`, an efficient collection of tools for summarizing work


## v0.1.0 [2023-08-01]
### cache
- [x] `SafeMemoryCache`: memory based secure and efficient caching.
### task
- [x] `Worker`: efficient task queue, decoupled through chan, producer consumer model, open to the public waiting for completion interface.
- [x] `Worker Group`: expand efficient `Worker Group` based on `Worker`, open hash classification and waiting interfaces to the public.
- [x] `Delay` add delayed tasks,adding too many delayed tasks to the current version may lead to memory overflow errors, as it is related to the operation of adding a quantity online. This is because when there are too many coroutines, each coroutine creates a timer, which consumes a large amount of memory resources.
### util
- [x] `String` convert any to string.
- [x] `IsImplements` using reflection to determine interface.

## v0.1.1 [2023-08-02]
attempt to resolve the issue of module not being found.

## V0.1.2 [2023-08-06]
### container/queue
- [x] `PriorityQueue`: Priority queue, determined by `compare` for large and small root heaps, supporting generics.

### container/set
- [x] `Set`: Borrowing the API encapsulated by the map data structure, supporting generics.

### task
- [x] `DelayPool`: Use the `PriorityQueue` data structure and `Worker` to implement a delayed task pool, avoiding creating too many chans at the same time for `tiem. After`.

### cache
- [x] `SafeMemoryCache`: Supports generics.

## TODO
- [x] `Delay` when optimizing for too many delayed tasks, 'time.After(duration)' may cause memory overflow errors.[Optimized for `DelayPool` in v0.1.2]
- [ ] `Set`: Concurrent operations are not supported.