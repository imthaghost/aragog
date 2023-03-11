package main

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial("amqp://ghost:hitherelol@localhost:5672/")
	if err != nil {
		log.Println(err)
	}


	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		log.Println(err)
	}


	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = channelRabbitMQ.QueueDeclare(
		"hello", // queue name
		true,            // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		log.Println(err)
	}
	go forever(channelRabbitMQ)
	select {} // block forever
	//client := pb.NewAragogProtobufClient("http://localhost:8080", &http.Client{})
	//// health check
	//health, err := client.HealthCheck(context.Background(), &pb.HealthReq{})
	//if err == nil {
	//	log.Println(health.Status)
	//}
}
func forever(channelRabbitMQ *amqp.Channel) {
	for {
		// Create a message to publish.
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%v+\n", time.Now())),
		}

		// Attempt to publish a message to the queue.
		err := channelRabbitMQ.Publish(
			"",              // exchange
			"hello", // queue name
			false,           // mandatory
			false,           // immediate
			message,         // message to publish
		)
		if err != nil {
			log.Println(err)
		}

		log.Println("Sent")
		time.Sleep(time.Second)
	}
}