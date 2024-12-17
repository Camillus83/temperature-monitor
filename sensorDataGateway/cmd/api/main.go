package main

import (
	"fmt"
	"sensor-data-gateway/internal/sensor"
)

func main() {

	dataGeneratorCfg := sensor.DataGeneratorCfg{
		MaxTemp:  sensor.MaxTemperature,
		MinTemp:  sensor.MinTemperature,
		Interval: sensor.Interval,
	}

	sensorIds := []string{"sensor1", "sensor2"}

	dataGenerator := sensor.DataGenerator{
		Cfg:       dataGeneratorCfg,
		SensorIds: sensorIds,
	}

	abc := dataGenerator.GenerateData()

	for _, v := range abc {
		fmt.Println(v)
	}

}
