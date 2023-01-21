// Antonio Pujol y Aisha Gandarova

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

	missatge, err := ch.QueueDeclare( // cola para los sushis
		"",    // name
		true,  // durable  // maybe cambiar esto luego
		true,  // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
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

	rand.Seed(time.Now().UTC().UnixNano())
	var peces = rand.Intn(10)
	fmt.Println("Bon vesper, vinc a sopar de sushi")
	fmt.Println("Avuir menajare ", peces, " peces")

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

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	finaliza := make(chan bool)
	var flag = false
	var counter = 1

	/*go func() {

		for m := range msgMissatge {
			if m.RoutingKey == "" {
				flag = true
				fmt.Println("Cliente1 Menjar", string(m.Body))
				//m.Ack(false)
			}
			fmt.Println("Cliente1 Menjar post if", string(m.Body))
		}
	}()*/

	go func() {

		fmt.Println("Cliente1 Comença ", flag, counter)
		for p := range msgMissatge {
			p.Ack(false)
			fmt.Println("Cliente1 Menjar", string(p.Body))
			for m := range msgSushis {
				if m.RoutingKey == plat.Name {
					log.Println("El client ha agafat ", string(m.Body))

					m.Ack(false)

					fmt.Printf("Al plat hi ha %d peces al plat.\n", plat.Messages)

					if counter == peces {

						ch.QueueDelete(missatge.Name, false, false, true)

						finaliza <- true
						//break
					}

					counter++

					time.Sleep(1 * time.Second)
				}

			}
		}

		/*var dlvry = <-msgMissatge
		//dlvry.Ack(true)
		if string(dlvry.Body) == "menjar" {
			fmt.Println("holaa menjaaa")
			flag = true
			for i := 0; i < peces; i++ {
				dlvry = <-msgSushis
				if flag {
					fmt.Println("holaa ", string(dlvry.Body))
					//dlvry.Ack(true)
				}
			}
			finaliza <- true
		}*/
	}()

	<-finaliza
}
