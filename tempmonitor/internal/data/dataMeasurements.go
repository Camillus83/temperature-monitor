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

func (dm DataMeasurementModel) GetLastMeasurement(sensorId string) (*DataMeasurement, error) {
	query := `
		SELECT t1.sensorId, t1.timestamp, t1.temperature
		FROM data_measurement
		WHERE sensorId = $1
		ORDER BY timestamp DESC
		LIMIT 1
	`
	var dataMeasurement DataMeasurement
	err := dm.DB.QueryRow(query, sensorId).Scan(
		&dataMeasurement.SensorId,
		&dataMeasurement.Timestamp,
		&dataMeasurement.Value,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &dataMeasurement, nil
}

func (dm DataMeasurementModel) GetAllLastMeasurement() ([]*DataMeasurement, error) {
	query := `
		SELECT t1.sensorId, t1.timestamp, t1.temperature
		FROM data_measurement t1
		INNER JOIN (
			SELECT 
				sensorId,
				MAX(timestamp) AS MaxTimestamp
			FROM
				data_measurement
			GROUP BY
				sensorId
		) t2
	ON t1.sensorId = t2.sensorId AND t1.timestamp = t2.MaxTimestamp
	`
	rows, err := dm.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dataMeasurements := []*DataMeasurement{}

	for rows.Next() {
		var dataMeasurement DataMeasurement
		err := rows.Scan(
			&dataMeasurement.SensorId,
			&dataMeasurement.Timestamp,
			&dataMeasurement.Value,
		)
		if err != nil {
			return nil, err
		}
		dataMeasurements = append(dataMeasurements, &dataMeasurement)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dataMeasurements, nil

}
