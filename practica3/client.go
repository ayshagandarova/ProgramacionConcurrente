// Antonio Pujol y Aisha Gandarova

package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
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

	permis, err := ch.QueueDeclare( // cola para los sushis
		"permis", // name
		true,     // durable  // maybe cambiar esto luego
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

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

	msgPermis, err := ch.Consume( // va a leer los mensajes de la cola encarrec
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
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	finaliza := make(chan bool)

	go func() {
		rand.Seed(time.Now().UTC().UnixNano())
		var peces = rand.Intn(15)
		peces++ // para que no sea 0
		fmt.Println("Bon vesper, vinc a sopar de sushi")
		fmt.Println("Avui menajare ", peces, " peces")
		quedan := 0
		for i := 0; i < peces; i++ {
			for p := range msgPermis {
				p.Ack(false)
				quedan, err = strconv.Atoi(string(p.Body))
				break
			}
			quedan--
			for s := range msgSushis {
				s.Ack(false)
				log.Println("El client ha agafat ", string(s.Body))
				fmt.Println("Al plat hi ha ", quedan, " peces")
				time.Sleep(2000)
				break
			}
			time.Sleep(5000)
			if quedan > 0 {
				permisSeguent := strconv.Itoa(quedan)
				err = ch.Publish(
					"",          // exchange
					permis.Name, // routing key
					false,       // mandatory
					false,       // immediate
					amqp.Publishing{
						DeliveryMode: amqp.Persistent,
						ContentType:  "text/plain",
						Body:         []byte(permisSeguent),
					})
				failOnError(err, "Failed to publish the message")
			}
		}

		finaliza <- true

	}()

	<-finaliza
}
