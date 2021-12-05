package event

// type HandlerFunc func(context.Context, cloudevents.Event) ek_analytic_service.Response

// // Consumer ...
// type Consumer struct {
// 	ctx          context.Context
// 	consumerName string
// 	topic        string
// 	handler      HandlerFunc
// 	responses    chan ek_analytic_service.Response
// }

// // AddConsumer ...
// func (kafka *Kafka) AddConsumer(consumerName, topic, groupID string, handler HandlerFunc) {
// 	if kafka.consumers[consumerName] != nil {
// 		panic(errors.New("consumer with the same name already exists: " + consumerName))
// 	}

// 	kafka.consumers[consumerName] = &Consumer{
// 		ctx:          kafka.ctx,
// 		consumerName: consumerName,
// 		topic:        topic,
// 		handler:      handler,
// 	}
// }

// // Setup is run at the beginning of a new session, before ConsumeClaim.
// func (c *Consumer) Setup(_ sarama.ConsumerGroupSession) error {
// 	return nil
// }

// // Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// // but before the offsets are committed for the very last time.
// func (c *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
// 	return nil
// }

// // ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// // Once the Messages() channel is closed, the Handler must finish its processing
// // loop and exit.
// func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
// 	count := 0
// 	for message := range claim.Messages() {
// 		count += 1
// 		event := util.MessageToEvent(message)
// 		resp := c.handler(c.ctx, event)

// 		select {
// 		case c.responses <- resp:
// 		case <-c.ctx.Done():
// 			return c.ctx.Err()
// 		case <-time.After(time.Second * 5):
// 		}

// 	}

// 	fmt.Println(count)
// 	if count == 0 {
// 		return errors.New("Consumer closed")
// 	} else {
// 		return nil
// 	}
// }
