package main

import (
	"net/http"

	"tempMonitor/internal/data"
	"tempMonitor/views/pages"
)

func (app *application) dashboardHandler(w http.ResponseWriter, r *http.Request) {
	dataMeasurements, err := app.models.DataMeasurements.GetAllLastMeasurement()
	if err != nil {
		app.logger.Error("err", "err", err)

	}
	app.logger.Info("oki")

	for _, dm := range dataMeasurements {
		app.logger.Info("dataMeasurement", "data", *dm)
	}
	var sensorData []data.DataMeasurement
	for _, dm := range dataMeasurements {
		sensorData = append(sensorData, *dm)
	}

	props := pages.DashboardPageProps{
		SensorData: sensorData,
	}

	err = pages.Home(props).Render(r.Context(), w)
	if err != nil {

	}

}
