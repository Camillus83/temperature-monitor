package data

import "database/sql"

type Models struct {
	DataMeasurements interface {
		Insert(dataMeasurement *DataMeasurement) error
		GetAllLastMeasurement() ([]*DataMeasurement, error)
		GetLastMeasurement(sensorId string) (*DataMeasurement, error)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		DataMeasurements: DataMeasurementModel{DB: db},
	}
}
