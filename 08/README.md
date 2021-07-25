# 第九周：Go语言实践-网络编程
1. 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用
固定长度（fix length）
每次读取固定长度的消息，如果当前读取到的消息不足指定长度，那么就会等待下一个消息到达后进行补足

基于分隔符（delimiter based）
通过用户指定的分隔符对数据进行粘包和拆包处理

自定义长度帧解码器（length field based frame decoder）
按照参数指定的包长度偏移量数据对接收到的数据进行解码，从而得到目标消息体数据

2. 实现一个从 socket connection 中解码出 goim 协议的解码器。
