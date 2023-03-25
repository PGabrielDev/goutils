package main

import (
	"fmt"
	"github.com/PGabrielDev/pggm-goutils/pkg/rabbitmq"
)

func main() {
	ch, _ := rabbitmq.OpenChannel()

	err := rabbitmq.Producer(ch, "orders.ex.direct", "Ola a todos que estou ouvindo essa mensagem")
	if err != nil {
		fmt.Println(err)
	}

}
