-- +goose Up
ALTER TABLE users 
    ADD country_code VARCHAR(4);

-- +goose Down
ALTER TABLE users 
    DROP country_code;