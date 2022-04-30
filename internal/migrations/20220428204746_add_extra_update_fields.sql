-- +goose Up
ALTER TABLE users 
    ADD last_name  VARCHAR,
    ADD interests  VARCHAR,
    ADD bio        VARCHAR,
    ADD gender     VARCHAR;

-- +goose Down
ALTER TABLE users 
    DROP last_name,
    DROP interests,
    DROP bio,       
    DROP gender;    

    
