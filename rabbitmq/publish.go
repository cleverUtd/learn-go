package main

import (
	"encoding/json"
	"fmt"

	"sync"

	"time"

	"runtime"

	"github.com/streadway/amqp"
)

const amqpURI = "amqp://zclau:zclau@10.0.0.11:5672/"

var (
	amqpConnection *amqp.Connection
)

func publish(amqpChannel *amqp.Channel, queueName string, body *string) error {
	// if err := channel.ExchangeDeclare(
	// 	exchange,     // name
	// 	exchangeType, // type
	// 	true,         // durable
	// 	false,        // auto-deleted
	// 	false,        // internal
	// 	false,        // noWait
	// 	nil,          // arguments
	// ); err != nil {
	// 	return fmt.Errorf("Exchange Declare: %s", err)
	// }

	if err := amqpChannel.Publish(
		"",        // publish to an exchange
		queueName, // routing to 0 or more queues
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(*body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}

	return nil
}

func connectRabbitMQ() *amqp.Connection {
	connection, err := amqp.Dial(amqpURI)
	for err != nil {
		fmt.Println("connect error.", err)
		connection, err = amqp.Dial(amqpURI)
	}
	return connection
}

func main() {
	runtime.GOMAXPROCS(3)

	amqpConnection = connectRabbitMQ()

	requests := 100000
	clients := 50

	bidlogChan := make(chan *string, requests)
	defer close(bidlogChan)

	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < clients; i++ {
		go func(i int) {
			amqpChannel, err := amqpConnection.Channel()
			if err != nil {
				fmt.Println("amqpChannel error.")
				return
			}
			defer amqpChannel.Close()

			for log := range bidlogChan {
				publish(amqpChannel, "task.bid.log.1", log)
				wg.Done()
			}
		}(i)
	}

	bidlog := constructBidlog()
	body, err := json.Marshal(*bidlog)
	if err != nil {
		fmt.Print("err: ", err)
		return
	}

	bidlogString := string(body)
	start := time.Now()

	for i := 0; i < requests; i++ {
		bidlogChan <- &bidlogString
	}

	wg.Wait()
	elapse := time.Now().Sub(start).Seconds()

	fmt.Printf("%0.2f secs elapsed, %0.2f qps", elapse, float64(requests)/elapse)
}

func constructBidlog() *Bidlog {
	bidlog := &Bidlog{}
	bidlog.BidId = "111111"
	bidlog.BidPrice = 9000
	bidlog.TaskId = "111"
	bidlog.ImpId = "testtest"
	bidlog.IsWon = 0
	bidlog.WinPrice = 0
	bidlog.ConType = "2"
	bidlog.Os = "ios"
	bidlog.Osv = "ios"
	bidlog.Brand = "ios"
	return bidlog
}

type Bidlog struct {
	Id        int64
	BidId     string
	ImpId     string
	TaskId    string
	BidPrice  int
	WinPrice  int
	IsWon     int
	Device    string
	Age       int
	Gender    string
	Location  string
	ConType   string
	Os        string
	Osv       string
	Brand     string
	Geo       string
	Ip        string
	Idfa      string
	WaxUserid string
}
