-- public.cannabis definition
-- Drop table
DROP TABLE IF EXISTS auth;
CREATE TABLE auth (
    id serial PRIMARY KEY,
    nickname VARCHAR NOT NULL UNIQUE,
    account_id VARCHAR NOT NULL UNIQUE,
    jwt VARCHAR NOT NULL UNIQUE,
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
