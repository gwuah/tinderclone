-- +goose Up
CREATE EXTENSION postgis;
alter table users
alter column location type geography
USING location::geography;;

-- +goose Down
alter table column location type string;