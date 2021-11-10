## 使用



参数

| 参数          | 含义                                                         |
| ------------- | ------------------------------------------------------------ |
| -bench regexp | 性能测试，支持表达式对测试函数进行筛选。                     |
| -bench        | . 则是对所有的benchmark函数测试，指定名称则只执行具体测试方法而不是全部 |
| -benchmem     | 性能测试的时候显示测试函数的内存分配的统计信息               |
| －count n     | 运行测试和性能多少此，默认一次                               |
| -run regexp   | 只运行特定的测试函数， 比如-run ABC只测试函数名中包含ABC的测试函数 |
| -timeout t    | 测试时间如果超过t, panic,默认10分钟                          |
| -v            | 显示测试的详细信息，也会把Log、Logf方法的日志显示出来        |
| -cpu          |           todo                                                   |



输出说明

| Benchmark_test-12 | **Benchmark_test** 是测试的函数名 **-12** 表示GOMAXPROCS（线程数）的值为12 |
| ----------------- | ------------------------------------------------------------ |
| 第二列           | 表示一共执行了**多少**次，即**b.N**的值                   |
| 152.0 ns/op       | 表示平均每次操作花费了**152.0纳秒**                          |
| 248B/op           | 表示每次操作申请了**248Byte**的内存申请                      |
| 5 allocs/op       | 表示每次操作申请了**5**次内存                                |






### 参数支持传入一个正则表达式，匹配到的用例
``` shell
go test -bench="Base64$"  
```
###  提升准确度
对于性能测试来说，提升测试准确度的一个重要手段就是增加测试的次数。我们可以使用 -benchtime 和 -count 两个参数达到这个目的

benchmark 的默认时间是 1s，那么我们可以使用 -benchtime 指定为 5s
```shell
go test -bench="Base64$"  -benchtime=5s .
```

除了是时间外，-benchtime还可以是具体的次数。例如，执行 30 次可以用 -benchtime=30x
```shell
go test -bench="Base64$"  -benchtime=50x .
```
-count 参数可以用来设置 benchmark 的轮数。例如，进行 3 轮 benchmark
```shell
go test -bench="Base64$"  -benchtime=50x  -count=3   .
```

### 内存分配情况
-benchmem 参数可以度量内存分配的次数。内存分配次数也性能也是息息相关的，例如不合理的切片容量，将导致内存重新分配，带来不必要的开销。
```shell
go test -bench="Base64$"  -benchtime=50x  -count=3  -benchmem .
```
### 计时启停
```shell
b.ResetTimer() // 重置定时器
b.StopTimer() //暂停计时以及使用 
b.StartTimer() //开始计时
```

