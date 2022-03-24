-- +goose Up
CREATE EXTENSION postgis;
alter table users
add first_name  VARCHAR,
add dob         DATE,
add longitude   float,
add latitude    float,
add location    geography;
create index on users using gist (location);

-- +goose Down
alter table users
drop first_name,
drop dob,
drop location;
DROP EXTENSION IF EXISTS postgis cascade;