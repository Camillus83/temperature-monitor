package main

import (
	"encoding/json"
	"tempMonitor/internal/data"
)

func (app *application) listenToRabbitMQ(queueName string) {
	ch, err := app.rabbitmq.Channel()
	if err != nil {
		app.logger.Error("Failed to open a channel", "err", err)
		return
	}

	defer ch.Close()

	_, err = ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		app.logger.Error("Failed to declare a queue", "err", err)
		return
	}

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		app.logger.Error("Failed to register a consumer", "err", err)
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			app.logger.Info("Received a message", "body", string(d.Body))
			var sensorData []data.DataMeasurement
			err := json.Unmarshal(d.Body, &sensorData)
			if err != nil {
				app.logger.Error("Failed to unmarshal message", "err", err)
				continue
			}

			for _, data := range sensorData {
				err := app.models.DataMeasurements.Insert(&data)
				if err != nil {
					app.logger.Error("Failed to insert data measurement", "err", err)
					continue
				}

			}
		}
	}()

	app.logger.Info("listening for messages on queue", "queue", queueName)
	<-forever

}
