package data

import "database/sql"

type Models struct {
	DataMeasurements interface {
		Insert(dataMeasurement *DataMeasurement) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		DataMeasurements: DataMeasurementModel{DB: db},
	}
}
