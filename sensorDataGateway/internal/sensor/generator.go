package sensor

import (
	"math/rand"
	"sensor-data-gateway/internal/data"

	"time"
)

const (
	MinTemperature = 15.0
	MaxTemperature = 30.0
	Interval       = 1 * time.Minute
)

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

func (d DataGenerator) GenerateData() []data.DataMeasurement {
	var measurementData []data.DataMeasurement
	for _, v := range d.SensorIds {

		measurement := data.DataMeasurement{
			SensorId:  v,
			Timestamp: time.Now(),
			Value:     generateTemperature(),
		}
		measurementData = append(measurementData, measurement)
	}
	return measurementData
}
