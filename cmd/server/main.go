package main

import (
	"fmt"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("Error dialing ")
	}
	defer conn.Close()

	fmt.Printf("Succesfully dialed")

	channel, err := conn.Channel()
	if err != nil {
		fmt.Printf("Error creating a channel")
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	select {
	case <-signalChan:
		fmt.Printf("Exiting...")
	}
}
