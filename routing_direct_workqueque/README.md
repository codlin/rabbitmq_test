# 测试说明

测试在标准routing模式下（routing类型为direct），在声明队列时指定队列名字，以测试此种情况是否和work queue等效

## 模式特点

同work queue

## Routing模式的适用场景

同work queue

## 如何实施

1. 交换机的声明和route direct模式相同
2. 在声明队列时指定名字

```go
q, err := ch.QueueDeclare(
		"route_work_queue", // 此处指定队列名字，测试此种情况是否和work queue等同。
		true,
		false,
		false,
		false,
		nil,
	)
```

### 测试用例
1. 启动多个消费者，然后再启动生产者，是否生产者的消息可以按照work queue的模式公平分发给不同的worker
2. 只启动生产者，不启动消费者，生产者发送完消息后，在启动消费者，验证是否生产者的消息可以按照work queue的模式公平分发给不同的worker

### 测试结论

1. **在标准routing模式下，消费者指定队列为相同的名字，则效果相当于work queue模式**

## 总结

work queue没有声明交换机，实际上使用的系统的默认交换机，所以本质上应该和routing模式时一样的