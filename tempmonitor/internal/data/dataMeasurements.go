package data

import (
	"database/sql"
	"time"
)

type DataMeasurement struct {
	SensorId  string    `json:"sensorId"`
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}

type DataMeasurementModel struct {
	DB *sql.DB
}

func (dm DataMeasurementModel) Insert(dataMeasurement *DataMeasurement) error {
	query := `
		INSERT INTO data_measurement (sensorId, temperature, timestamp)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	args := []any{dataMeasurement.SensorId, dataMeasurement.Value, dataMeasurement.Timestamp}
	return dm.DB.QueryRow(query, args...).Scan(&dataMeasurement.SensorId)
}
