# 测试说明

用来测试pub/sub模式

## pub/sub模式的特点

1. 可以将一个消息同时发给多个消费者
2. 如果没有消费者，生产者的消息将会被丢弃

## Routing模式的适用场景

适用于一个事件或消息需要被多个消费者消费的场景，如用户登录事件，日志消息等

## 如何实施

声明交换机是需要指明routing类型为fanout

```go
err = ch.ExchangeDeclare(
		"msg_fanout", // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
```

## 测试用例

1. 启动多个消费者，然后再启动生产者，是否生产者的消息可以被所有消费者所
2. 只启动生产者，不启动消费者，生产者发送完消息后，在启动消费者，验证是否生产者的消息可以按照work queue的模式公平分发给不同的worker

## 测试结论

**在route direct模式下，消费者指定队列为相同的名字，则效果相当于work queue模式**

## 总结

pub/sub模式相当于广播模式，同一交换机上的订阅者都可以接收到同一个消息的副本