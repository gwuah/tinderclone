-- +goose Up
ALTER TABLE users 
    ADD otp_created_at TIMESTAMP;

-- +goose Down
ALTER TABLE users 
    DROP otp_created_at;
