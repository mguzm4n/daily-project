ALTER TABLE users
ADD UNIQUE (email),
ADD COLUMN IF NOT EXISTS password_hash bytea NOT NULL DEFAULT ''::bytea,
ADD COLUMN IF NOT EXISTS activated bool NOT NULL DEFAULT false;