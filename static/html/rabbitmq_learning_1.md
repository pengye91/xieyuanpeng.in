# RabbitMQ学习 1

基本情况

### 信道（`channel`）的重要性

### Message Acknowledgement

In order to make sure a message is never lost, RabbitMQ supports message acknowledgments. An ack(nowledgement) is sent back by the consumer to tell RabbitMQ that a particular message had been received, processed and that RabbitMQ is free to delete it.

`Ack`机制通过让`Consumer`处理完任务后返回一个`Ack`来通知`broker`：

- 如果收到`Ack`，则`broker`可以将任务删除。
- 如果`Consumer`链接断了，则`broker`立即将这个任务发送给另一个`Consumer`。

**重点：**

在`RabbitMQ`中的消息没有`timeouts`的概念，`broker`会一直等待，直到对应的`Consumer`死掉了。

**显示调用`basic_ack`非常重要，如果忘记的话，`RabbitMQ`会一直保留这个任务，然后占用内存会越来越大。**

使用一下命令可以查看所有没有被确认的消息：

	sudo rabbitmqctl list_queues name messages_ready messages_unacknowledged


### 消息持久

`message ack`可以确保`消费者`挂掉的情况消息也可以被处理。

`message durability`可以确保`broker`挂掉的情况，消息也可以被正常处理。

消息持久化分成两个部分：

- 队列`queue`的持久
	
	- 在声明队列的时候，就声明持久化：
		
			channel.queue_declare(queue='task_queue', durable=True)
	
		队列的持久化声明需要同时在`消费者`处和`生产者`处都进行。

- 消息`message`的持久化

	`生产者`在发布消息的时候，给消息加上`delivery_mode = 2`的属性。
	
	**注意**
	
	在消息上加`delivery_mode`属性并不能完全保证所有的消息都被持久化，因为`broker`会先将消息缓存。
	更强的保证可以用`publisher confirms`模式。
	

### 确认机制

消息传递确认机制总的来说可以分为两类(其实就是从消息的接受方向发送方确认)：

- 从`消费者`到`broker`的确认机制：
	**前文的`Message Ack`**

- 从`broker`到`生产者`的确认机制：
	**publisher confirm**

### Consumer confirmation

**针对的都是Channel**

#### Delivery Tags

`Broker`将消息发送到`消费者`的时候会生成一个在本`channel`内唯一的`delivery_tag`。

客户端通过这个`delivery_tag`来识别一次发送。

#### multiple

在参数中设置`multiple = True`可以使在该delivery_tag之前所有没有ack的消息被Ack掉。

#### QoS

由于天然的异步性质，在消费者返回`ACK`的时候可能会有一些`ACK`在路上。

为了克服这种不确定性，可以在一条`Channel`上设置一个允许未被`ACK`的消息数量 using `basic.qos`方法。

一旦`unack`的消息数量超过这个值，`broker`就不会继续发送了，直到有消息被确认了。

**注意**
在运行的过程中尽量不要改动`QoS`的值，可能会引起竞争。

### Publisher Confirmation

概念和`Consumer Confirmation`类似，只是发送`basic.ack`的变成了`broker`。

**针对的都是Channel**

当使用了`队列持久化`技术和`消息持久化`技术的时候， `Publisher Ack` 只有在消息被写入磁盘的时候才会被`broker`发送回`Publisher`。


