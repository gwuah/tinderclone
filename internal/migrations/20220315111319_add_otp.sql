-- +goose Up
alter table users add raw_otp VARCHAR(5);

-- +goose Down
alter table users drop raw_otp;
