## async
1. 并行执行函数，统一阻塞等待
2. 并行执行函数，异步回调执行（非阻塞）

## What are the functions?
* async
    * func
      * Parallel（并行执行fns）
        * Add（添加要并行的函数fns）
        * ContinueWith（当并行任务执行完后，以非阻塞方式执行callbacks）
        * Wait（阻塞等待执行完成）


## 同步
```go
var count = 0
_ = Parallel(func() { count += 1}).Add(func() { count += 2 }).Wait() // 阻塞等待，直到两个函数执行完
count *= 2  // 由于阻塞，所以这里最后执行
// count = 6
```

## 异步
```go
var count = 0
Parallel(func() { count += 1}).Add(func() { count += 2 }).ContinueWith(func() {
    count += 3
})

count = 10  // 由于异步，这里会优先执行
time.Sleep(10 * time.Millisecond)
// count = 16
```