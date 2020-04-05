# 测试说明

测试在标准routing模式下（routing类型为direct），多个消费者绑定相同的routing key，以测试此种情况是否和pub/sub等效

## 模式特点

同pub/sub

## Routing模式的适用场景

同pub/sub

## 如何实施

1. 交换机的声明和route direct模式相同
2. 在声明队列时**不**指定名字
3. 不同的worker绑定不同的routing key

```go
q, err := ch.QueueDeclare(
		"", // 此处不指定队列名字
		true,
		false,
		false,
		false,
		nil,
	)

// worker1绑定key_direct1
ch.QueueBind(q.Name, "key_direct1", "routing_direct_pubsub", false, nil)

// worker2绑定key_direct2
ch.QueueBind(q.Name, "key_direct2", "routing_direct_pubsub", false, nil)
```

### 测试用例

1. 验证当前设置下是否同pub/sub等效

### 测试结论

1. **在标准routing模式下，消费者不指定队列名字，绑定不同的routing key，则效果相当于pub/sub模式**

## 总结

同pub/sub模式