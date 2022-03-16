-- +goose Up
alter table users add country_code VARCHAR(4);

-- +goose Down
alter table users drop country_code;