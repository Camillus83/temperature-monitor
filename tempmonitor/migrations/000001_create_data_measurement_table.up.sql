CREATE TABLE data_measurement (
    id SERIAL PRIMARY KEY,
    temperature DECIMAL(5, 2) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);