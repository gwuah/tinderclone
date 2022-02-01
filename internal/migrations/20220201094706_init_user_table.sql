-- +goose Up
CREATE TABLE IF NOT EXISTS users 
(
    id serial PRIMARY KEY,
	phone_number VARCHAR ( 50 ) NOT NULL,
	created_at TIMESTAMP NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS users ;
