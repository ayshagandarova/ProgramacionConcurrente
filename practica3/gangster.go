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

	missatge, err := ch.QueueDeclare( // cola para los sushis
		"missatge_gangster", // name
		true,                // durable  // maybe cambiar esto luego
		true,                // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	/*counter, err := ch.QueueDeclare( // cola para los sushis
		"counter", // name
		true,      // durable  // maybe cambiar esto luego
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")*/

	err = ch.ExchangeDeclare(
		"permis", // name
		"fanout", // type
		true,     // durable
		true,     // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	fmt.Println("Bon vesper, vinc a sopar de sushi")
	fmt.Println("Ho vull tot!")

	msgSushis, err := ch.Consume( // va a leer los mensajes de la cola de sushis
		plat.Name, // queue
		"",        // consumer
		true,      // auto-ack  // usamos mensajes ack manualmente
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	msgMissatge, err := ch.Consume( // va a leer los mensajes de la cola encarrec
		missatge.Name, // queue
		"",            // consumer
		false,         // auto-ack  // usamos mensajes ack manualmente
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	failOnError(err, "Failed to register a consumer")

	err = ch.QueueBind(
		missatge.Name,
		"",
		"permis",
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	finaliza := make(chan bool)
	var flag = false
	var counter = 1
	/*go func() {
		for m := range msgMissatge {
			if m.RoutingKey == "" {
				flag = true
				fmt.Println("Gangster Menjar", string(m.Body))
				//m.Ack(false)
			}
			fmt.Println("Gangster Menjar post if", string(m.Body))
		}
	}()*/

	go func() {

		fmt.Println("Gangster Comença ", flag, counter)

		for p := range msgMissatge {
			p.Ack(false)
			for m := range msgSushis {

				m.Ack(true)

				//plat.purge(short reserved-1, queue-name queue, no-wait no-wait)
				purged, err := ch.QueuePurge(plat.Name, true)
				failOnError(err, "Failed to purge a queue")

				fmt.Println("he menjat ", string(m.Body))

				fmt.Println("Gangster Consumeix tot ", string(purged))
				//m.Ack(false)

				ch.QueueDelete(missatge.Name, false, false, false)

				finaliza <- true

			}
		}

	}()
	<-finaliza

}
