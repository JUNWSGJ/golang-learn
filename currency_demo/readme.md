Concurrency is not Parallelism.


并发：同一个CPU通过时分共享运行多个任务。
并行：多个CPU同时执行多个任务，每个任务运行在单独的CPU。


## 1. Keep yourself busy or do the work  yourself


## 2. 把异步执行函数的决定权交给该函数的调用方。
设计一个查找指定目录下的所有文件的API：ListDirectory 三种设计:

### 第一种，返回string的切片包含当前目录下所有文件  
`func ListDirectory(dir string) ([]string, error)`  
缺点： 当目录下存在大量文件时，扫描比较耗时，会导致当前线程卡住

### 第二种，返回一个string 的通道 
`func ListDirectory(dir string) ([]string, error)`  
缺点：这种方式一般是函数内部切启动一个goroutine， 调用者无法控制这个goroutine的启停

### 第三种，传入一个回调函数去处理结果，
`func ListDirectory(dir string) ([]string, error)`  
由调用者去控制什么时候启动或结束goroutine，推荐的做法


## 3. 