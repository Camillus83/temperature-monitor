package data

import "time"

type DataMeasurement struct {
	SensorId  string    `json:"sensorId"`
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}
