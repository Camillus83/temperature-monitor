package sensor

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func SimulateTemperatureMeasurement(dg DataGenerator, ch *amqp091.Channel, queueName string) {

	ticker := time.NewTicker(dg.Cfg.Interval)
	for range ticker.C {
		data := dg.GenerateData()
		js, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			fmt.Println("err")
			continue
		}

		err = ch.Publish(
			"",
			queueName,
			false,
			false,
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        js,
			},
		)

		if err != nil {
			log.Fatalf("Error publishing")
			continue
		}

	}

}
