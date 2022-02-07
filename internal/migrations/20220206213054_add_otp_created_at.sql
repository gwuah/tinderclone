-- +goose Up
alter table users add otp_created_at TIMESTAMP;

-- +goose Down
alter table users drop otp_created_at;
