// Antonio Pujol y Aisha Gandarova

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

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
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	permis, err := ch.QueueDeclare( // cola para los permisos
		"permis", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgPermis, err := ch.Consume( // msgPermis contiene los mensajes de la cola de permisos
		permis.Name, // queue
		"",          // consumer
		false,       // auto-ack  // usamos mensajes ack manualmente
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	failOnError(err, "Failed to register a consumer")

	err = ch.Qos(
		1,    // prefetch count
		0,    // prefetch size
		true, // global
	)
	failOnError(err, "Failed to set QoS")

	finaliza := make(chan bool)

	fmt.Println("Bon vesper, vinc a sopar de sushi")
	fmt.Println("Ho vull tot!")

	go func() {
		elem := 0 // guardamos cuantos sushis ha consumido el gangster del plato
		for p := range msgPermis {
			p.Ack(false)
			elem, err = strconv.Atoi(string(p.Body))
			break
		}
		_, err := ch.QueueDelete(plat.Name, false, false, false)
		failOnError(err, "Failed to delete queue")

		fmt.Println("Gangster Consumeix tot ", elem)
		fmt.Println("Agafa totes les peces")
		fmt.Println("Romp el plat")
		fmt.Println("Men vaig")
		finaliza <- true

	}()

	<-finaliza

}
