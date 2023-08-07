CREATE TABLE IF NOT EXISTS notes (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    content text NOT NULL,
    version integer NOT NULL DEFAULT 1
);