// Antonio Pujol y Aisha Gandarova
// enlace video:

package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

const peces = 10
const tipusSushis = 3

var sushis [tipusSushis]int
var nomSushis = [tipusSushis]string{
	"niguiri de salmó",
	"shashimi de tonyina",
	"maki de cranc",
}

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

	// Se identifica
	fmt.Println("El cuiner de shushi ja és aquí")
	fmt.Println("El cuiner prepararà un plat amb:")
	rand.Seed(time.Now().UTC().UnixNano())
	sushis[0] = rand.Intn(peces)
	sushis[1] = rand.Intn(peces - sushis[0])
	sushis[2] = rand.Intn(peces - sushis[0] - sushis[1])

	fmt.Println(sushis[0], " peces amb ", nomSushis[0])
	fmt.Println(sushis[1], " peces amb ", nomSushis[1])
	fmt.Println(sushis[2], " peces amb ", nomSushis[2])

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

	msgs, err := ch.Consume( // va a leer los mensajes de la cola encarrec
		encarrec.Name, // queue
		"",            // consumer
		false,         // auto-ack  // usamos mensajes ack manualmente
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			var contador = 0
			for i := 0; i < tipusSushis; i++ {
				for j := 0; j < sushis[i]; j++ {
					contador++
					err = ch.Publish(
						"",        // exchange
						plat.Name, // routing key
						false,     // mandatory
						false,     // immediate
						amqp.Publishing{
							ContentType: "text/plain",
							Body:        []byte(strconv.Itoa(contador)),
						})
					failOnError(err, "Failed to publish a message")

					log.Printf("Posa dins el plat ", nomSushis[j])
				}
			}
			if contador == 10 {
				err = ch.Publish(
					"",        // exchange
					plat.Name, // routing key
					false,     // mandatory
					false,     // immediate
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte("menjar"),
					})
				failOnError(err, "Failed to publish a message")
			}
			d.Ack(false) //confirma una entrega una vez haya acabado la tarea
		}
	}()

}
