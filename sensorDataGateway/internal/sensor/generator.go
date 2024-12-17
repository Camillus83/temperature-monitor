package sensor

import (
	"math/rand"
	"time"
)

const (
	MinTemperature = 15.0
	MaxTemperature = 30.0
	Interval       = 1 * time.Minute
)

type dataMeasurement struct {
	sensorId  string
	timestamp time.Time
	value     float64
}

type DataGeneratorCfg struct {
	MaxTemp  float64
	MinTemp  float64
	Interval time.Duration
}

type DataGenerator struct {
	Cfg       DataGeneratorCfg
	SensorIds []string
}

func generateTemperature() float64 {
	return MinTemperature + rand.Float64()*(MaxTemperature-MinTemperature)
}

func (d DataGenerator) GenerateData() []dataMeasurement {
	var data []dataMeasurement
	for _, v := range d.SensorIds {
		measurement := dataMeasurement{
			sensorId:  v,
			timestamp: time.Now(),
			value:     generateTemperature(),
		}
		data = append(data, measurement)
	}
	return data
}
