// Antonio Pujol y Aisha Gandarova
// enlace video:

package main

import (
	"fmt"
	"log"
	"math/rand"
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

func main() {
	// conecta con el servidor RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// y define el canal necesario
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Inicialización de los valores y prints informativos
	fmt.Println("El cuiner de shushi ja és aquí")
	fmt.Println("El cuiner prepararà un plat amb:")
	rand.Seed(time.Now().UTC().UnixNano())
	sushis[0] = rand.Intn(peces)
	sushis[1] = rand.Intn(peces - sushis[0])
	sushis[2] = peces - sushis[0] - sushis[1]

	fmt.Println(sushis[0], " peces amb ", nomSushis[0])
	fmt.Println(sushis[1], " peces amb ", nomSushis[1])
	fmt.Println(sushis[2], " peces amb ", nomSushis[2])

	plat, err := ch.QueueDeclare( // cola para los sushis
		"plat", // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	permis, err := ch.QueueDeclare( // cola para los permisos (roundRobin)
		"permis", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	var contador = 0
	for i := 0; i < tipusSushis; i++ {
		for j := 0; j < sushis[i]; j++ {
			contador++
			err = ch.Publish( // enviamos las piezas de sushi a la cola de sushis
				"",        // exchange
				plat.Name, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(nomSushis[i]),
				})
			failOnError(err, "Failed to publish a message")

			log.Println("Posa dins el plat ", nomSushis[i])
		}
	}
	if contador == 10 {
		err = ch.Publish( // enviamos el permiso a la cola de permisos
			"",          // exchange
			permis.Name, // routing key
			false,       // mandatory
			false,       // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("10"), // El permiso contiene el número de sushis que hay en la cola
			})
		failOnError(err, "Failed to publish a message")
	}

}
