// Antonio Pujol y Aisha Gandarova

package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Empty struct{} //struct sense camps zero bytes

func main() {
	// conecta con el servidor RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// y define el canal necesario
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	plat, err := ch.QueueDeclare( // cola para los sushis
		"plat", // name
		true,   // durable  // maybe cambiar esto luego
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	fmt.Println("Bon vesper, vinc a sopar de sushi")
	fmt.Println("Ho vull tot!")

	msgSushis, err := ch.Consume( // va a leer los mensajes de la cola de sushis
		plat.Name, // queue
		"",        // consumer
		false,     // auto-ack  // usamos mensajes ack manualmente
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	finaliza := make(chan bool)
	go func() {
		for d := range msgSushis {
			if string(d.Body) == "menjar" {
				fmt.Println("holaa menjaaa")
				finaliza <- true
			} else {
				fmt.Println("holaa ", d.Body)
				finaliza <- true
			}
		}
	}()
	<-finaliza

}
