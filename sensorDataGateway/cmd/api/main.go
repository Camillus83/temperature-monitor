package main

import (
	"log"
	"sensor-data-gateway/internal/sensor"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a RabbitMQ channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"temperature_measurements",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to declare queue")
	}

	dataGenerator := sensor.DataGenerator{
		Cfg: sensor.DataGeneratorCfg{
			MaxTemp:  sensor.MaxTemperature,
			MinTemp:  sensor.MinTemperature,
			Interval: time.Second * 5,
		},
		SensorIds: []string{"sensor1", "sensor2"},
	}

	sensor.SimulateTemperatureMeasurement(dataGenerator, ch, q.Name)

}
