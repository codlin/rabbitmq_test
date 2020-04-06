package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.Qos(1, 0, false)
	failOnError(err, "Failed to set Qos")

	err = ch.ExchangeDeclare(
		"logs_topic_2_workerqueue", // name
		"topic",                    // type
		true,                       // durable
		false,                      // auto-deleted
		false,                      // internal
		false,                      // no-wait
		nil,                        // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	body := bodyFrom(os.Args)
	for i := 0; i < 10; i++ {
		err = ch.Publish(
			"logs_topic_2_workerqueue", // exchange
			severityFrom(os.Args),      // routing key
			false,                      // mandatory
			false,                      // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(strconv.Itoa(i) + "-->" + body),
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [%d] Sent %s", i, body)
	}
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}
