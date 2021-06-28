1. 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。 尝试举例其应用。

TCP传输是基于字节流的，只保证了数据传输的可靠和有序。

处理粘包和半包问题，本质上是为了找出消息的边界。

* fix length 应用层约定好每次传输的消息都是固定长度，如果需要发送的消息达不到固定长度可以采用补0的方式。 这种方式简单， 但浪费空间。

* delimiter based 应用层约定好消息之间的分隔符，用分隔符来划分消息边界。 这种方式实现简单，也没有浪费空间，缺点是内容本身如果出现分隔符时需要转义，所以需要扫描内容。

* length field based frame decoder 固定长度字段存内容的元数据信息，比如说内容长度等信息。 应用层先解析固定长度的字段获取长度，然后读取后续内容，能够精确定位消息内容数据，也不需要转义。

这种方式使用最为广泛，像goim

2. 实现一个从 socket connection 中解码出 goim 协议的解码器。

* server.go 启动socket_server代码
* client.go 启动socket_client代码
* proto/proto.go goim协议消息和字节流的转换








