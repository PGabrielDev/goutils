package main

import (
	"fmt"
	"github.com/PGabrielDev/pggm-goutils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, _ := rabbitmq.OpenChannel()
	defer ch.Close()

	msgsOrder := make(chan amqp.Delivery)
	msgsOrders2 := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgsOrder, "orders", "orders")
	go rabbitmq.Consume(ch, msgsOrders2, "orders2", "oders2")

	go func() {
		for msg := range msgsOrder {
			fmt.Println(string(msg.Body))
			msg.Ack(false)
		}
	}()

	for msg := range msgsOrders2 {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}

}
