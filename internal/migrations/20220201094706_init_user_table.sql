-- +goose Up
CREATE TABLE IF NOT EXISTS users 
(
	id 				UUID 		PRIMARY KEY DEFAULT gen_random_uuid(),	
	phone_number 	VARCHAR(50) NOT NULL,
	otp 			VARCHAR 	NOT NULL,
	created_at 		TIMESTAMP 	NOT NULL
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +goose Down
DROP TABLE IF EXISTS users ;
