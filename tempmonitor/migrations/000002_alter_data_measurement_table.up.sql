ALTER TABLE data_measurement
ADD COLUMN created_at timestamp(0) with time zone NOT NULL DEFAULT NOW();

ALTER TABLE data_measurement
ADD COLUMN sensorId text NOT NULL;