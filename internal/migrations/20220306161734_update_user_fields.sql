-- +goose Up
CREATE EXTENSION postgis;
alter table users
add first_name  VARCHAR,
add dob         DATE,
add location    geography;

-- +goose Down
alter table users
drop first_name,
drop dob,
drop location;
DROP EXTENSION IF EXISTS postgis cascade;