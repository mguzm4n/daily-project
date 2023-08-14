CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    firstname varchar(100) NOT NULL,
    lastname varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

INSERT INTO users (firstname, lastname, email) VALUES ('marcelo', 'guzman', 'admin@mail.com');
ALTER TABLE notes ADD user_id bigint REFERENCES users (id);
UPDATE notes SET user_id = 1 WHERE user_id IS NULL;