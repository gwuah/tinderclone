-- +goose Up
alter table users
add first_name  VARCHAR,
add dob         DATE,
add location    VARCHAR ;

-- +goose Down
alter table users
drop first_name,
drop dob,
drop location;