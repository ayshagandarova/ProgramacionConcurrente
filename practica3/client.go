// Laura Cavero y Aisha Gandarova
// enlace video: https://www.dropbox.com/s/m9cqa4cllo5tew2/Pr%C3%A1ctica3.mp4?dl=0

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Empty struct{} //struct sense camps zero bytes

func main() {
	// cconecta con el servidor RabbitMQ
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

	encarrec, err := ch.QueueDeclare( // cola para pedir sushi al cuiner
		"encarrec", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	rand.Seed(time.Now().UTC().UnixNano())
	var peces = rand.Intn(20)
	fmt.Println("Bon vesper, vinc a sopar de sushi")
	fmt.Println("Avuir menajare ", peces, " peces")

	err = ch.Publish(
		"",            // exchange
		encarrec.Name, // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("vull"),
		})

	msgSushis, err := ch.Consume( // va a leer los mensajes de la cola encarrec
		plat.Name, // queue
		"",        // consumer
		false,     // auto-ack  // usamos mensajes ack manualmente
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgSushis {
			if string(d.Body) == "menjar" {

			} else {

			}
		}
	}()

}
