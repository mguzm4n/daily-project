CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    firstname varchar(100) NOT NULL,
    lastname varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS notes (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    content text NOT NULL,
    version integer NOT NULL DEFAULT 1,
    user_id bigint NOT NULL REFERENCES users (id)
);

INSERT INTO users (firstname, lastname, email) VALUES ('marcelo', 'guzman', 'admin@mail.com');

