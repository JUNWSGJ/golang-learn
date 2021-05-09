context 主要是用来跨API传递deadline，取消信号和其它值

这是go定义的Context接口
```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```
