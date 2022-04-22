-- +goose Up
ALTER TABLE users 
    ADD raw_otp VARCHAR(5);

-- +goose Down
ALTER TABLE users 
    DROP raw_otp;
