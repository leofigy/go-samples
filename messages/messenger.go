package messages

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Messenger struct {
	URL          string
	QueueName    string
	IsAutoDelete bool
	IsDurable    bool
	IsExclusive  bool
	NoWait       bool
	Args         amqp.Table
	Timeout      time.Duration
}

func (m *Messenger) Send() (chan string, error) {

	log.Println("initializing connection")

	conn, err := amqp.Dial(m.URL)

	if err != nil {
		return nil, err
	}

	log.Println("initializing queue")

	ch, err := conn.Channel()

	if err != nil {
		conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(
		m.QueueName,    // name
		m.IsDurable,    // durable
		m.IsAutoDelete, // delete when unused
		m.IsExclusive,  // exclusive
		m.NoWait,       // no-wait
		m.Args,         // arguments
	)

	if err != nil {
		ch.Close()
		return nil, err
	}

	readChan := make(chan string)

	go func() {
		defer conn.Close()
		defer ch.Close()

		for body := range readChan {
			log.Println("waiting to get data")
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

			err = ch.PublishWithContext(ctx,
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				})
			if err != nil {
				log.Println(err)
				continue
			}

		}
	}()

	return readChan, nil
}

func (m *Messenger) Read() {
	log.Println("initializing connection")

	conn, err := amqp.Dial(m.URL)

	if err != nil {
		panic(err)
	}

	log.Println("initializing queue")

	ch, err := conn.Channel()

	if err != nil {
		conn.Close()
		panic(err)
	}

	q, err := ch.QueueDeclare(
		m.QueueName,    // name
		m.IsDurable,    // durable
		m.IsAutoDelete, // delete when unused
		m.IsExclusive,  // exclusive
		m.NoWait,       // no-wait
		m.Args,         // arguments
	)

	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		panic(err)
	}
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
