-- +goose Up
ALTER TABLE users 
    ADD last_name     VARCHAR,
    ADD interests     VARCHAR,
    ADD bio           VARCHAR,
    ADD gender        VARCHAR,
    add profile_photo VARCHAR;

-- +goose Down
ALTER TABLE users 
    DROP last_name,
    DROP interests,
    DROP bio,       
    DROP gender
    DROP profile_photo;


    
