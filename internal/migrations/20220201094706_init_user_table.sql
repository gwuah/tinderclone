-- +goose Up
CREATE TABLE IF NOT EXISTS users 
(
	id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
	phone_number VARCHAR ( 50 ) NOT NULL,
	otp VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS users ;
