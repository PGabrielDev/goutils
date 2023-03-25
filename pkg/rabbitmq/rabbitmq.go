package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queue, consumer_name string) error {
	msgs, err := ch.Consume(
		queue,
		consumer_name,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}
	for msg := range msgs {
		out <- msg
	}
	return nil
}

func Producer(ch *amqp.Channel, ex, mensagem string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	err := ch.PublishWithContext(ctx,
		ex,
		"",
		false,
		false,
		amqp.Publishing{
			Body: []byte(mensagem),
		})
	if err != nil {
		return err
	}
	return nil
}
