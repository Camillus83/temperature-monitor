package views

import (
	"strconv"
	"time"
)

func FormatTimestamp(timestamp time.Time) string {
	return timestamp.Format("15:04 02.01.2006")
}

func FormatTemperature(temperatue float64) string {
	return strconv.FormatFloat(temperatue, 'f', 2, 64)
}
