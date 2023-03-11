package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/imthaghost/aragog/config"
	"github.com/imthaghost/aragog/internal/bus"
	"github.com/imthaghost/aragog/internal/clients/monopoly"
	errs "github.com/imthaghost/aragog/internal/errors"
	"github.com/imthaghost/aragog/internal/logger"
	"time"

	"github.com/kr/pretty"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Event struct {
	EventType  string                 `json:"event_type"`
	EventData  interface{}            `json:"event_data"`
	Properties map[string]interface{} `json:"properties"`
}
type RabbitMQ struct {
	Conn         *amqp.Connection
	Config       *config.Config
	ErrorMonitor errs.Service
	Logger       logger.Service
	Monopoly     monopoly.ClientWrapper
}

func (r *RabbitMQ) Setup() {
	cfg := r.Config.RabbitMQ

	// build URI
	uri := fmt.Sprintf("amqp://%s:%s@%s/", cfg.Username, cfg.Password, cfg.Host)

	// create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(uri)
	if err != nil {
		r.ErrorMonitor.Report(err)
		r.Logger.Error("Failed to connect to RabbitMQ")

		return
	}

	// good connection
	r.Conn = connectRabbitMQ
	r.Logger.Msg("Connected to RabbitMQ")

}

// Ping will check to see if the RabbitMQ instance is alive.
func (r *RabbitMQ) Ping() error {
	if r.Conn == nil {
		return errors.New("Empty connection")
	}

	// Ping the RabbitMQ instance.
	return nil
}

// Consume will start consuming messages from a given RabbitMQ queue
func (r *RabbitMQ) Consume() {
	err := r.Ping()
	if err != nil {
		r.ErrorMonitor.Report(err)
		r.Logger.Error("Empty Connection")

		return
	}
	// Opening a channel to our RabbitMQ instances over
	// the connection we have already established.
	channelRabbitMQ, err := r.Conn.Channel()
	if err != nil {
		r.ErrorMonitor.Report(err)
		r.Logger.Error("Failed to open a channel to RabbitMQ")
	}

	// Subscribing to QueueService1 for getting messages.
	messages, err := channelRabbitMQ.Consume(
		"aragog", // queue name
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no local
		false,    // no wait
		nil,      // arguments
	)
	if err != nil {
		r.ErrorMonitor.Report(err)
		r.Logger.Error("Failed to consume from RabbitMQ")
	}

	// Build a welcome message.
	r.Logger.Msg("Waiting for messages")

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)
	// errorChan := make(chan error)
	event := &Event{}

	go func() {
		for message := range messages {
			r.Logger.Msg(fmt.Sprintf(" > Received message: %s\n", string(message.Body)))
			err := json.Unmarshal(message.Body, &event)
			if err != nil {
				r.ErrorMonitor.Report(err)
				r.Logger.Error("Failed to unmarshal message")
			}

			// TODO add user to tradingview
			_, _ = pretty.Println(event)
			username := event.Properties.(string)

			_, _ = pretty.Println(username)

			time.Sleep(time.Second * 1)
			resp, err := r.Monopoly.SendMessage("hello")
			if err != nil {
				r.ErrorMonitor.Report(err)
				r.Logger.Error("Failed to send message to monopoly")

				return
			}

			r.Logger.Msg(resp)

		}
	}()

	<-forever
}

// Shutdown will close the RabbitMQ connection
func (r *RabbitMQ) Shutdown() error {
	err := r.Conn.Close()
	if err != nil {

		return err
	}

	return nil
}

// NewService will create a new RabbitMQ client
func NewService(cfg *config.Config, ls logger.Service, errMonitor errs.Service, monopolyClient monopoly.ClientWrapper) bus.Service {
	rabbit := &RabbitMQ{
		Config:       cfg,
		ErrorMonitor: errMonitor,
		Logger:       ls,
		Monopoly:     monopolyClient,
	}

	rabbit.Setup()

	return rabbit
}
