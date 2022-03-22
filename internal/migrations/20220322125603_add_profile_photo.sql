-- +goose Up
alter table users add profile_photo VARCHAR;

-- +goose Down
alter table users drop profile_photo;