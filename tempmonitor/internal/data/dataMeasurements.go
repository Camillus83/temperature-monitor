package data

import "time"

type DataMeasurement struct {
	sensorId  string
	timestamp time.Time
	value     float64
}
