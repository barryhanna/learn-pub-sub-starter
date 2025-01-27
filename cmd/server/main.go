package main

import (
	"fmt"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")
	connectionString := "amqp://guest:guest@localhost:61613/"
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		fmt.Println("Problem connecting to RabbitMQ")
		return
	}
	defer conn.Close()
	fmt.Println("Successfully connected.")
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
