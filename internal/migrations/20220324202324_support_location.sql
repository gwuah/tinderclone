-- +goose Up
CREATE EXTENSION postgis;

ALTER TABLE users
    ADD first_name  VARCHAR,
    ADD dob         DATE,
    ADD location    geography(point);

CREATE INDEX users_location ON users USING GIST (location);

-- +goose Down
ALTER TABLE users
    DROP first_name,
    DROP dob,
    DROP location;

DROP INDEX users_location;

DROP EXTENSION IF EXISTS postgis cascade;